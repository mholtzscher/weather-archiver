package v1

import (
	"context"

	"connectrpc.com/connect"
	apiv1 "github.com/mholtzscher/weather-archiver/gen/api/v1"
	"github.com/mholtzscher/weather-archiver/internal/dal"
	"github.com/rs/zerolog/log"
)

func (s *FormulaDataServer) CreateSeason(
	ctx context.Context,
	request *connect.Request[apiv1.CreateSeasonRequest],
) (*connect.Response[apiv1.CreateSeasonResponse], error) {
	id, err := s.DB.CreateSeason(ctx, dal.CreateSeasonParams{
		SeasonYear: request.Msg.Year,
		Series:     request.Msg.Series,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to insert to db")
		return nil, mapPgErrorsToReturnCodes(err)
	}
	return &connect.Response[apiv1.CreateSeasonResponse]{
		Msg: &apiv1.CreateSeasonResponse{
			SeasonId: id,
		},
	}, nil
}
