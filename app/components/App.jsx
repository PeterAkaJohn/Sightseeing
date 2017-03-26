import React, {Component} from 'react';
import { BrowserRouter as Router, Route, Link } from 'react-router-dom'
import BaseContainer from './layout/BaseLayout.jsx';

class App extends Component {
  render() {
    return(
      <Router>
          <Route path="/" component={BaseContainer}></Route>
      </Router>
    )
  }
}

export default App
