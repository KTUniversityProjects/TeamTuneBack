/**
 * The global state selectors
 */

 /**
  * Main selectors
  */

 import { createSelector } from 'reselect';

 const selectMain = (state) => state.get('projects');
 const makeSelectProjects = () => createSelector(
  selectMain,
   (mainState) => mainState.get('projects')
 );


 export {
   selectMain,
   makeSelectProjects,
 };
