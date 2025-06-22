package integration

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	apiv1 "github.com/mholtzscher/weather-archiver/gen/api/v1"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/assert"
)

func TestCreateSeason(t *testing.T) {
	helper := CreateIntegrationTestHelper(t)
	client := helper.Client

	t.Run("create season should return season id", func(t *testing.T) {
		result, err := client.CreateSeason(context.Background(), connect.NewRequest(&apiv1.CreateSeasonRequest{
			Year:   int32(gofakeit.IntRange(1900, 2100)),
			Series: gofakeit.BookAuthor(),
		}))
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.SeasonId)
	})

	t.Run("season should require year", func(t *testing.T) {
		result, err := client.CreateSeason(context.Background(), connect.NewRequest(&apiv1.CreateSeasonRequest{
			Series: gofakeit.FarmAnimal(),
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("season should be after 1900", func(t *testing.T) {
		result, err := client.CreateSeason(context.Background(), connect.NewRequest(&apiv1.CreateSeasonRequest{
			Year:   1900,
			Series: gofakeit.FarmAnimal(),
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("season should be before 2100", func(t *testing.T) {
		result, err := client.CreateSeason(context.Background(), connect.NewRequest(&apiv1.CreateSeasonRequest{
			Year:   2101,
			Series: gofakeit.Adjective(),
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("season should require series name", func(t *testing.T) {
		result, err := client.CreateSeason(context.Background(), connect.NewRequest(&apiv1.CreateSeasonRequest{
			Year: int32(gofakeit.IntRange(1900, 2100)),
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("should not allow duplicate season", func(t *testing.T) {
		request := connect.NewRequest(&apiv1.CreateSeasonRequest{
			Series: gofakeit.Sentence(3),
			Year:   int32(gofakeit.IntRange(1900, 2100)),
		})

		result, err := client.CreateSeason(context.Background(), request)
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.SeasonId)

		result, err = client.CreateSeason(context.Background(), request)
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeAlreadyExists, connect.CodeOf(err))
		assert.Nil(t, result)
	})
}

func TestGetSeasonById(t *testing.T) {
	helper := CreateIntegrationTestHelper(t)
	client := helper.Client

	t.Run("should return season when querying by id", func(t *testing.T) {
		year := int32(gofakeit.IntRange(1900, 2100))
		series := gofakeit.Sentence(3)

		result, err := client.CreateSeason(context.Background(), connect.NewRequest(&apiv1.CreateSeasonRequest{
			Year:   year,
			Series: series,
		}))
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.SeasonId)

		actual, err := client.GetSeasonById(context.Background(), connect.NewRequest(&apiv1.GetSeasonByIdRequest{
			SeasonId: result.Msg.SeasonId,
		}))
		assert.Nil(t, err)
		assert.Equal(t, year, actual.Msg.Season.Year)
		assert.Equal(t, series, actual.Msg.Season.Series)
	})

	t.Run("should return not found when season id does not exist", func(t *testing.T) {
		_, err := client.GetSeasonById(context.Background(), connect.NewRequest(&apiv1.GetSeasonByIdRequest{
			SeasonId: gofakeit.Int32(),
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeNotFound, connect.CodeOf(err))
	})

	t.Run("season id should be greater than 0", func(t *testing.T) {
		_, err := client.GetSeasonById(context.Background(), connect.NewRequest(&apiv1.GetSeasonByIdRequest{
			SeasonId: -1,
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
	})

	t.Run("should return validation error when id is not in request", func(t *testing.T) {
		_, err := client.GetSeasonById(context.Background(), connect.NewRequest(&apiv1.GetSeasonByIdRequest{}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
	})
}

func TestGetAllSeasons(t *testing.T) {
	helper := CreateIntegrationTestHelper(t)
	client := helper.Client

	t.Run("should return all seasons", func(t *testing.T) {
		result, err := client.CreateSeason(context.Background(), connect.NewRequest(&apiv1.CreateSeasonRequest{
			Year:   int32(gofakeit.IntRange(1900, 2100)),
			Series: gofakeit.Sentence(3),
		}))
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.SeasonId)

		result, err = client.CreateSeason(context.Background(), connect.NewRequest(&apiv1.CreateSeasonRequest{
			Year:   int32(gofakeit.IntRange(1900, 2100)),
			Series: gofakeit.Sentence(3),
		}))
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.SeasonId)

		actual, err := client.GetAllSeasons(context.Background(), connect.NewRequest(&apiv1.GetAllSeasonsRequest{}))
		assert.Nil(t, err)
		assert.Len(t, actual.Msg.Seasons, 2)
	})
}
