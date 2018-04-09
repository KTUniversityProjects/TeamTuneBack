import {
  CHANGE_USERNAME,
  CHANGE_PASSWORD,
} from './constants';

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
