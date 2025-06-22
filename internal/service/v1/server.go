package v1

import (
	apiv1 "github.com/mholtzscher/weather-archiver/gen/api/v1/apiv1connect"
	"github.com/mholtzscher/weather-archiver/internal/dal"
)

func NewWeatherServer(db dal.Querier) *WeatherServer {
	return &WeatherServer{
		DB: db,
	}
}

type WeatherServer struct {
	apiv1.UnimplementedWeatherServiceHandler
	DB dal.Querier
}
