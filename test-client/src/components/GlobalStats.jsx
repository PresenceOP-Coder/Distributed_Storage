import React from 'react';
import { Card, Statistic, Progress, Row, Col } from 'antd';
import { DatabaseOutlined, ClusterOutlined, FileTextOutlined } from '@ant-design/icons';

const GlobalStats = () => {
  return (
    <div className="mb-6">
      <h2 className="text-xl font-semibold mb-4 text-gray-800">Global System View</h2>
      <Row gutter={[16, 16]}>
        <Col xs={24} md={8}>
          <Card bordered={false} className="shadow-sm rounded-lg hover:shadow-md transition-shadow">
            <Statistic
              title="Total Storage Used"
              value="4.2"
              suffix="TB / 10 TB"
              prefix={<DatabaseOutlined className="text-blue-500" />}
              valueStyle={{ color: '#1677ff', fontWeight: 600 }}
            />
            <div className="mt-4">
              <Progress percent={42} strokeColor="#1677ff" showInfo={false} />
            </div>
          </Card>
        </Col>

        <Col xs={24} md={8}>
          <Card bordered={false} className="shadow-sm rounded-lg hover:shadow-md transition-shadow">
            <Statistic
              title="Active Storage Nodes"
              value={5}
              suffix="/ 5 Online"
              prefix={<ClusterOutlined className="text-green-500" />}
              valueStyle={{ color: '#52c41a', fontWeight: 600 }}
            />
            <div className="mt-4">
               <Progress percent={100} strokeColor="#52c41a" steps={5} showInfo={false} />
            </div>
          </Card>
        </Col>

        <Col xs={24} md={8}>
          <Card bordered={false} className="shadow-sm rounded-lg hover:shadow-md transition-shadow">
            <Statistic
              title="Total Files Stored"
              value={1248}
              prefix={<FileTextOutlined className="text-purple-500" />}
              valueStyle={{ color: '#722ed1', fontWeight: 600 }}
            />
            <div className="mt-4">
              <p className="text-sm text-gray-400">Total objects currently distributed.</p>
            </div>
          </Card>
        </Col>
      </Row>
    </div>
  );
};

export default GlobalStats;
