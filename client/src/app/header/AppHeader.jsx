import React from 'react';
import injectSheet from 'react-jss';
import { Link } from 'react-router-dom';

import { PRIMARY_BLUE, OFF_WHITE } from 'common/constants/colors';

import MelloLogo from './MelloLogo';

const styles = {
  main: {
    height: 60,
    minHeight: 60,
    width: '100vw',
    display: 'flex',
    justifyContent: 'center',
    background: PRIMARY_BLUE,
    position: 'fixed',
    top: 0,
  },
  content: {
    height: '100%',
    width: 1300,
    padding: '0 10px',
    boxSizing: 'border-box',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'space-between',
  },
  right: {
    display: 'flex',
    alignItems: 'center',
  },
  link: {
    color: OFF_WHITE,
  },
};

class AppHeader extends React.Component {
  render() {
    const { classes } = this.props;
    return (
      <header className={classes.main}>
        <div className={classes.content}>
          <Link to="/" className={classes.link}>
            <MelloLogo />
          </Link>
          <div className={classes.right}>
            <Link to="/auth" className={classes.link}>
              LOG IN
            </Link>
          </div>
        </div>
      </header>
    );
  }
}

export default injectSheet(styles)(AppHeader);
