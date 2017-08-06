import React, { Component } from "react";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import BaseContainer from "./layout/BaseLayout.jsx";
import "bootstrap/dist/css/bootstrap.css";
import "../styles/style.css";

class App extends Component {
  render() {
    return (
      <Router>
        <Route path="/" component={BaseContainer} />
      </Router>
    );
  }
}

export default App;
