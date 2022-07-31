import React from "react";
import { CKEditor } from "@ckeditor/ckeditor5-react";
import ClassicEditor from "@ckeditor/ckeditor5-build-classic";
import axios from "axios";
import {Buffer} from 'buffer';

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

    render() {
        return (
            <CKEditor
                config={{
                    extraPlugins: [uploadPlugin]
                }}
                editor={ClassicEditor}
                onReady={(editor) => {}}
                onBlur={(event, editor) => {}}
                onFocus={(event, editor) => {}}
                onChange={(event, editor) => {}}
                {...this.props}
            />
        )
    }
}

export default PostEditor