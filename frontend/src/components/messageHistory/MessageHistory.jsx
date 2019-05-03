import React, { Component } from "react";
import "./messageHistory.scss";

class MessageHistory extends Component {

  constructor(props) {
    super(props)
  }
  render () {
    const messages = this.props.messages.map( (message, index) => (
        <p key={index}>{message.data} </p>
    ));

    return (
      <div className="messageHistory">
        <h2>Messages </h2>
        {messages}
      </div>
    );
  }
}

export default MessageHistory;
