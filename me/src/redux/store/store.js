import { configureStore } from '@reduxjs/toolkit';
import { createLogger } from 'redux-logger';
import { connectRouter } from 'connected-react-router'
import {BoardOptionSlice} from "../reducers/board/BoardOptionReducer";

const logger = createLogger();

const initialState = {};

export const store = configureStore({
    reducer: {
        boardOptionReducer: BoardOptionSlice.reducer
    },
    middleware: (getDefaultMiddleware) => getDefaultMiddleware().concat(logger),
    // devTools: process.env.NODE_ENV !== 'production',
    preloadedState: initialState,
    enhancers: (defaultEnhancers) => [...defaultEnhancers]
});

export default store;