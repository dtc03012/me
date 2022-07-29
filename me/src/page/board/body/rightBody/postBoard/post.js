import React from 'react';
import {Avatar, Chip, Grid, Paper, Typography} from "@mui/material";
import Box from "@mui/material/Box";
import FavoriteIcon from '@mui/icons-material/Favorite';
import {red} from "@mui/material/colors";

class Post extends React.Component {

    constructor(props) {
        super(props);
    }

    createAvatar = () => {
        if(this.props.postInfo.avatarInfo.avatarImgUrl != "") {
            return <Avatar alt={this.props.postInfo.userId} src={this.props.postInfo.avatarInfo.avatarImgUrl} sx={{ width: 50, height: 50 }}/>
        }
        return <Avatar sx={{bgcolor: this.props.postInfo.avatarInfo.avatarBgColor, width: 50, height: 50 }}> {this.props.postInfo.avatarInfo.avatarInitial} </Avatar>
    }

    createTagList = () => {
        return (this.props.postInfo.tags.map( (tag) => (
                <Grid item>
                    <Chip label={String(tag)} variant="outlined" sx={{
                        '& .MuiChip-label': {
                            fontSize: 14,
                            fontFamily: "Elice Digital Baeum",
                        },
                    }} />
                </Grid>
            )
        ))
    }

    convertTimeStampToReadableTime = (timeStamp) => {
        const date = new Date(timeStamp)
        const year = date.getFullYear()
        const month = date.getMonth() + 1
        const day = date.getDate()
        return year.toString() + "/" + month.toString() + "/" + day.toString()
    }

    render() {
        return (
            <Box>
                <Paper elevation={3} sx={{
                    p: 2.5,
                }}>
                    <Grid container direction="row" alignItems="center">
                        <Grid item container direction="column"  spacing={1} sx={{
                            pl: 1,
                        }} sm={1.5} xs={1.5}>
                            <Grid item sx={{
                                ml: 0.5,
                            }}>
                                {this.createAvatar()}
                            </Grid>
                            <Grid item>
                                <Typography sx={{
                                    fontSize: 13,
                                    fontWeight: 900,
                                    fontFamily: "Elice Digital Baeum",
                                }}>{this.props.postInfo.userId}</Typography>
                            </Grid>
                        </Grid>
                        <Grid item container direction="column" spacing={1.5} sx={{
                            pl: 2,
                        }} sm={9} xs={9}>
                            <Grid item>
                                <Typography sx={{
                                    fontSize: 30,
                                    fontWeight: 900,
                                    fontFamily: "Elice Digital Baeum",
                                }}>
                                    {this.props.postInfo.title}
                                </Typography>
                            </Grid>
                            <Grid item container direction="row" spacing={2}>
                                {this.createTagList()}
                            </Grid>
                        </Grid>
                        <Grid item container direction="column" spacing={1} sm={1.5} xs={1.5}>
                            <Grid item>
                                <Typography sx={{
                                    fontSize: 14,
                                    fontWeight: 600,
                                    fontFamily: "Elice Digital Baeum",
                                }}>
                                    {this.convertTimeStampToReadableTime(this.props.postInfo.timestamp)}
                                </Typography>
                            </Grid>
                            <Grid item>
                                <Typography sx={{
                                    fontSize: 14,
                                    fontWeight: 600,
                                    fontFamily: "Elice Digital Baeum",
                                }}>
                                    예상 읽는 시간: {this.props.postInfo.readExpTime}분
                                </Typography>
                            </Grid>
                            <Grid item container direction="row" spacing="5" alignItems="center">
                                <Grid item>
                                    <FavoriteIcon style={{
                                        color: 'red'
                                    }}/>
                                </Grid>
                                <Grid item>
                                    <Typography sx={{
                                        fontSize: 14,
                                        fontWeight: 600,
                                        fontFamily: "Elice Digital Baeum",
                                    }}>
                                        {this.props.postInfo.heartCnt}
                                    </Typography>
                                </Grid>
                            </Grid>
                        </Grid>
                    </Grid>
                </Paper>
            </Box>
        )
    }
}

export default Post;