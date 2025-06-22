package v1

import (
	"context"

	"connectrpc.com/connect"
	apiv1 "github.com/mholtzscher/weather-archiver/gen/api/v1"
	"github.com/rs/zerolog/log"
)

func (s *FormulaDataServer) GetSeasonById(
	ctx context.Context,
	request *connect.Request[apiv1.GetSeasonByIdRequest],
) (*connect.Response[apiv1.GetSeasonByIdResponse], error) {
	season, err := s.DB.GetSeasonById(ctx, request.Msg.SeasonId)
	if err != nil {
		log.Error().Err(err).Msg("failed to get season by id")
		return nil, mapPgErrorsToReturnCodes(err)
	}
	return &connect.Response[apiv1.GetSeasonByIdResponse]{
		Msg: &apiv1.GetSeasonByIdResponse{
			Season: &apiv1.Season{
				SeasonId: season.ID,
				Year:     season.SeasonYear,
				Series:   season.Series,
			},
		},
	}, nil
}
