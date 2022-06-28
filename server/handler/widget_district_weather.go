package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dtc03012/me/protobuf/proto/entity"
	"github.com/dtc03012/me/protobuf/proto/service/message"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	encodedAPI                 = os.Getenv("weatherAPI")
	weatherBaseDateList        = []string{"0200", "0500", "0800", "1100", "1400", "1700", "2000", "2300"}
	weatherBaseDateAPITimeList = []string{"0209", "0509", "0809", "1109", "1409", "1709", "2009", "2309"}
	weatherCategoryMap         = make(map[string][]*WeatherCategoryData)
)

const (
	unDefinedTemperature = "undefined temperature"
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

type WeatherCategoryData struct {
	FcstDate  string
	FcstTime  string
	FcstValue string
}

func getBaseDateFormat(time time.Time) string {
	year := time.Year()
	month := time.Month()
	day := time.Day()
	return fmt.Sprintf("%4d%02d%02d", year, month, day)
}

func getBaseTimeFormat(time time.Time) string {
	hour := time.Hour()
	minute := time.Minute()
	nowTime := fmt.Sprintf("%02d%02d", hour, minute)

	for i, _ := range weatherBaseDateAPITimeList {
		if nowTime < weatherBaseDateAPITimeList[i] {
			return weatherBaseDateList[i-1]
		}
	}

	return weatherBaseDateList[len(weatherBaseDateList)-1]
}

func getBaseDateTime() (string, string) {
	now := time.Now()
	if now.Hour() < 2 || (now.Hour() == 2 && now.Minute() < 9) {
		now = now.AddDate(0, 0, -1)
		return getBaseDateFormat(now), "2300"
	}
	return getBaseDateFormat(now), getBaseTimeFormat(now)
}

func fetchNowWeatherData(numOfRows int, pageNo int, nx int, ny int) (*Weather, error) {

	var (
		url      string
		baseDate string
		baseTime string
	)

	baseDate, baseTime = getBaseDateTime()

	url = "http://apis.data.go.kr/1360000/VilageFcstInfoService_2.0/getVilageFcst?serviceKey=" + encodedAPI
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

	if weather.Response == nil || weather.Response.Body == nil {
		return nil, errors.New("weather API error : weather can't be fetched from the API")
	}

	return weather, err
}

func setWeatherCategoryMap(weather *Weather) {
	for _, item := range weather.Response.Body.Items.ItemArr {
		weatherCategoryMap[item.Category] = append(weatherCategoryMap[item.Category], &WeatherCategoryData{
			FcstValue: item.FcstValue,
			FcstTime:  item.FcstTime,
			FcstDate:  item.FcstDate,
		})
	}
}

func getTemperature(weather *Weather) (string, string, string, error) {

	if len(weatherCategoryMap["TMP"]) == 0 {
		return unDefinedTemperature, unDefinedTemperature, unDefinedTemperature, errors.New("temperature error : now temperature isn't set")
	}

	if len(weatherCategoryMap["TMX"]) == 0 {
		return unDefinedTemperature, unDefinedTemperature, unDefinedTemperature, errors.New("temperature error : highest temperature isn't set")
	}

	if len(weatherCategoryMap["TMN"]) == 0 {
		return unDefinedTemperature, unDefinedTemperature, unDefinedTemperature, errors.New("temperature error : lowest temperature isn't set")
	}

	return weatherCategoryMap["TMX"][0].FcstValue, weatherCategoryMap["TMX"][0].FcstValue, weatherCategoryMap["TMN"][0].FcstValue, nil
}

func (m *MeServer) CheckDistrictWeather(ctx context.Context, req *message.CheckDistrictWeatherRequest) (*message.CheckDistrictWeatherResponse, error) {

	weather, err := fetchNowWeatherData(300, 1, 55, 127)
	if err != nil {
		return nil, err
	}

	setWeatherCategoryMap(weather)

	res := &message.CheckDistrictWeatherResponse{}

	nowTemperature, highestTemperature, lowestTemperature, err := getTemperature(weather)
	if err != nil {
		return nil, err
	}

	res.Temperature = &entity.Temperature{
		Now:     nowTemperature,
		Highest: highestTemperature,
		Lowest:  lowestTemperature,
	}

	return res, nil
}
