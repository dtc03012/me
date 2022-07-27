import React from "react";
import Box from "@mui/material/Box";
import CustomInput from "../../../../components/customInput";
import FormControl from "@material-ui/core/FormControl";
import InputLabel from "@material-ui/core/InputLabel";
import {Grid, MenuItem, Select, TextField} from "@mui/material";
import {createTheme, ThemeProvider} from "@mui/material/styles";

const Theme = createTheme({
    typography: {
        body1: {
            fontFamily: "Elice Digital Baeum",
            fontSize: 13,
            fontWeight: 700,
        }
    },
});


class TextSearchOption extends React.Component {

    options = [
        '제목+내용',
        '제목',
        '내용',
        '작성자',
        '댓글',
    ];

    constructor(props) {
        super(props);
        // Don't call this.setState() here!
        this.state = {
            text: "",
            option: "",
            open: false,
        };
        this.handleInputChange = this.handleInputChange.bind(this)
        this.handleOptionChange = this.handleOptionChange.bind(this)
        this.handleOptionOpen = this.handleOptionOpen.bind(this)
        this.handleOptionClose = this.handleOptionClose.bind(this)
    }

    handleInputChange = (event) => {
        this.setState({
            text: event.currentTarget.value
        })
    }

    handleOptionChange = (event) => {
        console.log(event)
        this.setState({
            option: event.target.value
        })
    };

    handleOptionClose = () => {
        this.setState({
            open: false,
        })
    };

    handleOptionOpen = () => {
        this.setState({
            open: true,
        })
    };

    render() {
        return (
            <Box sx={{
                display: 'flex',
            }}>
                <Grid container direction="row" alignItems="center">
                    <Grid item sx={{
                        pr: 3,
                    }}>
                        <CustomInput
                            labelText=""
                            id="text"
                            formControlProps={{
                                fullWidth: true
                            }}
                            handleChange={this.handleInputChange}
                            type="text"
                        />
                    </Grid>
                    <Grid item>
                        <ThemeProvider theme={Theme}>
                            <TextField
                                label="옵션"
                                select
                                value={this.state.option}
                                onChange={this.handleOptionChange}
                                InputProps={{
                                    style: {
                                        fontSize: 13,
                                        fontFamily: "Elice Digital Baeum",
                                    }
                                }}
                                InputLabelProps={{
                                    style: {
                                        fontSize: 13,
                                        fontFamily: "Elice Digital Baeum",
                                    }
                                }}
                                sx={{
                                    width: 110,
                                }}
                            >
                                {this.options.map((option) => (
                                    <MenuItem
                                        key={option}
                                        value={option}
                                    >
                                        {option}
                                    </MenuItem>
                                ))}
                            </TextField>
                        </ThemeProvider>
                    </Grid>
                </Grid>
            </Box>
        )
    }
}

export default TextSearchOption