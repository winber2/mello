import React from 'react';
import injectSheet from 'react-jss';

import { PRIMARY_BLUE } from 'common/constants/colors';

import MelloLogo from './MelloLogo';

const styles = {
  main: {
    height: 60,
    width: '100%',
    display: 'flex',
    justifyContent: 'center',
    background: PRIMARY_BLUE,
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
};

class AppHeader extends React.Component {
  render() {
    const { classes } = this.props;
    return (
      <header className={classes.main}>
        <div className={classes.content}>
          <MelloLogo />
          <div className={classes.right}>LOG IN</div>
        </div>
      </header>
    );
  }
}

export default injectSheet(styles)(AppHeader);
