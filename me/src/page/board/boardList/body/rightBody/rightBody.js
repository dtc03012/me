import React from 'react';
import Box from "@mui/material/Box";
import PostBoard from "./postBoard/postBoard";
import {Button, Grid} from "@mui/material";
import PostOption from "./postBoard/postOption";


class RightBody extends React.Component {
    render() {
        return (
            <Box sx={{
                p: 2,
                width: "55%",
                display: { xs: 'none', sm: 'none', md: 'block' },
            }}>
                <Grid container direction="column">
                    <Grid item>
                        <PostOption/>
                    </Grid>
                    <Grid item>
                        <PostBoard/>
                    </Grid>
                </Grid>
            </Box>
        )
    }
}

export default RightBody