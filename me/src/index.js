import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import PageRouter from "./pageRouter";
import reportWebVitals from './reportWebVitals';
import {Provider} from "react-redux";
import store from './redux/store/store'
import Header from "./page/common/header/header";
import Footer from "./page/common/footer/footer";
import Box from "@mui/material/Box";

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <div>
          <link href="https://fonts.googleapis.com/css2?family=Cinzel:wght@500&family=Open+Sans&display=swap" rel="stylesheet"/>
          <link href="https://fonts.googleapis.com/css2?family=Ubuntu&display=swap" rel="stylesheet"/>
          <link href="//font.elice.io/EliceDigitalBaeum.css" rel="stylesheet"/>
          <Provider store={store}>
              <Box sx={{
                  backgroundColor: '#f5f5f5',
              }}>
                  <Header/>
                  <PageRouter/>
                  <Footer/>
              </Box>
          </Provider>
    </div>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
