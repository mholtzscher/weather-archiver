package v1

import (
	"errors"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog/log"
)

var errNoResults = errors.New("no results found")

func mapPgErrorsToReturnCodes(err error) *connect.Error {
	var pgErr *pgconn.PgError
	if errors.Is(err, pgx.ErrNoRows) {
		return connect.NewError(connect.CodeNotFound, errNoResults)
	} else if errors.As(err, &pgErr) {
		log.Info().Str("code", pgErr.Code)
		switch pgErr.Code {
		case "23505":
			return connect.NewError(connect.CodeAlreadyExists, pgErr)
		default:
			return connect.NewError(connect.CodeUnknown, pgErr)
		}
	}
	return connect.NewError(connect.CodeUnknown, err)
}
