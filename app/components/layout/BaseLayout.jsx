//header, footer
import React, {Component} from 'react';
import HomeLayout from './HomeLayout.jsx';
import SearchLayout from './SearchLayout.jsx';
import BrowseLayout from './BrowseLayout.jsx';
import UserLayout from './UserLayout.jsx';
import {Switch, Link, Route, NoMatch} from 'react-router-dom';

class BaseLayout extends Component {
  render() {
    return (
      <div className="container-fluid">
        <header className="text-center">BaseLayout Header</header>
        <div className="row">
          <aside className="col-xs-4">
            <ul>
              <li><Link to="/">Home</Link></li>
              <li><Link to="/browse">Browse</Link></li>
              <li><Link to="/user">User</Link></li>
            </ul>
          </aside>
          <main className="col-xs-8">
            <Switch>
            <Route exact path="/" component={HomeLayout}></Route>
            <Route exact path="/browse" component={BrowseLayout}></Route>
            <Route exact path="/user" component={UserLayout}></Route>
            <Route component={NoMatch}/>
          </Switch>
        </main>
        </div>
        <div className="text-center">BaseLayout Footer</div>
      </div>
    );
  }
}

export default BaseLayout
