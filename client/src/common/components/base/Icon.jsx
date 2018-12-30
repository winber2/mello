import React from 'react';
import injectSheet from 'react-jss';
import classnames from 'classnames';

import MuiIcon from '@material-ui/core/Icon';
import {
  SMALL_SIZE,
  MEDIUM_SIZE,
  LARGE_SIZE,
  SMALL_ICON_FONT,
  MEDIUM_ICON_FONT,
  LARGE_ICON_FONT,
} from './constants';

const styles = {
  main: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
  },
  'size/large': {
    height: LARGE_SIZE,
    width: LARGE_SIZE,
    fontSize: LARGE_ICON_FONT,
  },
  'size/medium': {
    height: MEDIUM_SIZE,
    width: MEDIUM_SIZE,
    fontSize: MEDIUM_ICON_FONT,
  },
  'size/small': {
    height: SMALL_SIZE,
    width: SMALL_SIZE,
    fontSize: SMALL_ICON_FONT,
  },
};

class Icon extends React.Component {
  static defaultProps = {
    className: '',
    size: 'medium',
  };

  render() {
    const { className, classes, icon, size, ...rest } = this.props;
    const sizeClass = classes[`size/${size}`];

    return (
      <MuiIcon
        className={classnames(classes.main, sizeClass, className)}
        {...rest}
      >
        {icon}
      </MuiIcon>
    );
  }
}

export default injectSheet(styles)(Icon);
