import React from "react";
import Box from "@mui/material/Box";
import Introduction from "./introduction/introduction";

class RightBody extends React.Component {
    render() {
        return (
            <Box sx={{
                p: 2,
            }}>
                <Introduction/>
            </Box>
        )
    }
}

export default RightBody