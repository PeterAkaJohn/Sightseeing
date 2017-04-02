//header, footer, authentication, login/logout, user object
import React, {Component} from 'react';
import HomeContainer from './HomeLayout.jsx';
import SearchLayout from './SearchLayout.jsx';
import BrowseContainer from './BrowseLayout.jsx';
import UserContainer from './UserLayout.jsx';
import LocationContainer from './LocationLayout.jsx';
import HeaderContainer from '../header/Header.jsx';
import {Switch, Link, Route} from 'react-router-dom';

class BaseLayout extends Component {

  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div className="container-fluid">
        <HeaderContainer login={this.props.login.bind(this)} register={this.props.register.bind(this)} logout={this.props.logout.bind(this)} user={this.props.user}></HeaderContainer>
        <header className="text-center">BaseLayout header</header>
        <div className="row">
          <main className="col-xs-12">
            <Switch>
            <Route exact path="/" render={() => <HomeContainer {...this.props} {...this.state}/>}></Route>
            <Route exact path="/browse" render={() => <BrowseContainer {...this.props} {...this.state}/>}></Route>
            <Route exact path="/user" render={() => <UserContainer {...this.props} {...this.state}/>}></Route>
            <Route exact path="/browse/:id" render={({match}) => <LocationContainer id={match.params.id} {...this.state}/>}></Route>
          </Switch>
        </main>
        </div>
        <div className="text-center">BaseLayout Footer</div>
      </div>
    );
  }
}

class BaseContainer extends Component {
  constructor(props) {
    super(props);
    this.state = {
      user: { username: ''}
    }
    this.service = null;
    this.map = null;
  }

  login(user){
    console.log(user);
    fetch("https://localhost:8443/login", {
      method: "post",
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },

  //make sure to serialize your JSON body
    body: JSON.stringify({
      username: user.username,
      password: user.password
      })
    })
    .then( response => response.json())
    .then(json => {
      this.setState({user: json});
      console.log(this.state.user);
    })
    .catch(err => console.log(err));
  }

  register(user){
    console.log(user);
    fetch("https://localhost:8443/register", {
      method: "post",
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },

  //make sure to serialize your JSON body
    body: JSON.stringify({
      username: user.username,
      password: user.password
      })
    })
    .then( (response) => {
      console.log("Registered");
    }).catch(err => console.log(err));
  }

  logout() {
    console.log("Logout");
    this.setState({user: {}})
  }

  componentDidMount() {
    //load google maps and places api
  }

  //{...this.state} will be accessible to child component in this.props with the same name
  render() {
    return (
      <BaseLayout {...this.state}
        login={this.login.bind(this)}
        logout={this.logout.bind(this)}
        register={this.register.bind(this)}></BaseLayout>
    )
  }

}

export default BaseContainer;
