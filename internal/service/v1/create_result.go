package v1

import (
	"context"

	"connectrpc.com/connect"
	apiv1 "github.com/mholtzscher/weather-archiver/gen/api/v1"
	"github.com/mholtzscher/weather-archiver/internal/dal"
)

func (s *FormulaDataServer) CreateResult(
	ctx context.Context,
	request *connect.Request[apiv1.CreateResultRequest],
) (*connect.Response[apiv1.CreateResultResponse], error) {
	id, err := s.DB.CreateResult(ctx, dal.CreateResultParams{
		RaceID:   request.Msg.RaceId,
		DriverID: request.Msg.DriverId,
		TeamID:   request.Msg.TeamId,
		Position: request.Msg.Position,
		Points:   request.Msg.Points,
	})
	if err != nil {
		return nil, mapPgErrorsToReturnCodes(err)
	}

	return &connect.Response[apiv1.CreateResultResponse]{
		Msg: &apiv1.CreateResultResponse{
			ResultId: id,
		},
	}, nil
}
