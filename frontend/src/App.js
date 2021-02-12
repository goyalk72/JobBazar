import React, { Component } from 'react';
import './App.css';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link,
  Redirect
} from "react-router-dom";

import axios from 'axios';
import { Form, Input, Button, Row, Col, Avatar, message } from 'antd';
import {MailOutlined, LockOutlined } from '@ant-design/icons'
import 'antd/dist/antd.css';
import Profile from "./component/profile/profile";
import Register from "./component/register/register";

class App extends Component{

  constructor(){
    super();
    this.state = {
      error:{
        status: false,
        msg: ""
      },
      email: "",
      mobile: "",
      password: "",
      reProfile: false,
      reRegister: false
    }
  }

  validate = (values) => {
    const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    const phoneNum = /^[789]{1}[0-9]{9}$/;
    if (re.test(String(values["loginToken"]).toLowerCase())){
      this.setState({
        email: values["loginToken"],
        mobile: "",
        password: values["password"],
        error:{
          status: false,
          msg: ""
        },
        
      });
    } else if (phoneNum.test(values["loginToken"])){
            this.setState({
              email: "",
              mobile: values["loginToken"],
              password: values["password"],
              error:{
                status: false,
                msg: ""
              }
            });
          } else {
            this.setState({
              error:{
                status: true,
                msg: "Enter correct email id or mobile"
              }
            });
            message.error(this.state.error.msg)
          }

  }

  renderRedirect = () => {
    if (this.state.reProfile){
      return (
      <React.Fragment>
        <Route path="/profile" component={Profile} />
        <Redirect to={{
          pathname: "/profile/userinfo",
          state: { 
            email: this.state.email,
            mobile: this.state.mobile,
            // password: this.state.password,
            // isAuth: true
          },
        }} />
      </React.Fragment>);
    } else if (this.state.reRegister){
      return (
        <React.Fragment>
          <Route exact path="/register" component={Register} />
          <Redirect to={{
            pathname: "/register",
            state: { 
              email: this.state.email,
              mobile: this.state.mobile,
            },
          }} />
        </React.Fragment>
        );
    }
  }

  responseHandler = (res) => {
    if (res.isNewUser) {
      this.setState({
        reProfile: false,
        reRegister: true
      });
      
    }
    else {
      if (res.isAuth){
        this.setState({
          reProfile: true,
          reRegister: false
        });
      } else{
        this.setState({
          error:{
            status: true,
            msg: "Invalid user login",
            reProfile: false,
            reRegister: false
          }
        });
        message.error(this.state.error.msg)
      } 
    }
  }

  onFinish = (values) => {
    this.validate(values);
    if (!this.state.error.status){
      axios.get("http://127.0.0.1:3030/verifyLogin", {
        params:{
          email: this.state.email,
          mobile: this.state.mobile,
          password: this.state.password
        }
      }).then(res => {
        this.responseHandler(res.data);
      })
      .catch(err => console.error(err));
    }
  };

  onFinishFailed = (errorInfo) => {
    console.error(errorInfo)
  };

  render(){
    const layout = {
      wrapperCol: {
        span: 24,
      },
    };
    return(
      <Router>
        {this.renderRedirect()} 
        <Route exact path="/">
          <div style={{width: "100%", height: "100vh", backgroundColor: "white"}}>
            <Row type={"flex"} align={"middle"} justify={"center"} style={{ textAlign: "center"}}>
              <Col style={{padding: "30px"}} span={24}>
                <p style={{fontSize: "clamp(3rem, -0.3750rem + 15.0000vw, 7.5rem)", color: "#001529", fontFamily: "Fredoka One"}}>JOBBAZAR</p>
              </Col>
            </Row>
            <Row type={"flex"} align={"middle"} justify={"center"}>
              <Col style={{backgroundColor: "#001529", borderRadius: "10px", padding: "30px"}} xxxl={3} xxl={5} xl={6} lg={8} md={12} sm={14} xs={22}>
                  <Form
                  {...layout}
                  name="basic"
                  onFinish={this.onFinish}
                  onFinishFailed={this.onFinishFailed}
                  >
                    <Form.Item
                      name="loginToken"
                      rules={[
                        {
                          required: true,
                          message: 'Please input email or mobile number',
                        },
                      ]}
                    >
                      <Input
                        style={{borderRadius: "5px"}}
                        size='large'
                        prefix={
                            <MailOutlined style={{color: "rgba(0,0,0,.25)"}}/>
                        }
                        placeholder='Email or Mobile Number'
                      />
                    </Form.Item>
                    <Form.Item
                      name="password"
                      rules={[
                        {
                          required: true,
                          message: 'Please input your password!',
                        },
                      ]}
                    >
                      <Input.Password
                        style={{borderRadius: "5px"}} 
                        size='large'
                        prefix={
                          <LockOutlined style={{color: "rgba(0,0,0,.25)"}}/>
                        }
                        type='password'
                        placeholder='Password'
                      />
                    </Form.Item>
                    <Form.Item>
                      <Button type="primary" htmlType="submit" size={"large"} style={{borderRadius: "5px", width: "100%"}}>
                        Login
                      </Button>
                    </Form.Item>
                  </Form>
              </Col>
            </Row>
          </div>
      </Route>
    </Router>
  );
  }
}

export default App;
