import React, { useState, useEffect } from 'react';
import './index.scss';
import api from '../../../api';
import { message } from 'antd';
import { Button } from 'antd';
import { dataCategoryList } from '../../../config'
import { PlusOutlined, DownOutlined } from '@ant-design/icons';
import { Badge, Table, Drawer } from 'antd';
import { moduleActive, templateInfos } from '../../../store/atom';
import { useRecoilState } from 'recoil';
import IconCopy from "./../../../static/img/copy.png";
import { copyToClipboard } from "./../../../utils/tools";

const DataType = {
  name: '',
  templateName: '',
  dataCategory: '',
  subCategory: '',
  lowerBound: [0, 0],
  upperBound: [0, 0],
  createdAt: '',
  description: '',
  link: ''
}


export default function List() {

  const [activeTabStr, setActiveTabStr] = useRecoilState(moduleActive);

  const [claimList, setClaimList] = useState([]);

  const [open, setOpen] = useState(false);

  const [templateInfo, setTemplateInfo] = useRecoilState<any>(templateInfos);

  const [drawerRecords, setDrawerRecords] = useState(DataType);

  const getclaimList = async () => {
    const res: any = await api.offer.list({
      sort:'createdAt,DESC'
    });
    res?.result?.data.map((t: any) => {
      t.dataCategory = t.template.dataCategory
    })
    setClaimList(res?.result?.data);
  };

  useEffect(() => {
    getclaimList();
  }, []);

  const showDrawer = (record: any) => {
    console.log('record', record)
    setOpen(true);
    setDrawerRecords(record);
  };

  const onClose = () => {
    setOpen(false);
  };

  const setRevocation = (record: any) => {
    let classfications = JSON.parse(record.preClaims);
    classfications.map((t: any, i: number) => {
      t.lowerBoundType = t.lowerBound;
      t.upperBoundType = t.upperBound;
    })
    setTemplateInfo({
      ...record.template,
      link: record.link,
      claimId: record.id,
      expirationDate: JSON.parse(record.preClaims)[0]['expirationDate'],
      classfications: JSON.stringify(classfications)
    });
    setActiveTabStr('revocation');
  }

  const getTrigger = (lowerBound: number, upperBound: number) => {
    if (lowerBound == 0 && upperBound == 0) {
      return 'Never'
    }
    if (lowerBound == 1 && upperBound == 0) {
      return 'The lower bound is unmet'
    }
    if (lowerBound == 0 && upperBound == 1) {
      return 'The upper bound is unmet'
    }
    if (lowerBound == 1 && upperBound == 1) {
      return 'Either bound is unmet'
    }
  }

  const expandedRowRender = (record: any, index: number) => {

    const columns = [
      { title: 'Claim Name', width: 200, dataIndex: 'name', key: 'name', render: (text: string, record: any) => <span className="templateName">{text}</span> },
      { title: 'Lower Bound', width: 200, dataIndex: 'lowerBound', key: 'lowerBound', render: (text: any) => text && text.length > 0 ? <span>{text[0] === 0 ? '>0 (default)' : text[1] === 1 ? '≥' : '>'}{text[0] === 1 ? text[2] : ''}</span> : '-' },
      { title: 'Upper Bound', width: 200, dataIndex: 'upperBound', key: 'upperBound', render: (text: any) => <span>{text[0] === 0 ? '-' : text[1] === 1 ? '≤' : '<'}{text[2]}</span>},
      { title: 'Revocation Trigger', width: 200, dataIndex: 'description', key: 'description', render: (text: string, record: any) => <span>{getTrigger(record.lowerBound[3], record.upperBound[3])}</span> },
    ];

    let renderData: Array<any> = [];

    if (record.preClaims) {
      let expandData = JSON.parse(record.preClaims);
      let templateClass = JSON.parse(record.template.classfications);
      expandData.map((t: any, i: number) => {
        t.createdAt = record.createdAt;
        t.dataCategory = record.dataCategory;
        t.subCategory = t.subcategory ? t.subcategory : record.template.subCategory;
        t.templateName = record.template.name;
        t.description = templateClass[i]['description'];
        t.link = record.link;
      })
      console.log(expandData)
      renderData = [...expandData];
    } else {
      renderData = []
    }

    if (!Array.isArray(renderData)) {
      renderData = []
    }

    return <Table columns={columns} dataSource={renderData} pagination={false} onRow={(record) => ({
      onClick: () => showDrawer(record)
    })} />;
  };

  const columns = [
    {
      title: 'From Template', width: 200, dataIndex: 'name', key: 'name', render: (text: string, record: any) => (
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
    { title: 'Offer Time', width: 200, dataIndex: 'createdAt', key: 'createdAt', render: (text: string) => <span>{text ? text.split('T')[0] : '--'}</span> },
    { title: 'Action', width: 200, key: 'operation', render: (text: string, record: any) => <span className="revocation" onClick={() => setRevocation(record)}>Set Revocation</span> },
  ];

  return (
    <div className="list-page">
      <div className="list-des">
        <div className="list-title">Claims</div>
        <div>
          {claimList.length} Offers
        </div>
      </div>
      <div className="list-table">
        <Table
          rowKey="id"
          columns={columns}
          expandable={{ expandedRowRender }}
          dataSource={claimList}
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
            <span>{dataCategoryList[Number(drawerRecords.dataCategory)]}</span>
          </div>
          {
            drawerRecords.dataCategory === '1' && (
              <div>
                <span>NFT Contract:</span>
                <span>{drawerRecords.subCategory || "-"}</span>
                <span><img
                  alt=""
                  src={IconCopy}
                  onClick={() => copyToClipboard(drawerRecords.subCategory)}
                  className="copyIcon"
                /></span>
              </div>
            )
          }
          {
            drawerRecords.dataCategory === '2' && (
              <div>
                <span>Space ID:</span>
                <span>{drawerRecords.subCategory || "-"}</span>
                <span><img
                  alt=""
                  src={IconCopy}
                  onClick={() => copyToClipboard(drawerRecords.subCategory)}
                  className="copyIcon"
                /></span>
              </div>
            )
          }
          <div>
            <span>Lower Bound:</span>
            <span>{drawerRecords.lowerBound[0] === 0 ? '>0 (default)' : drawerRecords.lowerBound[1] === 1 ? '≥' : '>'}{drawerRecords.lowerBound[0] === 1 ? drawerRecords.lowerBound[2] : ''}</span>
          </div>
          <div>
            <span>Upper Bound:</span>
            <span>{drawerRecords.upperBound[0] === 0 ? '-' : drawerRecords.upperBound[1] === 1 ? '≤' : '<'}{drawerRecords.upperBound[2]}</span>
          </div>
          <div>
            <span>Creation Date:</span>
            <span>{drawerRecords.createdAt ? drawerRecords.createdAt.split('T')[0] : '-'}</span>
          </div>
          <div>
            <span>Class Hash:</span>
            <span>hash</span>
            <span><img
              alt=""
              src={IconCopy}
              onClick={() => copyToClipboard('...')}
              className="copyIcon"
            /></span>
          </div>
          <div>
            <span>Description:</span>
            <span>{drawerRecords.description}</span>
            <span><img
              alt=""
              src={IconCopy}
              onClick={() => copyToClipboard(drawerRecords.description)}
              className="copyIcon"
            /></span>
          </div>
          <div>
            <span>Link:</span>
            <span>{drawerRecords.link}</span>
            <span><img
              alt=""
              src={IconCopy}
              onClick={() => copyToClipboard(drawerRecords.link)}
              className="copyIcon"
            /></span>
          </div>
        </div>
      </Drawer>
    </div>
  );
}
