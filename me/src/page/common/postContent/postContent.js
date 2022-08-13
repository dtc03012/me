import React, {useEffect, useState} from "react";
import Box from "@mui/material/Box";
import {Button, Grid, Typography} from "@mui/material";
import createDOMPurify from 'dompurify'
import axios from "axios";
import {setCookie, getCookie} from "../../../util/cookie";
import {v4} from 'uuid';
import Comment from "./comment";
import ReplyComment from "./replyComment";

const DOMPurify = createDOMPurify(window)

export default function PostContent(props) {

    const numOfPage = 5
    const numOfComment = 8

    const isIncrement = false
    const [title, setTitle] = useState("")
    const [writer, setWriter] = useState("")
    const [content, setContent] = useState("")
    const [likeCnt, setLikeCnt] = useState(0)
    const [timeToReadMinute, setTimeToReadMinute] = useState(0)
    const [tags, setTags] = useState([])
    const [views, setViews] = useState(0)
    const [createAt, setCreateAt] = useState("")
    const [commentList, setCommentList] = useState([])
    const [totalCommentCount, setTotalCommentCount] = useState(0)

    const search = window.location.search
    const urlSearchParams = new URLSearchParams(search)
    let paramPostId = urlSearchParams.get("postId");
    let paramCommentPageId= urlSearchParams.get("commentPage");

    let commentPageId = 1
    if(paramCommentPageId != null && !isNaN(Number(paramCommentPageId))) {
        commentPageId = parseInt(paramCommentPageId)
    }

    let postId = 1
    if(paramPostId != null && !isNaN(Number(paramPostId))) {
        postId = parseInt(paramPostId)
    }

    function addLeadingZeros(num, totalLength) {
        return String(num).padStart(totalLength, '0');
    }

    const convertTimeStampToDate = (timestamp) => {
        if(timestamp === "") {
            return "Loading..."
        }

        let date = new Date(timestamp)
        let year = date.getUTCFullYear()
        let month = date.getUTCMonth() + 1
        let day = date.getUTCDate()
        let hour = date.getUTCHours()
        let minute = date.getUTCMinutes()
        return String(year) + ". " + String(month) + ". " + String(day) + ". " + addLeadingZeros(hour,2) + ":" + addLeadingZeros(minute,2)
    }

    const reviseContent = (content) => {
        content = content.replaceAll('<img', '<img width="100%" height="100%"')
        return content
    }

    useEffect(() => {
        let href = window.location.href
        let url = "/v2/fetch-board-post?id=" + postId

        axios.get(url).then(
            response => {
                response.data.post.content = reviseContent(response.data.post.content)

                setTitle(response.data.post.title)
                setWriter(response.data.post.writer)
                setContent(response.data.post.content)
                setLikeCnt(response.data.post.likeCnt)
                setTimeToReadMinute(response.data.post.timeToReadMinute)
                setTags(response.data.post.tags)
                setViews(response.data.post.views)
                setCreateAt(response.data.post.createAt)
            }
        ).catch( err => {
            console.log(err)
        })

        if(getCookie("uuid") === "") {
            let uuid = v4()
            setCookie("uuid", uuid)
        }

        url = "/v2/increment-board-view?id=" + postId + "&uuid=" + getCookie("uuid")
        axios.put(url).then(
        ).catch( err => {
            console.log(err)
        })

        url = "/v2/fetch-board-comment-list?post_id=" + postId + "&row=" + commentPageId.toString() + "&size=" + numOfComment.toString()
        axios.get(url).then(
            response => {
                if(response.data.commentList) {
                    setCommentList(response.data.commentList.reverse())
                }
                setTotalCommentCount(response.data.totalCommentCount)
            }
        ).catch( err => {
            console.log(err)
        })
    }, [])

    return (
        <Box>
            <Grid container direction="column" justifyContent="center" alignItems="center"spacing={1} sx={{
                p: 4,
            }}>
                <Grid item minWidth="60%">
                    <Typography sx={{
                        fontSize: 45,
                        fontWeight: 900,
                        fontFamily: "Elice Digital Baeum",
                    }}>
                        {title}
                    </Typography>
                </Grid>
                <Grid item display="flex" minWidth="60%">
                    <Typography sx={{
                        fontSize: 16,
                        fontWeight: 300,
                        fontFamily: "Elice Digital Baeum",
                    }}>
                        {writer}
                    </Typography>
                    <Typography sx={{
                        fontSize: 16,
                        fontWeight: 300,
                        fontFamily: "Elice Digital Baeum",
                        color: `#d3d3d3`,
                        ml: 2,
                        mr: 2,
                    }}>
                        |
                    </Typography>
                    <Typography sx={{
                        fontSize: 16,
                        fontWeight: 300,
                        fontFamily: "Elice Digital Baeum",
                    }}>
                        {convertTimeStampToDate(createAt)}
                    </Typography>
                    <Typography sx={{
                        fontSize: 16,
                        fontWeight: 300,
                        fontFamily: "Elice Digital Baeum",
                        color: `#d3d3d3`,
                        ml: 2,
                        mr: 2,
                    }}>
                        |
                    </Typography>
                    <Typography sx={{
                        fontSize: 16,
                        fontWeight: 300,
                        fontFamily: "Elice Digital Baeum",
                    }}>
                        조회수:  {views}
                    </Typography>
                </Grid>
                <Grid item minWidth="60%">
                    <Box sx={{
                        backgroundColor: `#d3d3d3`,
                        width: '100%',
                        height: '3px',
                    }}>
                    </Box>
                </Grid>
                <Grid item width="60%" minHeight="600px" sx={{
                    mt: 3,
                }}>
                    <div>
                        { <div className="article" dangerouslySetInnerHTML={{ __html: DOMPurify.sanitize(content) }} /> }
                    </div>
                </Grid>
                <Grid item minWidth="60%" sx={{
                    mt: 3,
                }}>
                    <Box sx={{
                        backgroundColor: `#d3d3d3`,
                        width: '100%',
                        height: '3px',
                    }}>
                    </Box>
                </Grid>
                <Grid item minWidth="60%" sx={{
                    mt: 3,

                }}>
                    <Typography sx={{
                        fontSize: 30,
                        fontWeight: 600,
                        fontFamily: "Elice Digital Baeum",
                    }}>
                        전체 댓글 ({totalCommentCount})
                    </Typography>
                </Grid>
                <Grid item minWidth="60%">
                    {commentList.map((comment) => (
                        <Grid item>
                            <Comment commentInfo={comment}/>
                        </Grid>
                    ))}
                </Grid>
                <Grid container item justifyContent={"center"} sx={{
                    mt: 2,
                }}>
                    {[...Array(Math.min(numOfPage,Math.floor((Math.max(((totalCommentCount-1)/numOfComment)+1, 1))))).keys()].map( e => (
                        <Grid item>
                            <Button title={(e+1).toString()} sx={{
                                fontSize: 18,
                                fontFamily: "Elice Digital Baeum",
                                fontWeight: 900,
                                color: (e+1 === commentPageId ? 'red' : 'black'),
                            }} href={"/board/post?postId="+postId+"&commentPage="+(e+1).toString()}>
                                {(e+1).toString()}
                            </Button>
                        </Grid>
                    ))}
                </Grid>
                <Grid item minWidth="60%"  sx={{
                    mt: 3,
                }}>
                    <ReplyComment postId={postId} parerntCid={0}/>
                </Grid>
            </Grid>
        </Box>
    );
}