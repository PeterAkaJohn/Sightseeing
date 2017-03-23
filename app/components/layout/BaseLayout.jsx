//header, footer, authentication, login/logout, user object
import React, {Component} from 'react';
import HomeContainer from './HomeLayout.jsx';
import SearchLayout from './SearchLayout.jsx';
import BrowseContainer from './BrowseLayout.jsx';
import UserContainer from './UserLayout.jsx';
import LocationContainer from './LocationLayout.jsx';
import Header from '../header/Header.jsx';
import {Switch, Link, Route, NoMatch} from 'react-router-dom';

var user = {
  name: "User"
}
var junk = [
  {
    id: 1,
    name: "Bob"
  },
  {
    id: 2,
    name: "Bob"
  },
  {
    id: 3,
    name: "Bob"
  },
  {
    id: 4,
    name: "Bob"
  }
]


class BaseLayout extends Component {

  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div className="container-fluid">
        <Header></Header>
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
    this.state = { junklist: junk, user }
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

export default BaseContainer
