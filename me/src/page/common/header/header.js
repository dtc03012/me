import React from 'react';
import Logo from "./logo/logo";
import Box from '@mui/material/Box';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import MenuWindow from "./menu/menuWindow";
import {createTheme, ThemeProvider} from "@mui/material";

const theme = createTheme({
    palette: {
        background: {
            paper: '#f5f5f5',
        },
        text: {
            primary: '#212121',
        }
    }
})

function GetMenuStyle(props) {
    if(window.length < 900) {
        return <MenuWindow/>
    }
    return <MenuWindow/>
}

class Header extends React.Component {

    render() {
        return (
            <ThemeProvider theme={theme}>
                <Box component="span">
                    <AppBar position="static" sx={{
                        backgroundColor: 'background.paper',
                        color: 'text.primary',
                    }}>
                        <Toolbar sx={{display: 'flex', justifyContent: 'space-between'}}>
                            <Logo/>
                            <GetMenuStyle />
                        </Toolbar>
                    </AppBar>
                </Box>
            </ThemeProvider>
        )
    }
}

export default Header