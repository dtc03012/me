import React from 'react';
import Box from "@mui/material/Box";
import PostContent from "./postContent/postContent";
import Header from "../common/header/header";
import {Grid} from "@mui/material";

class ViewPost extends React.Component {

    render() {
        return (
            <Box minHeight="1000px" sx={{
                backgroundColor: '#f5f5f5',
            }}>
                <Grid container direction="column">
                    <Grid item>
                        <Header/>
                    </Grid>
                    <Grid item>
                        <PostContent/>
                    </Grid>
                </Grid>
            </Box>
        )
    }
}

export default ViewPost