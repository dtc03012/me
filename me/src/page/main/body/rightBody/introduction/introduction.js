import React from "react"
import Box from "@mui/material/Box";
import {Grid, Paper, Typography} from "@mui/material";
import CampaignIcon from '@mui/icons-material/Campaign';
import {ThemeProvider, createTheme} from "@mui/material/styles";
import Weather from "./widget/weather";

const boardTheme = createTheme({
    typography: {
        body1: {
            fontFamily: 'Open+Sans',
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
                                <Typography>
                                    안녕하세요! 반갑습니다!
                                </Typography>
                            </ThemeProvider>
                        </Grid>
                    </Grid>
                </Paper>
                <Grid container spacing="20">
                    <Grid item xs="6">
                        <Paper elevation={3} sx={{
                            mt: 2,
                        }}>
                            <Weather/>
                        </Paper>
                    </Grid>
                    <Grid item xs="6">
                        <Paper elevation={3} sx={{
                            mt: 2,
                        }}>
                            <Weather/>
                        </Paper>
                    </Grid>
                </Grid>
            </Box>
        )
    }
}

export default Introduction