package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dtc03012/me/protobuf/proto/entity/widget"
	"github.com/dtc03012/me/protobuf/proto/service/message"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	encodedAPI                 = os.Getenv("WEATHER_API")
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

func getTimestamppb(baseDate string, baseTime string) *timestamppb.Timestamp {
	var (
		year, _   = strconv.Atoi(baseDate[:4])
		month, _  = strconv.Atoi(baseDate[4:6])
		day, _    = strconv.Atoi(baseDate[6:])
		hour, _   = strconv.Atoi(baseTime[:2])
		minute, _ = strconv.Atoi(baseTime[2:])
	)

	t := time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)

	return timestamppb.New(t)
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

func setTemperature(res *message.FetchDistrictWeatherResponse, weather *Weather, weatherCategoryMap map[string][]*weatherCategoryData) error {

	if len(weatherCategoryMap["TMP"]) == 0 {
		return errors.New("temperature error : now temperature isn't set")
	}

	if len(weatherCategoryMap["TMX"]) == 0 {
		return errors.New("temperature error : highest temperature isn't set")
	}

	if len(weatherCategoryMap["TMN"]) == 0 {
		return errors.New("temperature error : lowest temperature isn't set")
	}

	res.Temperature = &widget.Temperature{
		Now:     weatherCategoryMap["TMP"][0].FcstValue,
		Highest: weatherCategoryMap["TMX"][0].FcstValue,
		Lowest:  weatherCategoryMap["TMN"][0].FcstValue,
	}
	return nil
}

func setSkyCondition(res *message.FetchDistrictWeatherResponse, weather *Weather, weatherCategoryMap map[string][]*weatherCategoryData) error {

	if len(weatherCategoryMap["SKY"]) == 0 {
		return errors.New("weather error : weather isn't set")
	}

	res.Sky = make([]*widget.Sky, 0, len(weatherCategoryMap["SKY"]))

	for _, skyData := range weatherCategoryMap["SKY"] {
		if skyData.FcstValue == "1" {
			res.Sky = append(res.Sky, &widget.Sky{
				SkyCondition: widget.Sky_SUNNY,
				ForecastTime: getTimestamppb(skyData.FcstDate, skyData.FcstTime),
			})
		} else if skyData.FcstValue == "3" {
			res.Sky = append(res.Sky, &widget.Sky{
				SkyCondition: widget.Sky_CLOUDY,
				ForecastTime: getTimestamppb(skyData.FcstDate, skyData.FcstTime),
			})
		} else if skyData.FcstValue == "4" {
			res.Sky = append(res.Sky, &widget.Sky{
				SkyCondition: widget.Sky_OVERCAST,
				ForecastTime: getTimestamppb(skyData.FcstDate, skyData.FcstTime),
			})
		} else {
			return errors.New("sky error : unexpected sky code")
		}
	}

	return nil
}

func setPrecipitationCondition(res *message.FetchDistrictWeatherResponse, weather *Weather, weatherCategoryMap map[string][]*weatherCategoryData) error {

	if len(weatherCategoryMap["PTY"]) == 0 {
		return errors.New("precipitation error : precipitation isn't set")
	}

	res.Precipitation = make([]*widget.Precipitation, 0, len(weatherCategoryMap["PTY"]))

	for _, precipitationData := range weatherCategoryMap["PTY"] {
		if precipitationData.FcstValue == "0" {
			res.Precipitation = append(res.Precipitation, &widget.Precipitation{
				PrecipitationCondition: widget.Precipitation_NONE,
				ForecastTime:           getTimestamppb(precipitationData.FcstDate, precipitationData.FcstTime),
			})
		} else if precipitationData.FcstValue == "1" {
			res.Precipitation = append(res.Precipitation, &widget.Precipitation{
				PrecipitationCondition: widget.Precipitation_RAINY,
				ForecastTime:           getTimestamppb(precipitationData.FcstDate, precipitationData.FcstTime),
			})
		} else if precipitationData.FcstValue == "2" {
			res.Precipitation = append(res.Precipitation, &widget.Precipitation{
				PrecipitationCondition: widget.Precipitation_RAINY_SNOW,
				ForecastTime:           getTimestamppb(precipitationData.FcstDate, precipitationData.FcstTime),
			})
		} else if precipitationData.FcstValue == "3" {
			res.Precipitation = append(res.Precipitation, &widget.Precipitation{
				PrecipitationCondition: widget.Precipitation_SNOW,
				ForecastTime:           getTimestamppb(precipitationData.FcstDate, precipitationData.FcstTime),
			})
		} else if precipitationData.FcstValue == "4" {
			res.Precipitation = append(res.Precipitation, &widget.Precipitation{
				PrecipitationCondition: widget.Precipitation_SHOWER,
				ForecastTime:           getTimestamppb(precipitationData.FcstDate, precipitationData.FcstTime),
			})
		} else {
			return errors.New("precipitation error : unexpected precipitation code")
		}
	}

	return nil
}

func (m *MeServer) FetchDistrictWeather(ctx context.Context, req *message.FetchDistrictWeatherRequest) (*message.FetchDistrictWeatherResponse, error) {

	weather, err := fetchNowWeatherData(500, 1, int(req.Nx), int(req.Ny))
	if err != nil {
		return nil, err
	}

	weatherCategoryMap := make(map[string][]*weatherCategoryData)
	setWeatherCategoryMap(weather, weatherCategoryMap)

	res := &message.FetchDistrictWeatherResponse{}

	err = setTemperature(res, weather, weatherCategoryMap)
	if err != nil {
		return nil, err
	}

	err = setSkyCondition(res, weather, weatherCategoryMap)
	if err != nil {
		return nil, err
	}

	err = setPrecipitationCondition(res, weather, weatherCategoryMap)
	if err != nil {
		return nil, err
	}

	return res, nil
}
