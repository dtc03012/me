import React from "react";
import { CKEditor } from "@ckeditor/ckeditor5-react";
import ClassicEditor from "@ckeditor/ckeditor5-build-classic";
import Box from "@mui/material/Box";
import {Button, Chip, Grid, Link, TextField} from "@mui/material";
import "./postEditor.css"
import axios from "axios";

function uploadAdapter(loader) {
    return {
        upload: () => {
            return new Promise((resolve, reject) => {
                const body = new FormData();
                loader.file.then((file) => {
                    body.append("file", file);
                    fetch(`/file/upload-file`, {
                        method: "post",
                        body: body,
                        // mode: "no-cors"
                    })
                        .then((res) => {
                            return res.json()
                        })
                        .then((res) => {
                            resolve({
                                default: `/file/get-file/${res.filename}`
                            });
                        })
                        .catch((err) => {
                            reject(err);
                        });
                })
            });
        }
    };
}
function uploadPlugin(editor) {
    editor.plugins.get("FileRepository").createUploadAdapter = (loader) => {
        return uploadAdapter(loader);
    };
}

class PostEditor extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            title: "",
            editorContent: "",
            tags: [],
            inputTag: "",
            likeCnt: 0,
            timeToReadMinute: 0,
        }

        this.handleTagTextFieldChange = this.handleTagTextFieldChange.bind(this)
        this.handleTagTextFieldBlur = this.handleTagTextFieldBlur.bind(this)
        this.handleTimeToReadMinuteTextFieldChange = this.handleTimeToReadMinuteTextFieldChange.bind(this)
        this.handleTagTextFieldKeyDown = this.handleTagTextFieldKeyDown.bind(this)
        this.handleTagTextFieldDelete = this.handleTagTextFieldDelete.bind(this)
    }
    createTagList = () => {
        return (this.state.tags.map( (tag) => (
                <Grid item key={String(tag)}>
                    <Chip label={String(tag)} variant="outlined" onDelete={() => this.handleTagTextFieldDelete(String(tag))} sx={{
                        '& .MuiChip-label': {
                            fontSize: 14,
                            fontFamily: "Elice Digital Baeum",
                        },
                    }} />
                </Grid>
            )
        ))
    }

    handleTitleChange = (event) => {
        this.setState({
            title: event.target.value
        })
    }

    addInputTagToTags = () => {
        if(this.state.inputTag === "") return

        let isDup = false
        this.state.tags.forEach((tag) => {
            if(tag === this.state.inputTag) {
                isDup = true
            }
        })

        if(isDup) {
            this.setState({
                inputTag: "",
            })
            return
        }

        let newTags = [...this.state.tags, this.state.inputTag]

        this.setState({
            tags: newTags,
            inputTag: "",
        })
    }

    handleTagTextFieldBlur = (event) => {
        this.addInputTagToTags()
    }

    handleTagTextFieldChange = (event) => {
        this.setState({
            inputTag: event.target.value
        })
    }

    handleTagTextFieldKeyDown = (event) => {
        if(event.keyCode === 13){
            this.addInputTagToTags()
        }
    }

    handleTagTextFieldDelete = (tag) => {
        this.setState({
            tags: this.state.tags.filter(item => item !== tag)
        })
    }

    uploadPost = () => {
        axios.post("/v2/upload-board-post",{
            data: {
                title: this.state.title,
                writer: 'admin',
                content: this.state.editorContent,
                timeToReadMinute: this.state.timeToReadMinute,
                likeCnt: this.state.likeCnt,
                tags: this.state.tags,
            }
        }).then(
            response => {
                console.log("good")
            }
        ).catch((err) => {
            console.log("bad")
        })
    }

    handleTimeToReadMinuteTextFieldChange = (event) => {
        this.setState({
            timeToReadMinute: event.target.value
        })
    }

    render() {
        return (
            <Box sx={{
                pt: 7,
                pl: 50,
                pr: 50,
                pb: 30,
            }}>
                <Grid container direction="column" spacing={3} >
                    <Grid item>
                        <TextField id="title" variant="standard" fullWidth inputProps={{
                            style: {
                                fontFamily: "Elice Digital Baeum",
                                fontSize: 27,
                                fontWeight: 500,
                            }
                        }} value={this.state.title} onChange={this.handleTitleChange} />
                    </Grid>
                    <Grid item classes="ck-editor">
                        <CKEditor
                            config={{
                                extraPlugins: [uploadPlugin],
                            }}
                            editor={ClassicEditor}
                            onReady={(editor) => {}}
                            onBlur={(event, editor) => {}}
                            onFocus={(event, editor) => {}}
                            onChange={(event, editor) => {
                                this.setState({
                                    editorContent: editor.getData()
                                })
                            }}
                        />
                    </Grid>
                    <Grid item container direction="column" spacing="10">
                        <Grid item>
                            <Box sx={{
                                pl: 1
                            }}>
                                <TextField id="tag" variant="standard" label="태그 입력" value={this.state.inputTag} inputProps={{
                                    style: {
                                        fontFamily: "Elice Digital Baeum",
                                        fontSize: 15,
                                        fontWeight: 500,
                                    }
                                }}
                                           InputLabelProps={{
                                               style: {
                                                   fontFamily: "Elice Digital Baeum",
                                                   fontSize: 15,
                                                   fontWeight: 500,
                                               }
                                           }}
                                           onBlur={this.handleTagTextFieldBlur}
                                           onChange={this.handleTagTextFieldChange}
                                           onKeyDown={this.handleTagTextFieldKeyDown}
                                />
                            </Box>
                        </Grid>
                        <Grid item container direction="row" spacing="5">
                            {this.createTagList()}
                        </Grid>
                    </Grid>
                    <Grid item>
                        <Box sx={{
                            pl: 1
                        }}>
                            <TextField id="timeToReadMinute" variant="standard" label="읽는 데 걸리는 시간(분)" type="number" tyvalue={this.state.timeToReadMinute} inputProps={{
                                style: {
                                    fontFamily: "Elice Digital Baeum",
                                    fontSize: 15,
                                    fontWeight: 500,
                                }
                            }}
                                       InputLabelProps={{
                                           style: {
                                               fontFamily: "Elice Digital Baeum",
                                               fontSize: 15,
                                               fontWeight: 500,
                                           }
                                       }}

                                       onChange={this.handleTimeToReadMinuteTextFieldChange}
                            />
                        </Box>
                    </Grid>
                    <Grid item>
                        <Button variant="contained" color="success" sx={{
                            fontSize: 15,
                            fontFamily: "Elice Digital Baeum",
                        }} onClick={() => this.uploadPost()}>
                            <Link href="/board" underline="none" color="inherit">
                                완료
                            </Link>
                        </Button>
                    </Grid>
                </Grid>
            </Box>
        )
    }
}

export default PostEditor