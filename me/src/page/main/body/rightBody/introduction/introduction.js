import React from "react"
import Box from "@mui/material/Box";
import {Grid, Paper, Typography} from "@mui/material";
import CampaignIcon from '@mui/icons-material/Campaign';
import {ThemeProvider, createTheme} from "@mui/material/styles";
import Weather from "./widget/weather";
import CoinChart from "./widget/coinChart";

const boardTheme = createTheme({
    typography: {
        body1: {
            fontFamily: "Elice Digital Baeum",
            fontSize: 20,
            fontWeight: 700,
        }
    },
});

class Introduction extends React.Component {
    render() {
        return (
            <Box sx={{
                p: 2,
            }}>
                <Paper elevation={3} sx={{
                    p: 2,
                }}>
                    <Grid container spacing={2}
                          direction="row"
                          alignItems="center"
                          justify="center">
                        <Grid item>
                            <CampaignIcon
                                sx = {{
                                    fontSize: 35,
                                }}
                            />
                        </Grid>
                        <Grid item>
                            <ThemeProvider theme={boardTheme}>
                                <Typography variant="body1">
                                    안녕하세요! 반갑습니다!
                                </Typography>
                            </ThemeProvider>
                        </Grid>
                    </Grid>
                </Paper>
                <Grid container spacing="20">
                    <Grid item container spacing="20">
                        <Grid item xs>
                            <Weather districtName = "서울특별시 강남구"
                                     districtNx = "61"
                                     districtNy = "126"/>
                        </Grid>
                        <Grid item xs>
                            <Weather districtName = "대구광역시 중구"
                                     districtNx = "89"
                                     districtNy = "90"/>
                        </Grid>
                        <Grid item xs>
                            <Weather districtName = "부산광역시 기장군"
                                     districtNx = "100"
                                     districtNy = "77"/>
                        </Grid>
                    </Grid>
                    <Grid item container spacing="20">
                        <Grid item xs="6">
                            <CoinChart chartName = "BTC/USDT"/>
                        </Grid>
                        <Grid item xs="6">
                            <CoinChart chartName = "ETH/USDT"/>
                        </Grid>
                    </Grid>
                </Grid>
            </Box>
        )
    }
}

export default Introduction