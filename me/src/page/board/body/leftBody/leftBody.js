import React from 'react';
import Box from "@mui/material/Box";
import SearchOption from "./searchOption/searchOption";


class LeftBody extends React.Component {
    render() {
        return (
            <Box sx={{
                p: 2,
                width: "25%",
                display: { xs: 'none', sm: 'none', md: 'block' },
            }}>
                <SearchOption/>
            </Box>
        )
    }
}

export default LeftBody