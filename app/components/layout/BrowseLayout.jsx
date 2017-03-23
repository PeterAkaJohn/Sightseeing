//browse locations stored in db
import React, {Component} from 'react';

class BrowseLayout extends Component {
  render() {
    return (
      <div className="text-center">
        <header>Browse header</header>
        <div>Browse Content</div>
        {this.props.user.name}
      </div>
    )
  }
}

//doesn't need nothing from previous route, just user object from BaseContainer
class BrowseContainer extends Component {
  constructor() {
    super();
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
      <BrowseLayout user={this.props.user}></BrowseLayout>
    )
  }
}


export default BrowseContainer;
