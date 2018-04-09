/**
 * Homepage selectors
 */

import { createSelector } from 'reselect';

const selectLogin = (state) => state.get('login');

const makeSelectUsername = () => createSelector(
  selectLogin,
  (loginState) => loginState.get('username')
);
const makeSelectPassword = () => createSelector(
  selectLogin,
  (loginState) => loginState.get('password')
);

export {
  selectLogin,
  makeSelectUsername,
  makeSelectPassword,
};
