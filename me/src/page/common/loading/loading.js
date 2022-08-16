import React from 'react';
import {Grid, Typography} from "@mui/material";


export default function Loading() {

    return (
        <Grid container minHeight={"800px"} justifyContent={"center"}>
            <Typography sx={{
                mt: 2,
                fontSize: "40px",
                fontWeight: '900',
                fontFamily: "Elice Digital Baeum",
            }}>
                Loading....
            </Typography>
        </Grid>
    )
}