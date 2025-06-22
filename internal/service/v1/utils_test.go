package v1

import (
	"errors"
	"testing"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
)

func TestMapPgErrorsToReturnCodes(t *testing.T) {
	errorTests := []struct {
		scenario string
		actual   error
		expected connect.Error
	}{
		{"not a PgError", errors.New("hi mom"), *connect.NewError(connect.CodeUnknown, errors.New("test"))},
		{"constraint violation", &pgconn.PgError{Code: "23505"}, *connect.NewError(connect.CodeAlreadyExists, errors.New("test"))},
		{"unsupported pg error", &pgconn.PgError{Code: "23505asdfasdf"}, *connect.NewError(connect.CodeUnknown, errors.New("test"))},
	}

	for _, e := range errorTests {
		t.Run(e.scenario, func(t *testing.T) {
			result := mapPgErrorsToReturnCodes(e.actual)
			assert.Equal(t, e.expected.Code(), result.Code())
		})
	}
}
