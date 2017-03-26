//header, footer, authentication, login/logout, user object
import React, {Component} from 'react';
import HomeContainer from './HomeLayout.jsx';
import SearchLayout from './SearchLayout.jsx';
import BrowseContainer from './BrowseLayout.jsx';
import UserContainer from './UserLayout.jsx';
import LocationContainer from './LocationLayout.jsx';
import HeaderContainer from '../header/Header.jsx';
import {Switch, Link, Route, NoMatch} from 'react-router-dom';

var user = {
  name: "User"
}


class BaseLayout extends Component {

  constructor(props) {
    super(props);
    this.state = {
      user: user
    }
  }
  login(){
    console.log("Login");
  }

  register(){
    console.log("Register");
  }

  logout() {
    console.log("Logout");
  }

  render() {
    return (
      <div className="container-fluid">
        <HeaderContainer login={this.login.bind(this)} register={this.register.bind(this)} logout={this.logout.bind(this)} user={this.state.user}></HeaderContainer>
        <header className="text-center">BaseLayout header</header>
        <div className="row">
          <main className="col-xs-12">
            <Switch>
            <Route exact path="/" render={() => <HomeContainer {...this.props} />}></Route>
            <Route exact path="/browse" render={() => <BrowseContainer {...this.props} />}></Route>
            <Route exact path="/user" render={() => <UserContainer {...this.props} />}></Route>
            <Route exact path="/browse/:id" render={({match}) => <LocationContainer id={match.params.id} />}></Route>
            <Route component={NoMatch}/>
          </Switch>
        </main>
        </div>
        <div className="text-center">BaseLayout Footer</div>
      </div>
    );
  }
}

//base container only needs login/logout functionality, user object, can pass user object look above at / route
class BaseContainer extends Component {
  constructor() {
    super();
    this.state = { user }
  }

  componentDidMount() {
    // $.ajax({
    //   url: "/my-comments.json",
    //   dataType: 'json',
    //   success: function(comments) {
    //     this.setState({comments: comments});
    //   }.bind(this)
    // }); use fetch
  }
  //{...this.state} will be accessible to child component in this.props with the same name
  render() {
    return (
      <BaseLayout {...this.state}></BaseLayout>
    )
  }

}

export default BaseContainer;
