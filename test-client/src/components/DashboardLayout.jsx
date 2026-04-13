import React, { useState } from 'react';
import { Layout, Menu, Button, theme } from 'antd';
import {
  DashboardOutlined,
  FolderOpenOutlined,
  ClusterOutlined,
  CodeOutlined,
  SettingOutlined,
  MenuUnfoldOutlined,
  MenuFoldOutlined,
  UserOutlined
} from '@ant-design/icons';
import GlobalStats from './GlobalStats';
import NodeHealthGrid from './NodeHealthGrid';
import FileUploader from './FileUploader';
import FileExplorerTable from './FileExplorerTable';

const { Header, Sider, Content } = Layout;

const DashboardLayout = () => {
  const [collapsed, setCollapsed] = useState(false);
  const [activeTab, setActiveTab] = useState('1');
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();

  const handleMenuClick = (e) => {
    setActiveTab(e.key);
  };

  const renderContent = () => {
    switch (activeTab) {
      case '1':
        return (
          <div className="space-y-6">
            <GlobalStats />
            <NodeHealthGrid />
            <div className="bg-white p-6 rounded-lg shadow-sm">
              <h3 className="text-lg font-medium mb-2 text-gray-800">System Activity</h3>
              <p className="text-gray-500">System is running optimally. Distributed node latency is under 15ms.</p>
            </div>
          </div>
        );
      case '2':
        return (
          <div className="space-y-6">
            <FileUploader />
            <FileExplorerTable />
          </div>
        );
      default:
        return (
          <div className="flex flex-col items-center justify-center py-20 text-gray-400">
            <CodeOutlined className="text-4xl mb-4" />
            <p>This module is under construction.</p>
          </div>
        );
    }
  };

  return (
    <Layout className="min-h-screen">
      <Sider trigger={null} collapsible collapsed={collapsed} theme="light" className="shadow-sm border-r border-gray-100">
        <div className="h-16 flex items-center justify-center border-b border-gray-100">
          <span className={`font-bold text-blue-600 transition-all ${collapsed ? 'text-lg text-center' : 'text-xl'}`}>
            {collapsed ? 'DFS' : 'Gravity DFS'}
          </span>
        </div>
        <Menu
          theme="light"
          mode="inline"
          defaultSelectedKeys={['1']}
          onClick={handleMenuClick}
          className="border-r-0 mt-4 h-[calc(100vh-80px)]"
          items={[
            {
              key: '1',
              icon: <DashboardOutlined />,
              label: 'Dashboard',
            },
            {
              key: '2',
              icon: <FolderOpenOutlined />,
              label: 'My Files',
            },
            {
              key: '3',
              icon: <ClusterOutlined />,
              label: 'Storage Nodes',
            },
            {
              key: '4',
              icon: <CodeOutlined />,
              label: 'System Logs',
            },
            {
              key: '5',
              icon: <SettingOutlined />,
              label: 'Settings',
            },
          ]}
        />
      </Sider>
      <Layout className="bg-gray-50">
        <Header
          className="flex items-center justify-between px-4 shadow-sm z-10"
          style={{
            padding: 0,
            background: colorBgContainer,
          }}
        >
          <Button
            type="text"
            icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
            onClick={() => setCollapsed(!collapsed)}
            style={{
              fontSize: '16px',
              width: 64,
              height: 64,
            }}
          />
          <div className="pr-6 flex items-center space-x-4">
            <Button type="primary" className="bg-blue-600 font-medium rounded-full px-6 shadow-sm">Deploy Node</Button>
            <div className="w-8 h-8 rounded-full bg-blue-100 flex items-center justify-center text-blue-600 border border-blue-200">
              <UserOutlined />
            </div>
          </div>
        </Header>
        <Content
          className="m-6 p-6 min-h-[280px]"
          style={{
            borderRadius: borderRadiusLG,
          }}
        >
          {renderContent()}
        </Content>
      </Layout>
    </Layout>
  );
};

export default DashboardLayout;
