import React from 'react';
import Box from "@mui/material/Box";


class RightBody extends React.Component {
    render() {
        return (
            <Box sx={{
                p: 2,
                width: "75%",
                display: { xs: 'none', sm: 'none', md: 'block' },
            }}>
                <h1> rightBody </h1>
            </Box>
        )
    }
}

export default RightBody