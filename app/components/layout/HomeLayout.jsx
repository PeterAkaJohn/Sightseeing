//add map, location list
import React, {Component} from 'react';
import {Link} from 'react-router-dom';

class HomeLayout extends Component {
  render() {
    return (
      <div className="home text-center">
        <header>Home header</header>
        <div>Main Content, use the components to create the homepage</div>
        <div>
          <ul>
            {this.props.junk.map(junkElement => {
              return <li key={junkElement.id}><Link to={"/browse/"+ junkElement.id}>junkElement.name</Link></li>
            })}
          </ul>
      </div>
      <div>{this.props.user.name}</div>
      </div>
    );
  }
}

//locations, aware of user implement all the functions needed in MapSection, LocationSection
class HomeContainer extends Component {
  constructor() {
    super();
  }

  render() {
    return (
      <HomeLayout junk={this.props.junklist} user={this.props.user}></HomeLayout>
    )
  }
}

export default HomeContainer;
