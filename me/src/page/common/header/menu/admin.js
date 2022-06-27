import React from 'react';
import axios from "axios";
import {Button, Fade, IconButton, Paper, Popper, Tooltip, Typography} from "@mui/material";
import KeyIcon from "@mui/icons-material/Key";
import CustomInput from "../../../components/customInput";
import Box from "@mui/material/Box";

function CheckAdminLogin() {
    if(localStorage.getItem("admin")) {
        axios.get('http://localhost:9000/v2/check-admin').then(
            response => {
                if(response.data == localStorage.getItem("admin")) {
                    return true
                }
            }
        )
    }
    return false
}


class Admin extends React.Component {

    constructor(props) {
        super(props);
        // Don't call this.setState() here!
        this.state = {
            // isAdmin: CheckAdminLogin(),
            isOpen: false,
            anchorEl: null,
            loginId: "",
        };
        this.handleClick = this.handleClick.bind(this);
        this.handleChange = this.handleChange.bind(this);
    }

    handleClick = (event) => {
        this.setState({
            anchorEl: event.currentTarget,
            isOpen: !this.state.isOpen,
            loginId: "",
        })
    }

    handleChange = (event) => {
        this.setState({
            loginId: event.currentTarget.value
        })
    }

    render() {
        this.canBeOpen = this.state.isOpen && Boolean(this.state.anchorEl);
        this.id = this.canBeOpen ? 'transition-popper' : undefined;
        console.log(this.state.loginId)
        return (
            <div>
                <IconButton aria-describedby={this.id} type="button" onClick={this.handleClick}>
                    <Tooltip title="Admin">
                        <KeyIcon sx={{ p : 2 }} />
                    </Tooltip>
                </IconButton>
                <Popper id={this.id} open={this.state.isOpen} anchorEl={this.state.anchorEl} transition>
                    {({ TransitionProps }) => (
                        <Fade {...TransitionProps} timeout={350}>
                            <Paper sx={{ border: 1, p: 1}}>
                                <Box sx={{
                                    paddingLeft: 1,
                                    paddingRight: 1,
                                }}>
                                    <CustomInput
                                        labelText="Admin id"
                                        id="email"
                                        formControlProps={{
                                            fullWidth: true
                                        }}
                                        handleChange={this.handleChange}
                                        type="text"
                                    />
                                </Box>
                                <Button type="button"  sx={{
                                    fontFamily: 'Cinzel',
                                    fontWeight: 900,
                                    color: '#212121'
                                }}>
                                    Login
                                </Button>
                            </Paper>
                        </Fade>
                    )}
                </Popper>
            </div>
        )
    }
}

export default Admin