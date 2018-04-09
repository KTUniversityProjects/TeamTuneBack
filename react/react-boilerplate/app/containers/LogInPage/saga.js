/**
 * Gets the repositories of the user from Github
 */

import { call, put, select, takeLatest } from 'redux-saga/effects';

import request from 'utils/request';
import {makeSelectPassword, makeSelectUsername} from 'containers/LoginPage/selectors';
import {LOGIN} from "../App/constants";
import {loginSuccess, requestError} from "../App/actions";

/**
 * Login request handler
 */
export function* loginRequest() {

  // Select username and password from store
  const username = yield select(makeSelectUsername());
  const password = yield select(makeSelectPassword());
  const requestURL = `http://localhost:1338`;

  try {
    // Call our request helper (see 'utils/request')
    const response = yield call(request, requestURL, {username: username, password: password});
    yield put(loginSuccess(response));
  } catch (err) {
    yield put(requestError(err));
  }
}

/**
 * Root saga manages watcher lifecycle
 */
export default function* checkLoginState() {
  yield takeLatest(LOGIN, loginRequest);
}
