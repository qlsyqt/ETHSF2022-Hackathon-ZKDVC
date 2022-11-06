import React, { useState, useEffect } from 'react';
import './index.scss';
import { ArrowLeftOutlined } from '@ant-design/icons';
import { Checkbox, Select, Button, message } from 'antd';
import { moduleActive, templateInfos, routerNm } from '../../../store/atom';
import { useRecoilState } from 'recoil';
import { RouteComponentProps, useHistory } from 'react-router-dom';
import IconCopy from "./../../../static/img/copy.png";
import { copyToClipboard } from "./../../../utils/tools";
import { dataCategoryList } from '../../../config';
import api from '../../../api';
const { Option } = Select;

const defaultClassficationItem = {
  name: '',
  lowerBoundType: [0, 0, null, 1],
  upperBoundType: [0, 0, null, 1],
  description: '',
}

export default function CreateTemplate() {

  const [, setActiveTabStr] = useRecoilState(moduleActive);

  const [templateInfo, setTemplateInfo] = useRecoilState<any>(templateInfos);

  const [classfications, setClassfications] = useState<any>([defaultClassficationItem]);

  const [routerName, setRouterName] = useRecoilState<any>(routerNm);

  const [triggerValue, setTriggerValue] = useState('0');

  const history = useHistory();

  useEffect(() => {
    let classItems: any = JSON.parse(templateInfo.classfications);
    classItems.map((t: any) => {
      if (!t.lowerBoundType[3] && t.lowerBoundType[3] !== 0) {
        t.lowerBoundType[3] = 0;
        t.upperBoundType[3] = 0;
        t.triggerValue = '0';
      } else {
        if (t.lowerBoundType[3] === 0 &&
          t.upperBoundType[3] === 0) {
          t.triggerValue = '0';
        }
        if (t.lowerBoundType[3] === 1 &&
          t.upperBoundType[3] === 0) {
          t.triggerValue = '1';
        }
        if (t.lowerBoundType[3] === 0 &&
          t.upperBoundType[3] === 1) {
          t.triggerValue = '2';
        }
        if (t.lowerBoundType[3] === 1 &&
          t.upperBoundType[3] === 1) {
          t.triggerValue = '3';
        }
      }
    })
    setClassfications(classItems);
  }, []);

  const handleChange = (index: number, value: string) => {
    setTriggerValue(value);
    setClassfications((prev: any) => {
      let classCations = [...prev];
      classCations.map((t: any, i: number) => {
        if (index === i) {
          if (value === '0') {
            t.lowerBoundType[3] = 0;
            t.upperBoundType[3] = 0;
          }
          if (value === '1') {
            t.lowerBoundType[3] = 1;
            t.upperBoundType[3] = 0;
          }
          if (value === '2') {
            t.lowerBoundType[3] = 0;
            t.upperBoundType[3] = 1;
          }
          if (value === '3') {
            t.lowerBoundType[3] = 1;
            t.upperBoundType[3] = 1;
          }
          t.triggerValue = value;
        }
      });
      return [...classCations];
    })
  };

  const fallback = () => {
    if (history.location.pathname === '/home/template') {
      setActiveTabStr('setLink')
    } else {
      setActiveTabStr('claimList')
    }
  };

  const finishOffer = async () => {

    let preClaims: any = [];

    classfications.map((t: any) => {
      preClaims.push({
        name: t.name,
        datacategory: templateInfo.dataCategory,
        subcategory: templateInfo.subCategory,
        lowerBound: t.lowerBoundType, // 是否选择，是否包含 , 具体数值 ,  是否triger
        upperBound: t.upperBoundType,
        createDate: templateInfo.createdAt,
        expirationDate: templateInfo.expirationDate, // 过期时间
      })
    })

    let parms: object = {
      name: templateInfo.name,
      template: templateInfo.id,
      link: templateInfo.link,
      preClaims: JSON.stringify(preClaims),
    };

    const res: any = await api.offer.patch(templateInfo.claimId, parms);
      if (res.code === 200) {
        message.success('Offered claim');
        setTemplateInfo({});
        setActiveTabStr('claimList');
        setRouterName('claim');
      }
  }

  return (
    <div className="revocation-con">
      <div className="revocation-top">
        <div className="revocation-return" onClick={() => fallback()}><ArrowLeftOutlined /></div>
        <div className="revocation-des">
          <div>Set Revocation</div>
          <div>
            Set revocation triggers for claims with the support of dynamic verifications
          </div>
        </div>
      </div>
      <div className="revocation-form">
        <div className="revocation-form-title border-title">Revocation Settings</div>
        <div className="revocation-claims-item">
          <div className="revocation-form-title">Basics</div>
          <div className="revocation-base-item">
            <div className="revocation-base-info">
              <div className="info-common-style">
                <span>Template name:</span>
                <span>{templateInfo.name}</span>
              </div>
            </div>
            <div className="revocation-base-info">
              <div className="info-common-style">
                <span>Data Category:</span>
                <span>{dataCategoryList[Number(templateInfo.dataCategory)]}</span>
              </div>
              {
                templateInfo.dataCategory === '1' && (
                  <div className="info-common-style">
                    <span>NFT Contract:</span>
                    <span>{templateInfo.subCategory}</span>
                    <span><img
                      alt=""
                      src={IconCopy}
                      onClick={() => copyToClipboard(templateInfo.subCategory)}
                      className="copyIcon"
                    /></span>
                  </div>
                )
              }
              {
                templateInfo.dataCategory === '2' && (
                  <div className="info-common-style">
                    <span>Space ID:</span>
                    <span>{templateInfo.subCategory}</span>
                  </div>
                )
              }
            </div>
            <div className="revocation-base-info">
              <div className="info-common-style">
                <span>Creation Date:</span>
                <span>{templateInfo.createdAt ? templateInfo.createdAt.split('T')[0] : '--'}</span>
              </div>
              <div className="info-common-style">
                <span>Expiration Date:</span>
                <span>{templateInfo.expirationDate}</span>
              </div>
            </div>
            <div className="revocation-base-info">
              <div className="info-common-style">
                <span>Verification Link:</span>
                <span>{templateInfo.link}</span>
                <span><img
                  alt=""
                  src={IconCopy}
                  onClick={() => copyToClipboard(templateInfo.link)}
                  className="copyIcon"
                /></span>
              </div>
            </div>
          </div>
        </div>
        {classfications.map((item: any, index: number) =>
          <div className="revocation-claims-item" key={index}>
            <div className="revocation-form-title">Claims #{index + 1}</div>
            <div className="revocation-base-item">
              <div className="revocation-base-info">
                <div className="info-common-style">
                  <span>Class name:</span>
                  <span>{item.name}</span>
                </div>
              </div>
              <div className="revocation-base-info">
                <div className="info-common-style">
                  <span>Class Hash:</span>
                  <span>......</span>
                  <span><img
                    alt=""
                    src={IconCopy}
                    onClick={() => copyToClipboard('1')}
                    className="copyIcon"
                  /></span>
                </div>
              </div>
              <div className="revocation-base-info">
                {
                  item.lowerBoundType[0] === 1 && (
                    <div className="info-common-style">
                      <span>Lower Bound:</span>
                      <span>{item.lowerBoundType[1] === 0 ? '>' : '≥'}{item.lowerBoundType[2]}</span>
                    </div>
                  )
                }
                {
                  item.upperBoundType[0] === 1 && (
                    <div className="info-common-style">
                      <span>Upper Bound:</span>
                      <span>{item.upperBoundType[1] === 0 ? '<' : '≤'}{item.upperBoundType[2]}</span>
                    </div>
                  )
                }
              </div>
              <div className="revocation-trigger">
                <div>
                  Revocation Trigger
                </div>
                <div>
                  <Select value={item.triggerValue} style={{ width: '100%' }} onChange={(e) => handleChange(index, e)}>
                    <Option value="0">Never</Option>
                    <Option value="1">The lower bound is unmet</Option>
                    <Option value="2">The upper bound is unmet</Option>
                    <Option value="3">Either bound is unmet</Option>
                  </Select>
                </div>
              </div>
            </div>
          </div>
        )}

        <div className="button-group">
          <div><Button type="primary" size="large" onClick={() => finishOffer()}>
            Done
          </Button></div>
        </div>
      </div>
    </div>
  );
}
