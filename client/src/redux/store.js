import { applyMiddleware, createStore } from 'redux';
import { enableBatching } from 'redux-batched-actions';
import { createLogger } from 'redux-logger';
import thunk from 'redux-thunk';

import reducer from './reducer.js';

const middleware = [ thunk ];

if (process.env.NODE_ENV !== 'test') {
  middleware.push(createLogger());
}

export default createStore(
  enableBatching(reducer),
  applyMiddleware(...middleware)
);
