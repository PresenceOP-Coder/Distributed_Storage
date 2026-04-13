import React, { useState } from 'react';
import { Upload, message, Progress } from 'antd';
import { InboxOutlined } from '@ant-design/icons';

const { Dragger } = Upload;

const FileUploader = () => {
  const [uploadState, setUploadState] = useState({
    status: 'idle', // idle, chunking, distributing, done
    progress: 0,
    fileName: ''
  });

  const props = {
    name: 'file',
    multiple: true,
    action: 'https://run.mocky.io/v3/435e224c-44fb-4773-9faf-380c5e6a2188',
    showUploadList: false,
    onChange(info) {
      if (info.file.status !== 'uploading') {
        const fakeProgress = () => {
          setUploadState({ status: 'chunking', progress: 20, fileName: info.file.name });
          setTimeout(() => setUploadState(prev => ({ ...prev, progress: 50, status: 'distributing'})), 1000);
          setTimeout(() => setUploadState(prev => ({ ...prev, progress: 100, status: 'done'})), 2500);
          setTimeout(() => {
            message.success(`${info.file.name} chunked and distributed successfully.`);
            setUploadState({ status: 'idle', progress: 0, fileName: '' });
          }, 3500);
        };
        fakeProgress();
      }
    },
    onDrop(e) {
      console.log('Dropped files', e.dataTransfer.files);
    },
  };

  return (
    <div className="mb-8">
      <h2 className="text-xl font-semibold mb-4 text-gray-800">Upload to Cluster</h2>
      <div className="bg-white p-6 rounded-lg shadow-sm border border-gray-100">
        {uploadState.status === 'idle' ? (
          <Dragger {...props} className="bg-gray-50 border-gray-300 hover:border-blue-500 transition-colors">
            <p className="ant-upload-drag-icon text-blue-500">
              <InboxOutlined />
            </p>
            <p className="ant-upload-text font-medium text-gray-600">Click or drag file to this area to upload</p>
            <p className="ant-upload-hint text-gray-400">
              Files will be automatically chunked and distributed across active storage nodes.
            </p>
          </Dragger>
        ) : (
          <div className="py-8 px-4 flex flex-col items-center justify-center border-2 border-dashed border-blue-200 rounded-lg bg-blue-50">
            <h3 className="mb-2 text-lg font-medium text-blue-800">
              {uploadState.status === 'chunking' ? 'Chunking File...' : uploadState.status === 'distributing' ? 'Distributing across nodes...' : 'Completed!'}
            </h3>
            <p className="mb-4 text-gray-500 font-mono text-sm">{uploadState.fileName}</p>
            <Progress 
              percent={uploadState.progress} 
              status={uploadState.progress === 100 ? 'success' : 'active'} 
              strokeColor={{
                '0%': '#108ee9',
                '100%': '#87d068',
              }}
              className="w-full max-w-md"
            />
          </div>
        )}
      </div>
    </div>
  );
};

export default FileUploader;
