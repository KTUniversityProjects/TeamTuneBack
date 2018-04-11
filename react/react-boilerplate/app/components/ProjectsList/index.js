import React from 'react';
import PropTypes from 'prop-types';

import List from 'components/List';
import ListItem from 'components/ListItem';
import LoadingIndicator from 'components/LoadingIndicator';
import ProjectListItem from 'containers/ProjectListItem';

function ProjectsList({ loading, error, projects }) {
  if (loading) {
    return <List component={LoadingIndicator} />;
  }

    if (projects) {
      return <List items={projects} component={ProjectListItem} />;
    }

    const ErrorComponent = () => (
      <ListItem item={'Something went wrong, please try again!'} />
    );
    return <List component={ErrorComponent} />;

}

ProjectsList.propTypes = {
  loading: PropTypes.bool,
  error: PropTypes.any,
  projects: PropTypes.any,
};

export default ProjectsList;
