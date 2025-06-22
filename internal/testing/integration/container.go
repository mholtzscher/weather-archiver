package integration

import (
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	weatherarchiver "github.com/mholtzscher/weather-archiver"
	"github.com/mholtzscher/weather-archiver/gen/api/v1/apiv1connect"
	"github.com/mholtzscher/weather-archiver/internal/dal"
	srvV1 "github.com/mholtzscher/weather-archiver/internal/service/v1"
	"github.com/pressly/goose/v3"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type IntegrationTestHelper struct {
	Client           apiv1connect.FormulaDataServiceClient
	Container        *postgres.PostgresContainer
	ConnectionString string
}

func CreateIntegrationTestHelper(t *testing.T) *IntegrationTestHelper {
	t.Helper()
	ctx := context.Background()
	container, connStr := createPostgresContainer(t, ctx)
	t.Cleanup(func() {
		err := container.Terminate(ctx)
		if err != nil {
			t.Fatal(err)
		}
	})

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		err := conn.Close(ctx)
		if err != nil {
			t.Fatal(err)
		}
	})

	validator, err := validate.NewInterceptor()
	if err != nil {
		t.Fatal(err)
	}

	queries := dal.New(conn)
	fdServer := srvV1.NewFormulaDataServer(queries)

	mux := http.NewServeMux()
	mux.Handle(
		apiv1connect.NewFormulaDataServiceHandler(
			fdServer,
			connect.WithInterceptors(validator),
		),
	)
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	t.Cleanup(server.Close)

	return &IntegrationTestHelper{
		Client:           apiv1connect.NewFormulaDataServiceClient(server.Client(), server.URL),
		Container:        container,
		ConnectionString: connStr,
	}
}

func runMigrations(t *testing.T, db *sql.DB) {
	goose.SetBaseFS(weatherarchiver.MigrationsFileSystem)

	if err := goose.SetDialect("postgres"); err != nil {
		t.Fatal(err)
	}

	if err := goose.Up(db, "sql/migrations"); err != nil {
		t.Fatal(err)
	}
}

func createPostgresContainer(t *testing.T, ctx context.Context) (*postgres.PostgresContainer, string) {
	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:16-alpine"),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatal(err)
	}

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	conn, err := sql.Open("pgx", connStr)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	runMigrations(t, conn)

	return pgContainer, connStr
}
