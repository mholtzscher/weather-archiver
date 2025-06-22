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

func TestGetResultsByRace(t *testing.T) {
	mockDB := mocks.NewMockQuerier(t)
	service := NewFormulaDataServer(mockDB)

	t.Run("should get single result", func(t *testing.T) {
		raceResult := dal.Result{
			ID:       42,
			RaceID:   gofakeit.Int32(),
			DriverID: gofakeit.Int32(),
			TeamID:   gofakeit.Int32(),
			Position: gofakeit.Int32(),
			Points:   gofakeit.Float64(),
		}
		mockDB.On("GetResultsByRaceId", mock.Anything, int32(42)).Return([]dal.Result{raceResult}, nil).Once()
		request := &connect.Request[apiv1.GetResultsByRaceRequest]{
			Msg: &apiv1.GetResultsByRaceRequest{
				RaceId: 42,
			},
		}
		result, err := service.GetResultsByRace(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Msg.Results, 1)
		assert.Equal(t, raceResult.RaceID, result.Msg.Results[0].RaceId)
		assert.Equal(t, raceResult.DriverID, result.Msg.Results[0].DriverId)
		assert.Equal(t, raceResult.TeamID, result.Msg.Results[0].TeamId)
		assert.Equal(t, raceResult.Position, result.Msg.Results[0].Position)
		assert.Equal(t, raceResult.Points, result.Msg.Results[0].Points)
	})

	t.Run("should get multiple results", func(t *testing.T) {
		raceResult := dal.Result{
			ID:       42,
			RaceID:   gofakeit.Int32(),
			DriverID: gofakeit.Int32(),
			TeamID:   gofakeit.Int32(),
			Position: gofakeit.Int32(),
			Points:   gofakeit.Float64(),
		}
		raceResult2 := dal.Result{
			ID:       43,
			RaceID:   gofakeit.Int32(),
			DriverID: gofakeit.Int32(),
			TeamID:   gofakeit.Int32(),
			Position: gofakeit.Int32(),
			Points:   gofakeit.Float64(),
		}

		mockDB.On("GetResultsByRaceId", mock.Anything, int32(42)).Return([]dal.Result{raceResult, raceResult2}, nil).Once()
		request := &connect.Request[apiv1.GetResultsByRaceRequest]{
			Msg: &apiv1.GetResultsByRaceRequest{
				RaceId: 42,
			},
		}
		result, err := service.GetResultsByRace(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Msg.Results, 2)
	})

	t.Run("should return error when get result returns an error", func(t *testing.T) {
		mockDB.On("GetResultsByRaceId", mock.Anything, int32(42)).Return([]dal.Result{}, assert.AnError).Once()
		request := &connect.Request[apiv1.GetResultsByRaceRequest]{
			Msg: &apiv1.GetResultsByRaceRequest{
				RaceId: 42,
			},
		}
		result, err := service.GetResultsByRace(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
