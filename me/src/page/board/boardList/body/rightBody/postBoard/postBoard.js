import React from 'react'
import Box from "@mui/material/Box";
import Post from "./post";
import {Grid} from "@mui/material";
import {deepOrange} from "@mui/material/colors";
import axios from "axios";


class PostBoard extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            postInfo: [],
        }
    }

    componentDidMount() {
        let url = "/v2/fetch-board-post-list?row=1&size=6"
        axios.get(url).then(
            response => {
                let newPostInfo = []
                response.data.data.forEach((data) => {
                    newPostInfo = [...newPostInfo, {
                        id: data.id,
                        writer: data.writer,
                        avatarInfo: {
                            avatarImgUrl: "",
                            avatarInitial: 'A',
                            avatarBgColor: deepOrange[500],
                        },
                        title: data.title,
                        content: data.content,
                        likeCnt: data.likeCnt,
                        timeToReadMinute: data.timeToReadMinute,
                        tags: data.tags,
                        views: data.views,
                        createAt: data.createAt,
                    }]
                })

                this.setState({
                    postInfo: newPostInfo
                })
            }
        ).catch((err) => {
            console.log("bad")
        })
    }

    render() {
        return (
            <Box sx={{
                p: 4,
            }}>
                <Grid container spacing={2}>
                    {this.state.postInfo.map((postInfo) => (
                        <Grid item xs={12} sm={12} >
                           <Post postInfo={postInfo}/>
                        </Grid>
                    ))}
                </Grid>
            </Box>
        )
    }
}

export default PostBoard