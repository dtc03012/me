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
import {CardActionArea, CardMedia, createTheme, Grid, Link, Paper, Tooltip, Typography} from "@mui/material";
import {ThemeProvider, withStyles} from "@mui/styles";

const styles = theme => {
    return ({
        typography_name: {
            color: 'inherit',
            letterSpacing: '.1rem',
            textAlign: 'center',
            textDecoration: 'underline',
        },
        typography_detail: {
            color: 'inherit',
            letterSpacing: '.1rem',
            textAlign: 'center',
            textDecoration: 'none',
        }
    });
};

const nameFont = createTheme({
    typography: {
        fontFamily: 'Cinzel',
        fontSize: '30px',
    },
});

const infoFont = createTheme({
    typography: {
        fontFamily: 'Cinzel',
        fontSize: '14px',
    },
});

class Profile extends React.Component {

    render() {
        const { classes } = this.props
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
                        <ThemeProvider theme={nameFont}>
                            <Grid item>
                                    <Typography className={classes.typography_name} >
                                        Kim Tae Hun
                                    </Typography>
                           </Grid>
                        </ThemeProvider>
                        <ThemeProvider theme={infoFont}>
                            <Grid item>
                                    <Typography className={classes.typography_detail}>
                                        software engineer
                                    </Typography>
                                    <Typography className={classes.typography_detail}>
                                        dtc03012
                                    </Typography>
                                    <Box sx={{
                                        display: 'flex',
                                    }}>
                                        <LocationOnIcon/>
                                        <Typography className={classes.typography_detail}>
                                            Seoul, Republic of Korea
                                        </Typography>
                                    </Box>
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

export default withStyles(styles, { withTheme: true })(Profile);