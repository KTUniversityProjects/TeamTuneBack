/*
 * HomeReducer
 *
 * The reducer takes care of our data. Using actions, we can change our
 * application state.
 * To add a new action, add it to the switch statement in the reducer function
 *
 * Example:
 * case YOUR_ACTION_CONSTANT:
 *   return state.set('yourStateVariable', true);
 */
import { fromJS } from 'immutable';

import {
  CHANGE_PASSWORD,
  CHANGE_USERNAME,
} from './constants';

// The initial state of the App
const initialState = fromJS({
  username: '',
  password: '',
});

function loginReducer(state = initialState, action) {
  switch (action.type) {
    case CHANGE_USERNAME:
      // Delete prefied '@' from the github username

      return state.set('username', action.name);
    case CHANGE_PASSWORD:

      // Delete prefixed '@' from the github username
      return state.set('password', action.name);
    default:
      return state;
  }
}

export default loginReducer;
