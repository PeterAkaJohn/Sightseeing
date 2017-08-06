import React, {Component} from 'react';
import {Link} from 'react-router-dom';
import {Modal, Button} from 'react-bootstrap';

class Header extends Component {
  constructor(props){
    super(props);
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
              <li><a>{this.props.user.username}</a></li>
              <li onClick={this.props.openLogin.bind(this)}><a>Login</a></li>
              <li onClick={this.props.openRegister.bind(this)}><a>Register</a></li>
              <li onClick={this.props.onClickLogout.bind(this)}><a>Logout</a></li>
            </ul>
          </div>
        </div>
        <Modal show={this.props.showLogin} onHide={this.props.closeLogin.bind(this)}>
          <Modal.Header closeButton>
            <Modal.Title>Modal heading Login</Modal.Title>
          </Modal.Header>
          <Modal.Body>
            <form onSubmit={this.props.handleLogin.bind(this)}>
              <div className="form-group">
                <label htmlFor="username">Username:</label>
                <input className="form-control" value={this.props.username} onChange={this.props.handleChangeUsername.bind(this)} id="username" type="text" />
              </div>
              <div className="form-group">
                  <label htmlFor="password">Password:</label>
                  <input className="form-control" value={this.props.password} onChange={this.props.handleChangePassword.bind(this)} id="password" type="text" />
              </div>
              <Button className="btn btn-success" type="submit">Login</Button>
            </form>
          </Modal.Body>
          <Modal.Footer>
            <Button onClick={this.props.closeLogin.bind(this)}>Close</Button>
          </Modal.Footer>
        </Modal>
        <Modal show={this.props.showRegister} onHide={this.props.closeRegister.bind(this)}>
          <Modal.Header closeButton>
            <Modal.Title>Modal heading Register</Modal.Title>
          </Modal.Header>
          <Modal.Body>
            <form onSubmit={this.props.handleRegister.bind(this)}>
              <div className="form-group">
                <label htmlFor="username">Username:</label>
                <input className="form-control" value={this.props.username} onChange={this.props.handleChangeUsername.bind(this)} id="username" type="text" />
              </div>
              <div className="form-group">
                  <label htmlFor="password">Password:</label>
                  <input className="form-control" value={this.props.password} onChange={this.props.handleChangePassword.bind(this)} id="password" type="text" />
              </div>
              <Button className="btn btn-success" type="submit">Register</Button>
            </form>
          </Modal.Body>
          <Modal.Footer>
            <Button onClick={this.props.closeRegister.bind(this)}>Close</Button>
          </Modal.Footer>
        </Modal>
      </nav>
    )
  }
}

class HeaderContainer extends Component {
  constructor() {
    super();
    this.state = {
      showRegister: false,
      showLogin: false,
      username: '',
      password: ''
    }
  }

  onClickLogout(e){
    e.preventDefault();
    const {logout} = this.props;
    logout();
  }

  handleLogin(event){
    const {login} = this.props;
    login({username: this.state.username, password:this.state.password});
    this.setState({username: '', password: '', showLogin: false})
    event.preventDefault();
  }

  handleRegister(event){
    const {register} = this.props;
    register({username: this.state.username, password:this.state.password});
    this.setState({username: '', password: '', showRegister: false})
    event.preventDefault();
  }

  handleChangeUsername(event){
    this.setState({username: event.target.value});
  }

  handleChangePassword(event){
    this.setState({password: event.target.value});
  }

  closeLogin(e) {
    e.preventDefault();
    this.setState({ showLogin: false })
  }

  openLogin(e) {
    e.preventDefault();
    this.setState({showLogin: true})
  }

  closeRegister(e) {
    e.preventDefault();
    this.setState({
      showRegister: false
    })
  }

  openRegister(e) {
    e.preventDefault();
    this.setState({
      showRegister: true
    })
  }

  render() {
    return (
      <Header {...this.props} {...this.state}
        onClickLogout={this.onClickLogout.bind(this)}
        handleLogin={this.handleLogin.bind(this)}
        handleRegister={this.handleRegister.bind(this)}
        handleChangeUsername={this.handleChangeUsername.bind(this)}
        handleChangePassword={this.handleChangePassword.bind(this)}
        closeLogin={this.closeLogin.bind(this)}
        openLogin={this.openLogin.bind(this)}
        closeRegister={this.closeRegister.bind(this)}
        openRegister={this.openRegister.bind(this)}></Header>
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
  user: React.PropTypes.object.isRequired,
  showLogin: React.PropTypes.bool.isRequired,
  showRegister: React.PropTypes.bool.isRequired,
  openLogin: React.PropTypes.func.isRequired,
  closeLogin: React.PropTypes.func.isRequired,
  openRegister: React.PropTypes.func.isRequired,
  closeRegister: React.PropTypes.func.isRequired
}

export default HeaderContainer
