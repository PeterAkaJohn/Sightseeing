import React, {Component} from 'react';

class MapDescription extends Component {
  render() {
    return (
      <div>text MapDescription</div>
    )
  }
}

class MapDescriptionContainer extends Component {
  constructor() {
    super();
  }

  render() {
    return (
      <MapDescription {...this.props}></MapDescription>
    )
  }
}

MapDescription.propTypes = {
  user: React.PropTypes.object.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired,
  favoriteLocation: React.PropTypes.func.isRequired
}

MapDescriptionContainer.propTypes = {
  user: React.PropTypes.object.isRequired,
  setLocation: React.PropTypes.func.isRequired,
  activeLocation: React.PropTypes.object.isRequired,
  closeDescription: React.PropTypes.func.isRequired,
  favoriteLocation: React.PropTypes.func.isRequired
}

export default MapDescriptionContainer;
