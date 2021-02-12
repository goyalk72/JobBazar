import axios from 'axios';
import React, { Component } from 'react';
import Job from './job/job';
import { Input, Alert } from 'antd';

const { Search } = Input;

class SearchJobs extends Component {
    constructor(props){
        super(props);
        this.state = {
            jobs : [],
            query: "",
            loading: true,
            alert: false
        }
    }

    onSearch = (value) => {
        this.setState({
            query: value,
            loading:true
        });
        axios.get("http://127.0.0.1:3030/searchjobs", {
            params:{
                query: this.state.query,
                email: this.props.email,
                mobile: this.props.mobile
            }
        })
        .then( res => {
            this.setState({
                jobs : [...res.data],
                loading: false
            })
        }).catch(err => {
            console.log(err);
        });
     }

    componentDidMount() {
        this.props.startLoading()
        axios.get("http://127.0.0.1:3030/searchjobs", {
            params:{
                query: this.state.query,
                email: this.props.email,
                mobile: this.props.mobile
            }
        })
        .then( res => {
            this.setState({
                jobs : [...res.data],
                loading: false
            })
            this.props.stopLoading()
        }).catch(err => {
            this.props.stopLoading()
            console.log(err);
        });
    }

    onClickHandler = (jobid) => {
        axios.get("http://127.0.0.1:3030/applyjob", {
            params: {
                job: jobid,
                email: this.props.email,
                mobile: this.props.mobile
            }
        }).then(res => {
            this.setState({
                alert: true,
                alertmsg: res.data.ispresent
            })
            this.createAlert()
        }
        ).catch(err => console.error(err));
    }

    onSaveHandler = (jobid) => {
        axios.get("http://127.0.0.1:3030/savejob", {
            params: {
                job: jobid,
                email: this.props.email,
                mobile: this.props.mobile
            }
        }).then(res => {
            this.setState({
                alert: true,
                alertmsg: res.data.ispresent
            })
            this.createAlert()
        }
        ).catch(err => console.error(err));
    }

    createAlert() {
        if (this.state.alert){
            if (this.state.alertmsg){
                alert("You have already applied for this Job")
            }else {
                alert(" Application successful ")
            }
        }
    }

    render() {
        let jobs = this.state.loading ? null : this.state.jobs.map(job => (
            <Job
            onClickHandler={() => this.onClickHandler(job.job.id)}
            onSaveHandler={() => this.onSaveHandler(job.job.id)}
            key={job.job.id}
            companyname={job.job.companyname}
            title={job.job.title}
            location={job.job.location}
            contact={job.job.contact}
            description={job.job.description}
            type={job.job.type}
            salary={job.job.salary}
            date={job.job.date}
            tokens={job.job.tokens}
            num = {job.job.applicants ? job.job.applicants.length : 0}
            issaved={job.issaved}
            isapplied={job.isapplied}
             />
        ));
        return (
            <div>
                <Search placeholder="Type a keyword to search" onSearch={this.onSearch} style={{ width: 200 }} />
                {jobs}
            </div>
        );
    }


}

export default SearchJobs;