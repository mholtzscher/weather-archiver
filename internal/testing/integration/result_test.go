package integration

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	apiv1 "github.com/mholtzscher/weather-archiver/gen/api/v1"
	"google.golang.org/genproto/googleapis/type/date"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/assert"
)

func TestCreateResult(t *testing.T) {
	helper := CreateIntegrationTestHelper(t)
	client := helper.Client

	season, _ := client.CreateSeason(context.Background(), connect.NewRequest(&apiv1.CreateSeasonRequest{
		Year:   int32(gofakeit.IntRange(1900, 2100)),
		Series: gofakeit.BookAuthor(),
	}))

	d := gofakeit.Date()
	race, _ := client.CreateRace(context.Background(), connect.NewRequest(&apiv1.CreateRaceRequest{
		SeasonId: season.Msg.SeasonId,
		Name:     gofakeit.FarmAnimal(),
		Location: gofakeit.City(),
		Date:     &date.Date{Year: int32(d.Year()), Month: int32(d.Month()), Day: int32(d.Day())},
	}))

	driver, _ := client.CreateDriver(context.Background(), connect.NewRequest(&apiv1.CreateDriverRequest{
		FirstName:    gofakeit.FirstName(),
		LastName:     gofakeit.LastName(),
		PlaceOfBirth: gofakeit.City(),
		DateOfBirth: &date.Date{
			Year:  int32(d.Year()),
			Month: int32(d.Month()),
			Day:   int32(d.Day()),
		},
	}))

	team, _ := client.CreateTeam(context.Background(), connect.NewRequest(&apiv1.CreateTeamRequest{
		Name: gofakeit.Company(),
		Base: gofakeit.Country(),
	}))

	t.Run("create race should return race id", func(t *testing.T) {
		result, err := client.CreateResult(context.Background(), connect.NewRequest(&apiv1.CreateResultRequest{
			RaceId:   race.Msg.RaceId,
			DriverId: driver.Msg.DriverId,
			TeamId:   team.Msg.TeamId,
			Position: 1,
			Points:   25,
		}))
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.ResultId)
	})

	t.Run("race should require race id", func(t *testing.T) {
		result, err := client.CreateResult(context.Background(), connect.NewRequest(&apiv1.CreateResultRequest{
			DriverId: driver.Msg.DriverId,
			TeamId:   team.Msg.TeamId,
			Position: 1,
			Points:   25,
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("race should require driver id", func(t *testing.T) {
		result, err := client.CreateResult(context.Background(), connect.NewRequest(&apiv1.CreateResultRequest{
			RaceId:   race.Msg.RaceId,
			TeamId:   team.Msg.TeamId,
			Position: 1,
			Points:   25,
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("race should require team id", func(t *testing.T) {
		result, err := client.CreateResult(context.Background(), connect.NewRequest(&apiv1.CreateResultRequest{
			RaceId:   race.Msg.RaceId,
			DriverId: driver.Msg.DriverId,
			Position: 1,
			Points:   25,
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("race should require position greater than 0", func(t *testing.T) {
		result, err := client.CreateResult(context.Background(), connect.NewRequest(&apiv1.CreateResultRequest{
			RaceId:   race.Msg.RaceId,
			DriverId: driver.Msg.DriverId,
			TeamId:   team.Msg.TeamId,
			Points:   25,
			Position: 0,
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("race should require position less than 20", func(t *testing.T) {
		result, err := client.CreateResult(context.Background(), connect.NewRequest(&apiv1.CreateResultRequest{
			RaceId:   race.Msg.RaceId,
			DriverId: driver.Msg.DriverId,
			TeamId:   team.Msg.TeamId,
			Points:   25,
			Position: 21,
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("race should require points non negative", func(t *testing.T) {
		result, err := client.CreateResult(context.Background(), connect.NewRequest(&apiv1.CreateResultRequest{
			RaceId:   race.Msg.RaceId,
			DriverId: driver.Msg.DriverId,
			TeamId:   team.Msg.TeamId,
			Position: 1,
			Points:   -1,
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("race should require points be less than 26", func(t *testing.T) {
		result, err := client.CreateResult(context.Background(), connect.NewRequest(&apiv1.CreateResultRequest{
			RaceId:   race.Msg.RaceId,
			DriverId: driver.Msg.DriverId,
			TeamId:   team.Msg.TeamId,
			Position: 1,
			Points:   27,
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("should not allow duplicate result - race and position", func(t *testing.T) {
		driver2, _ := client.CreateDriver(context.Background(), connect.NewRequest(&apiv1.CreateDriverRequest{
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			PlaceOfBirth: gofakeit.City(),
			DateOfBirth: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		}))

		request := connect.NewRequest(&apiv1.CreateResultRequest{
			RaceId:   race.Msg.RaceId,
			DriverId: driver2.Msg.DriverId,
			TeamId:   team.Msg.TeamId,
			Position: 2,
			Points:   25,
		})

		result, err := client.CreateResult(context.Background(), request)
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.ResultId)

		result, err = client.CreateResult(context.Background(), request)
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeAlreadyExists, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	// TODO: verify this test after implementing table truncate
	t.Run("should not allow duplicate result - race and driver", func(t *testing.T) {
		driver2, _ := client.CreateDriver(context.Background(), connect.NewRequest(&apiv1.CreateDriverRequest{
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			PlaceOfBirth: gofakeit.City(),
			DateOfBirth: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		}))

		request := connect.NewRequest(&apiv1.CreateResultRequest{
			RaceId:   race.Msg.RaceId,
			DriverId: driver2.Msg.DriverId,
			TeamId:   team.Msg.TeamId,
			Position: 3,
			Points:   25,
		})

		result, err := client.CreateResult(context.Background(), request)
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.ResultId)

		result, err = client.CreateResult(context.Background(), request)
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeAlreadyExists, connect.CodeOf(err))
		assert.Nil(t, result)
	})
}

func TestGetResultById(t *testing.T) {
	helper := CreateIntegrationTestHelper(t)
	client := helper.Client

	season, _ := client.CreateSeason(context.Background(), connect.NewRequest(&apiv1.CreateSeasonRequest{
		Year:   int32(gofakeit.IntRange(1900, 2100)),
		Series: gofakeit.BookAuthor(),
	}))

	d := gofakeit.Date()
	race, _ := client.CreateRace(context.Background(), connect.NewRequest(&apiv1.CreateRaceRequest{
		SeasonId: season.Msg.SeasonId,
		Name:     gofakeit.FarmAnimal(),
		Location: gofakeit.City(),
		Date:     &date.Date{Year: int32(d.Year()), Month: int32(d.Month()), Day: int32(d.Day())},
	}))

	driver, _ := client.CreateDriver(context.Background(), connect.NewRequest(&apiv1.CreateDriverRequest{
		FirstName:    gofakeit.FirstName(),
		LastName:     gofakeit.LastName(),
		PlaceOfBirth: gofakeit.City(),
		DateOfBirth: &date.Date{
			Year:  int32(d.Year()),
			Month: int32(d.Month()),
			Day:   int32(d.Day()),
		},
	}))

	team, _ := client.CreateTeam(context.Background(), connect.NewRequest(&apiv1.CreateTeamRequest{
		Name: gofakeit.Company(),
		Base: gofakeit.Country(),
	}))

	t.Run("should return race result when querying by id", func(t *testing.T) {
		request := &apiv1.CreateResultRequest{
			RaceId:   race.Msg.RaceId,
			DriverId: driver.Msg.DriverId,
			TeamId:   team.Msg.TeamId,
			Position: 1,
			Points:   25,
		}
		result, err := client.CreateResult(context.Background(), connect.NewRequest(request))
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.ResultId)

		actual, err := client.GetResultById(context.Background(), connect.NewRequest(&apiv1.GetResultByIdRequest{
			ResultId: result.Msg.ResultId,
		}))
		assert.Nil(t, err)
		assert.Equal(t, request.RaceId, actual.Msg.Result.RaceId)
		assert.Equal(t, request.DriverId, actual.Msg.Result.DriverId)
		assert.Equal(t, request.TeamId, actual.Msg.Result.TeamId)
		assert.Equal(t, request.Position, actual.Msg.Result.Position)
		assert.Equal(t, request.Points, actual.Msg.Result.Points)
	})

	t.Run("should return not found when result id does not exist", func(t *testing.T) {
		_, err := client.GetResultById(context.Background(), connect.NewRequest(&apiv1.GetResultByIdRequest{
			ResultId: gofakeit.Int32(),
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeNotFound, connect.CodeOf(err))
	})

	t.Run("result id should be greater than 0", func(t *testing.T) {
		_, err := client.GetResultById(context.Background(), connect.NewRequest(&apiv1.GetResultByIdRequest{
			ResultId: -1,
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
	})

	t.Run("should return validation error when id is not in request", func(t *testing.T) {
		_, err := client.GetResultById(context.Background(), connect.NewRequest(&apiv1.GetResultByIdRequest{}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
	})
}

func TestGetResultsByRace(t *testing.T) {
	helper := CreateIntegrationTestHelper(t)
	client := helper.Client

	season, _ := client.CreateSeason(context.Background(), connect.NewRequest(&apiv1.CreateSeasonRequest{
		Year:   int32(gofakeit.IntRange(1900, 2100)),
		Series: gofakeit.BookAuthor(),
	}))

	d := gofakeit.Date()
	race, _ := client.CreateRace(context.Background(), connect.NewRequest(&apiv1.CreateRaceRequest{
		SeasonId: season.Msg.SeasonId,
		Name:     gofakeit.FarmAnimal(),
		Location: gofakeit.City(),
		Date:     &date.Date{Year: int32(d.Year()), Month: int32(d.Month()), Day: int32(d.Day())},
	}))

	driver, _ := client.CreateDriver(context.Background(), connect.NewRequest(&apiv1.CreateDriverRequest{
		FirstName:    gofakeit.FirstName(),
		LastName:     gofakeit.LastName(),
		PlaceOfBirth: gofakeit.City(),
		DateOfBirth: &date.Date{
			Year:  int32(d.Year()),
			Month: int32(d.Month()),
			Day:   int32(d.Day()),
		},
	}))

	team, _ := client.CreateTeam(context.Background(), connect.NewRequest(&apiv1.CreateTeamRequest{
		Name: gofakeit.Company(),
		Base: gofakeit.Country(),
	}))

	t.Run("should return race results when querying by race", func(t *testing.T) {
		request := &apiv1.CreateResultRequest{
			RaceId:   race.Msg.RaceId,
			DriverId: driver.Msg.DriverId,
			TeamId:   team.Msg.TeamId,
			Position: 1,
			Points:   25,
		}
		result, err := client.CreateResult(context.Background(), connect.NewRequest(request))
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.ResultId)

		actual, err := client.GetResultsByRace(context.Background(), connect.NewRequest(&apiv1.GetResultsByRaceRequest{
			RaceId: race.Msg.RaceId,
		}))
		assert.Nil(t, err)
		assert.Len(t, actual.Msg.Results, 1)
		assert.Equal(t, request.RaceId, actual.Msg.Results[0].RaceId)
		assert.Equal(t, request.DriverId, actual.Msg.Results[0].DriverId)
		assert.Equal(t, request.TeamId, actual.Msg.Results[0].TeamId)
		assert.Equal(t, request.Position, actual.Msg.Results[0].Position)
		assert.Equal(t, request.Points, actual.Msg.Results[0].Points)
	})

	t.Run("should return no results when race id does not exist", func(t *testing.T) {
		actual, err := client.GetResultsByRace(context.Background(), connect.NewRequest(&apiv1.GetResultsByRaceRequest{
			RaceId: 100,
		}))
		assert.Nil(t, err)
		assert.Len(t, actual.Msg.Results, 0)
	})

	t.Run("race id should be greater than 0", func(t *testing.T) {
		_, err := client.GetResultsByRace(context.Background(), connect.NewRequest(&apiv1.GetResultsByRaceRequest{
			RaceId: -1,
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
	})

	t.Run("should return validation error when id is not in request", func(t *testing.T) {
		_, err := client.GetResultsByRace(context.Background(), connect.NewRequest(&apiv1.GetResultsByRaceRequest{}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
	})
}
