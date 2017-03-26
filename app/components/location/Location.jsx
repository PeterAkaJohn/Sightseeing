import React, {Component} from 'react';
import {Link} from 'react-router-dom';

class Location extends Component {
  render() {
    return (
      <li>
        <Link to={"/browse/"+ this.props.location.id}>{this.props.location.name}</Link>
      </li>
    )
  }
}

class LocationContainer extends Component {
  constructor() {
    super();
  }

  render() {
    return (
      <Location {...this.props}></Location>
    )
  }
}

Location.propTypes = {
  user: React.PropTypes.object.isRequired,
  location: React.PropTypes.object.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired
}

LocationContainer.propTypes = {
  user: React.PropTypes.object.isRequired,
  location: React.PropTypes.object.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired
}


export default LocationContainer
