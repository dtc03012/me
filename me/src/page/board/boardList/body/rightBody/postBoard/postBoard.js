import React from 'react'
import Box from "@mui/material/Box";
import Post from "./post";
import {Grid} from "@mui/material";
import {deepOrange} from "@mui/material/colors";


class PostBoard extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            postInfo: [{
                postId: 123,
                userId: "dtc03012",
                avatarInfo: {
                    avatarImgUrl: "",
                    avatarInitial: 'A',
                    avatarBgColor: deepOrange[500],
                },
                title: "오늘의 일기 2022-07-30",
                tags: ['일기', '하루일과'],
                timestamp: Date.now(),
                heartCnt: 30,
                readExpTime: 3,
            }, {
                postId: 124,
                userId: "dtc03012",
                avatarInfo: {
                    avatarImgUrl: "",
                    avatarInitial: 'A',
                    avatarBgColor: deepOrange[500],
                },
                title: "코딩이 즐거운 이유?",
                tags: ['코딩', '일기'],
                timestamp: Date.now(),
                heartCnt: 24,
                readExpTime: 5,
            }],
        }
    }

    render() {
        return (
            <Box sx={{
                p: 4,
            }}>
                <Grid container direction="column" spacing={2}>
                    {this.state.postInfo.map((postInfo) => (
                        <Grid item>
                           <Post postInfo={postInfo}/>
                        </Grid>
                    ))}
                </Grid>
            </Box>
        )
    }
}

export default PostBoard