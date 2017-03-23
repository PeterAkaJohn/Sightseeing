// location detail, needs id from url to make an api call to the server
import React, {Component} from 'react';

class LocationLayout extends Component {
  render() {
    return (
      <div>
        {this.props.location.name}
        <br></br>{this.props.id}
      </div>
    )
  }
}

//api calls need id from url parameter
class LocationContainer extends Component {
  constructor() {
    super();
    this.state = { location: {name: "location"}}
  }

  render() {
    return (
      <LocationLayout location={this.state.location} id={this.props.id}></LocationLayout>
    )
  }
}

export default LocationContainer
