import {
  CHANGE_USERNAME,
  CHANGE_PASSWORD,
  CHANGE_PASSWORD_CONFIRM,
  CHANGE_EMAIL,
  REQUEST_ERROR,
  SIGNUP_SUCCESS,
} from './constants';

import {SIGNUP} from "./constants";
/**
 * Load the repositories, this action starts the request saga
 *
 * @return {object} An action object with a type of LOGIN
 */
export function signup() {
  return {
    type: SIGNUP,
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

export function changePasswordConfirm(name) {
  return {
    type: CHANGE_PASSWORD_CONFIRM,
    name,
  };
}

export function changeEmail(name) {
  return {
    type: CHANGE_EMAIL,
    name,
  };
}

/**
 * Dispatched when the repositories are loaded by the request saga
 *
 * @param  {array} response The response dta
 *
 * @return {object}      An action object with a type of LOGIN_SUCCESS passing the response
 */
export function signupSuccess(response) {
  return {
    type: SIGNUP_SUCCESS,
    response,
  };
}

/**
 * Dispatched when loading the repositories fails
 *
 * @param  {object} error The error
 *
 * @return {object}       An action object with a type of REQUEST_ERROR passing the error
 */
export function requestError(error) {
  return {
    type: REQUEST_ERROR,
    error,
  };
}
