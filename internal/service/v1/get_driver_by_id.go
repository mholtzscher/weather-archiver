package v1

import (
	"context"

	"connectrpc.com/connect"
	apiv1 "github.com/mholtzscher/weather-archiver/gen/api/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/type/date"
)

func (s *FormulaDataServer) GetDriverById(
	ctx context.Context,
	request *connect.Request[apiv1.GetDriverByIdRequest],
) (*connect.Response[apiv1.GetDriverByIdResponse], error) {
	driver, err := s.DB.GetDriverById(ctx, request.Msg.DriverId)
	if err != nil {
		log.Error().Err(err).Msg("failed to get driver by id")
		return nil, mapPgErrorsToReturnCodes(err)
	}
	return &connect.Response[apiv1.GetDriverByIdResponse]{
		Msg: &apiv1.GetDriverByIdResponse{
			Driver: &apiv1.Driver{
				DriverId:     driver.ID,
				FirstName:    driver.FirstName,
				LastName:     driver.LastName,
				PlaceOfBirth: driver.PlaceOfBirth,
				DateOfBirth: &date.Date{
					Year:  int32(driver.DateOfBirth.Time.Year()),
					Month: int32(driver.DateOfBirth.Time.Month()),
					Day:   int32(driver.DateOfBirth.Time.Day()),
				},
			},
		},
	}, nil
}
