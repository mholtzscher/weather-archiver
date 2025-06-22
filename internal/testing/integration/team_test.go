package integration

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	apiv1 "github.com/mholtzscher/weather-archiver/gen/api/v1"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/assert"
)

func TestCreateTeam(t *testing.T) {
	helper := CreateIntegrationTestHelper(t)
	client := helper.Client

	t.Run("create team should return team id", func(t *testing.T) {
		result, err := client.CreateTeam(context.Background(), connect.NewRequest(&apiv1.CreateTeamRequest{
			Name: gofakeit.Company(),
			Base: gofakeit.Country(),
		}))
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.TeamId)
	})

	t.Run("team should require name", func(t *testing.T) {
		result, err := client.CreateTeam(context.Background(), connect.NewRequest(&apiv1.CreateTeamRequest{
			Base: gofakeit.Country(),
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("team should require base", func(t *testing.T) {
		result, err := client.CreateTeam(context.Background(), connect.NewRequest(&apiv1.CreateTeamRequest{
			Name: gofakeit.Company(),
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("should not allow duplicate team", func(t *testing.T) {
		request := connect.NewRequest(&apiv1.CreateTeamRequest{
			Name: gofakeit.Company(),
			Base: gofakeit.Country(),
		})

		result, err := client.CreateTeam(context.Background(), request)
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.TeamId)

		result, err = client.CreateTeam(context.Background(), request)
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeAlreadyExists, connect.CodeOf(err))
		assert.Nil(t, result)
	})
}
