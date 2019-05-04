import React, { Component } from "react";
import "./messageHistory.scss";
import Message from "../message/Message";

class MessageHistory extends Component {

  render () {
    const messages = this.props.messages.map( (message, index)  => <Message key={index} message={message} />);


    return (
      <div className="message-history">
        <h2>Messages </h2>
        {messages}
      </div>
    );
  }
}

export default MessageHistory;
