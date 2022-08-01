import React from 'react';
import PostEditor from "./postEditor/postEditor";
import {Grid} from "@mui/material";
import Header from "../common/header/header";
import Box from "@mui/material/Box";


class WritePost extends React.Component {

    render() {
        return (
            <Box display="flex" sx={{
                backgroundColor: '#f5f5f5',
            }}>
                <Grid container direction="column">
                    <Grid item>
                        <Header/>
                    </Grid>
                    <Grid item>
                        <PostEditor/>
                    </Grid>
                </Grid>
            </Box>
        )
    }
}

export default WritePost