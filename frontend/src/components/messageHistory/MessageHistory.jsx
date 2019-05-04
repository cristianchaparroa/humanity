import React, { Component } from "react";
import "./messageHistory.scss";
import Message from "../message/Message";

class MessageHistory extends Component {

  render () {

    const messages = this.props.messages.map(
      message => <Message key={message.id} message={message} />
    ).reverse().filter((message,idx) => idx < 50);

    return (
      <div className="message-history">
        <h2>Messages </h2>
        {messages}
      </div>
    );
  }
}

export default MessageHistory;
