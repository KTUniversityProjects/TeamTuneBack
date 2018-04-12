import {
  CHANGE_USERNAME,
  CHANGE_PASSWORD,
  REQUEST_ERROR,
  LOGIN_SUCCESS,
} from './constants';
import {LOGIN} from "./constants";


/**
 * Load the repositories, this action starts the request saga
 *
 * @return {object} An action object with a type of LOGIN
 */
export function login() {
  return {
    type: LOGIN,
  };
}

/**
 * Changes the input field of the form
 *
 * @param  {name} name The new text of the input field
 *
 * @return {object}    An action object with a type of CHANGE_USERNAME
 */
 export function changePassword(name) {
   return {
     type: CHANGE_PASSWORD,
     name,
   };
 }

export function changeUsername(name) {
  return {
    type: CHANGE_USERNAME,
    name,
  };
}
