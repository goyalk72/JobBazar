import { Component } from 'react';
import { Descriptions, Card, Tag, Table, Collapse } from 'antd';

const { Panel } = Collapse;

const columns = [
    {
      title: 'Fisrtname',
      dataIndex: 'firstname',
      key: 'firstname',
    },
    {
      title: 'Lastname',
      dataIndex: 'lastname',
      key: 'lastname',
    },
    {
      title: 'Location',
      dataIndex: 'location',
      key: 'location',
    },
    {
        title: 'Contact',
        dataIndex: 'contact',
        key: 'contact',
    },
    {
      title: 'Skills',
      dataIndex: 'skills',
      key: 'skills',
      render: skills => (
        <>
          {skills.map(skill => {
            return (
              <Tag color="orange" key={skill}>
                {skill.toUpperCase()}
              </Tag>
            );
          })}
        </>
      ),
    },
    {
        title: 'Experience',
        dataIndex: 'experience',
        key: 'experience',
        render: experience => (
          <>
            {experience.map(exp => {
              return (
                <Tag color="green" key={exp}>
                  {exp.toUpperCase()}
                </Tag>
              );
            })}
          </>
        ),
      },

  ];

class Job extends Component {
    constructor(props){
        super(props);
    }

    render() {
        let tokens = this.props.tokens.map(item =>
            <Tag key={item} color='blue' style={{margin: 5}}>{item}</Tag>
        );
        return(
                <Card type="inner" title={this.props.title}>
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
                      <Descriptions.Item>
                          {tokens}
                      </Descriptions.Item>
                    </Descriptions>
                    <Collapse defaultActiveKey={['1']}>
                        <Panel header="View Applicants Details" key="1">
                            <Table columns={columns} dataSource={this.props.applicants} />
                        </Panel>
                    </Collapse>
                </Card>
                
        )
    }
}

export default Job;