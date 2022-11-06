import React, { useState, useEffect } from 'react';
import './index.scss';
import { ArrowLeftOutlined } from '@ant-design/icons';
import { Checkbox, Input, Button, DatePicker, message } from 'antd';
import { moduleActive, templateInfos } from '../../../store/atom';
import { useRecoilState } from 'recoil';
import { useHistory } from 'react-router-dom';
import IconCopy from "./../../../static/img/copy.png";
import { copyToClipboard } from "./../../../utils/tools";
import { dataCategoryList } from '../../../config';
import api from '../../../api';
import moment from 'moment';

const defaultClassficationItem = {
  name: '',
  lowerBoundType: [0, 0, null],
  upperBoundType: [0, 0, null],
  description: '',
}

const dateFormat = 'YYYY/MM/DD';

export default function CreateTemplate() {
  const [, setActiveTabStr] = useRecoilState(moduleActive);

  const [templateInfo, setTemplateInfo] = useRecoilState<any>(templateInfos);

  const [classfications, setClassfications] = useState<any>([defaultClassficationItem])

  const [expirationDate, setExpirationDate] = useState("")

  const [subCategory, setSubCategory] = useState("")

  const history = useHistory();

  const onChange = (index: number, field: string, subIndex: number, value: string) => {
    const reg = /^-?\d*(\.\d*)?$/;
    if (reg.test(value) || value === '' || value === '-') {
      setClassfications((prev: any) => {
        prev[index][field][subIndex] = Number(value);
        return [...prev]
      })
    }
  }

  useEffect(() => {
    let classItems: any = JSON.parse(templateInfo.classfications);
    classItems.map((t: any) => {
      t.lowerBoundType[2] = t.lowerBoundType[2] == undefined ? null : t.lowerBoundType[2];
      t.upperBoundType[2] = t.upperBoundType[2] == undefined ? null : t.upperBoundType[2];
      if (t.lowerBoundType[0] === 0) {
        t.lowerBoundType[2] = 0
      }
    })
    setClassfications(classItems);
  }, []);

  const maxArrObj = (arr: any, key: any) => {
    return Math.max.apply(Math, arr.map((o: any) => {
      return o[key]
    }));
  }

  const isIntersect = (arr1: any, arr2: any, isContain: any) => {
    let start = [Math.min(...arr1), Math.min(...arr2)];//区间的两个最小值
    let end = [Math.max(...arr1), Math.max(...arr2)];//区间的两个最大值
    if (isContain == 0) {
      return Math.max(...start) < Math.min(...end);//最大值里的最小值 是否 小于等于 最大值的最小值
    } else {
      return Math.max(...start) <= Math.min(...end);//最大值里的最小值 是否 小于等于 最大值的最小值
    }
  }

  const onSubChange = (e: any) => {
    setSubCategory(e);
  }

  const toLink = async () => {
    let preClaims: any = [];
    let errorMsg: string = "";
    let isClassError: boolean = false;
    let upperBoundValue: any = [];
    // let maxObj: any = {
    //   value: 0,
    //   isbH: 0
    // };

    classfications.map((t: any, i: number) => {

      // if (t.lowerBoundType[2] > maxObj.value) {
      //   maxObj = {
      //     value: t.lowerBoundType[2],
      //     isbH: t.lowerBoundType[1]
      //   }
      // }

      if ((t.lowerBoundType[0] === 1 && t.lowerBoundType[2] === null) ||
        (t.upperBoundType[0] === 1 && t.upperBoundType[2] === null)) {
        errorMsg = 'Please fill in all values';
      } else if (t.lowerBoundType[0] === 1 && t.upperBoundType[0] === 1) {
        if (t.lowerBoundType[2] >= t.upperBoundType[2]) {
          errorMsg = 'The upper bound must be greater than the lower bound';
        }
      }

      classfications.map((r: any, h: number) => {
        if (i !== h) {
          if (t.upperBoundType[0] == 1 && r.upperBoundType[0] == 1) {
            const s = [t.lowerBoundType[2], t.upperBoundType[2]];
            const e = [r.lowerBoundType[2], r.upperBoundType[2]];
            const b = t.lowerBoundType[1], h = t.upperBoundType[1],
              y = r.lowerBoundType[1], o = r.upperBoundType[1];
            const sMax = Math.max(...s), sMin = Math.min(...s);
            const eMax = Math.max(...e), eMin = Math.min(...e);
            if (sMax <= eMin) {
              if ((h == 1 && y == 1)) {
                if (isIntersect(s, e, 1)) {
                  isClassError = true;
                }
              }
              if ((h == 0 && y == 0)) {
                if (isIntersect(s, e, 0)) {
                  isClassError = true;
                }
              }
              if ((h == 0 && y == 1) || (h == 1 && y == 0)) {
                if (isIntersect(s, e, 0)) {
                  isClassError = true;
                }
              }
            } else if (eMax <= sMin) {
              if ((o == 1 && b == 1)) {
                if (isIntersect(s, e, 1)) {
                  isClassError = true;
                }
              }
              if ((o == 0 && b == 0)) {
                if (isIntersect(s, e, 0)) {
                  isClassError = true;
                }
              }
              if ((o == 0 && b == 1) || (o == 1 && b == 0)) {
                if (isIntersect(s, e, 0)) {
                  isClassError = true;
                }
              }
            } else {
              if (isIntersect(s, e, 0)) {
                isClassError = true;
              }
            }
          }
        }
      })

      upperBoundValue.push({
        value: t.upperBoundType[2],
        isbH: t.upperBoundType[1]
      })

      if (t.lowerBoundType.length == 3 && t.upperBoundType.length == 3) {
        t.lowerBoundType.push(0);
        t.upperBoundType.push(0);
      }

      preClaims.push({
        name: t.name,
        datacategory: templateInfo.dataCategory,
        subcategory: templateInfo.subCategory ? templateInfo.subCategory : subCategory,
        lowerBound: t.lowerBoundType, // 是否选择，是否包含 , 具体数值 ,  是否triger
        upperBound: t.upperBoundType,
        createDate: templateInfo.createdAt,
        expirationDate: templateInfo.expirationDate, // 过期时间
      })

    })

    let s = classfications.filter((t: any) => {
      return t.lowerBoundType[0] == 0 && t.upperBoundType[0] == 0
    })

    if (errorMsg) {
      message.error(errorMsg);
      return false;
    }

    if (isClassError) {
      message.error('The classes cannot overlap (even the same number).');
      return false;
    }

    if (s.length > 1) {
      message.error('The classes cannot overlap (even the same number).');
      return false;
      // if (maxObj.isbH == 0) {
      //   if (s[0].lowerBoundType[1] == 0) {
      //     if (s[0].lowerBoundType[2] < maxObj.value) {
      //       message.error('The classes cannot overlap (even the same number).');
      //       return false;
      //     }
      //   }
      // }
    }

    let parms: object = {
      name: templateInfo.name,
      template: templateInfo.id,
      preClaims: JSON.stringify(preClaims),
    };

    const res: any = await api.offer.create(parms);

    if (res.code === 200) {

      let parms1: object = {
        name: templateInfo.name,
        template: templateInfo.id,
        preClaims: JSON.stringify(preClaims),
        link: `${window.location.href.split('/home/template')[0]}/page/${res.result.id}`
      };

      const res1: any = await api.offer.patch(res.result.id, parms1);

      if (res1.code === 200) {

        setTemplateInfo((prev: any) => {
          let obj = { ...prev };
          if (subCategory) {
            obj.subCategory = subCategory
          }
          return {
            ...obj,
            classfications: JSON.stringify(classfications),
            expirationDate,
            link: `${window.location.href.split('/home/template')[0]}/page/${res.result.id}`,
            claimId: res.result.id
          }
        })

        setActiveTabStr('setLink');

      }
    }
  }

  return (
    <div className="claim-con">
      <div className="claim-top">
        <div className="claim-return" onClick={() => setActiveTabStr('templateList')}><ArrowLeftOutlined /></div>
        <div className="claim-des">
          <div>Offer claims</div>
          <div>
            Define the attributes and bounds of the template to generate
            a verification link for a group of claims (classes) offering
          </div>
        </div>
      </div>
      <div className="claim-form">
        <div className="claim-form-title">Template for claims</div>
        <div className="claim-base-item">
          <div className="claim-base-info">
            <div className="info-common-style">
              <span>Template name:</span>
              <span>{templateInfo.name}</span>
            </div>
          </div>
          <div className="claim-base-info">
            <div className="info-common-style">
              <span>Data Category:</span>
              <span>{dataCategoryList[Number(templateInfo.dataCategory)]}</span>
            </div>
            {
              templateInfo.dataCategory === '1' && templateInfo.subCategory && (
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
              templateInfo.dataCategory === '2' && templateInfo.subCategory && (
                <div className="info-common-style">
                  <span>Space ID:</span>
                  <span>{templateInfo.subCategory}</span>
                </div>
              )
            }
          </div>
          {
            templateInfo.dataCategory === '1' && !templateInfo.subCategory && (
              <div>
                <div className="info-common-style">
                  <span>NFT Contract:</span>
                </div>
                <div><Input value={subCategory} onChange={e => onSubChange(e.target.value)} /></div>
              </div>
            )
          }
          {
            templateInfo.dataCategory === '2' && !templateInfo.subCategory && (
              <div>
                <div className="info-common-style">
                  <span>Space ID:</span>
                </div>
                <div><Input value={subCategory} onChange={e => onSubChange(e.target.value)} /></div>
              </div>
            )
          }
        </div>
        {classfications.map((item: any, index: number) =>
          <div className="claim-claims-item" key={index}>
            <div className="claim-form-title">Claim #{index + 1}</div>
            <div className="info-common-style">
              <span>Class Name:</span>
              <span>{item.name}</span>
            </div>
            <div className="info-common-style">
              <span>Class Hash:</span>
              <span>hash</span>
              <span><img
                alt=""
                src={IconCopy}
                onClick={() => copyToClipboard('1')}
                className="copyIcon"
              /></span>
            </div>
            <div className="claim--bound">
              <div>
                <div>Lower Bound:</div>
                <div><Input value={item.lowerBoundType[2]} disabled={item.lowerBoundType[0] === 0} onChange={e => onChange(index, 'lowerBoundType', 2, e.target.value)} /></div>
                <div><Checkbox checked={item.lowerBoundType[1] === 1} disabled>Include lower Bound(≥)</Checkbox></div>
              </div>
              {
                item.upperBoundType[0] === 1 &&
                <div>
                  <div>Upper Bound:</div>
                  <div><Input value={item.upperBoundType[2]} onChange={e => onChange(index, 'upperBoundType', 2, e.target.value)} /></div>
                  <div><Checkbox checked={item.upperBoundType[1] === 1} disabled>Include Upper Bound(≤)</Checkbox></div>
                </div>
              }
            </div>
          </div>
        )}
        {
          templateInfo.isExpirable &&
          <div>
            <div className="claim-form-title">Claims expiration date</div>
            <div><DatePicker format={dateFormat} value={expirationDate ? moment(expirationDate, dateFormat) : null} onChange={(date: any, dateString: string) => setExpirationDate(dateString)} /></div>
          </div>
        }
        <div className="button-group">
          <div><Button type="primary" size="large" onClick={() => toLink()}>
            Generate Verification Link
          </Button></div>
        </div>
      </div>
    </div>
  );
}
