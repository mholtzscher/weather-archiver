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

func TestCreateTeam(t *testing.T) {
	mockDB := mocks.NewMockQuerier(t)
	service := NewFormulaDataServer(mockDB)

	t.Run("should create driver", func(t *testing.T) {
		mockDB.On("CreateTeam", mock.Anything, dal.CreateTeamParams{
			Name: "Toyota",
			Base: "Tokyo",
		}).Return(int32(1), nil).Once()

		request := &connect.Request[apiv1.CreateTeamRequest]{
			Msg: &apiv1.CreateTeamRequest{
				Name: "Toyota",
				Base: "Tokyo",
			},
		}

		result, err := service.CreateTeam(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, int32(1), result.Msg.TeamId)
	})

	t.Run("should return error when create driver fails", func(t *testing.T) {
		mockDB.On("CreateTeam", mock.Anything, mock.Anything).Return(int32(0), assert.AnError).Once()

		request := &connect.Request[apiv1.CreateTeamRequest]{
			Msg: &apiv1.CreateTeamRequest{
				Name: "Toyota",
				Base: "Tokyo",
			},
		}
		result, err := service.CreateTeam(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
