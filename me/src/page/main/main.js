import React from 'react';
import Header from "../common/header/header";
import Box from "@mui/material/Box";
import Body from "./body/body";

class Main extends React.Component {
    render() {
        return (
            <Box>
                <Header/>
                <Body/>
            </Box>
        )
    }
}

export default Main;