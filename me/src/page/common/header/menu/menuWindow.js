import React from "react";
import Box from "@mui/material/Box";
import KeyIcon from '@mui/icons-material/Key';
import {Button, IconButton, Tooltip, Link} from "@mui/material";


function CreateMenuButton(name) {
    const url = "/" + name.toLowerCase()

    return (
        <Button
            sx = {{
            '&:hover': {
              bgcolor: 'gray',
                color: 'white',
            },
            paddingRight: 2, paddingLeft: 2,
            fontSize: '18px', fontWeight: 570, fontFamily: 'Cantarell',
            color: 'text.primary',
        }}>
            <Link href={url} underline="none" color="inherit">
                {name}
            </Link>
        </Button>
    )
}

class MenuWindow extends React.Component {
    render() {
        return (
            <Box sx={{ display: { md: 'flex'},  paddingRight: 10 }}>
                {CreateMenuButton('Home')}
                {CreateMenuButton('Diary')}
                {CreateMenuButton('Board')}
                {CreateMenuButton('Projects')}
                {CreateMenuButton('Photo')}
                <Tooltip title="Admin">
                    <IconButton>
                        <KeyIcon sx={{ p : 2 }} />
                    </IconButton>
                </Tooltip>
            </Box>
        )
    }
}

export default MenuWindow