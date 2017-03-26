import React, {Component} from 'react';
import {Link} from 'react-router-dom';

class Header extends Component {
  onClickLogin(e){
    e.preventDefault();
    const {login} = this.props;
    login();
  }
  onClickRegister(e){
    e.preventDefault();
    const {register} = this.props;
    register();
  }
  onClickLogout(e){
    e.preventDefault();
    const {logout} = this.props;
    logout();
  }

  render() {
    return (
      <nav className="navbar navbar-inverse navbar-fixed-top">
        <div className="container-fluid">
          <div className="navbar-header">
            <button type="button" className="navbar-toggle collapsed" data-toggle="collapse" data-target="#bs-example-navbar-collapse-1" aria-expanded="false">
              <span className="sr-only">Toggle navigation</span>
              <span className="icon-bar"></span>
              <span className="icon-bar"></span>
              <span className="icon-bar"></span>
            </button>
            <Link className="navbar-brand" to="/">Brand</Link>
          </div>
          <div className="collapse navbar-collapse" id="bs-example-navbar-collapse-1">
            <ul className="nav navbar-nav">
              <li className="active"><Link to="/browse">Browse <span className="sr-only">(current)</span></Link></li>
              <li><Link to="/user">User</Link></li>
              <li className="dropdown">
                <Link to="/" className="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">Dropdown <span className="caret"></span></Link>
                <ul className="dropdown-menu">
                  <li><Link to="/">Action</Link></li>
                  <li><Link to="/">Another action</Link></li>
                  <li><Link to="/">Something else here</Link></li>
                  <li role="separator" className="divider"></li>
                  <li><Link to="/">Separated link</Link></li>
                  <li role="separator" className="divider"></li>
                  <li><Link to="/">One more separated link</Link></li>
                </ul>
              </li>
            </ul>
            <form className="navbar-form navbar-left">
              <div className="form-group">
                <input type="text" className="form-control" placeholder="Search" />
              </div>
              <button type="submit" className="btn btn-default">Submit</button>
            </form>
            <ul className="nav navbar-nav navbar-right">
              <li onClick={this.onClickLogin.bind(this)}><a>Login</a></li>
              <li onClick={this.onClickRegister.bind(this)}><a>Register</a></li>
              <li onClick={this.onClickLogout.bind(this)}><a>Logout</a></li>
            </ul>
          </div>
        </div>
      </nav>
    )
  }
}

class HeaderContainer extends Component {
  constructor() {
    super()
  }

  render() {
    return (
      <Header {...this.props}></Header>
    )
  }
}

HeaderContainer.propTypes = {
  login: React.PropTypes.func.isRequired,
  logout: React.PropTypes.func.isRequired,
  register: React.PropTypes.func.isRequired,
  user: React.PropTypes.object.isRequired
}

Header.propTypes = {
  login: React.PropTypes.func.isRequired,
  logout: React.PropTypes.func.isRequired,
  register: React.PropTypes.func.isRequired,
  user: React.PropTypes.object.isRequired
}

export default HeaderContainer
