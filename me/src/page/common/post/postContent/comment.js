import React, {useEffect, useState} from "react"
import {Avatar, Button, Grid, IconButton, Paper, Popper, TextField, Typography} from "@mui/material";
import HighlightOffIcon from '@mui/icons-material/HighlightOff';
import Box from "@mui/material/Box";
import axios from 'axios';
import {useNavigate} from "react-router-dom";

export default function Comment(props) {

    const navigate = useNavigate()

    const [comment, setComment] = useState(props.commentInfo)
    const [password, setPassword] = useState("")
    const [deleteError, setDeleteError] = useState(false)
    const [deleteButtonAnchorEl, setDeleteButtonAnchorEl] = useState(null);
    const [deleteButtonPopperOpen, setDeleteButtonPopperOpen] = useState(false);

    const handlePasswordChange = (event) => {
        console.log(event.target.value)
        setPassword(event.target.value)
    }
    const handleDeleteCommentButtonClick = (event) => {
        setDeleteButtonAnchorEl(event.currentTarget);
        setDeleteButtonPopperOpen(true)
    };

    const handleRequestDeleteCommentButtonClick = (event) => {
        let url = "/v2/delete-board-comment"
        axios.post(url,{
            commentId: props.commentInfo.id,
            password: password,
        }).then(response => {
            window.location.reload();
        }).catch( err => {
            setDeleteError(true)
            console.log(err)
        })
    }

    const handleCloseDeleteCommentButtonClick = (event) => {
        setDeleteButtonPopperOpen(false)
        setDeleteButtonAnchorEl(null)
        setDeleteError(false)
    };

    let open = Boolean(deleteButtonAnchorEl);
    let id = open ? 'simple-popper' : undefined;

    const createAvatar = () => {
        return <Avatar sx={{bgcolor: '#82b1ff', width: 50, height: 50 }}> {comment.writer.charAt(0)} </Avatar>
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


    const createComment = () => {
        if(comment.isExist) {
            return (
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
                                }}>{comment.writer}
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
                                    {convertTimeStampToDate(comment.createAt)}
                                </Typography>
                            </Grid>
                            <Grid item>
                                <Typography sx={{
                                    fontSize: 14,
                                    fontWeight: 900,
                                    fontFamily: "Elice Digital Baeum",
                                }}>{comment.comment}</Typography>
                            </Grid>
                        </Grid>
                    </Grid>
                    <Grid item>
                        <IconButton type="button" onClick={handleDeleteCommentButtonClick}>
                            <HighlightOffIcon sx={{
                                fontSize: "40px",
                                color: 'red',
                            }}/>
                        </IconButton>
                        <Popper open={deleteButtonPopperOpen} anchorEl={deleteButtonAnchorEl}>
                            <Paper elevation={3} sx={{
                                border: 2,
                                p: 1,
                            }}>
                                <Typography sx={{
                                    fontSize: 14,
                                    fontWeight: 500,
                                    fontFamily: "Elice Digital Baeum"
                                }}>
                                    비밀번호 입력
                                </Typography>
                                <Box alignItems={"center"} sx={{
                                    display: 'flex'
                                }}>
                                    <TextField type="password" sx={{
                                        mt: 1,
                                    }} inputProps={{style: {fontSize: 20, padding: 4, width: "70px"}}} onChange={handlePasswordChange}>

                                    </TextField>
                                    <Button variant="contained" color="success" sx={{
                                        fontSize: 13,
                                        ml: 2,
                                        fontFamily: "Elice Digital Baeum",
                                    }} style={{
                                        maxHeight: '30px',
                                    }} onClick={handleRequestDeleteCommentButtonClick}>
                                        확인
                                    </Button>
                                    <Button variant="contained" color="error" sx={{
                                        fontSize: 13,
                                        ml: 2,
                                        fontFamily: "Elice Digital Baeum",
                                    }} style={{
                                        maxHeight: '30px',
                                    }} onClick={handleCloseDeleteCommentButtonClick}>
                                        취소
                                    </Button>
                                </Box>
                                {deleteError ?
                                    <Typography sx={{
                                        mt: 1,
                                        color: 'red',
                                        fontSize: 14,
                                        fontWeight: 500,
                                        fontFamily: "Elice Digital Baeum"
                                    }}> 비밀번호가 틀립니다 </Typography> : undefined
                                }
                            </Paper>
                        </Popper>
                    </Grid>
                </Grid>
            )
        }
        return (
            <Typography sx={{
                p: 3,
                fontSize: 15,
                fontFamily: "Elice Digital Baeum",
                fontWeight: 900,
            }}>
                삭제된 댓글입니다.
            </Typography>
        )
    }

    return (
        <Paper elevation={3} sx={{
            mt: 2,
        }}>
            {createComment()}
        </Paper>
    )
}