/*
 * AppReducer
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
  LOGIN_SUCCESS,
  REQUEST_ERROR,
} from './constants';

// The initial state of the App
const initialState = fromJS({

});

function appReducer(state = initialState, action) {
  switch (action.type) {

    //Login request success
    case LOGIN_SUCCESS:
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


    default:
      return state;
  }
}

export default appReducer;
