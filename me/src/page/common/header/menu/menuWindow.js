import React from "react";
import Box from "@mui/material/Box";
import KeyIcon from '@mui/icons-material/Key';
import {Button, IconButton, Tooltip} from "@mui/material";

function CreateMenuButton(name) {
    return (
        <Button sx = {{
            p: 2,
            fontSize: '18px',
            color: 'text.primary',
            fontFamily: 'BlinkMacSystemFont'
        }}>
            {name}
        </Button>
    )
}

class MenuWindow extends React.Component {
    render() {
        return (
            <Box sx={{ display: { xs: 'none', md: 'flex'},  paddingRight: 10 }}>
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