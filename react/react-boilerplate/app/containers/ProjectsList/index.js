import React from 'react';
import PropTypes from 'prop-types';

import List from 'components/List';
import ListItem from 'components/ListItem';
import LoadingIndicator from 'components/LoadingIndicator';
import ProjectListItem from 'components/ProjectListItem';
import { connect } from 'react-redux';
import { createStructuredSelector } from 'reselect';
import reducer from './reducer';
import saga from './saga';
import { compose } from 'redux';
import injectReducer from 'utils/injectReducer';
import injectSaga from 'utils/injectSaga';
import {LOAD_PROJECTS} from './constants';

import {makeSelectProjects} from "./selectors";
//import reducer from './reducer';

export class ProjectsList extends React.Component {

   componentDidMount() {
       this.props.onPageLoad();
   }

  render() {
    const ErrorComponent = () => (
      <ListItem item={'Something went wrong, please try again!'} />
    );
    return <List component={ErrorComponent} />;

  }
}

ProjectsList.propTypes = {
  projects: PropTypes.any,
};

export function mapDispatchToProps(dispatch) {
  return {
    onPageLoad: (evt) => {
      dispatch({
        type: LOAD_PROJECTS
      });
    },
  };
}

const mapStateToProps = createStructuredSelector({
  projects: makeSelectProjects()
});


const withConnect = connect(mapStateToProps, mapDispatchToProps);

const withSaga = injectSaga({ key: 'projects', saga });
const withReducer = injectReducer({ key: 'projects', reducer });

export default compose(
  withReducer,
  withSaga,
  withConnect,
)(ProjectsList);
