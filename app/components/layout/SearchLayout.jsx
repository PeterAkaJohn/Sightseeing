//search locations by city
import React, {Component} from 'react';

class SearchLayout extends Component {
  render() {
    return (
      <div className="search">
        <header>Search Header</header>
        <div>{this.props.children}</div>
      </div>
    )
  }
}

export default SearchLayout;
