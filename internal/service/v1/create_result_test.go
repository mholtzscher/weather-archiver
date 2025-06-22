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

func TestCreateResult(t *testing.T) {
	mockDB := mocks.NewMockQuerier(t)
	service := NewFormulaDataServer(mockDB)

	t.Run("should create race", func(t *testing.T) {
		mockDB.On("CreateResult", mock.Anything, dal.CreateResultParams{
			RaceID:   1,
			DriverID: 2,
			TeamID:   3,
			Position: 4,
			Points:   5,
		}).Return(int32(1), nil).Once()

		request := &connect.Request[apiv1.CreateResultRequest]{
			Msg: &apiv1.CreateResultRequest{
				RaceId:   1,
				DriverId: 2,
				TeamId:   3,
				Position: 4,
				Points:   5,
			},
		}

		result, err := service.CreateResult(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, int32(1), result.Msg.ResultId)
	})

	t.Run("should return error when create race fails", func(t *testing.T) {
		mockDB.On("CreateResult", mock.Anything, mock.Anything).Return(int32(0), assert.AnError).Once()

		request := &connect.Request[apiv1.CreateResultRequest]{
			Msg: &apiv1.CreateResultRequest{
				RaceId:   1,
				DriverId: 2,
				TeamId:   3,
				Position: 4,
				Points:   5,
			},
		}
		result, err := service.CreateResult(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
