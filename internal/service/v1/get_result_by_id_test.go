package v1

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/brianvoe/gofakeit/v7"
	apiv1 "github.com/mholtzscher/weather-archiver/gen/api/v1"
	"github.com/mholtzscher/weather-archiver/internal/dal"
	"github.com/mholtzscher/weather-archiver/internal/testing/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetResultById(t *testing.T) {
	mockDB := mocks.NewMockQuerier(t)
	service := NewFormulaDataServer(mockDB)

	t.Run("should get result", func(t *testing.T) {
		raceResult := dal.Result{
			ID:       42,
			RaceID:   gofakeit.Int32(),
			DriverID: gofakeit.Int32(),
			TeamID:   gofakeit.Int32(),
			Position: gofakeit.Int32(),
			Points:   gofakeit.Float64(),
		}
		mockDB.On("GetResultById", mock.Anything, int32(42)).Return(raceResult, nil).Once()
		request := &connect.Request[apiv1.GetResultByIdRequest]{
			Msg: &apiv1.GetResultByIdRequest{
				ResultId: 42,
			},
		}
		result, err := service.GetResultById(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, raceResult.RaceID, result.Msg.Result.RaceId)
		assert.Equal(t, raceResult.DriverID, result.Msg.Result.DriverId)
		assert.Equal(t, raceResult.TeamID, result.Msg.Result.TeamId)
		assert.Equal(t, raceResult.Position, result.Msg.Result.Position)
		assert.Equal(t, raceResult.Points, result.Msg.Result.Points)
	})

	t.Run("should return error when get result returns an error", func(t *testing.T) {
		mockDB.On("GetResultById", mock.Anything, int32(42)).Return(dal.Result{}, assert.AnError).Once()
		request := &connect.Request[apiv1.GetResultByIdRequest]{
			Msg: &apiv1.GetResultByIdRequest{
				ResultId: 42,
			},
		}
		result, err := service.GetResultById(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
