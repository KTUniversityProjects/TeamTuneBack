/**
 * Gets the repositories of the user from Github
 */

import { put } from 'redux-saga/effects';
import { push } from 'react-router-redux';
import request from 'utils/request';
import {SESSIONID} from "../App/constants";

/**
 * Root saga manages watcher lifecycle
 */
export default function* checkLoginState() {
  const sessionID = sessionStorage.getItem(SESSIONID);
  if(sessionID == null)
  {
       yield put(push('/'));
  }
}
