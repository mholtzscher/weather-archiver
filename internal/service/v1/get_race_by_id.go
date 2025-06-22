package v1

import (
	"context"

	"connectrpc.com/connect"
	apiv1 "github.com/mholtzscher/weather-archiver/gen/api/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/type/date"
)

func (s *FormulaDataServer) GetRaceById(
	ctx context.Context,
	request *connect.Request[apiv1.GetRaceByIdRequest],
) (*connect.Response[apiv1.GetRaceByIdResponse], error) {
	race, err := s.DB.GetRaceById(ctx, request.Msg.RaceId)
	if err != nil {
		log.Error().Err(err).Msg("failed to get race by id")
		return nil, mapPgErrorsToReturnCodes(err)
	}
	return &connect.Response[apiv1.GetRaceByIdResponse]{
		Msg: &apiv1.GetRaceByIdResponse{
			Race: &apiv1.Race{
				RaceId:   race.ID,
				SeasonId: race.SeasonID,
				Name:     race.Name,
				Location: race.Location,
				Date: &date.Date{
					Year:  int32(race.Date.Time.Year()),
					Month: int32(race.Date.Time.Month()),
					Day:   int32(race.Date.Time.Day()),
				},
			},
		},
	}, nil
}
