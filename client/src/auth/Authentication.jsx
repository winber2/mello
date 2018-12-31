import React from 'react';
import injectSheet from 'react-jss';
import { Formik, Form } from 'formik';
import { withRouter } from 'react-router';

import MelloHeader from 'common/components/general/MelloHeader';
import Button from 'common/components/base/Button';
import Input from 'common/components/base/Input';

import { postUser, getToken, getUsers } from 'common/api/user';
import { getURLParams } from 'common/helpers/routeUtils';

import { SECONDARY_GREY, SECONDARY_BLUE } from 'common/constants/colors';

const styles = {
  main: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
    height: 500,
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
  account: {
    fontWeight: 300,
    color: SECONDARY_GREY,
    margin: 10,
    display: 'flex',
  },
  toggle: {
    textDecoration: 'underline',
    color: SECONDARY_BLUE,
    cursor: 'pointer',
    marginLeft: 5,
  },
};

const LOGIN = 'login';

function getLoginText(type) {
  return type === LOGIN ? 'LOG IN' : 'SIGN UP';
}

function getDisabledStatus({ email, username, password }) {
  return !(email && username && password);
}

class Authentication extends React.Component {
  onSubmit = (type, values) => {
    if (type === LOGIN) {
      getToken(values);
    } else {
      postUser(values);
    }
  };

  toggleLogin = type => {
    if (type === LOGIN) {
      this.props.history.push('/auth');
    } else {
      this.props.history.push('/auth?type=login');
    }
  };

  render() {
    const { classes, location } = this.props;
    const { type } = getURLParams(location.search);

    return (
      <div className={classes.main}>
        <MelloHeader>{getLoginText(type)}</MelloHeader>
        <Formik
          initialValues={{
            email: '',
            username: '',
            password: '',
          }}
          onSubmit={values => this.onSubmit(type, values)}
        >
          {props => (
            <Form onSubmit={props.handleSubmit} className={classes.form}>
              {type !== LOGIN && (
                <Input
                  type="text"
                  name="email"
                  error={props.errors.email}
                  onChange={props.handleChange}
                  value={props.values.email}
                  label={props.errors.email || 'E-mail'}
                />
              )}
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
              <span className={classes.account}>
                {type === LOGIN ? "Don't" : 'Already'} have an account?
                <div
                  tabIndex="0"
                  role="button"
                  className={classes.toggle}
                  onClick={() => this.toggleLogin(type)}
                >
                  {type === LOGIN ? 'Sign Up' : 'Log In'}
                </div>
              </span>
              <Button type="submit">{getLoginText(type)}</Button>
            </Form>
          )}
        </Formik>
        this should only work if you are logged in
        <Button onClick={() => getUsers().then(data => console.log(data))}>
          GET USERS
        </Button>
      </div>
    );
  }
}

export default withRouter(injectSheet(styles)(Authentication));
