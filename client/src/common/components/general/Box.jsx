import React from 'react';
import injectSheet from 'react-jss';
import classnames from 'classnames';

import { OFF_WHITE } from 'common/constants/colors';

const styles = {
  main: {
    boxShadow: '0 0 2px 2px rgba(0, 0, 0, 0.1)',
    backgroundColor: props => props.color,
  },
};

class Box extends React.Component {
  static defaultProps = {
    color: OFF_WHITE,
  };

  render() {
    const { classes, className, children, style } = this.props;
    return (
      <article className={classnames(classes.main, className)} style={style}>
        {children}
      </article>
    );
  }
}

export default injectSheet(styles)(Box);
