import React from 'react';
import axios from "axios";
import {v4} from 'uuid';
import {Button, Fade, IconButton, Paper, Popper, Tooltip, Typography} from "@mui/material";
import KeyIcon from "@mui/icons-material/Key";
import CustomInput from "../../../components/customInput";
import Box from "@mui/material/Box";

class Admin extends React.Component {

    constructor(props) {
        super(props);
        // Don't call this.setState() here!
        this.state = {
            isAdmin: this.CheckAdminLogin(),
            isOpen: false,
            anchorEl: null,
            loginID: "",
            loginFail: false,
        };
        this.handleClick = this.handleClick.bind(this);
        this.handleChange = this.handleChange.bind(this);
        this.Login = this.Login.bind(this);
        this.CheckAdminLogin = this.CheckAdminLogin.bind(this);
    }

    Login = (event) => {
        let uuid = v4()
        axios.post('/v2/login-admin', {
            password: this.state.loginID,
            uuid: uuid,
        }).then(
            response => {
                if(response.data.isAdmin){
                    localStorage.setItem('uuid', uuid)
                    this.setState({
                        isOpen: false,
                        loginId : "",
                    })
                }else {
                    this.setState({
                        loginFail : true,
                    })
                }
            }
        )
    }

    CheckAdminLogin = (event) =>  {
        let uuid = localStorage.getItem("uuid")
        console.log(uuid)
        if(uuid) {
            axios.post('/v2/find-admin-uuid', {
                uuid: uuid
            }).then(
                response => {
                    console.log(response)
                    if(response.data.isFind) {
                        return true
                    }
                }
            )
        }
        return false
    }

    handleClick = (event) => {
        this.setState({
            anchorEl: event.currentTarget,
            isOpen: !this.state.isOpen,
            loginID: "",
            loginFail: false,
        })
    }

    handleChange = (event) => {
        this.setState({
            loginID: event.currentTarget.value
        })
    }

    render() {
        this.canBeOpen = this.state.isOpen && Boolean(this.state.anchorEl);
        this.id = this.canBeOpen ? 'transition-popper' : undefined;

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
                                }} onClick={this.Login}>
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