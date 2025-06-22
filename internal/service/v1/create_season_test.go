package v1

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	apiv1 "github.com/mholtzscher/weather-archiver/gen/api/v1"
	"github.com/mholtzscher/weather-archiver/internal/dal"
	"github.com/mholtzscher/weather-archiver/internal/testing/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateSeason(t *testing.T) {
	mockDB := mocks.NewMockQuerier(t)
	service := NewFormulaDataServer(mockDB)

	t.Run("should create season", func(t *testing.T) {
		mockDB.On("CreateSeason", mock.Anything, dal.CreateSeasonParams{SeasonYear: 2024, Series: "mock-series"}).Return(int32(1), nil).Once()
		request := &connect.Request[apiv1.CreateSeasonRequest]{
			Msg: &apiv1.CreateSeasonRequest{
				Year:   2024,
				Series: "mock-series",
			},
		}
		result, err := service.CreateSeason(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, int32(1), result.Msg.SeasonId)
	})

	t.Run("should return error when create season fails", func(t *testing.T) {
		mockDB.On("CreateSeason", mock.Anything, mock.Anything).Return(int32(0), assert.AnError).Once()

		request := &connect.Request[apiv1.CreateSeasonRequest]{
			Msg: &apiv1.CreateSeasonRequest{
				Year:   2024,
				Series: "mock-series",
			},
		}
		result, err := service.CreateSeason(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
