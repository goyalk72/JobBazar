import { Component } from 'react';
import { Descriptions, Card, Tag , Button} from 'antd';

class Job extends Component {
    constructor(props){
        super(props);
    }

    render() {
        let tokens = this.props.tokens.map(item =>
            <Tag key={item} color='blue' style={{margin: 5}}>{item}</Tag>
        );
        return(
                <Card type="inner" 
                    title={this.props.title} 
                    extra={
                        <div>
                            <Button type="primary" onClick={this.props.onClickHandler} disabled={this.props.isapplied}> Apply </Button>
                            <Button onClick={this.props.onSaveHandler} disabled={this.props.issaved}> Save </Button>
                        </div> }>
                    <Descriptions
                    title=""
                    bordered
                    column={{ xxl: 4, xl: 3, lg: 3, md: 3, sm: 2, xs: 1 }}
                    >
                    <Descriptions.Item label="Company Name">{this.props.companyname}</Descriptions.Item>
                    <Descriptions.Item label="Location">{this.props.location}</Descriptions.Item>
                    <Descriptions.Item label="Contact">{this.props.contact}</Descriptions.Item>
                    <Descriptions.Item label="Description">{this.props.description}</Descriptions.Item>
                    <Descriptions.Item label="Type of Job">{this.props.type}</Descriptions.Item>
                    <Descriptions.Item label="Salary">{this.props.salary}</Descriptions.Item>
                    <Descriptions.Item label="Apply By">{this.props.date}</Descriptions.Item>
                    <Descriptions.Item label="Number of Applicants">{this.props.num}</Descriptions.Item>
                    <Descriptions.Item>
                        {tokens}
                    </Descriptions.Item>
                    </Descriptions>
                </Card>
        )
    }
}

export default Job;