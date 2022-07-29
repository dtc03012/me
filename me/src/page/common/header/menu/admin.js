import React from 'react';
import axios from "axios";
import {v4} from 'uuid';
import {Button, Fade, IconButton, Paper, Popper, Typography} from "@mui/material";
import KeyIcon from "@mui/icons-material/Key";
import LockIcon from '@mui/icons-material/Lock';
import CustomInput from "../../../components/customInput";
import Box from "@mui/material/Box";
import {setCookie, getCookie} from '../../../../util/cookie'

class Admin extends React.Component {

    constructor(props) {
        super(props);
        // Don't call this.setState() here!
        this.state = {
            isOpen: false,
            anchorEl: null,
            loginID: "",
            loginFail: false,
        };
        this.handleClick = this.handleClick.bind(this);
        this.handleChange = this.handleChange.bind(this);
        this.login = this.login.bind(this);
        this.logout = this.logout.bind(this);
        this.checkAdminLogin = this.checkAdminLogin.bind(this);
        this.isLoginFail = this.isLoginFail.bind(this);
        this.adminIcon = this.adminIcon.bind(this);
    }

    login = (event) => {
        let uuid = v4()
        axios.post('/v2/login-admin', {
            password: this.state.loginID,
            uuid: uuid,
        }).then(
            response => {
                if(response.data.isAdmin){
                    setCookie('uuid', uuid)
                    this.setState({
                        isAdmin: true,
                        isOpen: false,
                        loginID: "",
                    })
                }else {
                    this.setState({
                        loginFail : true,
                    })
                }
            }
        )
    }

    logout = (event) => {
        setCookie('uuid','')
        this.setState({
            isAdmin: false,
            isOpen: false,
        })
    }

    checkAdminLogin = (event) =>  {
        let uuid = getCookie('uuid')
        if(uuid) {
            axios.post('/v2/find-admin-uuid', {
                uuid: uuid
            }).then(
                response => {
                    if(response.data.isFind) {
                        this.setState({
                            isAdmin: true
                        })
                    } else {
                        this.setState({
                            isAdmin: false
                        })
                    }
                }
            )
            return
        }
        this.setState({
            isAdmin: false
        })
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

    componentDidMount() {
        this.checkAdminLogin()
    }

    isLoginFail() {
        if(this.state.loginFail) {
            return (
                <Typography sx={{
                    fontFamily: 'Cinzel',
                    fontSize: 12,
                    fontWeight: 900,
                    color: '#f44336',
                    paddingRight: 1,
                }}>
                    Wrong!
                </Typography>
            )
        }
        return undefined
    }

    adminIcon() {
        if(this.state.isAdmin) {
            return (
                <LockIcon sx={{p: 2}}/>
            )
        }
        return (
            <KeyIcon sx={{p: 2}}/>
        )
    }
    render() {
        return (
            <div>
                <IconButton aria-describedby={this.id} type="button" onClick={this.handleClick}>
                    {this.adminIcon()}
                </IconButton>
                <Popper open={this.state.isOpen && !this.state.isAdmin} anchorEl={this.state.anchorEl} transition>
                    {({TransitionProps}) => (
                        <Fade {...TransitionProps} timeout={350}>
                            <Paper sx={{border: 1, p: 1}}>
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
                                <Box sx={{
                                    display: "flex",
                                    justifyContent: 'space-between',
                                    alignItems: 'center'
                                }}
                                >
                                    <Button type="button" sx={{
                                        fontFamily: 'Cinzel',
                                        fontWeight: 900,
                                        color: '#212121',
                                    }} onClick={this.login}>
                                        Login
                                    </Button>
                                    {this.isLoginFail()}
                                </Box>
                            </Paper>
                        </Fade>
                    )}
                </Popper>
                <Popper id={this.id} open={this.state.isOpen && this.state.isAdmin} anchorEl={this.state.anchorEl}
                        transition>
                    {({TransitionProps}) => (
                        <Fade {...TransitionProps} timeout={350}>
                            <Paper sx={{border: 1, p: 1}}>
                                <Button type="button" sx={{
                                    fontFamily: 'Cinzel',
                                    fontWeight: 900,
                                    color: '#ffffff',
                                    bgcolor: '#f44336',
                                    '&:hover': {
                                        bgcolor: '#000000',
                                        color: 'white',
                                    },
                                }} onClick={this.logout}>
                                    Logout
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