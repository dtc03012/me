import React from 'react';
import Box from "@mui/material/Box";
import PostBoard from "./postBoard/postBoard";


class RightBody extends React.Component {
    render() {
        return (
            <Box sx={{
                p: 2,
                width: "60%",
                display: { xs: 'none', sm: 'none', md: 'block' },
            }}>
                <PostBoard/>
            </Box>
        )
    }
}

export default RightBody