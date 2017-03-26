import React, {Component} from 'react';
import MapMarkerContainer from './MapMarker.jsx';
import MapDescriptionContainer from './MapDescription.jsx';
import MapUserContainer from './MapUser.jsx';

class Map extends Component {
  render() {
    return (
      <MapDescriptionContainer></MapDescriptionContainer>
      <MapUserContainer></MapUserContainer>
      <MapMarkerContainer></MapMarkerContainer>
    )
  }

}

class MapContainer extends Component {
  constructor() {
    super();
  }

  render() {
    return (
      <Map {...this.props}></Map>
    )
  }
}

Map.propTypes = {
  locations: React.PropTypes.array.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired,
  favoriteLocation: React.PropTypes.func.isRequired,
  user: React.PropTypes.object.isRequired
}

MapContainer.propTypes = {
  locations: React.PropTypes.array.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired,
  favoriteLocation: React.PropTypes.func.isRequired,
  user: React.PropTypes.object.isRequired
}


export default MapContainer
