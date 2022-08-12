import React from "react";
import Box from "@mui/material/Box";
import {Button, Link} from "@mui/material";
import Admin from "./admin";


function CreateMenuButton(name) {
    const url = "/" + name.toLowerCase()

    return (
        <Button
            href={url}
            sx = {{
            '&:hover': {
              bgcolor: 'gray',
                color: 'white',
            },
            paddingRight: 2, paddingLeft: 2,
            fontSize: '18px', fontWeight: 860, fontFamily: 'Cinzel',
            color: 'text.primary',
        }}>
            {name}
        </Button>
    )
}

class MenuWindow extends React.Component {
    render() {
        return (
            <Box sx={{ display: { md: 'flex'},  paddingRight: 10 }}>
                {CreateMenuButton('Home')}
                {/*{CreateMenuButton('Diary')}*/}
                {CreateMenuButton('Board')}
                {/*{CreateMenuButton('Projects')}*/}
                {/*{CreateMenuButton('Photo')}*/}
                <Admin/>
            </Box>
        )
    }
}

export default MenuWindow