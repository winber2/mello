import React from 'react';
import injectSheet from 'react-jss';
import classnames from 'classnames';

import TextField from '@material-ui/core/TextField';

const styles = {
  main: {
    width: '100%',
  },
  input: {
    padding: '5px 0 5px 0',
  },
};

class Input extends React.Component {
  static defaultProps = {
    className: '',
    size: 'medium',
  };

  render() {
    const { classes, children, className, size, ...rest } = this.props;
    return (
      <TextField
        className={classnames(classes.main, className)}
        InputLabelProps={{ shrink: true }}
        inputProps={{ className: classes.input }}
        {...rest}
      />
    );
  }
}

export default injectSheet(styles)(Input);
