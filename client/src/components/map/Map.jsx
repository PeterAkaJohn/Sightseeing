import React, { Component } from "react";
import MapMarkerContainer from "./MapMarker.jsx";
import scriptLoader from "react-async-script-loader";

class Map extends Component {
  render() {
    return (
      <div id="map" style={{ height: "100vh", width: "100vw" }}>
        {this.props.locations.map(location => {
          if (this.props.map != null && this.props.bounds != null) {
            return (
              <MapMarkerContainer
                key={location.id}
                location={location}
                {...this.props}
              />
            );
          }
        })}
      </div>
    );
  }
}

@scriptLoader([
  "https://maps.googleapis.com/maps/api/js?key=AIzaSyDSSIeGc4in6qVsJE6BwSSLBS5RlqkHRm4"
])
class MapContainer extends Component {
  constructor() {
    super();
    this.map = null;
    this.bounds = null;
  }

  componentWillReceiveProps({ isScriptLoaded, isScriptLoadSucceed }) {
    if (isScriptLoaded && !this.props.isScriptLoaded) {
      // load finished
      if (isScriptLoadSucceed) {
        let { currentPosition } = this.props;
        this.bounds = new window.google.maps.LatLngBounds();
        this.map = new window.google.maps.Map(document.getElementById("map"), {
          center: currentPosition,
          zoom: 10
        });
      } else {
        console.log("error not success");
      }
    } else {
      console.log("not loaded");
    }
  }

  render() {
    return (
      <Map
        {...this.props}
        {...this.state}
        map={this.map}
        bounds={this.bounds}
      />
    );
  }
}

Map.propTypes = {
  locations: React.PropTypes.array.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired,
  favoriteLocation: React.PropTypes.func.isRequired,
  user: React.PropTypes.object.isRequired
};

MapContainer.propTypes = {
  locations: React.PropTypes.array.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired,
  favoriteLocation: React.PropTypes.func.isRequired,
  user: React.PropTypes.object.isRequired
};

export default MapContainer;
