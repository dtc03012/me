import React from "react";
import Box from "@mui/material/Box";
import {Grid, Paper, Typography} from "@mui/material";
import TextSearchOption from "./textSearchOption";
import TopicSearchOption from "./topicSearchOption";
import TagSearchOption from "./tagSearchOption";
import {createTheme, ThemeProvider} from "@mui/material/styles";

const Theme = createTheme({
    typography: {
        body1: {
            fontFamily: "Elice Digital Baeum",
            fontSize: 13,
            fontWeight: 700,
        },
        body2: {
            fontFamily: "Elice Digital Baeum",
            fontSize: 20,
            fontWeight: 1000,
        }
    },
    palette: {
        background: {
            belowAppbarBox: '#212121',
        },
    }
});

class SearchOption extends React.Component {
    render() {
        return (
            <Box sx={{
                p: 6,
            }}>
                <Paper elevation={3} >
                    <Grid spacing={2} container
                          direction="column"
                          sx={{
                              pl: 4,
                          }}
                    >
                        <Grid item sx={{
                            width: '80%',
                            justifyContent: 'center'
                        }}>
                            <ThemeProvider theme={Theme}>
                                <Typography variant="body2" sx={{
                                    pb: 1,
                                }}>
                                    검색 옵션
                                </Typography>
                                <Box sx={{
                                    backgroundColor: 'background.belowAppbarBox',
                                    width: '100%',
                                    height: 2,
                                }}/>
                            </ThemeProvider>
                        </Grid>
                        <Grid item>
                            <TextSearchOption/>
                        </Grid>
                        <Grid item>
                            <TopicSearchOption/>
                        </Grid>
                        <Grid item>
                            <TagSearchOption/>
                        </Grid>
                        <Grid item sx={{
                            pb: 1,
                        }}>
                            <ThemeProvider theme={Theme}>
                                <Grid item>
                                    <Typography variant="body1">
                                        기능 추가 문의는 dtc03012@gmail.com 로 해주세요!
                                    </Typography>
                                </Grid>
                            </ThemeProvider>
                        </Grid>
                    </Grid>
                </Paper>

            </Box>
        )
    }
}

export default SearchOption