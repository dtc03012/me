import React from 'react';
import Main from "./page/main/main";
import {Route, Routes, BrowserRouter} from 'react-router-dom'
import Board from "./page/board/board";
import WritePost from "./page/board/writePost";

export default class PageRouter extends React.Component {
    render() {
        return (
            <BrowserRouter>
                <Routes>
                    <Route exact path="/" element={<Main/>} />
                    <Route exact path="/home" element={<Main/>} />
                    <Route exact path="/board" element={<Board/>} />
                    <Route exact path="/board/write" element={<WritePost/>} />
                </Routes>
            </BrowserRouter>
        );
    }
}