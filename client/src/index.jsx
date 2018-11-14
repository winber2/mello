import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';
import { HashRouter as Router, Route } from 'react-router-dom';

import jss from 'jss';
import preset from 'jss-preset-default';
import { ThemeProvider } from 'react-jss';

import App from 'app/App';
import store from 'redux/store';

// Install default jss plugins
jss.setup(preset());

// custom colors for jss
const theme = {

};

ReactDOM.render(
  <Provider store={store}>
    <ThemeProvider theme={theme}>
      <Router>
        <Route path="/" component={App} />
      </Router>
    </ThemeProvider>
  </Provider>, document.getElementById('root')
);
