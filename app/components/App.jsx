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
