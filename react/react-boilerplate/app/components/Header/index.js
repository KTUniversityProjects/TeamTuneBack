import React from 'react';
import { FormattedMessage } from 'react-intl';

import NavBar from './NavBar';
import Img from './Img';
import HeaderLink from './HeaderLink';
import Banner from './Banner.jpg';
import messages from './messages';

class Header extends React.Component { // eslint-disable-line react/prefer-stateless-function
  render() {
    return (
      <div>
      <Img src={Banner} alt="react-boilerplate - Logo" />
        <NavBar>
          <HeaderLink to="/">
            <FormattedMessage {...messages.home} />
          </HeaderLink>
          <HeaderLink to="/about">
            <FormattedMessage {...messages.about} />
          </HeaderLink>
          <HeaderLink to="/login">
            <FormattedMessage {...messages.login} />
          </HeaderLink>
          <HeaderLink to="/signup">
            <FormattedMessage {...messages.signup} />
          </HeaderLink>
        </NavBar>
      </div>
    );
  }
}

export default Header;
