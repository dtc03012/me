import React from "react";
import axios from "axios";
import Box from "@mui/material/Box";
import PostEditor from "./postEditor/postEditor";
import {useNavigate} from "react-router-dom";

export default function WritePost(props) {

    const navigate = useNavigate()
    
    const requestPost = (post) => {
        axios.post("/v2/upload-board-post",{
            post: {
                title: post.title,
                password: post.password,
                writer: post.writer,
                isNotice: post.isNotice,
                content: post.content,
                timeToReadMinute: post.timeToReadMinute,
                tags: post.tags,
            }
        }).then(
            navigate("/board/lists")
        ).catch((err) => {
            console.log(err)
        })
    }

    return (
        <Box>
            <PostEditor requestPost={requestPost}/>
        </Box>
    )
}