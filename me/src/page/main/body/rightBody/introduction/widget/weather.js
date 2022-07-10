import React from 'react'
import Box from "@mui/material/Box";
import {Button, Grid, Paper, Typography} from "@mui/material";
import {createTheme, ThemeProvider} from "@mui/material/styles";
import axios from "axios";

const skyCondition = {
    SUNNY: "SUNNY",
    CLOUDY: "CLOUDY",
    OVERCAST: "OVERCAST",
}

const precipitationCondition = {
    NONE: "NONE",
    RAINY: "RAINY",
    RAINY_SNOW: "RAINY_SNOW",
    SNOW: "SNOW",
    SHOWER: "SHOWER",
}

const unexpectedEndOfTime = -1

class Weather extends React.Component {

    Theme = createTheme({
        typography: {
            body1: {
                fontFamily: "Elice Digital Baeum",
                fontSize: 18,
            },
        },
    });

    constructor(props) {
        super(props);
        this.state = {
            districtName: props.districtName,
            districtNx: props.districtNx,
            districtNy: props.districtNy,
            weatherData: {
                temperature: {
                    highest: undefined,
                    lowest: undefined,
                    now: undefined,
                },
                sky: [{
                    condition: undefined,
                    forecastTime: undefined,
                }],
                precipitation: [{
                    condition: undefined,
                    forecastTime: undefined,
                }]
            },
        }
        this.fetchWeatherData = this.fetchWeatherData.bind(this)
    }

    fetchWeatherData() {
        let url = "/v2/fetch-district-weather/" + this.state.districtNx + "/" + this.state.districtNy
        axios.get(url).then(
            response => {
                this.setState(  {
                    weatherData: {
                        temperature: {
                            highest: response.data.temperature.highest,
                            lowest: response.data.temperature.lowest,
                            now: response.data.temperature.now + "º",
                        },
                        sky : response.data.sky.map(skyData => (
                            {
                                condition: skyData.skyCondition,
                                forecastTime: skyData.forecastTime,
                            }
                        )),
                        precipitation : response.data.precipitation.map(precipitationData => (
                            {
                                condition: precipitationData.precipitationCondition,
                                forecastTime: precipitationData.forecastTime,
                            }
                        ))
                    }
                })
            }
        )
    }

    analyzeCurrentWeather() {

        if(this.state.weatherData.sky.length === 0 || this.state.weatherData.precipitation.length === 0) {
            return "현재 기상 정보를 확인할 수 없습니다"
        }

        let message = "현재 날씨는 "
        if(this.state.weatherData.sky.length > 0) {
            if(this.state.weatherData.sky[0].condition === skyCondition.SUNNY) {
                message += "맑은 날씨이며 "
            }else if(this.state.weatherData.sky[0].condition === skyCondition.CLOUDY) {
                message += "구름이 가득하고 "
            }else if(this.state.weatherData.sky[0].condition === skyCondition.OVERCAST) {
                message += "흐리고 "
            }
        }

        if(this.state.weatherData.precipitation.length > 0) {
            if(this.state.weatherData.precipitation[0].condition === precipitationCondition.NONE) {
                message += "강수 확률은 없을 것으로 예상됩니다."
            }else if(this.state.weatherData.precipitation[0].condition === precipitationCondition.SNOW) {
                message += "눈이 올 것으로 예상됩니다."
            }else if(this.state.weatherData.precipitation[0].condition === precipitationCondition.RAINY_SNOW) {
                message += "눈과 비가 올 것으로 예상됩니다."
            }else if(this.state.weatherData.precipitation[0].condition === precipitationCondition.RAINY) {
                message += "비가 올 것으로 예상됩니다."
            }else if(this.state.weatherData.precipitation[0].condition === precipitationCondition.SHOWER) {
                message += "소나기가 올 것으로 예상됩니다."
            }
        }

        return message
    }

    createListForCreatingConditionMessage(conditionList, condition){
        let timeList = [], beginTime = -1, lastTime = -1

        conditionList.forEach(data => {
            const time = Date.parse(data.forecastTime)
            if(data.condition === condition){
                if(beginTime == -1) {
                    beginTime = time
                }
                lastTime = time
            }else {
                if(beginTime != -1) {
                    timeList.push({beginTime: beginTime, endTime: time})
                    beginTime = -1
                }
            }
        })

        if(beginTime != -1){
            timeList.push({beginTime: beginTime, endTime: lastTime})
        }

        return timeList
    }

