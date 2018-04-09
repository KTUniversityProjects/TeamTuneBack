/*
 * FeaturePage
 *
 * List all the features
 */
import React from 'react';
import { Helmet } from 'react-helmet';
import { FormattedMessage } from 'react-intl';

import H2 from 'components/H2';
import messages from './messages';
import CenteredSection from './CenteredSection';
import saga from './saga';
import injectSaga from 'utils/injectSaga';
import { compose } from 'redux';

export class Main extends React.Component { // eslint-disable-line react/prefer-stateless-function
  render() {
    return (
      <div>
      </div>
    );
  }
}

const withSaga = injectSaga({ key: 'loginMain', saga });

export default compose(
  withSaga,
)(Main);
