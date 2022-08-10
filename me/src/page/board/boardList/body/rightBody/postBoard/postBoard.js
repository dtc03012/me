import React from 'react'
import Box from "@mui/material/Box";
import Post from "./post";
import {Button, Grid, Link} from "@mui/material";
import {deepOrange} from "@mui/material/colors";
import axios from "axios";


class PostBoard extends React.Component {

    numOfPage = 5

    search = window.location.search;
    urlSearchParams = new URLSearchParams(this.search)
    pageId = this.urlSearchParams.get("page");
    pageId = this.urlSearchParams.get("page");
    queryOption = this.urlSearchParams.get("queryOption")
    queryString = this.urlSearchParams.get("queryString")
    tags = [...this.urlSearchParams.getAll("tags")]

    constructor(props) {
        super(props);
        this.state = {
            postInfo: [],
        }

        if(this.pageId === null || isNaN(Number(this.pageId))) this.pageId = 1
        if(this.queryString === null) this.queryString = ""
        if(this.queryOption === null) this.queryOption = ""
    }


    componentDidMount() {

        let url = ""

        console.log(this.queryString)
        console.log(this.queryOption)
        console.log(this.tags)

        if(this.queryOption === "") {
            url = "/v2/fetch-board-post-list?row=" + this.pageId.toString() + "&size=7"
        }else {
            url = "/v2/search-board-post-list?row=" + this.pageId.toString()
            url += "&size=7"
            url += "&option.search_option="+this.queryOption
            url += "&option.search_query="+this.queryString
            this.tags.map((tag) => {
                url += "&option.tags=" + tag
            })
        }

        console.log(url)

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

    createPageHref = (page) => {
        let url = "board?page=" + page.toString()
        if(this.queryOption !== "") {
            url += "&queryOption=" + this.queryOption
            url += "&queryString=" + this.queryString
            this.tags.map((tag) => {
                url += "&tags=" + tag
            })
        }

        return url
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
                    <Grid container item justifyContent={"center"}>
                        {[...Array(this.numOfPage).keys()].map( e => (
                            <Grid item>
                                <Button title={(e+1).toString()} sx={{
                                    fontSize: 18,
                                    fontFamily: "Elice Digital Baeum",
                                    fontWeight: 900,
                                    color: "black",
                                }} href={this.createPageHref(e+1)}>
                                    {(e+1).toString()}
                                </Button>
                            </Grid>
                        ))}
                    </Grid>
                </Grid>
            </Box>
        )
    }
}

export default PostBoard