    getHourMessage(date) {
        if(date.getUTCHours() < 12) {
            return "오전 " + date.getUTCHours()
        }
        date.setUTCHours(date.getUTCHours()-12)
        return "오후 " + date.getUTCHours()
    }

    generateExpectedSkyConditionMessage(date) {
        let skyConditionMessage = ""

        const overcastTimeList = this.createListForCreatingConditionMessage(this.state.weatherData.sky, skyCondition.OVERCAST)

        let isWritten = false
        overcastTimeList.forEach(timeRange => {
            let beginTime = new Date(timeRange.beginTime)
            let endTime = new Date(timeRange.endTime)

            if(beginTime.getUTCDate() === date.getDate()) {
                if(endTime.getUTCDate() === date.getDate()) {
                    skyConditionMessage += "   - " + this.getHourMessage(beginTime) + "시부터 ~ " + this.getHourMessage(endTime) + "시까지 흐립니다.\n"
                }else {
                    skyConditionMessage += "   - " + this.getHourMessage(beginTime) + "시부터 ~ 다음 날 " + this.getHourMessage(endTime) + "시까지 흐립니다.\n"
                }
                isWritten = true
            }
        })

        if(!isWritten){
            skyConditionMessage += "   - 하루종일 맑은 날씨입니다.\n"
        }

        return skyConditionMessage
    }

    generateExpectedPrecipitationConditionMessage(date) {
        let precipitationConditionMessage = ""

        const rainyTimeList = this.createListForCreatingConditionMessage(this.state.weatherData.precipitation, precipitationCondition.RAINY)

        if(rainyTimeList.length > 0) {
            rainyTimeList.forEach(timeRange => {
                let beginTime = new Date(timeRange.beginTime)
                let endTime = new Date(timeRange.endTime)

                if(beginTime.getUTCDate() === date.getDate()) {
                    if(endTime.getUTCDate() === date.getDate()) {
                        precipitationConditionMessage += "   - " + this.getHourMessage(beginTime) + "시부터 ~ " + this.getHourMessage(endTime) + "시까지 비가 옵니다.\n"
                    }else {
                        precipitationConditionMessage += "   - " + this.getHourMessage(beginTime) + "시부터 ~ 다음날 " + this.getHourMessage(endTime) + "시까지 비가 옵니다\n"
                    }
                }
            })
        }

        const rainySnowTimeList = this.createListForCreatingConditionMessage(this.state.weatherData.precipitation, precipitationCondition.RAINY_SNOW)

        if(rainySnowTimeList.length > 0) {
            rainySnowTimeList.forEach(timeRange => {
                let beginTime = new Date(timeRange.beginTime)
                let endTime = new Date(timeRange.endTime)

                if(beginTime.getUTCDate() === date.getDate()) {
                    if(endTime.getUTCDate() === date.getDate()) {
                        precipitationConditionMessage += "   - " + this.getHourMessage(beginTime) + "시부터 ~ " + this.getHourMessage(endTime) + "시까지 비와 눈이 옵니다.\n"
                    }else {
                        precipitationConditionMessage += "   - " + this.getHourMessage(beginTime) + "시부터 ~ 다음날 " + this.getHourMessage(endTime) + "시까지 비와 눈이 옵니다\n"
                    }
                }
            })
        }

        const snowTimeList = this.createListForCreatingConditionMessage(this.state.weatherData.precipitation, precipitationCondition.SNOW)

        if(snowTimeList.length > 0) {
            snowTimeList.forEach(timeRange => {
                let beginTime = new Date(timeRange.beginTime)
                let endTime = new Date(timeRange.endTime)

                if(beginTime.getUTCDate() === date.getDate()) {
                    if(endTime.getUTCDate() === date.getDate()) {
                        precipitationConditionMessage += "   - " + this.getHourMessage(beginTime) + "시부터 ~ " + this.getHourMessage(endTime) + "시까지 눈이 옵니다.\n"
                    }else {
                        precipitationConditionMessage += "   - " + this.getHourMessage(beginTime) + "시부터 ~ 다음날 " + this.getHourMessage(endTime) + "시까지 눈이 옵니다\n"
                    }
                }
            })
        }

        const showerTimeList = this.createListForCreatingConditionMessage(this.state.weatherData.precipitation, precipitationCondition.SHOWER)

        if(showerTimeList.length > 0) {
            showerTimeList.forEach(timeRange => {
                let beginTime = new Date(timeRange.beginTime)
                let endTime = new Date(timeRange.endTime)

                if(beginTime.getUTCDate() === date.getDate()) {
                    if(endTime.getUTCDate() === date.getDate()) {
                        precipitationConditionMessage += "   - " + this.getHourMessage(beginTime) + "시부터 ~ " + this.getHourMessage(endTime) + "시까지 소나기가 옵니다.\n"
                    }else {
                        precipitationConditionMessage += "   - " + this.getHourMessage(beginTime) + "시부터 ~ 다음날 " + this.getHourMessage(endTime) + "시까지 소나기가 옵니다\n"
                    }
                }
            })
        }

        return precipitationConditionMessage
    }

