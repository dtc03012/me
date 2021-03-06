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
            appbar: '#f5f5f5',
            belowAppbarBox: '#212121',
        },
        text: {
            primary: '#212121',
        }
    }
})

function GetMenuStyle(props) {
    if(window.length < 900) {
        // TODO (@dtc03012): show the mobile menu when window size is less than 900, or it is mobile case
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
                        backgroundColor: 'background.appbar',
                        color: 'text.primary',
                    }}>
                        <Toolbar sx={{
                            display: 'flex', justifyContent: 'space-between'
                        }}>
                            <Logo/>
                            <GetMenuStyle />
                        </Toolbar>
                        <Box sx={{
                            backgroundColor: 'background.belowAppbarBox',
                            width: '100%',
                            height: '20px',
                        }}>
                        </Box>
                    </AppBar>
                </Box>
            </ThemeProvider>
        )
    }
}

export default Header