import React, {Component} from 'react';

class MapMarker extends Component {
  render() {

  }
}

MapMarker.propTypes = {
  location: React.PropTypes.object.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired
}

export default MapMarker
