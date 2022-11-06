import React, { useState, useEffect } from 'react';
import './index.scss';
import api from '../../../api';
import { Button } from 'antd';
import { dataCategoryList } from '../../../config'
import { PlusOutlined } from '@ant-design/icons';
import { Table, Drawer } from 'antd';
import { moduleActive, templateInfos } from '../../../store/atom';
import { useRecoilState } from 'recoil';
import IconCopy from "./../../../static/img/copy.png";
import { copyToClipboard } from "./../../../utils/tools";

const DataType = {
  name: '',
  templateName: '',
  dataCategory: '0',
  subCategory: '',
  lowerBoundType: [0, 0],
  upperBoundType: [0, 0],
  createdAt: '',
  description: ''
}

export default function List() {

  const [activeTabStr, setActiveTabStr] = useRecoilState(moduleActive);

  const [templateList, setTemplateList] = useState([]);

  const [open, setOpen] = useState(false);

  const [templateInfo, setTemplateInfo] = useRecoilState<any>(templateInfos);

  const [drawerRecords, setDrawerRecords] = useState(DataType);

  const getTemplateList = async () => {
    const res: any = await api.template.list({
      sort:'createdAt,DESC'
    });
    setTemplateList(res?.result?.data);
  };

  useEffect(() => {
    getTemplateList();
  }, []);

  const showDrawer = (record: any) => {
    setDrawerRecords(record);
    setOpen(true);
  };

  const onClose = () => {
    setOpen(false);
  };

  const offer = (record: any) => {
    setTemplateInfo(record);
    setActiveTabStr('offerClaims');
  }

  const expandedRowRender = (record: any, index: number) => {
    const columns = [
      { title: 'Class Name', width: 200, dataIndex: 'name', key: 'name', render: (text: string, record: any) => <span className="templateName">{text}</span> },
      { title: 'Lower Bound', width: 200, dataIndex: 'lowerBoundType', key: 'lowerBoundType', render: (text: any) => <span>{text[0] === 0 ? '>' : text[1] === 1 ? '≥' : '>'}</span>},
      { title: 'Upper Bound', width: 200, dataIndex: 'upperBoundType', key: 'upperBoundType', render: (text: any) => <span>{text[0] === 0 ? '-' : text[1] === 1 ? '≤' : '<'}</span>},
      { title: 'Description', width: 200, dataIndex: 'description', key: 'description', render: (text: any) => <span>{text || '-'}</span> },
    ];

    let renderData: Array<any> = [];

    if (record.classfications) {
      let expandData = JSON.parse(record.classfications);
      expandData.map((t: any, index: number) => {
        t.createdAt = record.createdAt;
        t.dataCategory = record.dataCategory;
        t.subCategory = record.subCategory;
        t.templateName = record.name;
        t.key = index;
      })
      renderData = [...expandData];
    } else {
      renderData = []
    }

    if (!Array.isArray(renderData)) {
      renderData = []
    }

    return <Table columns={columns} dataSource={renderData} pagination={false} onRow={(record) => ({
      onClick: () => showDrawer(record)
    })}/>;
  };

  const columns = [
    {
      title: 'Template Names', width: 200, dataIndex: 'name', key: 'name', render: (text: string, record: any) => (
        <span>{text}</span>
      ),
    },
    {
      title: 'Category', width: 200, dataIndex: 'dataCategory', key: 'dataCategory', render: (category: string) => (
        <div className="tag" style={{
          backgroundColor: category === '0' ? '#fff' : category === '1' ? 'rgb(32,128,226)' : category === '3' ? 'green' : '#fff',
          color: category === '1' || category === '3' ? '#fff' : category === '2' ? 'rgb(247,186,71)' : 'blue'
        }}>
          {dataCategoryList[Number(category)]}
        </div>
      ),
    },
    { title: 'Create Time', width: 200, dataIndex: 'createdAt', key: 'createdAt', render: (text: string) => <span>{text ? text.split('T')[0] : '-'}</span> },
    { title: 'Action', width: 200, key: 'operation', render: (text: string, record: any) => <span className="offer" onClick={() => offer(record)}>Offer</span> },
  ];

  return (
    <div className="list-page">
      <div className="list-des">
        <div className="list-title">Templates</div>
        <div>
          {templateList.length} Templates
        </div>
      </div>
      <div className="list-btn">
        <Button type="primary" size="large" onClick={() => setActiveTabStr('creatTempalte')}>
          Create Template
        </Button>
      </div>
      <div className="list-table">
        <Table
          rowKey="id"
          columns={columns}
          expandable={{ expandedRowRender }}
          dataSource={templateList}
          pagination={false}
        />
      </div>
      <Drawer
        title="View Class"
        placement={'right'}
        width={400}
        onClose={onClose}
        visible={open}
      >
        <div className="drawer-des">
          <div>
            <span>Template name:</span>
            <span>{drawerRecords.templateName}</span>
          </div>
          <div>
            <span>Class name:</span>
            <span>{drawerRecords.name}</span>
          </div>
          <div>
            <span>Data Category:</span>
            <span>{dataCategoryList[Number(drawerRecords?.dataCategory)]}</span>
          </div>
          {
            drawerRecords.dataCategory === '1' && <div>
              <span>NFT Contract:</span>
              <span>{drawerRecords.subCategory || '-'}</span>
              <span><img
                alt=""
                src={IconCopy}
                onClick={() => copyToClipboard(drawerRecords.subCategory)}
                className="copyIcon"
              /></span>
            </div>
          }

          {
            drawerRecords.dataCategory == '2' && <div>
              <span>Space ID:</span>
              <span>{drawerRecords.subCategory || '-'}</span>
              <span><img
                alt=""
                src={IconCopy}
                onClick={() => copyToClipboard(drawerRecords.subCategory)}
                className="copyIcon"
              /></span>
            </div>
          }

          <div>
            <span>Lower Bound:</span>
            <span>{drawerRecords.lowerBoundType[0] === 0 ? '>' : drawerRecords.lowerBoundType[1] === 1 ? '≥' : '>'}</span>
          </div>
          <div>
            <span>Upper Bound:</span>
            <span>{drawerRecords.upperBoundType[0] === 0 ? '-' : drawerRecords.upperBoundType[1] === 1 ? '≤' : '<'}</span>
          </div>
          <div>
            <span>Creation Date:</span>
            <span>{drawerRecords.createdAt ? drawerRecords.createdAt.split('T')[0] : '-'}</span>
          </div>
          <div>
            <span>Description:</span>
            <span>{drawerRecords.description}</span>
          </div>
        </div>
      </Drawer>
    </div>
  );
}
