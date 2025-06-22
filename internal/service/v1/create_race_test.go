package v1

import (
	"context"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5/pgtype"
	apiv1 "github.com/mholtzscher/weather-archiver/gen/api/v1"
	"github.com/mholtzscher/weather-archiver/internal/dal"
	"github.com/mholtzscher/weather-archiver/internal/testing/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/genproto/googleapis/type/date"
)

func TestCreateRace(t *testing.T) {
	mockDB := mocks.NewMockQuerier(t)
	service := NewFormulaDataServer(mockDB)

	t.Run("should create race", func(t *testing.T) {
		mockDB.On("CreateRace", mock.Anything, dal.CreateRaceParams{
			SeasonID: 1,
			Name:     "Belgian Grand Prix",
			Location: "Spa-Francorchamps",
			Date: pgtype.Date{
				Time:  time.Date(1997, 9, 30, 0, 0, 0, 0, time.UTC),
				Valid: true,
			},
		}).Return(int32(1), nil).Once()

		request := &connect.Request[apiv1.CreateRaceRequest]{
			Msg: &apiv1.CreateRaceRequest{
				SeasonId: 1,
				Name:     "Belgian Grand Prix",
				Location: "Spa-Francorchamps",
				Date: &date.Date{
					Year:  1997,
					Month: 9,
					Day:   30,
				},
			},
		}

		result, err := service.CreateRace(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, int32(1), result.Msg.RaceId)
	})

	t.Run("should return error when create race fails", func(t *testing.T) {
		mockDB.On("CreateRace", mock.Anything, mock.Anything).Return(int32(0), assert.AnError).Once()

		request := &connect.Request[apiv1.CreateRaceRequest]{
			Msg: &apiv1.CreateRaceRequest{
				SeasonId: 1,
				Name:     "Belgian Grand Prix",
				Location: "Spa-Francorchamps",
				Date: &date.Date{
					Year:  1997,
					Month: 9,
					Day:   30,
				},
			},
		}
		result, err := service.CreateRace(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
