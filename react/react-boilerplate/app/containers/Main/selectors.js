/**
 * The global state selectors
 */

 /**
  * Main selectors
  */

 import { createSelector } from 'reselect';

 const selectMain = (state) => state.get('main');
 const makeSelectLoading = () => createSelector(
   selectMain,
   (mainState) => mainState.get('loading')
 );
 const makeSelectError = () => createSelector(
   selectMain,
   (mainState) => mainState.get('error')
 );
 const makeSelectProjects = () => createSelector(
  selectMain,
   (mainState) => mainState.get('projects')
 );


 export {
   selectMain,
   makeSelectLoading,
   makeSelectError,
   makeSelectProjects,
 };
