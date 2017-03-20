import React, {Component} from 'react';
import Map from './MapSection.jsx';

class MapSection extends Component {
  render() {

  }
}

MapSection.propTypes = {
  locations: React.PropTypes.array.isRequired,
  user: React.PropTypes.object.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired,
  favoriteLocation: React.PropTypes.func.isRequired
}

export default MapSection;
