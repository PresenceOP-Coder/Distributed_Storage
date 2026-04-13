import React from 'react';
import { Card, Progress, Badge, Row, Col } from 'antd';
import { HddOutlined } from '@ant-design/icons';

const nodes = [
  { id: 'Node-1', usage: 65, active: true },
  { id: 'Node-2', usage: 82, active: true },
  { id: 'Node-3', usage: 45, active: true },
  { id: 'Node-4', usage: 90, active: true },
  { id: 'Node-5', usage: 12, active: true },
];

const NodeHealthGrid = () => {
  return (
    <div className="mb-6">
      <h2 className="text-xl font-semibold mb-4 text-gray-800">Storage Nodes Health</h2>
      <Row gutter={[16, 16]}>
        {nodes.map((node) => (
          <Col xs={12} sm={8} md={6} lg={4} key={node.id}>
            <Card
              bordered={false}
              size="small"
              className="shadow-sm border border-gray-100 hover:shadow-md transition-shadow text-center"
            >
              <div className="flex flex-col items-center justify-center p-2">
                <HddOutlined className="text-2xl text-gray-600 mb-2" />
                <h3 className="font-medium text-gray-700">{node.id}</h3>
                <Badge 
                  status={node.active ? 'success' : 'error'} 
                  text={node.active ? 'Online' : 'Offline'} 
                  className="mb-3 mt-1" 
                />
                <div className="w-full relative px-2">
                  <Progress 
                    percent={node.usage} 
                    size="small" 
                    status={node.usage > 85 ? 'exception' : 'active'}
                    format={(percent) => <span className="text-xs">{percent}% full</span>}
                  />
                </div>
              </div>
            </Card>
          </Col>
        ))}
      </Row>
    </div>
  );
};

export default NodeHealthGrid;
