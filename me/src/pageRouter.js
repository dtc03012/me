import React from 'react';
import Main from "./page/main/main";
import {Route, Routes, BrowserRouter} from 'react-router-dom';
import Board from "./page/board/board";
import PostContent from "./page/common/post/postContent/postContent";
import WritePost from "./page/common/post/writePost";
import DeletePost from "./page/common/post/deletePost";
import UpdatePost from "./page/common/post/updatePost";
import Error404 from "./page/common/error/error404";

export default class PageRouter extends React.Component {
    render() {
        return (
            <BrowserRouter>
                <Routes>
                    <Route exact path="/" element={<Main/>} />
                    <Route exact path="/home" element={<Main/>} />
                    <Route path="/board/lists/*" element={<Board/>} />
                    <Route exact path="/board/write" element={<WritePost/>} />
                    <Route exact path="/board/post/*" element={<PostContent/>} />
                    <Route path="/board/delete/*" element={<DeletePost/>} />
                    <Route path="/board/update/*" element={<UpdatePost/>} />
                    <Route path="/error" element={<Error404/>} />
                    <Route element={<Error404/>} />
                </Routes>
            </BrowserRouter>
        );
    }
}