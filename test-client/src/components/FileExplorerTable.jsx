import React from 'react';
import { Table, Tag, Space, Button } from 'antd';
import { 
  FilePdfOutlined, 
  FileImageOutlined, 
  FileZipOutlined, 
  FileTextOutlined,
  DownloadOutlined,
  DeleteOutlined
} from '@ant-design/icons';

const getFileIcon = (type) => {
  switch(type) {
    case 'pdf': return <FilePdfOutlined className="text-red-500 mr-2 text-lg" />;
    case 'image': return <FileImageOutlined className="text-blue-500 mr-2 text-lg" />;
    case 'zip': return <FileZipOutlined className="text-yellow-500 mr-2 text-lg" />;
    default: return <FileTextOutlined className="text-gray-500 mr-2 text-lg" />;
  }
}

const data = [
  {
    key: '1',
    name: 'Q3_Financial_Report.pdf',
    type: 'pdf',
    size: '4.2 MB',
    status: 'Ready',
    chunks: 4,
  },
  {
    key: '2',
    name: 'user_avatars_backup.zip',
    type: 'zip',
    size: '1.2 GB',
    status: 'Distributed',
    chunks: 124,
  },
  {
    key: '3',
    name: 'landing_hero_bg.png',
    type: 'image',
    size: '2.4 MB',
    status: 'Chunking',
    chunks: 2,
  },
  {
    key: '4',
    name: 'server_logs_2023.txt',
    type: 'text',
    size: '840 KB',
    status: 'Ready',
    chunks: 1,
  },
];

const columns = [
  {
    title: 'File Name',
    dataIndex: 'name',
    key: 'name',
    render: (text, record) => (
      <div className="font-medium text-gray-700 flex items-center">
        {getFileIcon(record.type)}
        {text}
      </div>
    ),
  },
  {
    title: 'Size',
    dataIndex: 'size',
    key: 'size',
    render: (size) => <span className="text-gray-500">{size}</span>,
  },
  {
    title: 'Chunks',
    dataIndex: 'chunks',
    key: 'chunks',
    render: (chunks) => <span className="text-gray-400 font-mono text-sm">{chunks}</span>,
  },
  {
    title: 'Status',
    key: 'status',
    dataIndex: 'status',
    render: (_, { status }) => {
      let color = status === 'Ready' ? 'green' : status === 'Chunking' ? 'gold' : 'blue';
      return (
        <Tag color={color} key={status} className="rounded-full px-3">
          {status.toUpperCase()}
        </Tag>
      );
    },
  },
  {
    title: 'Action',
    key: 'action',
    render: (_, record) => (
      <Space size="middle">
        <Button disabled={record.status === 'Chunking'} type="text" icon={<DownloadOutlined className="text-blue-500"/>} className="hover:bg-blue-50" />
        <Button danger type="text" icon={<DeleteOutlined />} className="hover:bg-red-50" />
      </Space>
    ),
  },
];

const FileExplorerTable = () => {
  return (
    <div>
      <h2 className="text-xl font-semibold mb-4 text-gray-800">Cluster File Explorer</h2>
      <div className="bg-white p-2 rounded-lg shadow-sm border border-gray-100">
        <Table 
          columns={columns} 
          dataSource={data} 
          pagination={{ pageSize: 5 }} 
          className="ant-table-striped"
        />
      </div>
    </div>
  );
};

export default FileExplorerTable;
