import React, {useState, useEffect} from 'react';
import {Avatar, Chip, Grid, IconButton, Link, Paper, Typography} from "@mui/material";
import FavoriteIcon from '@mui/icons-material/Favorite';
import "./post.css";
import {getCookie, setCookie} from "../../../../../../util/cookie";
import {v4} from "uuid";
import axios from "axios";

export default function Post(props) {

    const [likes, setLikes] = useState(0)
    const [isLike, setIsLike] = useState(false)

    const createAvatar = () => {
        if(props.postInfo.avatarInfo.avatarImgUrl !== "") {
            return <Avatar alt={props.postInfo.writer} src={props.postInfo.avatarInfo.avatarImgUrl} sx={{ width: 50, height: 50 }}/>
        }
        return <Avatar sx={{bgcolor: props.postInfo.avatarInfo.avatarBgColor, width: 50, height: 50 }}> {props.postInfo.avatarInfo.avatarInitial} </Avatar>
    }

    const createTagList = () => {
        return (props.postInfo.tags.map( (tag) => (
                <Grid item>
                    <Chip label={String(tag)} variant="outlined" sx={{
                        '& .MuiChip-label': {
                            fontSize: 12,
                            fontFamily: "Elice Digital Baeum",
                        },
                    }} />
                </Grid>
            )
        ))
    }

    const handleLikeIconClick = (event) => {

        if(getCookie("uuid") === "") {
            let uuid = v4()
            setCookie("uuid", uuid)
        }

        let uuid = getCookie("uuid")

        if(isLike) {
            let url = "/v2/decrement-board-like?postId=" + props.postInfo.id
            url += "&uuid=" + uuid

            axios.delete(url).then( response => {
                setIsLike(!isLike)
                setLikes(likes-1)
            }).catch( err => {
                console.log(err)
            })
        } else {
            let url = "/v2/increment-board-like?postId=" + props.postInfo.id
            url += "&uuid=" + uuid

            axios.put(url).then( response => {
                setIsLike(!isLike)
                setLikes(likes+1)
            }).catch( err => {
                console.log(err)
            })
        }
    }

    useEffect(() => {
        setLikes(props.postInfo.likes)
        setIsLike(props.postInfo.isLike)
    }, [])

    const convertTimeStampToReadableTime = (timeStamp) => {
        const date = new Date(timeStamp)
        const year = date.getFullYear()
        const month = date.getMonth() + 1
        const day = date.getDate()
        return year.toString() + "/" + month.toString() + "/" + day.toString()
    }

    return (
        <Grid container>
            <Grid item xs={12} sm={12}>
                <Paper elevation={3} sx={{ p: 1,}}>
                    <Grid container direction="row" alignItems="center">
                        <Grid item sx={{
                            pl: 2,
                        }} xs={1.5} sm={1.5}>
                            <Grid container direction="column"  spacing={1} alignItems="center">
                                <Grid item>
                                    {createAvatar()}
                                </Grid>
                                <Grid item>
                                    <Typography sx={{
                                        fontSize: 12,
                                        fontWeight: 900,
                                        fontFamily: "Elice Digital Baeum",
                                    }}>{props.postInfo.writer}</Typography>
                                </Grid>
                            </Grid>
                        </Grid>
                        <Grid item sx={{
                            pl: 2,
                        }} xs={9} sm={9}>
                            <Grid container direction="column" spacing={1.5}>
                                <Grid item>
                                    <Typography sx={{
                                        fontSize: 25,
                                        fontWeight: 900,
                                        fontFamily: "Elice Digital Baeum",
                                    }}>
                                        <Link href={"/board/post?postId="+props.postInfo.id} color="inherit" className="title-link" underline="none">
                                            {props.postInfo.title}
                                        </Link>
                                    </Typography>
                                </Grid>
                                <Grid item>
                                    <Grid container direction="row" spacing={2}>
                                        {createTagList()}
                                    </Grid>
                                </Grid>
                            </Grid>
                        </Grid>
                        <Grid item xs={1.5} sm={1.5}>
                            <Grid container direction="column" spacing={0.5}>
                                <Grid item>
                                    <Typography sx={{
                                        pl: 1,
                                        fontSize: 13,
                                        fontWeight: 600,
                                        fontFamily: "Elice Digital Baeum",
                                    }}>
                                        {convertTimeStampToReadableTime(props.postInfo.createAt)}
                                    </Typography>
                                </Grid>
                                <Grid item>
                                    <Typography sx={{
                                        pl: 1,
                                        fontSize: 13,
                                        fontWeight: 600,
                                        fontFamily: "Elice Digital Baeum",
                                    }}>
                                        예상 읽는 시간: {props.postInfo.timeToReadMinute}분
                                    </Typography>
                                </Grid>
                                <Grid item>
                                    <Typography sx={{
                                        pl: 1,
                                        fontSize: 13,
                                        fontWeight: 600,
                                        fontFamily: "Elice Digital Baeum",
                                    }}>
                                        조회수: {props.postInfo.views}
                                    </Typography>
                                </Grid>
                                <Grid item>
                                    <Grid container direction="row" spacing="5" alignItems="center">
                                        <Grid item>
                                            <IconButton type="button" onClick={handleLikeIconClick}>
                                                <FavoriteIcon style={{
                                                    color: (isLike ? 'red' : 'black')
                                                }}/>
                                            </IconButton>
                                        </Grid>
                                        <Grid item>
                                            <Typography sx={{
                                                fontSize: 13,
                                                fontWeight: 600,
                                                fontFamily: "Elice Digital Baeum",
                                            }}>
                                                {likes}
                                            </Typography>
                                        </Grid>
                                    </Grid>
                                </Grid>
                            </Grid>
                        </Grid>
                    </Grid>
                </Paper>
            </Grid>
        </Grid>
    )
}