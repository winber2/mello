import React from 'react';
import injectSheet from 'react-jss';

import HomeBody from './body/HomeBody';

const styles = {
  main: {
    height: '100%',
    width: '100%',
  },
};

class Home extends React.Component {
  render() {
    const { classes } = this.props;
    return (
      <main className={classes.main}>
        <HomeBody />
      </main>
    );
  }
}

export default injectSheet(styles)(Home);
