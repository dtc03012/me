import React from "react";
import {Typography} from "@mui/material";
import Box from "@mui/material/Box";

class Logo extends React.Component {
    render() {
        return (
            <Box sx={{p: 2}}>
                <Typography
                    variant="h3"
                    noWrap
                    component="a"
                    href="/"
                    sx={{
                        mr: 2,
                        display: { xs: 'none', md: 'flex' },
                        fontFamily: 'Cinzel',
                        fontWeight: 700,
                        letterSpacing: '.3rem',
                        color: 'inherit',
                        textDecoration: 'none',
                    }}
                >
                    Taehun
                </Typography>
            </Box>
        )
    }
}

export default Logo