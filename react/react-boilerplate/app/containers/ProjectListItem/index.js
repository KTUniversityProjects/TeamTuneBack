/**
 * RepoListItem
 *
 * Lists the name and the issue count of a repository
 */

import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { createStructuredSelector } from 'reselect';
import { FormattedNumber } from 'react-intl';

import { makeSelectCurrentUser } from 'containers/App/selectors';
import ListItem from 'components/ListItem';
import ProjectLink from './ProjectLink';
import Wrapper from './Wrapper';

export default class ProjectListItem extends React.PureComponent { // eslint-disable-line react/prefer-stateless-function
  render() {
    const item = this.props.item;

    // Put together the content of the repository
    const content = (
      <Wrapper>
        <ProjectLink href={`/project/${item.id}`}>
          {item.name}
        </ProjectLink>
      </Wrapper>
    );

    // Render the content into a list item
    return (
      <ListItem key={`project-item-${item.id}`} item={content} />
    );
  }
}
ProjectListItem.propTypes = {
  item: PropTypes.object,
  currentUser: PropTypes.string,
};
