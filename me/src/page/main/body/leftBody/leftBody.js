import React from "react";
import Box from "@mui/material/Box";
import Profile from "./profile/profile";

class LeftBody extends React.Component {
    render() {
        return (
            <Box sx={{
                p: 2,
                width: "25%",
            }}>
                <Profile/>
            </Box>
        )
    }
}

export default LeftBody