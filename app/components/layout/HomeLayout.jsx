//add map, location list
import React, {Component} from 'react';
import {Link} from 'react-router-dom';
import LocationSectionContainer from '../location/LocationSection.jsx';

var locations = [
  {
    id:1,
    name:"example"
  },
  {
    id:2,
    name:"example"
  },
  {
    id:3,
    name:"example"
  },
  {
    id:4,
    name:"example"
  },
  {
    id:5,
    name:"example"
  },
  {
    id:6,
    name:"example"
  },
  {
    id:7,
    name:"example"
  }
]

class HomeLayout extends Component {
  render() {
    return (
      <div className="home text-center">
        <header>Home header</header>
        <div>Main Content, use the components to create the homepage</div>
        <div className="row">
          <div className="col-xs-12">
            <LocationSectionContainer {...this.props}></LocationSectionContainer>
          </div>
      </div>
      <div>{this.props.user.name}</div>
      </div>
    );
  }
}

//locations, aware of user implement all the functions needed in MapSection, LocationSection
class HomeContainer extends Component {
  constructor() {
    super();
    this.state = {
      locations: locations,
      activeLocation: locations[0]
    }
  }

  setLocation(location) {
    console.log("setLocation");
  }

  closeDescription(){
    console.log("closeDescription");
  }

  render() {
    return (
      <HomeLayout locations={this.state.locations}
         user={this.props.user}
         activeLocation={this.state.activeLocation}
         setLocation={this.setLocation.bind(this)}
         closeDescription={this.closeDescription.bind(this)}></HomeLayout>
    )
  }
}

HomeContainer.propTypes = {
  user: React.PropTypes.object.isRequired
}
HomeLayout.propTypes = {
  user: React.PropTypes.object.isRequired,
  locations: React.PropTypes.array.isRequired
}

export default HomeContainer;
