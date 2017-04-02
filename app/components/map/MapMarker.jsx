import React, {Component} from 'react';

class MapMarker extends Component {
  render() {
    return (
      false
    )
  }
}

class MapMarkerContainer extends Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {
      let {map, location, bounds, registerMarker} = this.props;
      this.registerMarker(location)
  }

  registerMarker(location){
    let {bounds, map} = this.props
    let position = new google.maps.LatLng(location.latitude, location.longitude);
    bounds.extend(position);
    let marker = new google.maps.Marker({
      position: position,
      map: map,
      title: 'Hello World!'
    });
    map.fitBounds(bounds);
    console.log("here");
  }

  render() {
    return (
      <MapMarker {...this.props}></MapMarker>
    )
  }
}

MapMarker.propTypes = {
  location: React.PropTypes.object.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired
}

MapMarkerContainer.propTypes = {
  location: React.PropTypes.object.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired
}

export default MapMarkerContainer;
