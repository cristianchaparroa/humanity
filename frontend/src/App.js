import React, {Component} from 'react';
import './App.css';

import {connect, sendMessage} from "./api";
import Header from "./components/header/Header";
import MessageHistory from "./components/messageHistory/MessageHistory";
import ChatInput from "./components/chatInput/ChatInput";


import * as uuid from "uuid";

class App extends Component {
   constructor(props) {
     super(props);
     console.log("App")
     console.log(props)

     this.state = {
       messages: [],
       userId : this.props.history.location.userId,
       nickname : this.props.history.location.nickname
     }

     this.send = this.send.bind(this);
   }

   componentDidMount() {

     let wsUrl = encodeURI('ws://localhost:8080/ws/room?userId=' + this.state.userId + '&nickname='+ this.state.nickname + '');
     console.log(wsUrl)
     this.connection = new WebSocket(wsUrl);

     let callback = (message) => {
       let ms = JSON.parse(message.data);
       this.setState(
         prevState => ({ messages: [...this.state.messages, ms] })
       )
      
     };

     connect(this.connection,callback);
   }

   send(event) {
     if(event.keyCode === 13) {

        let message = JSON.stringify({
          body: event.target.value,
          user_id: this.state.userId,
          nickname: this.state.nickname,
          id : uuid.v4()
        })

        sendMessage(this.connection,message);
        event.target.value = "";
     }

   }

   render() {
     return (
       <div className="App">
        <Header />
        <ChatInput send={this.send} />
        <MessageHistory messages={this.state.messages} />
       </div>
     );
   }
 }

export default App;
