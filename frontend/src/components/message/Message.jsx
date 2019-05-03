import React, {Component} from "react";
import "./message.scss";


class Message extends Component {

  constructor(props) {
    super(props);
    let msg = JSON.parse(this.props.message);
    this.state = {
      message: msg
    };
  }

  render() {
    return <div className="message">{this.state.message.body}</div>;
  }
}

export default Message;
