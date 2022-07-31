import React from "react";
import Box from "@mui/material/Box";
import {Button, Grid, Tab, Tabs, Typography} from "@mui/material";

class PostOption extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            option: "",
        }

        this.handleChange = this.handleChange.bind(this)
    }

    TabPanel(props) {
        const { children, value, index, ...other } = props;

        return (
            <div
                role="tabpanel"
                hidden={value !== index}
                id={`simple-tabpanel-${index}`}
                aria-labelledby={`simple-tab-${index}`}
                {...other}
            >
                {value === index && (
                    <Box sx={{ p: 3 }}>
                        <Typography>{children}</Typography>
                    </Box>
                )}
            </div>
        );
    }

    a11yProps(index) {
        return {
            id: `simple-tab-${index}`,
            'aria-controls': `simple-tabpanel-${index}`,
        };
    }

    handleChange = (event, newValue) => {
        this.setState({
            option: newValue
        })
    };

    render() {
        return (
            <Box sx={{
                display: 'flex',
                pl: 4,
                pr: 4,
                pt: 4,
            }} justifyContent="space-between">
                <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
                    <Tabs value={this.state.value} onChange={this.handleChange} aria-label="basic tabs example">
                        <Tab label="전체" {...this.a11yProps(0)} style={{
                            fontSize: 18,
                            fontWeight: 900,
                            fontFamily: "Elice Digital Baeum",
                        }}/>
                        <Tab label="인기글" {...this.a11yProps(1)} style={{
                            fontSize: 18,
                            fontWeight: 900,
                            fontFamily: "Elice Digital Baeum",
                        }}/>
                        <Tab label="공지" {...this.a11yProps(2)} style={{
                            fontSize: 18,
                            fontWeight: 900,
                            fontFamily: "Elice Digital Baeum",
                        }}/>
                    </Tabs>
                </Box>
                <Button variant="contained" color="success" sx={{
                    fontSize: 15,
                    fontFamily: "Elice Digital Baeum",
                }}>글쓰기</Button>
            </Box>
        )
    }
}

export default PostOption