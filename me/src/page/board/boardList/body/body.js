import React from "react";
import Box from "@mui/material/Box";
import LeftBody from "./leftBody/leftBody";
import RightBody from "./rightBody/rightBody";

class Body extends React.Component {
    render() {
        return (
            <Box display="flex" sx={{
                backgroundColor: '#f5f5f5',
            }}>
                <LeftBody/>
                <RightBody/>
            </Box>
        )
    }
}

export default Body