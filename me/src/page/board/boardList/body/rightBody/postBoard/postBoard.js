import React from 'react'
import Box from "@mui/material/Box";
import Post from "./post";
import {Button, Grid} from "@mui/material";
import {deepOrange} from "@mui/material/colors";
import axios from "axios";


class PostBoard extends React.Component {

    numOfPage = 5
    numOfPost = 7

    search = window.location.search;
    urlSearchParams = new URLSearchParams(this.search)
    paramPageId = this.urlSearchParams.get("page");
    pageId = 1;
    queryOption = this.urlSearchParams.get("queryOption")
    queryString = this.urlSearchParams.get("queryString")
    classificationOption = this.urlSearchParams.get("classificationOption")
    tags = [...this.urlSearchParams.getAll("tags")]

    constructor(props) {
        super(props);
        this.state = {
            postInfo: [],
            totalPostCount: 0,
        }

        if(this.paramPageId != null && !isNaN(Number(this.paramPageId))) {
            this.pageId = parseInt(this.paramPageId)
        }

        if(this.pageId == null || isNaN(Number(this.pageId))) this.pageId = 1
        if(this.queryString === null) this.queryString = ""
        if(this.queryOption === null) this.queryOption = ""
        if(this.classificationOption == null) this.classificationOption = ""
    }


    componentDidMount() {

        let url = ""
        url = "/v2/fetch-board-post-list?row=" + this.pageId.toString()
        url += "&size=" + this.numOfPost.toString()

        if(this.queryOption === ""){
            url += "&option.query_option=Undefined"
        }else {
            url += "&option.query_option="+this.queryOption
        }

        url += "&option.query="+this.queryString

        if(this.classificationOption === "") {
            url += "&option.classification_option=All"
        }else {
            url += "&option.classification_option=" + this.classificationOption
        }

        this.tags.forEach((tag) => {
            url += "&option.tags=" + tag
        })

        axios.get(url).then(
            response => {
                let newPostInfo = []
                response.data.postList.forEach((post) => {
                    newPostInfo = [...newPostInfo, {
                        id: post.id,
                        writer: post.writer,
                        avatarInfo: {
                            avatarImgUrl: "",
                            avatarInitial: 'A',
                            avatarBgColor: deepOrange[500],
                        },
                        title: post.title,
                        content: post.content,
                        likeCnt: post.likeCnt,
                        isNotice: post.isNotice,
                        timeToReadMinute: post.timeToReadMinute,
                        tags: post.tags,
                        views: post.views,
                        createAt: post.createAt,
                    }]
                })
                
                this.setState({
                    postInfo: newPostInfo,
                    totalPostCount: response.data.totalPostCount
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
            this.tags.forEach((tag) => {
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
                        {[...Array(Math.min(this.numOfPage,Math.floor((Math.max(((this.state.totalPostCount-1)/this.numOfPost)+1, 1))))).keys()].map( e => (
                            <Grid item>
                                <Button title={(e+1).toString()} sx={{
                                    fontSize: 18,
                                    fontFamily: "Elice Digital Baeum",
                                    fontWeight: 900,
                                    color: (e+1 === this.pageId ? 'red' : 'black'),
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