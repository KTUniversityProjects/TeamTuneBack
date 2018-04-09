/**
 * Gets the repositories of the user from Github
 */

import { call, put, select, takeLatest } from 'redux-saga/effects';

import request from 'utils/request';
import {makeSelectPassword, makeSelectUsername} from 'containers/LogInPage/selectors';
import {LOGIN} from "./constants";
import {loginSuccess, requestError} from "./actions";
import {SESSIONID} from "../App/constants";

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

    const sessionID = sessionStorage.getItem(SESSIONID);

      const requestURL = `http://localhost:1337`;
      const requestData = {
        session {
          id :sessionID
        }
      };
      
        try {
          // Call our request helper (see 'utils/request')
          const response = yield call(request, requestURL, requestData);
          yield put(signupSuccess(response));
        } catch (err) {
          yield put(requestError(err));
        }

  yield takeLatest(LOGIN, loginRequest);
}