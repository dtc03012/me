import React from "react";
import Box from "@mui/material/Box";
import {Button, Link} from "@mui/material";

class PostOption extends React.Component {

    search = window.location.search;
    urlSearchParams = new URLSearchParams(this.search)
    classificationOption = this.urlSearchParams.get("classificationOption")

    constructor(props) {
        super(props);

        if(this.classificationOption == null) this.classificationOption = "All"
    }

    render() {
        return (
            <Box sx={{
                display: 'flex',
                pl: 4,
                pr: 4,
                pt: 4,
            }} justifyContent="space-between">
                <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
                    <Button variant={this.classificationOption === "All" ? "contained" : ""} href={"/board/lists?classificationOption=All"} sx={{
                        fontSize: 18,
                        fontWeight: 900,
                        fontFamily: "Elice Digital Baeum",
                    }}>전체</Button>
                    <Button variant={this.classificationOption === "Popular" ? "contained" : ""} href={"/board/lists?classificationOption=Popular"} sx={{
                        fontSize: 18,
                        fontWeight: 900,
                        fontFamily: "Elice Digital Baeum",
                    }}>인기글</Button>
                    <Button variant={this.classificationOption === "Notice" ? "contained" : ""} href={"/board/lists?classificationOption=Notice"} sx={{
                        fontSize: 18,
                        fontWeight: 900,
                        fontFamily: "Elice Digital Baeum",
                    }}>공지</Button>
                </Box>
                <Link href="/board/write" underline="none" color="inherit">
                    <Button variant="contained" color="success" sx={{
                        fontSize: 15,
                        fontFamily: "Elice Digital Baeum",
                    }}>
                        글쓰기
                    </Button>
                </Link>
            </Box>
        )
    }
}

export default PostOption