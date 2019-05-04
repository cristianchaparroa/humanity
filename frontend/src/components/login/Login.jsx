import React, {Component} from "react";
import { Button, FormGroup, FormControl, FormLabel } from "react-bootstrap";
import "./login.scss";

import axios from 'axios';

import {withRouter} from 'react-router-dom'


class Login extends Component {

  constructor(props) {
    super(props)

    this.state = {
      email: "",
      password: ""
    };

    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange = event => {
   this.setState({
     [event.target.id]: event.target.value
   });
  }

  handleSubmit(event) {
    event.preventDefault();

    var self = this;
    var url = "http://localhost:8080/api/login";

    var payload={
      "email":this.state.email,
      "password":this.state.password
    }

    axios.post(url, payload)
      .then(function (response) {
        console.log(response);

        if(response.status === 200){
          console.log("login! ");

          self.props.history.push(`/rooms`)
        }
      }).catch(function (error) {
        console.log(error);
      });

  }

  render() {
    return (
      <div className="Login">

        <form onSubmit={this.handleSubmit}>

          <FormGroup controlId="email" >
            <FormLabel>Email</FormLabel>
            <FormControl
              autoFocus
              type="email"
              value={this.state.email}
              onChange={this.handleChange}
            />
          </FormGroup>

          <FormGroup controlId="password" >
            <FormLabel>Password</FormLabel>
            <FormControl
              value={this.state.password}
              onChange={this.handleChange}
              type="password"
            />
          </FormGroup>

          <Button
            block

            type="submit"
          >
            Login
          </Button>
        </form>
      </div>
    );
  }

}

export default Login;
