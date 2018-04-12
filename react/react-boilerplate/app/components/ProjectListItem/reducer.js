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
    DELETE_ITEM,
} from './constants';

// The initial state of the App
const initialState = fromJS({
  item: false
});

function mainReducer(state = initialState, action) {
  switch (action.type) {

    case DELETE_ITEM:
      return state
        .set('loading', false)
        .set('projects', action.projects)

    default:
      return state;
  }
}

export default mainReducer;
