import { configureStore } from '@reduxjs/toolkit';
import { createLogger } from 'redux-logger';
import { connectRouter } from 'connected-react-router'
import {tagOptionSlice} from "../reducers/board/tagOptionReducer";

const logger = createLogger();

const initialState = {};

export const store = configureStore({
    reducer: {
        tagOptionReducer: tagOptionSlice.reducer
    },
    middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(logger),
    // devTools: process.env.NODE_ENV !== 'production',
    preloadedState: initialState,
    enhancers: (defaultEnhancers) => [...defaultEnhancers]
});

export default store;