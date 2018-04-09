/**
 * Gets the repositories of the user from Github
 */

import { call, put, select, takeLatest } from 'redux-saga/effects';

import request from 'utils/request';
import {SESSIONID} from "../App/constants";


/**
 * Root saga manages watcher lifecycle
 */
export default function* checkLoginState() {
  console.log('asd');
    const sessionID = sessionStorage.getItem(SESSIONID);
    console.log(sessionID);
      const requestURL = `http://localhost:1337`;
      const requestData = {
        session: {
          id :sessionID
        }
      };

        try {
          // Call our request helper (see 'utils/request')
          const response = yield call(request, requestURL, requestData);
          console.log(response);
  ///  yield put(signupSuccess(response));
        } catch (err) {
        //  yield put(requestError(err));
        }

  yield takeLatest(SESSIONID, checkLoginState);
}
