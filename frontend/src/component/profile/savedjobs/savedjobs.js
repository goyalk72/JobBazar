import { Collapse } from 'antd';
import { Component } from 'react';
import Job from "../searchjobs/job/job";
import axios from 'axios';

class Savedjobs extends Component{
    constructor(props){
        super(props);
        this.state = {
            jobs: [],
            loading: true
        }
    }

    componentDidMount() {
        this.props.startLoading()
        axios.get("http://127.0.0.1:3030/viewsavedjobs", {
            params:{
                query: "",
                email: this.props.email,
                mobile: this.props.mobile
            }
        })
        .then( res => {
            console.log(res.data);
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
        console.log(jobid);
        axios.get("http://127.0.0.1:3030/applyjob", {
            params: {
                job: jobid,
                email: this.props.email,
                mobile: this.props.mobile
            }
        }).then(res => {
            console.log(res.data)
        }
        ).catch(err => console.error(err));
    }

    onSaveHandler = (jobid) => {
        console.log(jobid);
        axios.get("http://127.0.0.1:3030/savejob", {
            params: {
                job: jobid,
                email: this.props.email,
                mobile: this.props.mobile
            }
        }).then(res => {
            console.log(res.data)
        }
        ).catch(err => console.error(err));
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
                {jobs}
            </div>
        )
    }
}

export default Savedjobs;