import axios from 'axios';
import React, { Component } from 'react';
import Job from './job/job';

class Postedjobs extends Component {
    constructor(props){
        super(props);
        this.state = {
            jobs: [],
            loading: true
        }
    }

    componentDidMount() {
        this.props.startLoading();
        axios.get("http://127.0.0.1:3030/postedjobs", {
            params:{
                email: this.props.email,
                mobile: this.props.mobile
            }
        }).then( res => {
            this.setState({
                jobs: [...res.data],
                loading: false
            });
            this.props.stopLoading();
        }).catch(err => {
            this.props.stopLoading()
            console.error(err)
        })
    }

    render(){
        let jobss = this.state.loading ? null : this.state.jobs.map(job => (
            <Job
            key={job.id}
            companyname={job.companyname}
            title={job.title}
            location={job.location}
            contact={job.contact}
            description={job.description}
            type={job.type}
            salary={job.salary}
            date={job.date}
            tokens={job.tokens}
            applicants={job.applicants}
             />
        ));

        return (
            <div>
                {jobss}
            </div>
        );
    }
    
}

export default Postedjobs;