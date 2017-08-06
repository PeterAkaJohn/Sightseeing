//add map, location list
import React, {Component} from 'react';
import {Link} from 'react-router-dom';
import LocationSectionContainer from '../location/LocationSection.jsx';
import MapSectionContainer from '../map/MapSection.jsx';

const config = {
    client_id: 'HKAKJCNXZCC3MJCK0THWOWX33RPE5CLZFIEYNDOHNMOIPNUU',
    client_secret: 'JNBCBVJLSE23FV5EKY5R1VKM3XBQVVHB2VDII3RZK3KF4JZ4',
    version: '20170101',
    apiUrl: 'https://api.foursquare.com/'
};

const options = {
  enableHighAccuracy: true
};

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
        <div className="row">
          <div className="col-xs-12">
            <MapSectionContainer {...this.props}></MapSectionContainer>
          </div>
        </div>
        <div>{this.props.user.username}</div>
      </div>
    );
  }
}

//locations, aware of user implement all the functions needed in MapSection, LocationSection
class HomeContainer extends Component {
  constructor() {
    super();
    this.state = {
      locations: [],
      activeLocation: {},
      currentPosition: null
    }
  }

  setLocation(location) {
    console.log("setLocation");
  }

  closeDescription(){
    console.log("closeDescription");
  }

  favoriteLocation(){
    console.log("favoriteLocation");
  }

  loadLocationsWithFourSquare(){
    if (this.state.locations.length = 0) {
      let {currentPosition} = this.state;
      fetch(config.apiUrl + 'v2/venues/explore' + '?v=' + config.version + '&client_id=' + config.client_id + '&client_secret=' + config.client_secret + '&ll='+ currentPosition.lat.toString()+ ','+ currentPosition.lng.toString() + '&limit=30&radius=10000')
      .then(response => response.json())
      .then(json => {
        console.log(json);
        let {locations} = this.state;
        let venues = json.response.groups[0].items;
        console.log(venues);
        for (let venueObject of venues) {
          let venue = venueObject.venue;
          let id = venue.id;
          let latitude = venue.location.lat;
          let longitude = venue.location.lng;
          let name = venue.name;

          let location = {
            venue,
            id,
            latitude,
            longitude,
            name
          }

          locations.push(location)
          this.setState({locations})

        }
      }).catch(function(err){
        console.log(err);
      });
    }
  }

  componentDidMount(){
    if (navigator.geolocation) {
          navigator.geolocation.getCurrentPosition((position) => {
            const pos = {
              lat: position.coords.latitude,
              lng: position.coords.longitude
            };
            this.setState({currentPosition: pos})
            this.loadLocationsWithFourSquare();
          }, () => {
            console.log('navigator disabled');
          }, options);
        } else {
          // Browser doesn't support Geolocation
          console.log('navigator disabled');
        }
  }

  render() {
    return (
      <HomeLayout locations={this.state.locations}
         currentPosition={this.state.currentPosition}
         user={this.props.user}
         activeLocation={this.state.activeLocation}
         setLocation={this.setLocation.bind(this)}
         closeDescription={this.closeDescription.bind(this)}
         favoriteLocation={this.favoriteLocation.bind(this)}></HomeLayout>
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
