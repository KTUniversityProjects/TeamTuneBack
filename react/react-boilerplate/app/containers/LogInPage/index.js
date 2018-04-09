/*
 * FeaturePage
 *
 * List all the features
 */
import React from 'react';
import PropTypes from 'prop-types';
import { Helmet } from 'react-helmet';
import { FormattedMessage } from 'react-intl';
import { connect } from 'react-redux';
import { compose } from 'redux';
import { createStructuredSelector } from 'reselect';

import injectReducer from 'utils/injectReducer';
import injectSaga from 'utils/injectSaga';
import { login } from './actions';
import { changeUsername } from './actions';
import { changePassword } from './actions';
import { makeSelectUsername } from './selectors';
import { makeSelectPassword } from './selectors';

import reducer from './reducer';
import saga from './saga';

import Form from 'components/Form';
import Input from 'components/Input';
import H2 from 'components/H2';
import messages from './messages';
import CenteredSection from './CenteredSection';
import Button from 'components/Button';

export class LogInPage extends React.Component { // eslint-disable-line react/prefer-stateless-function

  render() {
    return (
      <div>
        <Helmet>
          <title>FeaturePage</title>
          <meta name="description" content="Log in - TeamTune" />
        </Helmet>
        <CenteredSection>
        <H2>
          <FormattedMessage {...messages.header} />
        </H2>
        <Form onSubmit={this.props.onSubmitForm}>
                <Input
                  id="username"
                  type="text"
                  placeholder="Username"
                  value={this.props.username}
                  onChange={this.props.onChangeUsername}
                /><br />
                <Input
                  id="password"
                  type="password"
                  placeholder="Password"
                  value={this.props.password}
                  onChange={this.props.onChangePassword}
                /><br />
                <Button
                  id="submit"
                  type="submit"
                  children="Login"
                  onClick={this.props.onSubmitForm}
                />
            </Form>
          </CenteredSection>
      </div>
    );
  }
}
LogInPage.propTypes = {
  onSubmitForm: PropTypes.func,
  username: PropTypes.string,
  password: PropTypes.string,
  onChangeUsername: PropTypes.func,
  onChangePassword: PropTypes.func,
};

export function mapDispatchToProps(dispatch) {
  return {
    onChangeUsername: (evt) => dispatch(changeUsername(evt.target.value)),
    onChangePassword: (evt) => dispatch(changePassword(evt.target.value)),
    onSubmitForm: (evt) => {
      if (evt !== undefined && evt.preventDefault) evt.preventDefault();
      dispatch(login());

    },
  };
}

const mapStateToProps = createStructuredSelector({
  username: makeSelectUsername(),
  password: makeSelectPassword(),
});

const withConnect = connect(mapStateToProps, mapDispatchToProps);

const withReducer = injectReducer({ key: 'login', reducer });
const withSaga = injectSaga({ key: 'login', saga });

export default compose(
  withReducer,
  withSaga,
  withConnect,
)(LogInPage);
