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

func TestGetDriverById(t *testing.T) {
	mockDB := mocks.NewMockQuerier(t)
	service := NewFormulaDataServer(mockDB)

	t.Run("should get driver", func(t *testing.T) {
		driver := dal.Driver{
			ID:           42,
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			PlaceOfBirth: gofakeit.City(),
			DateOfBirth: pgtype.Date{
				Time: time.Now(),
			},
		}
		mockDB.On("GetDriverById", mock.Anything, int32(42)).Return(driver, nil).Once()
		request := &connect.Request[apiv1.GetDriverByIdRequest]{
			Msg: &apiv1.GetDriverByIdRequest{
				DriverId: 42,
			},
		}
		result, err := service.GetDriverById(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, driver.FirstName, result.Msg.Driver.FirstName)
		assert.Equal(t, driver.LastName, result.Msg.Driver.LastName)
		assert.Equal(t, driver.PlaceOfBirth, result.Msg.Driver.PlaceOfBirth)
		assert.EqualValues(t, driver.DateOfBirth.Time.Year(), result.Msg.Driver.DateOfBirth.Year)
		assert.EqualValues(t, driver.DateOfBirth.Time.Month(), result.Msg.Driver.DateOfBirth.Month)
		assert.EqualValues(t, driver.DateOfBirth.Time.Day(), result.Msg.Driver.DateOfBirth.Day)
	})

	t.Run("should return error when get driver returns an error", func(t *testing.T) {
		mockDB.On("GetDriverById", mock.Anything, int32(42)).Return(dal.Driver{}, assert.AnError).Once()
		request := &connect.Request[apiv1.GetDriverByIdRequest]{
			Msg: &apiv1.GetDriverByIdRequest{
				DriverId: 42,
			},
		}
		result, err := service.GetDriverById(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
