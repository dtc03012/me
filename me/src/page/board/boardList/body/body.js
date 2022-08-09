import React from "react";
import Box from "@mui/material/Box";
import LeftBody from "./leftBody/leftBody";
import RightBody from "./rightBody/rightBody";

class Body extends React.Component {
    render() {
        return (
            <Box display="flex">
                <LeftBody/>
                <RightBody/>
            </Box>
        )
    }
}

export default Body