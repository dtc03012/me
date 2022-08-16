import React from 'react';
import Header from "../common/header/header";
import Box from "@mui/material/Box";
import Body from "./body/body";
import Footer from "../common/footer/footer";

class Main extends React.Component {
    render() {
        return (
            <Box>
                <Body/>
            </Box>
        )
    }
}

export default Main;