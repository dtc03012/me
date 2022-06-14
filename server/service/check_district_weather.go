package service

import (
	"context"
	"github.com/dtc03012/me/protobuf/proto/service/message"
	"log"
)

func (m *MeServer) CheckDistrictWeather(ctx context.Context, req *message.CheckDistrictWeatherRequest) (*message.CheckDistrictWeatherResponse, error) {
	log.Printf(req.District)
	return &message.CheckDistrictWeatherResponse{}, nil
}
