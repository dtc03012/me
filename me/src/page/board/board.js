import React from 'react';
import Box from "@mui/material/Box";
import Header from "../common/header/header";
import Body from "./body/body";

class Board extends React.Component {
    render() {
        return (
            <Box>
                <Header/>
                <Body/>
            </Box>
        );
    }
}

export default Board;