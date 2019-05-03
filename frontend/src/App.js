import React, {Component} from 'react';
import './App.css';

import {connect, sendMessage} from "./api";
import Header from "./components/header/Header";
import MessageHistory from "./components/messageHistory/MessageHistory";

 class App extends Component {
   constructor(props) {

     super(props);
     this.state = {
       messages: [],
     }
     connect();
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

   send() {
     console.log("hello");
     sendMessage("hello");
   }

   render() {
     return (
       <div className="App">
        <Header />
        <MessageHistory messages={this.state.messages} />
        <button onClick={this.send}>Hit</button>
       </div>
     );
   }
 }

export default App;
