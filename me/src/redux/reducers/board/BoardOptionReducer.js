import { createSlice } from '@reduxjs/toolkit'

export const BoardOptionSlice = createSlice({
    name: 'BoardOption',
    initialState: {
        selectedTag: [],
        inputTag: "",
    },
    reducers: {
        setSelectedTag: (state, action) => {
            state.selectedTag = action.payload.selectedTag
        },

        setInputTag: (state, action) => {
            state.inputTag = action.payload.inputTag
        },
    },
})

export const { setSelectedTag, setInputTag } = BoardOptionSlice.actions