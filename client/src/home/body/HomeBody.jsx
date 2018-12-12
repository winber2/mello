import React from 'react';
import injectSheet from 'react-jss';

const styles = {
  main: {
    height: '100%',
    width: '100%',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    color: 'white',
  },
};

class HomeBody extends React.Component {
  render() {
    const { classes } = this.props;
    return <content className={classes.main}>content</content>;
  }
}

export default injectSheet(styles)(HomeBody);
