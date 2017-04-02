//user page with favorites and a brief bio if provided
import React, { Component} from 'react'

class UserLayout extends Component {
  render() {
    return (
      <div className="home text-center">
        <header>User header</header>
        <div>User content, favorites</div>
        <div>{this.props.user.username}</div>
      </div>
    )
  }
}

//needs user object from baselayout, implment api calls for the user (retrieve favorite locations and user informations)
class UserContainer extends Component {
  constructor(props) {
    super(props);
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

  render() {
    return (
      <UserLayout user={this.props.user}></UserLayout>
    )
  }
}

export default UserContainer
