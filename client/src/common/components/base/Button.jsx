import React from 'react';
import injectSheet from 'react-jss';
import classnames from 'classnames';

import MuiButton from '@material-ui/core/Button';
import { SECONDARY_GREY } from 'common/constants/colors';
import { SMALL_FONT, MEDIUM_FONT, LARGE_FONT } from './constants';

const styles = {
  main: {
    display: 'flex',
    alignItems: 'center',
    border: `1px solid ${SECONDARY_GREY}`,
    borderRadius: 5,
    padding: '3px 12px 4px 12px',
  },
  'size/large': {
    fontSize: LARGE_FONT,
  },
  'size/medium': {
    fontSize: MEDIUM_FONT,
  },
  'size/small': {
    fontSize: SMALL_FONT,
  },
};

class Button extends React.Component {
  static defaultProps = {
    className: '',
    size: 'medium',
  };

  render() {
    const { classes, children, className, size, ...rest } = this.props;
    const sizeClass = classes[`size/${size}`];

    return (
      <MuiButton
        className={classnames(classes.main, sizeClass, className)}
        size={size}
        {...rest}
      >
        {children}
      </MuiButton>
    );
  }
}

export default injectSheet(styles)(Button);
