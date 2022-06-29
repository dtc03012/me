package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dtc03012/me/protobuf/proto/entity/widget"
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
	cacheWeatherDataMap        = make(map[cacheWeatherKey]*Weather)
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

type weatherCategoryData struct {
	FcstDate  string
	FcstTime  string
	FcstValue string
}

type cacheWeatherKey struct {
	Nx       int
	Ny       int
	baseDate string
	baseTime string
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

	cacheWeatherKey := cacheWeatherKey{
		Nx:       nx,
		Ny:       ny,
		baseTime: baseTime,
		baseDate: baseDate,
	}

	cacheWeather := cacheWeatherDataMap[cacheWeatherKey]

	if cacheWeather != nil {
		return cacheWeather, nil
	}

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

	cacheWeatherDataMap[cacheWeatherKey] = weather

	return weather, err
}

func setWeatherCategoryMap(weather *Weather, weatherCategoryMap map[string][]*weatherCategoryData) {
	for _, item := range weather.Response.Body.Items.ItemArr {
		weatherCategoryMap[item.Category] = append(weatherCategoryMap[item.Category], &weatherCategoryData{
			FcstValue: item.FcstValue,
			FcstTime:  item.FcstTime,
			FcstDate:  item.FcstDate,
		})
	}
}

func getTemperature(weather *Weather, weatherCategoryMap map[string][]*weatherCategoryData) (string, string, string, error) {

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

func getWeatherCondition(weather *Weather, weatherCategoryMap map[string][]*weatherCategoryData) (widget.WeatherCondition, error) {

	if len(weatherCategoryMap["SKY"]) == 0 {
		return widget.Weather_NONE, errors.New("weather error : weather isn't set")
	}

	weatherData := weatherCategoryMap["SKY"][0]

	if weatherData.FcstValue == "1" {
		return widget.Weather_SUNNY, nil
	} else if weatherData.FcstValue == "3" {
		return widget.Weather_CLOUDY, nil
	} else if weatherData.FcstValue == "4" {
		return widget.Weather_OVERCAST, nil
	}

	return widget.Weather_NONE, errors.New("weather error : unexpected weather code")
}

func getPrecipitationCondition(weather *Weather, weatherCategoryMap map[string][]*weatherCategoryData) (widget.PrecipitationCondition, error) {

	if len(weatherCategoryMap["PTY"]) == 0 {
		return widget.Precipitation_NONE, errors.New("precipitation error : precipitation isn't set")
	}

	precipitationData := weatherCategoryMap["PTY"][0]

	if precipitationData.FcstValue == "0" {
		return widget.Precipitation_RAINY, nil
	} else if precipitationData.FcstValue == "1" {
		return widget.Precipitation_RAINY_SNOW, nil
	} else if precipitationData.FcstValue == "2" {
		return widget.Precipitation_SNOW, nil
	} else if precipitationData.FcstValue == "3" {
		return widget.Precipitation_SHOWER, nil
	}

	return widget.Precipitation_NONE, errors.New("precipitation error : unexpected precipitation code")
}

func (m *MeServer) FetchDistrictWeather(ctx context.Context, req *message.FetchDistrictWeatherRequest) (*message.FetchDistrictWeatherResponse, error) {

	weather, err := fetchNowWeatherData(500, 1, int(req.Nx), int(req.Ny))
	if err != nil {
		return nil, err
	}

	weatherCategoryMap := make(map[string][]*weatherCategoryData)
	setWeatherCategoryMap(weather, weatherCategoryMap)

	res := &message.FetchDistrictWeatherResponse{}

	nowTemperature, highestTemperature, lowestTemperature, err := getTemperature(weather, weatherCategoryMap)
	if err != nil {
		return nil, err
	}

	res.Temperature = &widget.Temperature{
		Now:     nowTemperature,
		Highest: highestTemperature,
		Lowest:  lowestTemperature,
	}

	weatherCondition, err := getWeatherCondition(weather, weatherCategoryMap)
	if err != nil {
		return nil, err
	}

	res.Weather = &widget.Weather{
		WeatherCondition: weatherCondition,
	}

	precipitationCondition, err := getPrecipitationCondition(weather, weatherCategoryMap)
	if err != nil {
		return nil, err
	}

	res.Precipitation = &widget.Precipitation{
		PrecipitationCondition: precipitationCondition,
	}
	
	return res, nil
}
