import { createSlice } from '@reduxjs/toolkit'

export const tagOptionSlice = createSlice({
    name: 'tagOption',
    initialState: {
        selectedTag: [],
        inputTag: "",
    },
    reducers: {
        setSelectedTag: (state, action) => {
            state.selectedTag = action.payload.selectedTag
        },

        setInputTag: (state, action) => {
            console.log(action)
            state.inputTag = action.payload.inputTag
        },
    },
})

export const { setSelectedTag, setInputTag } = tagOptionSlice.actions