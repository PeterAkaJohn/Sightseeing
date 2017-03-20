import React, {Component} from 'react';
import LocationList from './LocationList.jsx';

class LocationSection extends Component {
  render() {

  }
}

LocationSection.propTypes = {
  locations: React.PropTypes.array.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired
}

export default LocationSection
