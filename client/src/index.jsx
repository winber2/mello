import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import { HashRouter as Router, Route } from 'react-router-dom';
import {
  PRIMARY_BLUE,
  SECONDARY_BLUE,
  PRIMARY_RED,
} from 'common/constants/colors';

import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles';
import jss from 'jss';
import preset from 'jss-preset-default';
import { ThemeProvider } from 'react-jss';

import App from 'app/App';
import store from 'redux/store';

// Install default jss plugins
jss.setup(preset());

// custom colors for jss
const theme = {
  // palette: {
  //   primary: {
  //     main: PRIMARY_BLUE,
  //   },
  //   secondary: {
  //     main: SECONDARY_BLUE,
  //   },
  //   action: {
  //     active: '',
  //     disabled: '',
  //   },
  //   error: {
  //     main: PRIMARY_RED,
  //   },
  // },
};

const muiTheme = {};

ReactDOM.render(
  <Provider store={store}>
    <ThemeProvider theme={theme}>
      <MuiThemeProvider theme={createMuiTheme(muiTheme)}>
        <Router>
          <Route path="/" component={App} />
        </Router>
      </MuiThemeProvider>
    </ThemeProvider>
  </Provider>,
  document.getElementById('root')
);
