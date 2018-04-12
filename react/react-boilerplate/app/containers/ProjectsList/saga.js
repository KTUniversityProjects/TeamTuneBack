/**
 * Gets the repositories of the user from Github
 */

import { call, put, takeLatest } from 'redux-saga/effects';

import request from 'utils/request';
import {SESSIONID} from "../App/constants";
import {LOAD_PROJECTS} from "./constants";
import {LOAD_PROJECTS2} from "./constants";
import loadProjects2 from './actions';

/**
 * Github repos request/response handler
 */
export function* getProjects() {
console.log('sadasdasd');
  const requestURL = `http://localhost:1338`;
  const sessionID = sessionStorage.getItem(SESSIONID);
  const requestData = {
    session: {
      id :sessionID
    }
  };
    try {
      // Call our request helper (see 'utils/request')
      const response = yield call(request,  requestURL, "POST", requestData);
      console.log(response);
      if(response.code == 0)
      {
        yield put({
          type:LOAD_PROJECTS2,
          projects:response.data
        });
      }
    } catch (err) {
      Console.log("Hujakst");
    }
}
/**
 * Root saga manages watcher lifecycle
 */
export default function* projectListInit() {
  yield takeLatest(LOAD_PROJECTS, getProjects);
}
