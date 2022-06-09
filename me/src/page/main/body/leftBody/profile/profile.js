import React from "react";
import ProfileImage from "./image/me.jpg"
import Box from "@mui/material/Box";
import LocationOnIcon from '@mui/icons-material/LocationOn';
import InstagramIcon from '@mui/icons-material/Instagram';
import FacebookIcon from '@mui/icons-material/Facebook';
import LinkedInIcon from '@mui/icons-material/LinkedIn';
import GitHubIcon from '@mui/icons-material/GitHub';
import EmailIcon from '@mui/icons-material/Email';
import HomeIcon from '@mui/icons-material/Home';
import {CardActionArea, CardMedia, Grid, Link, Paper, Tooltip, Typography} from "@mui/material";
import {ThemeProvider, createTheme} from "@mui/material/styles";

const Theme = createTheme({
    typography: {
        body1: {
            textAlign: 'center',
            fontFamily: 'Ubuntu',
            fontSize: 20,
            fontWeight: 900,
        },
        body2: {
            textAlign: 'center',
            fontFamily: 'Open+Sans',
            fontSize: 14,
            fontWeight: 500,
            letterSpacing: '.1rem',
        }
    },
});

class Profile extends React.Component {

    render() {
        return (
            <Box sx={{
                    p: 2,
                }}>
                <Paper elevation={3} >
                    <Grid container
                          direction="column"
                          alignItems="center"
                          justify="center"
                    >
                        {/*My Picture*/}
                        <Grid item>
                            <CardActionArea>
                                <CardMedia
                                    component="img"
                                    image={ProfileImage}
                                    title="It's me!"
                                />
                            </CardActionArea>
                        </Grid>

                        {/*My Information*/}
                        <ThemeProvider theme={Theme}>
                            <Grid item>
                                    <Typography variant="body1" sx={{
                                        textDecoration: "underline",
                                    }}>
                                        Kim Tae Hun
                                    </Typography>
                           </Grid>
                        </ThemeProvider>
                        <ThemeProvider theme={Theme}>
                            <Grid item>
                                    <Typography variant="body2" >
                                        software engineer <br/>
                                        dtc03012 <br/>
                                        <Box sx={{
                                            display: 'flex',
                                        }}>
                                            <LocationOnIcon/>
                                            <Typography variant="body2">
                                                Seoul, Republic of Korea
                                            </Typography>
                                        </Box>
                                    </Typography>
                            </Grid>
                        </ThemeProvider>

                        {/*My sns*/}
                        <Grid item sx={{
                            paddingTop: 2,
                            paddingBottom: 1,
                            width: '200px',
                        }}>
                            <Box display='flex' justifyContent='space-between'>
                                <Link href="https://www.instagram.com/ggggssddddrr/" color="inherit" target="_blank">
                                    <Tooltip title="Instagram" >
                                        <InstagramIcon/>
                                    </Tooltip>
                                </Link>
                                <Link href="https://www.facebook.com/profile.php?id=100003841952503" color="inherit" target="_blank">
                                    <Tooltip title="Facebook">
                                        <FacebookIcon/>
                                    </Tooltip>
                                </Link>
                                <Link href="https://www.linkedin.com/in/taehun-kim-b622431b7/" color="inherit" target="_blank">
                                    <Tooltip title="LinkedIn">
                                        <LinkedInIcon/>
                                    </Tooltip>
                                </Link>
                                <Link href="https://github.com/dtc03012" color="inherit" target="_blank">
                                    <Tooltip title="Github">
                                        <GitHubIcon/>
                                    </Tooltip>
                                </Link>
                                <Link href="https://dtc03012.tistory.com/" color="inherit" target="_blank">
                                    <Tooltip title="Tistory">
                                        <HomeIcon/>
                                    </Tooltip>
                                </Link>
                                <Link href="mailto:dtc03012@gmail.com" color="inherit">
                                    <Tooltip title="Gmail">
                                        <EmailIcon/>
                                    </Tooltip>
                                </Link>
                            </Box>
                        </Grid>
                    </Grid>
                </Paper>
            </Box>
        )
    }
}

export default Profile;