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

import ProjectsList from 'containers/ProjectsList';
import injectSaga from 'utils/injectSaga';

import saga from './saga';

export class Main extends React.Component { // eslint-disable-line react/prefer-stateless-function

  render() {
    return (
      <div>
      <ProjectsList />
      </div>
    );
  }
}
const withSaga = injectSaga({ key: 'main', saga });

export default compose(
  withSaga,
)(Main);
