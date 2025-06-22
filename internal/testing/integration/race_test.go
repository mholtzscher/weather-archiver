package integration

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	apiv1 "github.com/mholtzscher/weather-archiver/gen/api/v1"
	"google.golang.org/genproto/googleapis/type/date"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/assert"
)

func TestCreateRace(t *testing.T) {
	helper := CreateIntegrationTestHelper(t)
	client := helper.Client

	season, _ := client.CreateSeason(context.Background(), connect.NewRequest(&apiv1.CreateSeasonRequest{
		Year:   int32(gofakeit.IntRange(1900, 2100)),
		Series: gofakeit.BookAuthor(),
	}))

	t.Run("create race should return race id", func(t *testing.T) {
		d := gofakeit.Date()
		result, err := client.CreateRace(context.Background(), connect.NewRequest(&apiv1.CreateRaceRequest{
			SeasonId: season.Msg.SeasonId,
			Name:     gofakeit.FarmAnimal(),
			Location: gofakeit.City(),
			Date: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		}))
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.RaceId)
	})

	t.Run("race should require season id", func(t *testing.T) {
		d := gofakeit.Date()
		result, err := client.CreateRace(context.Background(), connect.NewRequest(&apiv1.CreateRaceRequest{
			Name:     gofakeit.FarmAnimal(),
			Location: gofakeit.City(),
			Date: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("race should require name", func(t *testing.T) {
		d := gofakeit.Date()
		result, err := client.CreateRace(context.Background(), connect.NewRequest(&apiv1.CreateRaceRequest{
			SeasonId: season.Msg.SeasonId,
			Location: gofakeit.City(),
			Date: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("race should require location", func(t *testing.T) {
		d := gofakeit.Date()
		result, err := client.CreateRace(context.Background(), connect.NewRequest(&apiv1.CreateRaceRequest{
			SeasonId: season.Msg.SeasonId,
			Name:     gofakeit.FarmAnimal(),
			Date: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("race should require date", func(t *testing.T) {
		result, err := client.CreateRace(context.Background(), connect.NewRequest(&apiv1.CreateRaceRequest{
			SeasonId: season.Msg.SeasonId,
			Name:     gofakeit.FarmAnimal(),
			Location: gofakeit.City(),
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("race should require date year", func(t *testing.T) {
		d := gofakeit.Date()
		result, err := client.CreateRace(context.Background(), connect.NewRequest(&apiv1.CreateRaceRequest{
			SeasonId: season.Msg.SeasonId,
			Name:     gofakeit.FarmAnimal(),
			Location: gofakeit.City(),
			Date: &date.Date{
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("race should require date month", func(t *testing.T) {
		d := gofakeit.Date()
		result, err := client.CreateRace(context.Background(), connect.NewRequest(&apiv1.CreateRaceRequest{
			SeasonId: season.Msg.SeasonId,
			Name:     gofakeit.FarmAnimal(),
			Location: gofakeit.City(),
			Date: &date.Date{
				Year: int32(d.Year()),
				Day:  int32(d.Day()),
			},
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("race should require date day", func(t *testing.T) {
		d := gofakeit.Date()
		result, err := client.CreateRace(context.Background(), connect.NewRequest(&apiv1.CreateRaceRequest{
			SeasonId: season.Msg.SeasonId,
			Name:     gofakeit.FarmAnimal(),
			Location: gofakeit.City(),
			Date: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
			},
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("should not allow duplicate race", func(t *testing.T) {
		d := gofakeit.Date()
		request := connect.NewRequest(&apiv1.CreateRaceRequest{
			SeasonId: season.Msg.SeasonId,
			Name:     gofakeit.FarmAnimal(),
			Location: gofakeit.City(),
			Date: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		})

		result, err := client.CreateRace(context.Background(), request)
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.RaceId)

		result, err = client.CreateRace(context.Background(), request)
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeAlreadyExists, connect.CodeOf(err))
		assert.Nil(t, result)
	})
}

func TestGetRaceById(t *testing.T) {
	helper := CreateIntegrationTestHelper(t)
	client := helper.Client

	season, _ := client.CreateSeason(context.Background(), connect.NewRequest(&apiv1.CreateSeasonRequest{
		Year:   int32(gofakeit.IntRange(1900, 2100)),
		Series: gofakeit.BookAuthor(),
	}))

	t.Run("should return race when querying by id", func(t *testing.T) {
		d := gofakeit.Date()
		request := &apiv1.CreateRaceRequest{
			SeasonId: season.Msg.SeasonId,
			Name:     gofakeit.Adjective(),
			Location: gofakeit.City(),
			Date: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		}

		result, err := client.CreateRace(context.Background(), connect.NewRequest(request))
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.RaceId)

		actual, err := client.GetRaceById(context.Background(), connect.NewRequest(&apiv1.GetRaceByIdRequest{
			RaceId: result.Msg.RaceId,
		}))
		assert.Nil(t, err)
		assert.Equal(t, request.SeasonId, actual.Msg.Race.SeasonId)
		assert.Equal(t, request.Name, actual.Msg.Race.Name)
		assert.Equal(t, request.Location, actual.Msg.Race.Location)
		assert.EqualValues(t, request.Date.Year, actual.Msg.Race.Date.Year)
		assert.EqualValues(t, request.Date.Month, actual.Msg.Race.Date.Month)
		assert.EqualValues(t, request.Date.Day, actual.Msg.Race.Date.Day)
	})

	t.Run("should return not found when race id does not exist", func(t *testing.T) {
		_, err := client.GetRaceById(context.Background(), connect.NewRequest(&apiv1.GetRaceByIdRequest{
			RaceId: gofakeit.Int32(),
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeNotFound, connect.CodeOf(err))
	})

	t.Run("race id should be greater than 0", func(t *testing.T) {
		_, err := client.GetRaceById(context.Background(), connect.NewRequest(&apiv1.GetRaceByIdRequest{
			RaceId: -1,
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
	})

	t.Run("should return validation error when id is not in request", func(t *testing.T) {
		_, err := client.GetRaceById(context.Background(), connect.NewRequest(&apiv1.GetRaceByIdRequest{}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
	})
}
