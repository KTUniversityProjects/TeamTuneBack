/**
 * Gets the repositories of the user from Github
 */

import { call, put, select, takeLatest } from 'redux-saga/effects';

import request from 'utils/request';
import {makeSelectPassword, makeSelectUsername, makeSelectPasswordConfirm, makeSelectEmail} from 'containers/SignUpPage/selectors';
import {SIGNUP} from "./constants";
import {signupSuccess, requestError} from "./actions";

/**
 * Login request handler
 */
export function* signupRequest() {

  // Select username and password from store
  const username = yield select(makeSelectUsername());
  const password = yield select(makeSelectPassword());
  const passwordConfirm = yield select(makeSelectPasswordConfirm());
  const email = yield select(makeSelectEmail());
  const requestURL = `http://localhost:1339`;

  try {
    // Call our request helper (see 'utils/request')
    const response = yield call(request, requestURL, {username: username, password: password, password2: passwordConfirm, email: email});
    yield put(signupSuccess(response));
  } catch (err) {
    yield put(requestError(err));
  }
}

/**
 * Root saga manages watcher lifecycle
 */
export default function* checkSignupState() {
  yield takeLatest(SIGNUP, signupRequest);
}
