import React, {Component} from 'react';
import './App.css';

import {connect, sendMessage} from "./api";
import Header from "./components/header/Header";
import MessageHistory from "./components/messageHistory/MessageHistory";
import ChatInput from "./components/chatInput/ChatInput";



class App extends Component {
   constructor(props) {

     super(props);
     this.state = {
       messages: [],
     }
   }

   componentDidMount() {

     let callback = (message) => {
       console.log("new message")
       this.setState(
         prevState => ({ messages: [...this.state.messages, message] })
       )
       console.log(this.state);
     };

     connect(callback);
   }

   send(event) {
     if(event.keyCode === 13) {
       sendMessage(event.target.value);
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
