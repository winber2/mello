import React from 'react';
import injectSheet from 'react-jss';

import { H1_FONT } from 'common/constants/styles';

const styles = {
  main: {
    fontSize: H1_FONT,
    fontWeight: 700,
    margin: 20,
  },
};

class MelloHeader extends React.Component {
  render() {
    const { classes, children } = this.props;
    return <header className={classes.main}>{children}</header>;
  }
}

export default injectSheet(styles)(MelloHeader);
