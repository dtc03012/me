import React from 'react';
import Main from "./page/main/main";
import {Route, Routes, BrowserRouter} from 'react-router-dom'
import Board from "./page/board/board";
import WritePost from "./page/board/writePost";
import ViewPost from "./page/board/viewPost";

export default class PageRouter extends React.Component {
    render() {
        return (
            <BrowserRouter>
                <Routes>
                    <Route exact path="/" element={<Main/>} />
                    <Route exact path="/home" element={<Main/>} />
                    <Route exact path="/board" element={<Board/>} />
                    <Route exact path="/board/write" element={<WritePost/>} />
                    <Route path="/board/post/*" element={<ViewPost/>} />
                </Routes>
            </BrowserRouter>
        );
    }
}