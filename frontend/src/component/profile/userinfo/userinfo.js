import axios from "axios";
import React, { Component } from "react";
import { Descriptions, Tag } from 'antd';

class UserInfo extends Component {
    constructor(props){
        super(props);
        this.state = {
            data : {},
            loading: true
        }
        
    }

    componentDidMount() {
        this.props.startLoading()
        axios.get("http://127.0.0.1:3030/userinfo", {
            params: {
                email : this.props.email,
                mobile: this.props.mobile
            }
        }).then(res => {
            this.setState({
                data : {...res.data},
                loading: false
             })
        this.props.stopLoading()
         }).catch(err => {
                console.error(err)
                this.props.stopLoading()
            });
    }

    render(){
        var skills = this.state.loading ? null : this.state.data.skills.map(item =>
            <Tag key={item} color='blue' style={{margin: 5}}>{item}</Tag>
        );
        var experience = this.state.loading ? null : this.state.data.experience.map(item =>
            <Tag key={item} color='blue' style={{margin: 5}}>{item}</Tag>
        );
        return(
            <div>
                <Descriptions
                title="User Information"
                bordered
                column={{ xxl: 4, xl: 3, lg: 3, md: 3, sm: 2, xs: 1 }}
                >
                    <Descriptions.Item label="First Name">{this.state.data.firstname}</Descriptions.Item>
                    <Descriptions.Item label="Last Name">{this.state.data.lastname}</Descriptions.Item>
                    <Descriptions.Item label="Date of Birth">{this.state.data.dob}</Descriptions.Item>
                    <Descriptions.Item label="Location">{this.state.data.location}</Descriptions.Item>
                    <Descriptions.Item label="Gender">{this.state.data.gender}</Descriptions.Item>
                    <Descriptions.Item label="Contact No">{this.state.data.contact}</Descriptions.Item>
                    <Descriptions.Item label="Email ID">{this.state.data.email}</Descriptions.Item>
                    <Descriptions.Item label="Skills">
                        {skills}
                    </Descriptions.Item>
                    <Descriptions.Item label="Experience">
                        {experience}
                    </Descriptions.Item>
                </Descriptions>
            </div>
        );
    }
}

export default UserInfo;