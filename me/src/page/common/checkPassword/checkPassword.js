import React, {useEffect, useState} from "react";
import {Button, Grid, Paper, TextField, Typography} from "@mui/material";
import Box from "@mui/material/Box";
import axios from 'axios';

export default function CheckPassword(props) {

    const [password, setPassword] = useState("")
    const [postId, setPostId] = useState(0)

    useEffect(() => {
        setPostId(props.postId)
    },[])

    const handlePasswordChange = (event) => {
        setPassword(event.target.value)
    }

    const handleCheckPasswordButtonClick = (event) => {

        let url = "/v2/check-board-post-password"
        axios.post(url,{
            "postId": postId,
            "password": password,
        }).then( response => {

        }).catch( err => {
            console.log(err)
        })
    }

    return (
        <Grid container minHeight="400px" justifyContent={"center"} sx={{
            mt: 4,
            mb: 4,
        }}>
            <Grid item width={"40%"}>
                <Paper elevation={3}>
                    <Box sx={{
                        backgroundColor: "#00b0ff",
                        height: "40px",
                    }}/>
                    <Grid container direction={"column"} sx={{
                        pl: 3,
                        pr: 3,
                        pb: 3,
                    }}>
                        <Grid item sx={{
                            mt: 3,
                        }}>
                            <Typography sx={{
                                fontSize: 16,
                                fontWeight: 300,
                                fontFamily: "Elice Digital Baeum",
                            }}>
                                비밀번호를 입력하세요.
                            </Typography>
                        </Grid>
                        <Grid item sx={{
                            mt: 2,
                            display: 'flex',
                            justifyContent: 'space-between',
                            alignItems: 'center',
                        }}>
                            <TextField id="password" type={"password"} variant="standard" value={password} inputProps={{
                                style: {
                                    fontFamily: "Elice Digital Baeum",
                                    fontSize: 15,
                                    fontWeight: 500,
                                }
                            }} onChange={handlePasswordChange}/>
                            <Button variant="contained" color="success" sx={{
                                fontSize: 15,
                                fontFamily: "Elice Digital Baeum",
                            }} onClick={handleCheckPasswordButtonClick}>
                                확인
                            </Button>
                        </Grid>
                    </Grid>
                </Paper>
            </Grid>
        </Grid>
    )
}