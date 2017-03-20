import React, {Component} from 'react';

class Header extends Component {
  render() {

  }
}

Header.propTypes = {
  login: React.PropTypes.func.isRequired,
  logout: React.PropTypes.func.isRequired,
  register: React.PropTypes.func.isRequired,
  user: React.PropTypes.object.isRequired
}

export default Header
