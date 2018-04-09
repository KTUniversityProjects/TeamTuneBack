
import {
  LOGIN,
  LOGIN_SUCCESS,
  REQUEST_ERROR,
} from './constants';

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
 * Dispatched when the repositories are loaded by the request saga
 *
 * @param  {array} response The response dta
 *
 * @return {object}      An action object with a type of LOGIN_SUCCESS passing the response
 */
export function loginSuccess(response) {
  return {
    type: LOGIN_SUCCESS,
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
