import React from 'react';
import Box from "@mui/material/Box";
import PostContent from "../common/postContent/postContent";
import Header from "../common/header/header";
import Footer from "../common/footer/footer";

class ViewPost extends React.Component {

    render() {
        return (
            <Box minHeight="1000px" sx={{
                backgroundColor: '#f5f5f5',
            }}>
                <Header/>
                <PostContent/>
                <Footer/>
            </Box>
        )
    }
}

export default ViewPost