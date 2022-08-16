import React from "react";
import Box from "@mui/material/Box";
import {Button, Grid, Paper, Typography} from "@mui/material";
import TextSearchOption from "./textSearchOption";
import TagSearchOption from "./tagSearchOption";
import {createTheme, ThemeProvider} from "@mui/material/styles";
import {useDispatch, useSelector} from "react-redux";
import {
    setSelectedTag,
    setInputTag,
} from "../../../../../../redux/reducers/board/BoardOptionReducer";

const Theme = createTheme({
    typography: {
        body1: {
            fontFamily: "Elice Digital Baeum",
            fontSize: 13,
            fontWeight: 700,
        },
        body2: {
            fontFamily: "Elice Digital Baeum",
            fontSize: 20,
            fontWeight: 1000,
        }
    },
    palette: {
        background: {
            belowAppbarBox: '#212121',
        },
    }
});

class SearchOption extends React.Component {

    search = window.location.search;
    urlSearchParams = new URLSearchParams(this.search)
    classificationOption = this.urlSearchParams.get("classificationOption")

    constructor(props) {
        super(props);
        this.state = {
            text: "",
            option: "",
            searchURL: "",
        }

        if(this.classificationOption == null) this.classificationOption = "All"

        this.handleInitButtonClick = this.handleInitButtonClick.bind(this)
    }

    getText = () => {
        return this.state.text
    }

    setText = (text) => {
        this.setState({
            text: text
        })
    }

    getOption = () => {
        return this.state.option
    }

    setOption = (option) => {
        this.setState({
            option: option
        })
    }

    handleInitButtonClick = (event) => {
        this.setState({
            text: "",
            option: "",
        })
        this.props.dispatch(setSelectedTag({selectedTag: []}))
        this.props.dispatch(setInputTag({inputTag: ""}))
    }

    optionMap = {
        "제목+내용": "TitleOrContent",
        "제목": "Title",
        "내용": "Content",
        "작성자": "Writer",
    }

    createSearchURL = () => {
        let url = "/board/lists?page=1"
        if(this.state.option !== "" || this.optionMap[this.state.option] !== undefined) {
            url += "&queryOption=" + this.optionMap[this.state.option]
        }

        if(this.state.text !== ""){
            url += "&queryString=" + this.state.text
        }

        url += "&classificationOption=" + this.classificationOption

        this.props.selectedTag.map((tag) => {
            url += "&tags=" + tag
        })

        return url
    }


    render() {
        return (
            <Box sx={{
                p: 6,
            }}>
                <Paper elevation={3} >
                    <Grid spacing={2} container
                          direction="column"
                          sx={{
                              pl: 4,
                          }}
                    >
                        <Grid item sx={{
                            width: '80%',
                            justifyContent: 'center'
                        }}>
                            <ThemeProvider theme={Theme}>
                                <Typography variant="body2" sx={{
                                    pb: 1,
                                }}>
                                    검색 옵션
                                </Typography>
                                <Box sx={{
                                    backgroundColor: 'background.belowAppbarBox',
                                    width: '100%',
                                    height: 2,
                                }}/>
                            </ThemeProvider>
                        </Grid>
                        <Grid item>
                            <TextSearchOption
                                getText={this.getText}
                                setText={this.setText}
                                getOption={this.getOption}
                                setOption={this.setOption}
                            />
                        </Grid>
                        <Grid item sx={{
                            width: '80%',
                            justifyContent: 'center'
                        }}>
                            <ThemeProvider theme={Theme}>
                                <Typography variant="body2" sx={{
                                    pb: 1,
                                }}>
                                    태그 옵션
                                </Typography>
                                <Box sx={{
                                    backgroundColor: 'background.belowAppbarBox',
                                    width: '100%',
                                    height: 2,
                                }}/>
                            </ThemeProvider>
                        </Grid>
                        <Grid item>
                            <TagSearchOption/>
                        </Grid>
                        <Grid item sx={{
                            mt: 3,
                            pr: 6,
                            display: 'flex',
                            justifyContent: 'space-between'
                        }}>
                            <Button variant="contained" color="success" sx={{
                                fontSize: 15,
                                fontFamily: "Elice Digital Baeum",
                            }} href={this.createSearchURL()}>검색</Button>
                            <Button variant="contained" color="error" onClick={this.handleInitButtonClick} sx={{
                                fontSize: 15,
                                fontFamily: "Elice Digital Baeum",
                            }}>초기화</Button>
                        </Grid>
                        <Grid item sx={{
                            mt: 2,
                            mb: 2,
                            pr: 4,
                        }}>
                            <ThemeProvider theme={Theme}>
                                <Grid item>
                                    <Typography variant="body1">
                                        기능 추가 문의는 dtc03012@gmail.com 로 해주세요!
                                    </Typography>
                                </Grid>
                            </ThemeProvider>
                        </Grid>
                    </Grid>
                </Paper>

            </Box>
        )
    }
}

export default () => {
    const dispatch = useDispatch();
    const selectedTag = useSelector((state) => state.boardOptionReducer.selectedTag);
    return (
        <SearchOption
            selectedTag={selectedTag}
            dispatch={dispatch}/>
    )
}