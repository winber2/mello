import React from 'react';
import injectSheet from 'react-jss';
import { Switch, Route } from 'react-router';
import { OFF_WHITE, PRIMARY_GREY } from 'common/constants/colors';

import PrivateRoute from 'auth/PrivateRoute';
import AppHeader from './header/AppHeader';
import Home from 'home/Home';

const styles = {
  main: {
    height: '100vh',
    width: '100vw',
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    backgroundColor: OFF_WHITE,
  },
  body: {
    flexGrow: 1,
    width: 1300,
    backgroundColor: PRIMARY_GREY,
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
