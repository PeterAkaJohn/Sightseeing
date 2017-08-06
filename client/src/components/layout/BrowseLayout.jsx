//browse locations stored in db
import React, {Component} from 'react';
import LocationSectionContainer from '../location/LocationSection.jsx';

let locations = [
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


class BrowseLayout extends Component {
  render() {
    return (
      <div className="text-center">
        <header>Browse header</header>
          <div className="row">
            <div className="col-xs-12">
              <LocationSectionContainer {...this.props}></LocationSectionContainer>
            </div>
        </div>
        {this.props.user.name}
      </div>
    )
  }
}

//doesn't need nothing from previous route, just user object from BaseContainer
class BrowseContainer extends Component {
  constructor() {
    super();
    this.state = {
      locations: locations
    }
  }

  componentDidMount() {
    fetch("https://localhost:8443/browse")
      .then(response => response.json())
      .then(json => {
        console.log(json);
        this.setState({
          locations: json
        });
      }).catch(function(err){
        console.log(err);
      });
  }

  render() {
    return (
      <BrowseLayout user={this.props.user} locations={this.state.locations}></BrowseLayout>
    )
  }
}


export default BrowseContainer;
