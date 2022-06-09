import React from "react"
import Box from "@mui/material/Box";
import {Grid, Paper, Typography} from "@mui/material";
import {ThemeProvider, createTheme} from "@mui/material/styles";

const boardTheme = createTheme({
    typography: {
        fontFamily: 'Ubuntu',
        fontSize: 16,
        body1: {
            fontSize: 24,
            fontWeight: 900,
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
                    <Grid container
                          direction="column"
                    >
                        <Grid item>
                            <ThemeProvider theme={boardTheme}>
                                <Typography>
                                    Hello World! <br/>
                                    I'm Taehun <br/>
                                    This is my first homepage <br/>
                                    Have a Good time in Here and Hope you have been well! <br/>
                                    i'm so honored to meet you <br/>
                                    Have a nice day~~! <br/>
                                </Typography>
                            </ThemeProvider>
                        </Grid>
                    </Grid>
                </Paper>
            </Box>
        )
    }
}

export default Introduction