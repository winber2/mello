import React from 'react';
import injectSheet from 'react-jss';

import { PRIMARY_GREY, SECONDARY_GREY } from 'common/constants/colors';

import Box from 'common/components/general/Box';

const styles = {
  main: {
    width: '100%',
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
    color: 'white',
  },
  box: {
    height: 400,
    minHeight: 400,
    width: 1250,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    margin: 20,
    backgroundColor: PRIMARY_GREY,
  },
  header: {
    width: 1300,
    height: 300,
    background: PRIMARY_GREY,
    borderRadius: '0 0 5px 5px',
  },
  body: {
    width: 1250,
    height: 800,
    background: SECONDARY_GREY,
  }
};

class HomeBody extends React.Component {
  render() {
    const { classes } = this.props;
    return (
      <content className={classes.main}>
        <header className={classes.header}>

        </header>
        <section className={classes.body}>
        </section>
        {/* <Box className={classes.box}>
          I am a box with content hello, lets put a search bar here
        </Box>
        <Box className={classes.box} style={{ height: 900, minHeight: 900 }}>
          I am another box that is bigger because im special
        </Box> */}
      </content>
    );
  }
}

export default injectSheet(styles)(HomeBody);
