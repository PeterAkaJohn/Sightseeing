//user page with favorites and a brief bio if provided
import React, { Component} from 'react'

class UserLayout extends Component {
  render() {
    return (
      <div className="home">
        <header>User header</header>
        <div>User content, favorites</div>
      </div>
    )
  }
}

export default UserLayout
