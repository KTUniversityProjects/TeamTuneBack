/**
 * Gets the repositories of the user from Github
 */

import { call, put, select, takeLatest } from 'redux-saga/effects';

import request from 'utils/request';
import {makeSelectPassword, makeSelectUsername} from 'containers/LogInPage/selectors';
import {LOGIN} from "./constants";
import { requestError} from "./actions";
import { push } from 'react-router-redux';

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

    if(response.code == 0)
    {
      sessionStorage.setItem('sessionID', response.data);
      yield put(push('/signup'));
    }
    else
    {
      //webservices/core/Responses.go
    }
  } catch (err) {
    yield put(requestError(err));
  }
}

/**
 * Root saga manages watcher lifecycle
 */
export default function* checkLoginState() {
  console.log('asdxxx');
  yield takeLatest(LOGIN, loginRequest);
}
