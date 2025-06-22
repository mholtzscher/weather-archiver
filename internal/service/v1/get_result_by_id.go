package v1

import (
	"context"

	"connectrpc.com/connect"
	apiv1 "github.com/mholtzscher/weather-archiver/gen/api/v1"
	"github.com/rs/zerolog/log"
)

func (s *FormulaDataServer) GetResultById(
	ctx context.Context,
	request *connect.Request[apiv1.GetResultByIdRequest],
) (*connect.Response[apiv1.GetResultByIdResponse], error) {
	result, err := s.DB.GetResultById(ctx, request.Msg.ResultId)
	if err != nil {
		log.Error().Err(err).Msg("failed to get result by id")
		return nil, mapPgErrorsToReturnCodes(err)
	}
	return &connect.Response[apiv1.GetResultByIdResponse]{
		Msg: &apiv1.GetResultByIdResponse{
			Result: &apiv1.Result{
				ResultId: result.ID,
				RaceId:   result.RaceID,
				DriverId: result.DriverID,
				TeamId:   result.TeamID,
				Position: result.Position,
				Points:   result.Points,
			},
		},
	}, nil
}
