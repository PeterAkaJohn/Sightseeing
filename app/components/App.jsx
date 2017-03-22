import React, {Component} from 'react';
import { BrowserRouter as Router, Route, Link, Switch, NoMatch } from 'react-router-dom'
import BaseLayout from './layout/BaseLayout.jsx';
import HomeLayout from './layout/HomeLayout.jsx';
import SearchLayout from './layout/SearchLayout.jsx';
import BrowseLayout from './layout/BrowseLayout.jsx';


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

class App extends Component {
  render() {
    return(
      <Router>
          <Route path="/" component={BaseLayout}></Route>
      </Router>
    )
  }
}

// <div className='container-fluid'>
//   <div className='row'>
//     {junk.map(ju =>{
//       return <PlaceHolder junk={ju} key={ju.id}></PlaceHolder>
//     })
//   }</div>
// </div>

// class PlaceHolder extends Component {
//   render() {
//     const {junk} = this.props;
//     return (
//       <div className="col-xs-12">
//         <h1>{junk.name}</h1>
//       </div>
//     )
//   }
// }

export default App
