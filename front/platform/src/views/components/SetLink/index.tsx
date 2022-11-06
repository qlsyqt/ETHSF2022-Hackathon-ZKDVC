import React, { useState, useEffect } from 'react';
import './index.scss';
import { ArrowLeftOutlined } from '@ant-design/icons';
import { Checkbox, Input, Button } from 'antd';
import { moduleActive, templateInfos, routerNm } from '../../../store/atom';
import { useRecoilState } from 'recoil';
import IconCopy from "./../../../static/img/copy.png";
import { copyToClipboard } from "./../../../utils/tools";
import { dataCategoryList } from '../../../config'

const defaultClassficationItem = {
  name: '',
  lowerBoundType: [0, 0, null],
  upperBoundType: [0, 0, null],
  description: '',
}

export default function CreateTemplate() {
  const [, setActiveTabStr] = useRecoilState(moduleActive);

  const [templateInfo, setTemplateInfo] = useRecoilState<any>(templateInfos);

  const [classfications, setClassfications] = useState<any>([defaultClassficationItem]);

  const [link, setLink] = useState("");

  const [routerName, setRouterName] = useRecoilState<any>(routerNm);

  useEffect(() => {
    setClassfications(JSON.parse(templateInfo.classfications));
    if (templateInfo.link) setLink(templateInfo.link);
  }, []);

  const toRevocation = () => {
    setTemplateInfo((prev: any) => {
      return {
        ...prev,
        link
      }
    });
    setActiveTabStr('revocation');
  }

  const fallback = () => {
    setActiveTabStr('claimList');
    setRouterName('claim');
    setTemplateInfo({});
  }

  return (
    <div className="link-con">
      <div className="link-top">
        <div className="link-return" onClick={() => fallback()}><ArrowLeftOutlined /></div>
        <div className="link-des">
          <div>Offer claims</div>
          <div>
            Define the attributes and bounds of the template to generate
            a verification link for a group of claims (classes) offering
          </div>
        </div>
      </div>
      <div className="link-form">
        <div className="link-form-title border-title">Claims Offering</div>
        {classfications.map((item: any, index: number) =>
          <div className="link-claims-item" key={index}>
            <div className="link-form-title">Claims #{index + 1}</div>
            <div className="link-base-item">
              <div className="link-base-info">
                <div className="info-common-style">
                  <span>Class name:</span>
                  <span>{item.name}</span>
                </div>
              </div>
              <div className="link-base-info">
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
              <div className="link-base-info">
                <div className="info-common-style">
                  <span>Lower Bound:</span>
                  <span>{item.lowerBoundType[1] === 0 ? '>' : '≥'}{item.lowerBoundType[2]}</span>
                </div>
                {
                  item.upperBoundType[0] === 1 && (
                    <div className="info-common-style">
                      <span>Upper Bound:</span>
                      <span>{item.upperBoundType[1] === 0 ? '<' : '≤'}{item.upperBoundType[2]}</span>
                    </div>
                  )
                }
              </div>
              <div className="link-base-info">
                <div className="info-common-style">
                  <span>Creation Date:</span>
                  <span>{templateInfo.createdAt ? templateInfo.createdAt.split('T')[0] : '--'}</span>
                </div>
                <div className="info-common-style">
                  <span>Expiration Date:</span>
                  <span>{templateInfo.expirationDate}</span>
                </div>
              </div>
            </div>
          </div>
        )}

        <div className="link-url">
          <div>Verification Link</div>
          <div>{templateInfo.link}
            <span>
              <img
                alt=""
                src={IconCopy}
                onClick={() => copyToClipboard(templateInfo.link)}
                className="copyIcon"
              />
            </span>
          </div>
        </div>
        <div className="button-group">
          <div><Button type="primary" size="large" onClick={() => toRevocation()}>
            Set Revocation
          </Button></div>
        </div>
      </div>
    </div>
  );
}
