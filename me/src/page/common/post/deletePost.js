import React, {useEffect, useState} from "react";
import CheckPassword from "./checkPassword";
import axios from "axios";
import {CheckStatus, CheckValidPostId} from "../../../util/checkValidParameter";
import Box from "@mui/material/Box";
import Loading from "../loading/loading";
import Error404 from "../error/error404";
import {useNavigate} from "react-router-dom";

export default function DeletePost(props) {

    const navigate = useNavigate()

    const search = window.location.search
    const urlSearchParams = new URLSearchParams(search)
    let paramPostId = urlSearchParams.get("postId")

    const [postId, setPostId] = useState(0)
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

    const requestCheckPassword = (password, callback) => {
        let url = "/v2/delete-board-post"
        axios.post(url, {
            postId: postId,
            password: password
        }).then( () => {
            navigate("/board/lists")
        }).catch( err => {
            console.log(err)
            callback()
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

    return (
        <Box>
            {createCheckPasswordComponent()}
        </Box>
    )
}