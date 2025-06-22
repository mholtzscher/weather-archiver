package v1

import (
	"context"

	"connectrpc.com/connect"
	apiv1 "github.com/mholtzscher/weather-archiver/gen/api/v1"
	"github.com/mholtzscher/weather-archiver/internal/dal"
)

func (s *FormulaDataServer) CreateTeam(
	ctx context.Context,
	request *connect.Request[apiv1.CreateTeamRequest],
) (*connect.Response[apiv1.CreateTeamResponse], error) {
	id, err := s.DB.CreateTeam(ctx, dal.CreateTeamParams{
		Name: request.Msg.Name,
		Base: request.Msg.Base,
	})
	if err != nil {
		return nil, mapPgErrorsToReturnCodes(err)
	}

	return &connect.Response[apiv1.CreateTeamResponse]{
		Msg: &apiv1.CreateTeamResponse{
			TeamId: id,
		},
	}, nil
}
