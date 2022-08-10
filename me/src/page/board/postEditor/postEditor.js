import React, {useState} from "react";
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

export default function PostEditor(props) {

    const [title, setTitle] = useState("");
    const [editorContent,setEditorContent] = useState("")
    const [tags, setTags] = useState([])
    const [inputTag, setInputTag] = useState("")
    const [likeCnt, setLikeCnt] = useState(0)
    const [timeToReadMinute, setTimeToReadMinute] = useState(0)

    const createTagList = () => {
        return (tags.map( (tag) => (
                <Grid item key={String(tag)}>
                    <Chip label={String(tag)} variant="outlined" onDelete={() => handleTagTextFieldDelete(String(tag))} sx={{
                        '& .MuiChip-label': {
                            fontSize: 14,
                            fontFamily: "Elice Digital Baeum",
                        },
                    }} />
                </Grid>
            )
        ))
    }

    const handleTitleChange = (event) => {
        setTitle(event.target.value)
    }

    const addInputTagToTags = () => {
        if(inputTag === "") return

        let isDup = false
        tags.forEach((tag) => {
            if(tag === inputTag) {
                isDup = true
            }
        })

        if(isDup) {
            setInputTag("")
            return
        }

        let newTags = [...tags, inputTag]

        setTags(newTags)
        setInputTag("")
    }

    const handleTagTextFieldBlur = (event) => {
        addInputTagToTags()
    }

    const handleTagTextFieldChange = (event) => {
        setInputTag(event.target.value)
    }

    const handleTagTextFieldKeyDown = (event) => {
        if(event.keyCode === 13){
            addInputTagToTags()
        }
    }

    const handleTagTextFieldDelete = (tag) => {
        setTags(tags.filter(item => item !== tag))
    }

    const uploadPost = () => {
        axios.post("/v2/upload-board-post",{
            data: {
                title: title,
                writer: 'admin',
                content: editorContent,
                timeToReadMinute: timeToReadMinute,
                likeCnt: likeCnt,
                tags: tags,
            }
        }).then(
            response => {
                console.log("good")
            }
        ).catch((err) => {
            console.log("bad")
        })
    }

    const handleTimeToReadMinuteTextFieldChange = (event) => {
        setTimeToReadMinute(event.target.value)
    }

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
                    }} value={title} onChange={handleTitleChange} />
                </Grid>
                <Grid item className="ck-editor__editable_inline" >
                    <CKEditor
                        config={{
                            extraPlugins: [uploadPlugin],
                        }}
                        editor={ClassicEditor}
                        onBlur={(event, editor) => {}}
                        onFocus={(event, editor) => {}}
                        onChange={(event, editor) => {
                            setEditorContent(editor.getData())
                        }}
                    />
                </Grid>
                <Grid item container direction="column" spacing="10">
                    <Grid item>
                        <Box sx={{
                            pl: 1
                        }}>
                            <TextField id="tag" variant="standard" label="태그 입력" value={inputTag} inputProps={{
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
                                       onBlur={handleTagTextFieldBlur}
                                       onChange={handleTagTextFieldChange}
                                       onKeyDown={handleTagTextFieldKeyDown}
                            />
                        </Box>
                    </Grid>
                    <Grid item container direction="row" spacing="5">
                        {createTagList()}
                    </Grid>
                </Grid>
                <Grid item>
                    <Box sx={{
                        pl: 1
                    }}>
                        <TextField id="timeToReadMinute" variant="standard" label="읽는 데 걸리는 시간(분)" type="number" tyvalue={timeToReadMinute} inputProps={{
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

                                   onChange={handleTimeToReadMinuteTextFieldChange}
                        />
                    </Box>
                </Grid>
                <Grid item>
                    <Button variant="contained" color="success" sx={{
                        fontSize: 15,
                        fontFamily: "Elice Digital Baeum",
                    }} onClick={() => uploadPost()}>
                        <Link href="/board" underline="none" color="inherit">
                            완료
                        </Link>
                    </Button>
                </Grid>
            </Grid>
        </Box>
    )
}