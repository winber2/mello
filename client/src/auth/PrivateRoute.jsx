import React from 'react';
import PropTypes from 'prop-types';
import { Route, Redirect } from 'react-router-dom';

const PrivateRoute = ({ component: Component, isAuthenticated, ...rest }) => (
  <Route {...rest} render={props =>
      isAuthenticated
        ? <Component {...props} />
        : <Redirect
            to={{
              pathname: '/auth',
              state: { from: props.location },
            }}
          />
  } />
);

PrivateRoute.propTypes = {
  component: PropTypes.func.isRequired,
  isAuthenticated: PropTypes.bool.isRequired,
  location: PropTypes.object,
};

export default PrivateRoute;
