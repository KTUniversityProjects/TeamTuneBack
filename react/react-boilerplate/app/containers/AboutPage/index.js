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

export default class AboutPage extends React.Component { // eslint-disable-line react/prefer-stateless-function

  // Since state and props are static,
  // there's no need to re-render this component
  shouldComponentUpdate() {
    return false;
  }

  render() {
    return (
      <div>
        <Helmet>
          <title>Feature Page</title>
          <meta name="description" content="About us - TeamTune" />
        </Helmet>
        <CenteredSection>
        <H2>
          <FormattedMessage {...messages.header} />
        </H2>
          <FormattedMessage {...messages.scaffoldingMessage}/>
          </CenteredSection>
      </div>
    );
  }
}
