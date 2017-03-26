import React, {Component} from 'react';

class MapMarker extends Component {
  render() {

  }
}

class MapMarkerContainer extends Component {
  constructor() {
    super();
  }

  render() {
    return (
      <MapMarker {...this.props}></MapMarker>
    )
  }
}

MapMarker.propTypes = {
  location: React.PropTypes.object.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired
}

MapMarkerContainer.propTypes = {
  location: React.PropTypes.object.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired
}

export default MapMarkerContainer;
