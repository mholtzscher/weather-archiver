package v1

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5/pgtype"
	apiv1 "github.com/mholtzscher/weather-archiver/gen/api/v1"
	"github.com/mholtzscher/weather-archiver/internal/dal"
)

func (s *FormulaDataServer) CreateDriver(
	ctx context.Context,
	request *connect.Request[apiv1.CreateDriverRequest],
) (*connect.Response[apiv1.CreateDriverResponse], error) {
	dobRaw := request.Msg.DateOfBirth
	dob := pgtype.Date{
		Time:             time.Date(int(dobRaw.GetYear()), time.Month(dobRaw.GetMonth()), int(dobRaw.GetDay()), 0, 0, 0, 0, time.UTC),
		InfinityModifier: 0,
		Valid:            true,
	}

	id, err := s.DB.CreateDriver(ctx, dal.CreateDriverParams{
		FirstName:    request.Msg.FirstName,
		LastName:     request.Msg.LastName,
		PlaceOfBirth: request.Msg.PlaceOfBirth,
		DateOfBirth:  dob,
	})
	if err != nil {
		return nil, mapPgErrorsToReturnCodes(err)
	}

	return &connect.Response[apiv1.CreateDriverResponse]{
		Msg: &apiv1.CreateDriverResponse{
			DriverId: id,
		},
	}, nil
}
