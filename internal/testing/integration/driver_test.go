package integration

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/brianvoe/gofakeit/v7"
	apiv1 "github.com/mholtzscher/weather-archiver/gen/api/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/type/date"
)

func TestCreateDriver(t *testing.T) {
	helper := CreateIntegrationTestHelper(t)
	client := helper.Client

	t.Run("create driver should return driver id", func(t *testing.T) {
		d := gofakeit.Date()
		result, err := client.CreateDriver(context.Background(), connect.NewRequest(&apiv1.CreateDriverRequest{
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			PlaceOfBirth: gofakeit.City(),
			DateOfBirth: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		}))
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.DriverId)
	})

	t.Run("driver should require first name", func(t *testing.T) {
		d := gofakeit.Date()
		result, err := client.CreateDriver(context.Background(), connect.NewRequest(&apiv1.CreateDriverRequest{
			LastName:     gofakeit.LastName(),
			PlaceOfBirth: gofakeit.City(),
			DateOfBirth: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("driver should require last name", func(t *testing.T) {
		d := gofakeit.Date()
		result, err := client.CreateDriver(context.Background(), connect.NewRequest(&apiv1.CreateDriverRequest{
			FirstName:    gofakeit.FirstName(),
			PlaceOfBirth: gofakeit.City(),
			DateOfBirth: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("driver should require place of birth", func(t *testing.T) {
		d := gofakeit.Date()
		result, err := client.CreateDriver(context.Background(), connect.NewRequest(&apiv1.CreateDriverRequest{
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			DateOfBirth: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("driver should require date of birth", func(t *testing.T) {
		result, err := client.CreateDriver(context.Background(), connect.NewRequest(&apiv1.CreateDriverRequest{
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			PlaceOfBirth: gofakeit.City(),
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("driver should require date of birth year", func(t *testing.T) {
		d := gofakeit.Date()
		result, err := client.CreateDriver(context.Background(), connect.NewRequest(&apiv1.CreateDriverRequest{
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			PlaceOfBirth: gofakeit.City(),
			DateOfBirth: &date.Date{
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("driver should require date of birth month", func(t *testing.T) {
		d := gofakeit.Date()
		result, err := client.CreateDriver(context.Background(), connect.NewRequest(&apiv1.CreateDriverRequest{
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			PlaceOfBirth: gofakeit.City(),
			DateOfBirth: &date.Date{
				Year: int32(d.Year()),
				Day:  int32(d.Day()),
			},
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("driver should require date of birth day", func(t *testing.T) {
		d := gofakeit.Date()
		result, err := client.CreateDriver(context.Background(), connect.NewRequest(&apiv1.CreateDriverRequest{
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			PlaceOfBirth: gofakeit.City(),
			DateOfBirth: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
			},
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("should not allow duplicate driver", func(t *testing.T) {
		d := gofakeit.Date()
		request := connect.NewRequest(&apiv1.CreateDriverRequest{
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			PlaceOfBirth: gofakeit.City(),
			DateOfBirth: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		})

		result, err := client.CreateDriver(context.Background(), request)
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.DriverId)

		result, err = client.CreateDriver(context.Background(), request)
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeAlreadyExists, connect.CodeOf(err))
		assert.Nil(t, result)
	})
}

func TestGetDriverById(t *testing.T) {
	helper := CreateIntegrationTestHelper(t)
	client := helper.Client

	t.Run("should return driver when querying by id", func(t *testing.T) {
		d := gofakeit.Date()
		request := &apiv1.CreateDriverRequest{
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			PlaceOfBirth: gofakeit.City(),
			DateOfBirth: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		}
		result, err := client.CreateDriver(context.Background(), connect.NewRequest(request))
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.DriverId)

		actual, err := client.GetDriverById(context.Background(), connect.NewRequest(&apiv1.GetDriverByIdRequest{
			DriverId: result.Msg.DriverId,
		}))
		assert.Nil(t, err)
		assert.Equal(t, request.FirstName, actual.Msg.Driver.FirstName)
		assert.Equal(t, request.LastName, actual.Msg.Driver.LastName)
		assert.Equal(t, request.PlaceOfBirth, actual.Msg.Driver.PlaceOfBirth)
		assert.EqualValues(t, request.DateOfBirth.Year, actual.Msg.Driver.DateOfBirth.Year)
		assert.EqualValues(t, request.DateOfBirth.Month, actual.Msg.Driver.DateOfBirth.Month)
		assert.EqualValues(t, request.DateOfBirth.Day, actual.Msg.Driver.DateOfBirth.Day)
	})

	t.Run("should return not found when driver id does not exist", func(t *testing.T) {
		_, err := client.GetDriverById(context.Background(), connect.NewRequest(&apiv1.GetDriverByIdRequest{
			DriverId: gofakeit.Int32(),
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeNotFound, connect.CodeOf(err))
	})

	t.Run("driver id should be greater than 0", func(t *testing.T) {
		_, err := client.GetDriverById(context.Background(), connect.NewRequest(&apiv1.GetDriverByIdRequest{
			DriverId: -1,
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
	})

	t.Run("should return validation error when id is not in request", func(t *testing.T) {
		_, err := client.GetDriverById(context.Background(), connect.NewRequest(&apiv1.GetDriverByIdRequest{}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
	})
}
