import React from 'react';
import PostEditor from "../common/postEditor/postEditor";
import Header from "../common/header/header";
import Box from "@mui/material/Box";
import Footer from "../common/footer/footer";


class WritePost extends React.Component {

    render() {
        return (
            <Box minHeight="1000px" sx={{
                backgroundColor: '#f5f5f5',
            }}>
                <Header/>
                <PostEditor/>
                <Footer/>
            </Box>
        )
    }
}

export default WritePost