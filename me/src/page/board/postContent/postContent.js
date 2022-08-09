import React from "react";
import Box from "@mui/material/Box";
import {Grid, Typography} from "@mui/material";
import createDOMPurify from 'dompurify'
import axios from "axios";
import {createTheme, ThemeProvider} from "@mui/material/styles";

const DOMPurify = createDOMPurify(window)

const theme = createTheme({
    image: {
        flex: 1,
        height: undefined,
        width: "50px",
    }
})

class PostContent extends React.Component {

    isIncrement = false

    constructor(props) {
        super(props);
        this.state = {
            id: "",
            title: "",
            writer: "",
            content: "",
            likeCnt: 0,
            timeToReadMinute: 0,
            tags: [],
            createAt: "",
        }

        // let href = window.location.href
        // let id = href.substring(href.lastIndexOf('/')+1)
        // let url = "/v2/fetch-board-post?id=" + id
        // axios.put(url).then( response => {
        //     this.isIncrement = true
        // }).catch((err) => {
        //     console.log(err)
        // })
    }

    addLeadingZeros(num, totalLength) {
        return String(num).padStart(totalLength, '0');
    }

    convertTimeStampToDate = (timestamp) => {
        let date = new Date(timestamp)
        let year = date.getUTCFullYear()
        let month = date.getUTCMonth() + 1
        let day = date.getUTCDay()
        let hour = date.getUTCHours()
        let minute = date.getUTCMinutes()
        return String(year) + ". " + String(month) + ". " + String(day) + ". " + this.addLeadingZeros(hour,2) + ":" + this.addLeadingZeros(minute,2)
    }

    reviseContent = (content) => {
        content = content.replaceAll('<img', '<img width="100%" height="100%"')
        return content
    }

    componentDidMount() {
        let href = window.location.href
        let id = href.substring(href.lastIndexOf('/')+1)
        let url = "/v2/fetch-board-post?id=" + id
        axios.get(url).then(
            response => {
                response.data.data.content = this.reviseContent(response.data.data.content)

                this.setState({
                    id: response.data.data.id,
                    title: response.data.data.title,
                    writer: response.data.data.writer,
                    content: response.data.data.content,
                    likeCnt: response.data.data.likeCnt,
                    timeToReadMinute: response.data.data.timeToReadMinute,
                    tags: response.data.data.tags,
                    views: response.data.data.views,
                    createAt: response.data.data.createAt
                })
            }
        ).catch((err) => {
            console.log(err)
        })
    }

    render() {
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
                            {this.state.title}
                        </Typography>
                    </Grid>
                    <Grid item display="flex" minWidth="60%">
                        <Typography sx={{
                            fontSize: 16,
                            fontWeight: 300,
                            fontFamily: "Elice Digital Baeum",
                        }}>
                            {this.state.writer}
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
                            {this.convertTimeStampToDate(this.state.createAt)}
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
                            조회수:  {this.state.views}
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
                            { <div className="article" dangerouslySetInnerHTML={{ __html: DOMPurify.sanitize(this.state.content) }} /> }
                        </div>
                    </Grid>
                </Grid>
            </Box>
        );
    }
}

export default PostContent