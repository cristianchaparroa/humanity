import React, {Component} from "react";
import "./message.scss";
import Moment from 'react-moment';

class Message extends Component {

  constructor(props) {
    super(props);
    this.state = {
      message: this.props.message
    };
  }

  render() {
    return <div className="message">
        <div>{this.state.message.nickname}</div>
        <div>{this.state.message.body}</div>
        <Moment format="YYYY/MM/DD HH:mm:ss" className="message-time">
          {this.state.message.time}
        </Moment>
    </div>;
  }
}

export default Message;
