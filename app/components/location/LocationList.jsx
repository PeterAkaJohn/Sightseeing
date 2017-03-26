import React, {Component} from 'react';
import LocationContainer from './Location.jsx';

class LocationList extends Component {
  render() {
    return (
      <ul>
        {this.props.locations.map(location => {
          return <LocationContainer key={location.id} location={location} {...this.props}></LocationContainer>
        })}
      </ul>
    )
  }
}

class LocationListContainer extends Component {
  constructor() {
    super();
  }
  render(){
    return (
      <LocationList {...this.props}></LocationList>
    )
  }
}

LocationList.propTypes = {
  user: React.PropTypes.object.isRequired,
  locations: React.PropTypes.array.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired
}

LocationListContainer.propTypes = {
  user: React.PropTypes.object.isRequired,
  locations: React.PropTypes.array.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired
}

export default LocationListContainer;