    generateWeatherMessage(date) {
        let message = ""

        message += this.generateExpectedSkyConditionMessage(date)
        message += this.generateExpectedPrecipitationConditionMessage(date)

        return message
    }

    analyzeExpectedWeather() {
        let message = ""

        let currentTime = new Date()
        let currentWeatherMessage = this.generateWeatherMessage(currentTime)

        if(currentWeatherMessage == "") {
            return ""
        }

        message = '● 오늘 기상 예보\n'
        message += currentWeatherMessage

        let tomorrowTime = new Date(currentTime)
        tomorrowTime.setUTCDate(tomorrowTime.getUTCDate() + 1)

        let tomorrowWeatherMessage = this.generateWeatherMessage(tomorrowTime)

        if(tomorrowWeatherMessage != "") {
            message += '\n● 내일 기상 예보\n'
            message += tomorrowWeatherMessage
        }

        return message
    }

    componentDidMount() {
        this.fetchWeatherData()
    }

    render() {
        return (
            <Paper elevation={3} sx={{
                mt: 2,
            }}>
                <Box sx = {{
                    p: 2
                }}>
                    <ThemeProvider theme={this.Theme}>
                        <Grid container direction="column">
                            <Grid item container spacing={2}>
                                <Grid item xs container>
                                    <Grid item container direction="column" >
                                        <Grid item>
                                            <Typography variant="body1" sx={{
                                                fontWeight: 500,
                                            }}>
                                                현재 위치
                                            </Typography>
                                        </Grid>
                                        <Grid item container spacing={2}>
                                            <Grid item sm>
                                                <Typography variant="body1" sx={{
                                                    fontWeight: 900,
                                                }}>
                                                    {this.state.districtName}
                                                </Typography>
                                            </Grid>
                                        </Grid>
                                    </Grid>
                                </Grid>
                                <Grid item>
                                    <Grid item xs container>
                                        <Grid item xs container direction="column" sx={{
                                            pr: 1
                                        }}>
                                            <Typography sx={{
                                                color: '#f44336',
                                                fontSize: 13,
                                                fontWeight: 600,
                                            }}>
                                                {this.state.weatherData.temperature.highest}
                                            </Typography>
                                            <Typography sx={{
                                                color: '#3d5afe',
                                                fontSize: 13,
                                                fontWeight: 600,
                                            }}>
                                                {this.state.weatherData.temperature.lowest}
                                            </Typography>
                                        </Grid>
                                        <Grid item>
                                            <Typography sx={{
                                                fontSize: 23,
                                                fontWeight: 600,
                                            }}>
                                                {this.state.weatherData.temperature.now}
                                            </Typography>
                                        </Grid>
                                    </Grid>
                                </Grid>
                            </Grid>
                            <Grid item sx={{
                                pt: 3
                            }}>
                                <Typography sx={{
                                    fontSize: 15,
                                    fontWeight: 600
                                }}>
                                    {this.analyzeCurrentWeather()}
                                </Typography>
                            </Grid>
                            <Grid item sx={{
                                pt: 4
                            }}>
                                <Typography sx={{
                                    fontSize: 15,
                                    fontWeight: 600,
                                    whiteSpace: "pre-wrap",
                                }}>
                                    {this.analyzeExpectedWeather()}
                                </Typography>
                            </Grid>
                        </Grid>
                    </ThemeProvider>
                </Box>
            </Paper>
        )
    }
}

export default Weather