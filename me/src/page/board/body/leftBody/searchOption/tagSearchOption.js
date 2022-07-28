import React from "react";
import Box from "@mui/material/Box";
import {Chip, Grid, IconButton} from "@mui/material";
import AddCircleOutlineIcon from '@mui/icons-material/AddCircleOutline';
import CustomInput from "../../../../components/customInput";

class TagSearchOption extends React.Component {

    constructor(props) {
        super(props);

        this.handleInputTagChange = this.handleInputTagChange.bind(this)
        this.handleDeleteTagClick = this.handleDeleteTagClick.bind(this)
        this.checkExistedTag = this.checkExistedTag.bind(this)
        this.createTagList = this.createTagList.bind(this)
        this.handleAddTagClick = this.handleAddTagClick.bind(this)
    }

    handleInputTagChange = (event) => {
        this.props.setInputTag(event.currentTarget.value)
    }

    handleDeleteTagClick = (tag) => {
        this.props.setSelectedTags(this.props.getSelectedTags().filter(item => item !== tag))
    }

    handleAddTagClick = (tag) => {
        if(this.props.getInputTag() !== "" && this.checkExistedTag(tag) === false) {
            this.props.setSelectedTags(this.props.getSelectedTags().concat(tag))
        }
        this.props.setInputTag("")
    }

    checkExistedTag = (tag) => {
        let exist = false;
        this.props.getSelectedTags().forEach(item => {
            if(item === tag) {
                exist = true
            }
        })

        return exist
    }

    createTagList = () => {
        return (this.props.getSelectedTags().map( (tag) => (
                <Grid item>
                    <Chip label={String(tag)} variant="outlined" onDelete={() => this.handleDeleteTagClick(String(tag))} sx={{
                        '& .MuiChip-label': {
                            fontSize: 13,
                            fontFamily: "Elice Digital Baeum",
                        },
                    }} />
                </Grid>
            )
        ))
    }

    render() {
        return (
            <Box sx={{
                display: 'flex',
                pr: 2
            }}>
                <Grid container spacing={2} direction="row">
                    <Grid container alignItems="center" spacing={1} direction="row" sx={{
                        pl: 2,
                    }}>
                        <Grid item>
                            <CustomInput
                                labelText=""
                                value={this.props.getInputTag()}
                                id="text"
                                formControlProps={{
                                    fullWidth: true
                                }}
                                handleChange={this.handleInputTagChange}
                                type="text"
                            />
                        </Grid>
                        <Grid item>
                            <IconButton type="button" onClick={() => this.handleAddTagClick(this.props.getInputTag())}>
                                <AddCircleOutlineIcon sx={{
                                    fontSize: 28,
                                }}/>
                            </IconButton>
                        </Grid>
                    </Grid>
                    <Grid container spacing={2} alignItems="center" sx={{
                        pl: 2,
                        pt: 2
                    }}>
                        {this.createTagList()}
                    </Grid>
                </Grid>
            </Box>
        )
    }
}

export default TagSearchOption