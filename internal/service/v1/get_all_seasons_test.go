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

func TestGetAllSeasons(t *testing.T) {
	mockDB := mocks.NewMockQuerier(t)
	service := NewFormulaDataServer(mockDB)

	t.Run("should get all seasons", func(t *testing.T) {
		season := dal.Season{
			ID:         42,
			SeasonYear: 2024,
			Series:     "Formula Test",
		}
		season2 := dal.Season{
			ID:         41,
			SeasonYear: 2023,
			Series:     "Formula Test",
		}
		mockDB.On("GetAllSeasons", mock.Anything).Return([]dal.Season{season, season2}, nil).Once()
		request := &connect.Request[apiv1.GetAllSeasonsRequest]{
			Msg: &apiv1.GetAllSeasonsRequest{},
		}
		result, err := service.GetAllSeasons(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result.Msg.Seasons, 2)
	})

	t.Run("should return error when get all seasons returns an error", func(t *testing.T) {
		mockDB.On("GetAllSeasons", mock.Anything).Return([]dal.Season{}, assert.AnError).Once()
		request := &connect.Request[apiv1.GetAllSeasonsRequest]{
			Msg: &apiv1.GetAllSeasonsRequest{},
		}
		result, err := service.GetAllSeasons(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
