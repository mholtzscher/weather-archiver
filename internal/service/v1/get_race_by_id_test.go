package v1

import (
	"context"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/jackc/pgx/v5/pgtype"
	apiv1 "github.com/mholtzscher/weather-archiver/gen/api/v1"
	"github.com/mholtzscher/weather-archiver/internal/dal"
	"github.com/mholtzscher/weather-archiver/internal/testing/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetRaceById(t *testing.T) {
	mockDB := mocks.NewMockQuerier(t)
	service := NewFormulaDataServer(mockDB)

	t.Run("should get race", func(t *testing.T) {
		race := dal.Race{
			ID:       42,
			SeasonID: gofakeit.Int32(),
			Name:     gofakeit.Name(),
			Location: gofakeit.City(),
			Date: pgtype.Date{
				Time: time.Now(),
			},
		}
		mockDB.On("GetRaceById", mock.Anything, int32(42)).Return(race, nil).Once()
		request := &connect.Request[apiv1.GetRaceByIdRequest]{
			Msg: &apiv1.GetRaceByIdRequest{
				RaceId: 42,
			},
		}
		result, err := service.GetRaceById(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, race.SeasonID, result.Msg.Race.SeasonId)
		assert.Equal(t, race.Name, result.Msg.Race.Name)
		assert.Equal(t, race.Location, result.Msg.Race.Location)
		assert.EqualValues(t, race.Date.Time.Year(), result.Msg.Race.Date.Year)
		assert.EqualValues(t, race.Date.Time.Month(), result.Msg.Race.Date.Month)
		assert.EqualValues(t, race.Date.Time.Day(), result.Msg.Race.Date.Day)
	})

	t.Run("should return error when get race returns an error", func(t *testing.T) {
		mockDB.On("GetRaceById", mock.Anything, int32(42)).Return(dal.Race{}, assert.AnError).Once()
		request := &connect.Request[apiv1.GetRaceByIdRequest]{
			Msg: &apiv1.GetRaceByIdRequest{
				RaceId: 42,
			},
		}
		result, err := service.GetRaceById(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
