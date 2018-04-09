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
  CHANGE_PASSWORD_CONFIRM,
  CHANGE_EMAIL,
  REQUEST_ERROR,
  SIGNUP_SUCCESS,
} from './constants';

// The initial state of the App
const initialState = fromJS({
  username: '',
  password: '',
  passwordConfirm:'',
  email:'',
});

function signupReducer(state = initialState, action) {
switch (action.type) {
    //Login request success
    case SIGNUP_SUCCESS:
      console.log(action.response);
      if(action.response.code == 0)
      {
        sessionStorage.setItem('sessionID', action.response.data);
      }
      return state;


    //Any request error
    case REQUEST_ERROR:
      return state
        .set('error', action.error)
        .set('loading', false);

    case CHANGE_USERNAME:
      return state.set('username', action.name);

    case CHANGE_PASSWORD:
      return state.set('password', action.name);

    case CHANGE_PASSWORD_CONFIRM:
      return state.set('passwordConfirm', action.name);

    case CHANGE_EMAIL:
      return state.set('email', action.name);
    default:
      return state;
  }
}

export default signupReducer;
