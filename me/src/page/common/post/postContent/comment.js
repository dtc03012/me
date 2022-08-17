import React from "react"
import {Avatar, Grid, Paper, Typography} from "@mui/material";

export default function Comment(props) {

    const createAvatar = () => {
        return <Avatar sx={{bgcolor: '#82b1ff', width: 50, height: 50 }}> {props.commentInfo.writer.charAt(0)} </Avatar>
    }

    function addLeadingZeros(num, totalLength) {
        return String(num).padStart(totalLength, '0');
    }

    const convertTimeStampToDate = (timestamp) => {
        if(timestamp === "") {
            return "Loading..."
        }

        let date = new Date(timestamp)
        let year = date.getFullYear()
        let month = date.getMonth() + 1
        let day = date.getDate()
        let hour = date.getHours()
        let minute = date.getMinutes()
        return String(year) + ". " + String(month) + ". " + String(day) + ". " + addLeadingZeros(hour,2) + ":" + addLeadingZeros(minute,2)
    }

    return (
        <Paper elevation={3} sx={{
            mt: 2,
        }}>
            <Grid container spacing={0.5} direction="row" sx={{
                p: 2,
            }}>
                <Grid item xs={1} sm={1}>
                    {createAvatar()}
                </Grid>
                <Grid item xs={10} sm={10}>
                    <Grid container spacing={0.5} direction="column">
                        <Grid item sx={{
                            display: 'flex'
                        }}>
                            <Typography sx={{
                                fontSize: 12,
                                fontWeight: 700,
                                color: "#616161",
                                fontFamily: "Elice Digital Baeum",
                            }}>{props.commentInfo.writer}
                            </Typography>
                            <Typography sx={{
                                fontSize: 12,
                                fontWeight: 600,
                                fontFamily: "Elice Digital Baeum",
                                color: `#d3d3d3`,
                                ml: 1,
                                mr: 1,
                            }}>
                                |
                            </Typography>
                            <Typography sx={{
                                fontSize: 12,
                                fontWeight: 700,
                                color: "#616161",
                                fontFamily: "Elice Digital Baeum",
                            }}>
                                {convertTimeStampToDate(props.commentInfo.createAt)}
                            </Typography>
                        </Grid>
                        <Grid item>
                            <Typography sx={{
                                fontSize: 14,
                                fontWeight: 900,
                                fontFamily: "Elice Digital Baeum",
                            }}>{props.commentInfo.comment}</Typography>
                        </Grid>
                    </Grid>
                </Grid>
            </Grid>
        </Paper>
    )
}