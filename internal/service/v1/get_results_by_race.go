package v1

import (
	"context"

	"connectrpc.com/connect"
	apiv1 "github.com/mholtzscher/weather-archiver/gen/api/v1"
	"github.com/rs/zerolog/log"
)

func (s *FormulaDataServer) GetResultsByRace(
	ctx context.Context,
	request *connect.Request[apiv1.GetResultsByRaceRequest],
) (*connect.Response[apiv1.GetResultsByRaceResponse], error) {
	results, err := s.DB.GetResultsByRaceId(ctx, request.Msg.RaceId)
	if err != nil {
		log.Error().Err(err).Msg("failed to get results by race id")
		return nil, mapPgErrorsToReturnCodes(err)
	}

	var resultsResponses []*apiv1.Result
	for _, result := range results {
		resultsResponses = append(resultsResponses, &apiv1.Result{
			ResultId: result.ID,
			RaceId:   result.RaceID,
			DriverId: result.DriverID,
			TeamId:   result.TeamID,
			Position: result.Position,
			Points:   result.Points,
		})
	}

	return &connect.Response[apiv1.GetResultsByRaceResponse]{
		Msg: &apiv1.GetResultsByRaceResponse{
			Results: resultsResponses,
		},
	}, nil
}
