import React from 'react';
import Main from "./page/main/Main";
import {Route, Routes, BrowserRouter} from 'react-router-dom'

export default class PageRouter extends React.Component {
    render() {
        return (
            <BrowserRouter>
                <Routes>
                    <Route exact path="/" element={<Main/>} />
                </Routes>
            </BrowserRouter>
        );
    }
}