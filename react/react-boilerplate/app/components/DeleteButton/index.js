import React, { Children } from 'react';
import PropTypes from 'prop-types';

import Wrapper from './Wrapper';
import A from './A';

function Button(props) {
  let button = (
    <A onClick={props.onClick}>
      {Children.toArray(props.children)}
    </A>
  );

  return (
    <Wrapper>
      {button}
    </Wrapper>
  );

}

Button.propTypes = {
  onClick: PropTypes.func,
  children: PropTypes.node.isRequired,
};

export default Button;
