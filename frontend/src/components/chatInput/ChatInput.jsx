import React, { Component} from "react";
import "./chatInput.scss";

class ChatInput extends Component {
  render() {
    return (
      <div className="chat-input">
        <input placeholder="Write your message" onKeyDown={this.props.send}/>
      </div>

    );
  }
}


export default ChatInput;
