diff git a/react/reactboilerplate/app/containers/App/reducer.js b/react/reactboilerplate/app/containers/App/reducer.js
deleted file mode 100644
index 34af859..0000000
 a/react/reactboilerplate/app/containers/App/reducer.js
+++ /dev/null
@@ 1,55 +0,0 @@
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
  loading: false,
  error: false,
  currentUser: false,
  userData: {
    repositories: false,
  },
});

function appReducer(state = initialState, action) {
  switch (action.type) {

    //Login request success
    case LOGIN_SUCCESS:
      console.log(action.response.message);
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
