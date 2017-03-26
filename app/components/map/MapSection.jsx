import React, {Component} from 'react';
import MapContainer from './Map.jsx';


class MapSection extends Component {
  render() {
    return (
      <MapContainer {...this.props}></MapContainer>
    )
  }
}

class MapSectionContainer extends Component {
  constructor() {
    super();
  }

  render() {
    return (
      <MapSection {...this.props}></MapSection>
    )
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

MapSectionContainer.propTypes = {
  locations: React.PropTypes.array.isRequired,
  user: React.PropTypes.object.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired,
  favoriteLocation: React.PropTypes.func.isRequired
}

export default MapSectionContainer;
