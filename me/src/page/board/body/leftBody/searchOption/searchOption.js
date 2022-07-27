import React from "react";
import Box from "@mui/material/Box";
import {Grid, Paper, Typography} from "@mui/material";
import NameSearchOption from "./nameSearchOption";
import TopicSearchOption from "./topicSearchOption";
import TagSearchOption from "./tagSearchOption";
import {createTheme, ThemeProvider} from "@mui/material/styles";

const Theme = createTheme({
    typography: {
        body1: {
            fontFamily: "Elice Digital Baeum",
            fontSize: 8,
            fontWeight: 700,
        }
    },
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
                          alignItems="center"
                          justify="center">
                        <Grid item>
                            <NameSearchOption/>
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