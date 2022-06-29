import React from 'react'
import Box from "@mui/material/Box";
import {Button, Grid, Typography} from "@mui/material";
import {createTheme, ThemeProvider} from "@mui/material/styles";
import axios from "axios";
import {object} from "prop-types";


class Weather extends React.Component {

    Theme = createTheme({
        typography: {
            body1: {
                fontFamily: "Open+Sans",
                fontSize: 14,
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
                }
            },
        }
        this.fetchWeatherData = this.fetchWeatherData.bind(this)
    }

    fetchWeatherData() {
        let url = "/v2/fetch-district-weather/" + this.state.districtNx + "/" + this.state.districtNy
        axios.get(url).then(
            response => {
                console.log(response)
                this.setState(  {
                    weatherData: {
                        temperature: {
                            highest: response.data.temperature.highest,
                            lowest: response.data.temperature.lowest,
                            now: response.data.temperature.now + "º",
                        }
                    }
                })
            }
        )
    }

    componentDidMount() {
        this.fetchWeatherData()
    }

    render() {
        return (
            <Box sx = {{
                p: 2
            }}>
                <ThemeProvider theme={this.Theme}>
                    <Grid container spacing={2}>
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
                                        fontFamily: "Open+Sans",
                                        fontWeight: 600,
                                    }}>
                                        {this.state.weatherData.temperature.highest}
                                    </Typography>
                                    <Typography sx={{
                                        color: '#3d5afe',
                                        fontSize: 13,
                                        fontFamily: "Open+Sans",
                                        fontWeight: 600,
                                    }}>
                                        {this.state.weatherData.temperature.lowest}
                                    </Typography>
                                </Grid>
                                <Grid item>
                                    <Typography sx={{
                                        fontSize: 23,
                                        fontFamily: "Open+Sans",
                                        fontWeight: 600,
                                    }}>
                                        {this.state.weatherData.temperature.now}
                                    </Typography>
                                </Grid>
                            </Grid>
                        </Grid>
                    </Grid>
                </ThemeProvider>
            </Box>
        )
    }
}

export default Weather