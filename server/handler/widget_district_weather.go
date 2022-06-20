package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dtc03012/me/protobuf/proto/service/message"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	encodedAPI = os.Getenv("weatherAPI")
)

type Item struct {
	Category  string `json:"category"`
	FcstDate  string `json:"fcstDate"`
	FcstTime  string `json:"fcstTime"`
	FcstValue string `json:"fcstValue"`
	Nx        int    `json:"nx"`
	Ny        int    `json:"ny"`
}

type Items struct {
	ItemArr []*Item `json:"item"`
}

type Body struct {
	Items *Items `json:"items"`
}

type WeatherResponse struct {
	Body *Body `json:"body"`
}

type Weather struct {
	Response *WeatherResponse `json:"response"`
}

func fetchNowWeatherData(numOfRows int, pageNo int, nx int, ny int) (*Weather, error) {

	var (
		url      string
		baseDate string
		baseTime string
	)

	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	if minute <= 45 {
		hour--
		if hour < 0 {
			hour += 24
			yesterday := now.AddDate(0, 0, -1)
			year = yesterday.Year()
			month = yesterday.Month()
			day = yesterday.Day()
		}
	}
	minute = 30

	baseDate = fmt.Sprintf("%4d%02d%02d", year, month, day)
	baseTime = fmt.Sprintf("%02d%02d", hour, minute)

	url = "http://apis.data.go.kr/1360000/VilageFcstInfoService_2.0/getUltraSrtNcst?serviceKey=" + encodedAPI
	url = url + "&numOfRows=" + strconv.Itoa(numOfRows)
	url = url + "&pageNo=" + strconv.Itoa(pageNo)
	url = url + "&base_date=" + baseDate
	url = url + "&base_time=" + baseTime
	url = url + "&dataType=JSON"
	url = url + "&nx=" + strconv.Itoa(nx)
	url = url + "&ny=" + strconv.Itoa(ny)

	res, err := http.Get(url)
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	weather := &Weather{}
	err = json.NewDecoder(res.Body).Decode(weather)

	if err != nil {
		return nil, err
	}

	return weather, err
}

func (m *MeServer) CheckDistrictWeather(ctx context.Context, req *message.CheckDistrictWeatherRequest) (*message.CheckDistrictWeatherResponse, error) {

	_, err := fetchNowWeatherData(100, 1, 55, 127)
	if err != nil {
		return nil, err
	}

	return &message.CheckDistrictWeatherResponse{}, nil
}
