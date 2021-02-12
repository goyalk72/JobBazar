import React, { Component } from 'react';
import { Layout, Menu, Spin } from 'antd';
import {
  UserOutlined,
  SaveOutlined,
  MailOutlined,
  FileSearchOutlined,
  FileAddOutlined
} from '@ant-design/icons';
import "./profile.css";
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link,
    Redirect
  } from "react-router-dom";
import UserInfo from "./userinfo/userinfo";
import SearchJobs from "./searchjobs/searchjobs";
import Postedjobs from "./postedjobs/postedjobs";
import Savedjobs from "./savedjobs/savedjobs";
import CreateJob from "./createjob/createjob";

const { Header, Sider, Content } = Layout;
class Profile extends Component{
    constructor(props){
        super(props);
        this.state = {
          loading: false,
          collapsed: false,
          email: this.props.location.state.email,
          mobile: this.props.location.state.mobile
        };
    }

    toggle = () => {
      this.setState({
        collapsed: !this.state.collapsed,
      });
    };

    startLoading = () => {
      this.setState({
        loading: true,
      });
    };

    stopLoading = () => {
      this.setState({
        loading: false,
      });
    };

    render(){
      const commonProps = {
        email: this.state.email,
        mobile: this.state.mobile,
        startLoading: this.startLoading,
        stopLoading: this.stopLoading
      }

      return (
        <Router>
          <Layout>
            <Sider trigger={null}
              width={250}
              style={{height: '100vh', position: 'fixed' }}
            >
              <div className="logo" style={{textAlign: "center", margin: "10px", marginBottom: "2%"}}>
                <p style={{fontSize: "clamp(1rem, -0.3750rem + 15.0000vw, 2.5rem)", color: "white", fontFamily: "Fredoka One"}}>JOBBAZAR</p>
              </div>
              <Menu theme="dark" mode="inline" defaultSelectedKeys={['1']}>
                <Menu.Item key="1" icon={<UserOutlined />}>
                  <span>User Info</span>
                    <Link to={{
                      pathname: `/profile/userinfo`
                    }}/>
                </Menu.Item>
                <Menu.Item key="2" icon={<SaveOutlined />}>
                  <span>Saved Jobs</span>
                  <Link to={{
                    pathname: `/savedjobs`
                  }}/>
                </Menu.Item>
                <Menu.Item key="3" icon={<FileSearchOutlined />}>
                  <span>All Jobs</span>
                  <Link to={{
                      pathname: `/searchjobs`
                  }}/>
                </Menu.Item>
                <Menu.Item key="4" icon={<MailOutlined />}>
                  <span>Posted Jobs</span>
                  <Link to={{
                      pathname: `/postedjobs`
                  }}/>
                </Menu.Item>
                <Menu.Item key="5" icon={<FileAddOutlined />}>
                  <span>Create Job</span>
                  <Link to={{
                      pathname: `/createjob`
                  }}/>
                </Menu.Item>
              </Menu>
            </Sider>
            <Layout className="site-layout" style={{marginLeft: 250}}>
              <Header className="site-layout-background" style={{ padding: 0 }}>
              </Header>
              <Spin spinning={this.state.loading} tip={"Loading....."}>
                <Content
                  className="site-layout-background"
                  style={{
                    margin: '24px 16px',
                    padding: 24,
                    minHeight: "80vh",
                  }}
                >
                  <Route exact path="/profile/userinfo" render={() => <UserInfo {...commonProps}/>}/>
                  <Route exact path="/savedjobs" render={() => <Savedjobs {...commonProps}/>} />
                  <Route exact path="/searchjobs" render={() => <SearchJobs {...commonProps}/>} />
                  <Route exact path="/postedjobs" render={() => <Postedjobs {...commonProps}/>} />
                  <Route exact path="/createjob" render={() => <CreateJob {...commonProps}/>} />
                </Content>
              </Spin>
            </Layout>
          </Layout>
        </Router>
      );
    }
}
export default Profile;