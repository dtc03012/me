import React, {useEffect, useState} from "react";
import Box from "@mui/material/Box";
import {Grid, Typography, unstable_useId} from "@mui/material";
import createDOMPurify from 'dompurify'
import axios from "axios";
import {createTheme, ThemeProvider} from "@mui/material/styles";
import {setCookie, getCookie} from "../../../util/cookie";
import {v4} from 'uuid';

const DOMPurify = createDOMPurify(window)

const theme = createTheme({
    image: {
        flex: 1,
        height: undefined,
        width: "50px",
    }
})

export default function PostContent(props) {

    const isIncrement = false
    const [id, setId] = useState("")
    const [title, setTitle] = useState("")
    const [writer, setWriter] = useState("")
    const [content, setContent] = useState("")
    const [likeCnt, setLikeCnt] = useState(0)
    const [timeToReadMinute, setTimeToReadMinute] = useState(0)
    const [tags, setTags] = useState([])
    const [views, setViews] = useState(0)
    const [createAt, setCreateAt] = useState("")

    function addLeadingZeros(num, totalLength) {
        return String(num).padStart(totalLength, '0');
    }

    const convertTimeStampToDate = (timestamp) => {
        let date = new Date(timestamp)
        let year = date.getUTCFullYear()
        let month = date.getUTCMonth() + 1
        let day = date.getUTCDay()
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
        let id = href.substring(href.lastIndexOf('/')+1)
        let url = "/v2/fetch-board-post?id=" + id
        axios.get(url).then(
            response => {
                console.log("haha")
                response.data.data.content = reviseContent(response.data.data.content)

                setId(response.data.data.id)
                setTitle(response.data.data.title)
                setWriter(response.data.data.writer)
                setContent(response.data.data.content)
                setLikeCnt(response.data.data.likeCnt)
                setTimeToReadMinute(response.data.data.timeToReadMinute)
                setTags(response.data.data.tags)
                setViews(response.data.data.views)
                setCreateAt(response.data.data.createAt)
            }
        ).catch( err => {
            console.log(err)
        })

        if(getCookie("uuid") === "") {
            let uuid = v4()
            setCookie("uuid", uuid)
        }

        url = "/v2/increment-board-view?id=" + id + "&uuid=" + getCookie("uuid")
        axios.put(url).then(
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
            </Grid>
        </Box>
    );
}