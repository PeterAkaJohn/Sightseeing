import React, {Component} from 'react';
import Location from './Location.jsx';

class LocationList extends Component {
  render() {

  }
}

LocationList.propTypes = {
  locations: React.PropTypes.array.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired
}

export default LocationList;
