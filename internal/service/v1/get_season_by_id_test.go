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

func TestGetSeasonById(t *testing.T) {
	mockDB := mocks.NewMockQuerier(t)
	service := NewFormulaDataServer(mockDB)

	t.Run("should get season", func(t *testing.T) {
		season := dal.Season{
			ID:         42,
			SeasonYear: 2024,
			Series:     "Formula Test",
		}
		mockDB.On("GetSeasonById", mock.Anything, int32(42)).Return(season, nil).Once()
		request := &connect.Request[apiv1.GetSeasonByIdRequest]{
			Msg: &apiv1.GetSeasonByIdRequest{
				SeasonId: 42,
			},
		}
		result, err := service.GetSeasonById(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.Nil(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should return error when get season returns an error", func(t *testing.T) {
		mockDB.On("GetSeasonById", mock.Anything, int32(42)).Return(dal.Season{}, assert.AnError).Once()
		request := &connect.Request[apiv1.GetSeasonByIdRequest]{
			Msg: &apiv1.GetSeasonByIdRequest{
				SeasonId: 42,
			},
		}
		result, err := service.GetSeasonById(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
