import React from 'react';
import injectSheet from 'react-jss';
import Button from 'common/components/base/Button';
import Input from 'common/components/base/Input';
import { Formik, Form } from 'formik';

import { postUser } from 'common/api/user';

const styles = {
  main: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
  },
  form: {
    width: 400,
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',

    '& > *': {
      margin: '10px 0',
    },
  },
};

function getDisabledStatus({ email, username, password }) {
  return !(email && username && password);
}

class Authentication extends React.Component {
  onSubmit = (values, actions) => {
    console.log(values, actions);
    postUser(values);
  };

  render() {
    const { classes } = this.props;
    return (
      <div className={classes.main}>
        <Formik
          initialValues={{
            email: '',
            username: '',
            password: '',
          }}
          onSubmit={this.onSubmit}
        >
          {props => (
            <Form onSubmit={props.handleSubmit} className={classes.form}>
              <Input
                type="text"
                name="email"
                error={props.errors.email}
                onChange={props.handleChange}
                value={props.values.email}
                label={props.errors.email || 'E-mail'}
              />
              <Input
                type="text"
                name="username"
                error={props.errors.username}
                onChange={props.handleChange}
                value={props.values.username}
                label={props.errors.username || 'Username'}
              />
              <Input
                type="password"
                name="password"
                error={props.errors.password}
                onChange={props.handleChange}
                value={props.values.password}
                label={props.errors.password || 'Password'}
              />
              <Button type="submit" disabled={getDisabledStatus(props.values)}>
                Submit
              </Button>
            </Form>
          )}
        </Formik>
      </div>
    );
  }
}

export default injectSheet(styles)(Authentication);
