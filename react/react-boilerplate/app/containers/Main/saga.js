/**
 * Gets the repositories of the user from Github
 */

import { call, put, takeLatest } from 'redux-saga/effects';

import request from 'utils/request';
import {SESSIONID} from "../App/constants";
import {LOAD_PROJECTS} from "./constants";
import {LOAD_PROJECTS2} from "./constants";
import { push } from 'react-router-redux';
import loadProjects2 from './actions';

/**
 * Github repos request/response handler
 */
export function* getProjects() {
console.log('sadasdasd');
  const requestURL = `http://localhost:1337`;
  const sessionID = sessionStorage.getItem(SESSIONID);
  const requestData = {
    session: {
      id :sessionID
    }
  };
    try {
      // Call our request helper (see 'utils/request')
      const response = yield call(request, requestURL, requestData);
      console.log(response);
      if(response.code == 0)
      {
        yield put({
          type:LOAD_PROJECTS2,
          projects:response.data
        });
      }
      else
      {
          sessionStorage.removeItem(SESSIONID);
          yield put(push('/login'));
      }
    } catch (err) {
    }
}
/**
 * Root saga manages watcher lifecycle
 */
export default function* checkLoginState() {
  const sessionID = sessionStorage.getItem(SESSIONID);
  if(sessionID == null)
  {
    yield put(push('/login'));
  }

  yield takeLatest(LOAD_PROJECTS, getProjects);
}
