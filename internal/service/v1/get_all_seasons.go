package v1

import (
	"context"

	"connectrpc.com/connect"
	apiv1 "github.com/mholtzscher/weather-archiver/gen/api/v1"
)

func (s *FormulaDataServer) GetAllSeasons(
	ctx context.Context,
	request *connect.Request[apiv1.GetAllSeasonsRequest],
) (*connect.Response[apiv1.GetAllSeasonsResponse], error) {
	seasons, err := s.DB.GetAllSeasons(ctx)
	if err != nil {
		return nil, err
	}

	seasonsMapped := make([]*apiv1.Season, len(seasons))
	for i, season := range seasons {
		seasonsMapped[i] = &apiv1.Season{
			SeasonId: season.ID,
			Year:     season.SeasonYear,
			Series:   season.Series,
		}
	}

	return &connect.Response[apiv1.GetAllSeasonsResponse]{
		Msg: &apiv1.GetAllSeasonsResponse{
			Seasons: seasonsMapped,
		},
	}, nil
}
