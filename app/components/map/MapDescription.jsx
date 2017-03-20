import React, {Component} from 'react';

class MapDescription extends Component {
  render() {

  }
}

MapDescription.propTypes = {
  user: React.PropTypes.object.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired,
  favoriteLocation: React.PropTypes.func.isRequired
}

export default MapDescription;
