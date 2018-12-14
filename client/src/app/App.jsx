import React from 'react';
import injectSheet from 'react-jss';
import { Switch, Route } from 'react-router';
import { OFF_WHITE, PRIMARY_GREY } from 'common/constants/colors';

import Home from 'home/Home';
import PrivateRoute from 'auth/PrivateRoute';
import AppHeader from './header/AppHeader';

const styles = {
  main: {
    height: 'min-content',
    minHeight: '100vh',
    width: '100vw',
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    backgroundColor: OFF_WHITE,
  },
  body: {
    marginTop: 60,
    height: 'calc(100% - 60px)',
    width: 1300,
  },
};

class App extends React.Component {
  render() {
    const { classes } = this.props;
    return (
      <div className={classes.main}>
        <AppHeader />
        <div className={classes.body}>
          <Switch>
            <PrivateRoute
              path={'/users'}
              component={Home}
              isAuthenticated
            />
            <Route
              path={'/'}
              component={Home}
            />
            {/* <Redirect
              path={'/auth'}
              component={Authentication}
            /> */}
          </Switch>
        </div>
      </div>
    );
  }
}

export default injectSheet(styles)(App);
