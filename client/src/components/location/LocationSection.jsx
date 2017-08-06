import React, {Component} from 'react';
import LocationListContainer from './LocationList.jsx';
import {Link} from 'react-router-dom';

class LocationSection extends Component {
  render() {
    return (
      <LocationListContainer {...this.props}></LocationListContainer>
    )
  }
}

class LocationSectionContainer extends Component {
  constructor(){
    super();
  }
  render(){
    return (
      <LocationSection {...this.props}></LocationSection>
    )
  }
}

LocationSection.propTypes = {
  user: React.PropTypes.object.isRequired,
  locations: React.PropTypes.array.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired
}

LocationSectionContainer.propTypes = {
  user: React.PropTypes.object.isRequired,
  locations: React.PropTypes.array.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired
}

export default LocationSectionContainer
