import React from 'react';
import Box from "@mui/material/Box";
import Header from "../common/header/header";
import Body from "./boardList/body/body";
import Footer from "../common/footer/footer";

class Board extends React.Component {
    render() {
        return (
            <Box sx={{
                backgroundColor: '#f5f5f5',
            }}>
                <Header/>
                <Body/>
                <Footer/>
            </Box>
        );
    }
}

export default Board;