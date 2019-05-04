import React, { Component } from "react";
import "./messageHistory.scss";
import Message from "../message/Message";

class MessageHistory extends Component {

  constructor(props) {
    super(props)
  }
  render () {
    const messages = this.props.messages.map( message => <Message message={message.data} />);


    return (
      <div className="message-history">
        <h2>Messages </h2>
        {messages}
      </div>
    );
  }
}

export default MessageHistory;
