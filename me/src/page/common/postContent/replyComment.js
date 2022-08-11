import React, {useState} from 'react';
import {Button, Grid, Link, TextField} from "@mui/material";
import axios from "axios";

export default function ReplyComment(props) {

    const [writer, setWriter] = useState("")
    const [password, setPassword] = useState("")
    const [commentContent, setCommentContent] = useState("")

    const handleWriterChange = (event) => {
        setWriter(event.currentTarget.value)
    }

    const handlePasswordChange = (event) => {
        setPassword(event.currentTarget.value)
    }

    const handleCommentContentChange = (event) => {
        setCommentContent(event.currentTarget.value)
    }

    const handleReplyButtonClick = (event) => {
        let url = "/v2/leave-board-comment"
        axios.post(url,{
            comment: {
                postId: props.postId,
                isExist: true,
                writer: writer,
                password: password,
                comment: commentContent,
                likeCnt: 0,
            }
        }).then( response => {
        }).catch( err => {
            console.log(err)
        })
    }

    return (
        <Grid container spacing={2} direction="column">
            <Grid item>
                <Grid container direction="row">
                    <Grid item>
                        <TextField
                            id="writer"
                            label="이름"
                            defaultValue=""
                            InputProps={{
                                style: {
                                    fontSize: 14,
                                    fontFamily: "Elice Digital Baeum",
                                }
                            }}
                            InputLabelProps={{
                                style: {
                                    fontSize: 15,
                                    fontFamily: "Elice Digital Baeum",
                                },
                            }}
                            onChange={handleWriterChange}
                            variant="standard"
                        />
                    </Grid>
                    <Grid item sx={{
                        ml: 3,
                    }}>
                        <TextField
                            pass
                            id="password"
                            label="비밀번호"
                            defaultValue=""
                            type="password"
                            InputProps={{
                                style: {
                                    fontSize: 14,
                                    fontFamily: "Elice Digital Baeum",
                                }
                            }}
                            InputLabelProps={{
                                style: {
                                    fontSize: 15,
                                    fontFamily: "Elice Digital Baeum",
                                },
                            }}
                            onChange={handlePasswordChange}
                            variant="standard"
                        />
                    </Grid>
                </Grid>
                <Grid item sx={{
                    mt: 3,
                }}>
                    <TextField
                        fullWidth
                        id="writer"
                        multiline
                        defaultValue=""
                        InputProps={{
                            style: {
                                fontSize: 15,
                                fontFamily: "Elice Digital Baeum",
                            },
                        }}
                        onChange={handleCommentContentChange}
                    />
                </Grid>
                <Grid item sx={{
                    mt: 3,
                }}>
                    <Link href={"/board/post?postId="+ props.postId.toString()} underline="none" color="inherit">
                        <Button variant="contained" color="success" sx={{
                            fontSize: 15,
                            fontFamily: "Elice Digital Baeum",
                        }} onClick={handleReplyButtonClick}>
                            댓글 작성
                        </Button>
                    </Link>
                </Grid>
            </Grid>
        </Grid>
    )
}