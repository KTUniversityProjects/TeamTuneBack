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
  LOAD_PROJECTS,
    LOAD_PROJECTS2,
} from './constants';

// The initial state of the App
const initialState = fromJS({
  loading: false,
  error: false,
  projects: false,
});

function mainReducer(state = initialState, action) {
  console.log("REDUCER");
    console.log(action.type);
  switch (action.type) {

    case LOAD_PROJECTS2:
    console.log(action.projects);
      return state
        .set('loading', false)
        .set('error', false)
        .set('projects', action.projects)

    case LOAD_PROJECTS:
      return state
        .set('loading', true)
        .set('error', false)
      //  .set('projects', action.projects)
        .set('projects', [])
        /*
    case LOAD_REPOS_SUCCESS:
      return state
        .setIn(['userData', 'repositories'], action.repos)
        .set('loading', false)
        .set('currentUser', action.username);
    case LOAD_REPOS_ERROR:
      return state
        .set('error', action.error)
        .set('loading', false);
        */
    default:
      return state;
  }
}

export default mainReducer;
