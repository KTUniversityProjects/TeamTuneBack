/*
 * FeaturePage
 *
 * List all the features
 */
import React from 'react';
import PropTypes from 'prop-types';
import { Helmet } from 'react-helmet';
import { FormattedMessage } from 'react-intl';
import { compose } from 'redux';
import { createStructuredSelector } from 'reselect';
import { connect } from 'react-redux';

import ProjectsList from 'components/ProjectsList';
import injectReducer from 'utils/injectReducer';
import injectSaga from 'utils/injectSaga';
import { call, put } from 'redux-saga/effects';
import request from 'utils/request';
import {SESSIONID} from "../App/constants";

import {makeSelectProjects} from "./selectors";
import {makeSelectError} from "./selectors";
import {makeSelectLoading} from "./selectors";
//import reducer from './reducer';
import saga from './saga';
import reducer from './reducer';

import H2 from 'components/H2';
import CenteredSection from './CenteredSection';
import {LOAD_PROJECTS} from "./constants";

export class Main extends React.Component { // eslint-disable-line react/prefer-stateless-function

 componentDidMount() {
     this.props.onPageLoad();
 }

  render() {
    const { loading, error, projects } = this.props;
    const projectsListProps = {
      loading,
      error,
      projects,
    };

    return (
      <div>
      <ProjectsList {...projectsListProps} />
      </div>
    );
  }
}
Main.propTypes = {
  loading: PropTypes.bool,
  error: PropTypes.oneOfType([
    PropTypes.object,
    PropTypes.bool,
  ]),
  projects: PropTypes.oneOfType([
    PropTypes.array,
    PropTypes.bool,
  ]),
  onPageLoad:PropTypes.func,
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
  projects: makeSelectProjects(),
    error: makeSelectError(),
      loading: makeSelectLoading(),
});


const withConnect = connect(mapStateToProps, mapDispatchToProps);

const withSaga = injectSaga({ key: 'main', saga });
const withReducer = injectReducer({ key: 'main', reducer });

export default compose(
  withReducer,
  withSaga,
    withConnect,
)(Main);
