import React, {Component} from 'react';
import MapMarker from './MapMarker.jsx';
import MapDescription from './MapDescription.jsx';
import MapUser from './MapUser.jsx';

class Map extends Component {
  constructor() {

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

export default Map
