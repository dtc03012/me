import React from 'react';
import Box from "@mui/material/Box";
import Header from "../common/header/header";
import Body from "./boardList/body/body";
import PostEditor from "./postEditor/postEditor";

class Board extends React.Component {
    render() {
        return (
            <Box minHeight="1000px" sx={{
                backgroundColor: '#f5f5f5',
            }}>
                <Header/>
                <Body/>
            </Box>
        );
    }
}

export default Board;