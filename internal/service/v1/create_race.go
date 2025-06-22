package v1

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5/pgtype"
	apiv1 "github.com/mholtzscher/weather-archiver/gen/api/v1"
	"github.com/mholtzscher/weather-archiver/internal/dal"
)

func (s *FormulaDataServer) CreateRace(
	ctx context.Context,
	request *connect.Request[apiv1.CreateRaceRequest],
) (*connect.Response[apiv1.CreateRaceResponse], error) {
	dobRaw := request.Msg.Date
	raceDate := pgtype.Date{
		Time:             time.Date(int(dobRaw.GetYear()), time.Month(dobRaw.GetMonth()), int(dobRaw.GetDay()), 0, 0, 0, 0, time.UTC),
		InfinityModifier: 0,
		Valid:            true,
	}

	id, err := s.DB.CreateRace(ctx, dal.CreateRaceParams{
		SeasonID: request.Msg.SeasonId,
		Name:     request.Msg.Name,
		Location: request.Msg.Location,
		Date:     raceDate,
	})
	if err != nil {
		return nil, mapPgErrorsToReturnCodes(err)
	}

	return &connect.Response[apiv1.CreateRaceResponse]{
		Msg: &apiv1.CreateRaceResponse{
			RaceId: id,
		},
	}, nil
}
