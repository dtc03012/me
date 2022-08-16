import React, {useEffect, useState} from "react";
import CheckPassword from "./checkPassword";
import axios from "axios";
import {CheckStatus, CheckValidPostId} from "../../../util/checkValidParameter";
import Box from "@mui/material/Box";
import Loading from "../loading/loading";
import Error404 from "../error/error404";
import {useNavigate} from "react-router-dom";
import PostEditor from "./postEditor/postEditor";

export default function UpdatePost(props) {

    const navigate = useNavigate()

    const search = window.location.search
    const urlSearchParams = new URLSearchParams(search)
    let paramPostId = urlSearchParams.get("postId")

    const [post, setPost] = useState({})
    const [postId, setPostId] = useState(0)
    const [checkPassword, setCheckPassword] = useState(false)
    const [isValidPostId, setIsValidPostId] = useState(CheckStatus.LOADING)

    useEffect(() => {
        CheckValidPostId(paramPostId).then( () => {
            setIsValidPostId(CheckStatus.SUCCESS)
            setPostId(parseInt(paramPostId))
        }).catch( err => {
            console.log(err)
            setIsValidPostId(CheckStatus.FAIL)
        })
    },[paramPostId])

    const requestCheckPassword = (password) => {
        let url = "/v2/check-post-password"
        axios.post(url,{
            postId: postId,
            password: password,
        }).then( response => {
            url = "/v2/fetch-board-post?postId=" + postId.toString()
            axios.get(url).then( response => {
                setPost(response.data.post)
                setCheckPassword(true)
            }).catch( err => {
                console.log(err)
                navigate("/error")
            })
        }).catch( err => {
            console.log(err)
            navigate("/error")
        })
    }

    const requestPost = (post) => {
        axios.post("/v2/update-board-post",{
            post: {
                id: postId,
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
        ).catch( err => {
            console.log(err)
        })
    }

    const createCheckPasswordComponent = () => {
        if(isValidPostId === CheckStatus.SUCCESS) {
            return (
                <CheckPassword requestCheckPassword={requestCheckPassword}/>
            )
        }else if(isValidPostId === CheckStatus.LOADING) {
            return (
                <Loading/>
            )
        }
        return (
            <Error404/>
        )
    }

    const createUpdatePostComponent = () => {
        return (
            <PostEditor post={post} requestPost={requestPost}/>
        )
    }

    return (
        <Box>
            { checkPassword ? createUpdatePostComponent() : createCheckPasswordComponent()}
        </Box>
    )
}