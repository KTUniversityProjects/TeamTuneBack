import { css } from 'styled-components';

const buttonStyles = css`
  box-sizing: border-box;
  text-decoration: none;
  -webkit-font-smoothing: antialiased;
  -webkit-touch-callout: none;
  user-select: none;
  cursor: pointer;
  font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;
  font-weight: bold;
  font-size: 16px;
  color: #41addd;

  &:active {
    color: #fff;
  }
  &:onhover {
    color: #0000;
  }
`;

export default buttonStyles;
