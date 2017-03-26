import React, {Component} from 'react';

class MapUser extends Component {
  render() {

  }
}

class MapUserContainer extends Component {
  constructor() {
    super();
  }

  render() {
    return (
      <MapUser {...this.props}></MapUser>
    )
  }
}

MapUser.propTypes = {
  user: React.PropTypes.object.isRequired
}

MapUserContainer.propTypes = {
  user: React.PropTypes.object.isRequired
}

export default MapUserContainer;
