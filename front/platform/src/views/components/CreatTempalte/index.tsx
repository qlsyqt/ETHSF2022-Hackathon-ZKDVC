import React, { useEffect, useState } from 'react';
import './index.scss';
import { dataCategoryList } from '../../../config'
import api from '../../../api';
import { ArrowLeftOutlined, InfoCircleOutlined } from '@ant-design/icons';
import { Checkbox, Radio, Input, Select, Button, message } from 'antd';
import { moduleActive, templateInfos } from '../../../store/atom';
import IconRemove from '../../../static/img/class-remove.png';
import { useRecoilState } from 'recoil';
const { Option } = Select;
const { TextArea } = Input;

const defaultClassficationItem = {
  name: '',
  lowerBoundType: [0, 0],
  upperBoundType: [0, 0],
  description: '',
}

export default function CreateTemplate() {
  const [activeTabStr, setActiveTabStr] = useRecoilState(moduleActive);

  const [templateInfo, setTemplateInfo] = useRecoilState(templateInfos);

  const [templateName, setTemplateName] = useState("");

  const [classfications, setClassfications] = useState<any>([defaultClassficationItem])

  const [dataCategory, setDataCategory] = useState('0');

  const [subCategory, setSubCategory] = useState('');

  const [isExpirable, setIsExpirable] = useState(false);

  const handleChange = (value: string) => {
    console.log(`selected ${value}`);
    setDataCategory(value)
  };

  const onClassificationChange = (index: number, field: string, value: string) => {
    setClassfications((prev: any) => {
      prev[index][field] = value;
      return [...prev]
    })
  }

  const onBoundChange = (index: number, field: string, subIndex: number, value: number) => {
    console.log('bond change', index, field, subIndex, value)
    setClassfications((prev: any) => {
      prev[index][field][subIndex] = value;
      if(subIndex === 0){
        prev[index][field][1] = 0;
      }
      return [...prev]
    })
  }

  const doCreate = async () => {
    let noInputChecked = classfications.some((t: any) => {
      // return (t.lowerBoundType[0] === 0 && t.upperBoundType[0] === 0) || !t.name;
      return !t.name;
    })

    if (noInputChecked || !templateName) {
      message.error('Please fill in required items.')
      return false;
    } else {
      const res: any = await api.template.create({
        name: templateName,
        dataCategory,
        isExpirable,
        subCategory,
        isAutoRevokable: true,
        classfications: JSON.stringify(classfications),
      })

      if (res.code === 200) {
        message.success('Created template')
        setActiveTabStr('templateList');
      }
      return res;
    }
  }

  const doCreateAndOffer = async () => {
    const res: any = await doCreate();
    if (res) {
      setActiveTabStr('offerClaims');
      setTemplateInfo(res.result);
    }
  };

  const addTemplate = async () => {
    setClassfications((prev: any) => [
      ...prev,
      {
        name: '',
        lowerBoundType: [0, 0],
        upperBoundType: [0, 0],
        description: '',
      }
    ])
  }

  const removeTemplate = async (index: number) => {
    setClassfications((prev: any) => {
      prev.splice(index, 1);
      return [...prev]
    })
  }

  const returnToList = () => {
    setActiveTabStr('templateList');
  }

  useEffect(() => {
    setSubCategory('');
  }, [dataCategory])

  useEffect(() => {
    setClassfications([{
      name: '',
      lowerBoundType: [0, 0],
      upperBoundType: [0, 0],
      description: '',
    }]);
  }, [])

  return (
    <div className="template-con">
      <div className="template-top">
        <div className="template-return" onClick={() => setActiveTabStr('templateList')}><ArrowLeftOutlined /></div>
        <div className="template-return" onClick={() => returnToList()}><ArrowLeftOutlined /></div>
        <div className="template-des">
          <div>Create Template</div>
          <div>Templates provide an easy way standardize a
            group of class (claims) width different data categories and bounds
          </div>
        </div>
      </div>
      <div className="template-form">
        <div className="template-form-title">Define Tempalate</div>
        <div className="template-form-input-item"><div>Template name<span className="require">*</span><span className="input-des">Only alphanumeric characters allowed. No spaces.</span></div>
          <div><Input placeholder="e.g. ENS-holding-number" value={templateName} onChange={e => setTemplateName(e.target.value)} /></div>
        </div>
        <div className="template-form-input-item">
          <div>Data Category<span className="require">*</span></div>
          <div>
            <Select value={dataCategory} style={{ width: '100%' }} onChange={handleChange}>
              {dataCategoryList.map((item: string, index: number) => <Option key={index} value={index.toString()}>{dataCategoryList[index]}</Option>)}
            </Select></div>
        </div>
        {dataCategory === "1" &&
          (<div className="template-form-input-item">
            <div>NFT Contract (Optional for Template)</div>
            <div>
              <Input placeholder="Enter NFT Contract" value={subCategory} onChange={e => setSubCategory(e.target.value)} />
            </div>
          </div>
          )}
        {dataCategory === "2" &&
          (<div className="template-form-input-item">
            <div>Space ID (Optional for Template)</div>
            <div>
              <Input placeholder="Enter Space ID" onChange={e => setSubCategory(e.target.value)} />
            </div>
          </div>
          )}
        {classfications.map((item: any, index: number) =>
          <div key={index} className="template-form-class-item">
            <div className="template-form-class-name class-item">Class #{index + 1}
              {
                index !== 0 &&
                <img src={IconRemove} className="icon-remove" onClick={() => removeTemplate(index)} />
              }
            </div>
            <div>
              <div className="template-form-input-item">
                <div>Class name<span className="require">*</span></div>
                <div><Input placeholder="e.g. Gold" value={item.name} onChange={(e) => onClassificationChange(index, 'name', e.target.value)} /></div>
              </div>
              <div className="template-form-bound">
                <div>Bounds
                  {/* <span className="require">*</span> 
                  (at least one) */}
                </div>
                <div className="template-form-bound-select">
                  <div>
                    <div>
                      <Checkbox checked={item.lowerBoundType[0] === 1} onChange={e => onBoundChange(index, 'lowerBoundType', 0, e.target.checked ? 1 : 0)}>Lower Bound{'(>)'}</Checkbox>
                    </div>
                    <div>
                      <Checkbox checked={item.upperBoundType[0] === 1} onChange={e => onBoundChange(index, 'upperBoundType', 0, e.target.checked ? 1 : 0)}>Upper Bound{'(<)'}</Checkbox>
                    </div>
                  </div>
                  <div>
                    <div>
                      {
                        item.lowerBoundType[0] === 1 &&
                        <Checkbox checked={item.lowerBoundType[1] === 1} onChange={e => onBoundChange(index, 'lowerBoundType', 1, e.target.checked ? 1 : 0)}>Include lower Bound(≥)</Checkbox>
                      }
                    </div>
                    <div>
                      {
                        item.upperBoundType[0] === 1 &&
                        <Checkbox checked={item.upperBoundType[1] === 1} onChange={e => onBoundChange(index, 'upperBoundType', 1, e.target.checked ? 1 : 0)}>Include upper Bound(≤)</Checkbox>
                      }
                    </div>
                  </div>
                </div>
              </div>
              <div className="template-form-input-item">
                <div>Classification description (Optional)</div>
                <div>
                  <TextArea rows={4} placeholder="Enter a description..." value={item.description} onChange={(e) => onClassificationChange(index, 'description', e.target.value)} />
                </div>
              </div>
            </div>
          </div>
        )}


        <div className="add-class" onClick={addTemplate}><span>+</span> Add Class (no intersection)</div>
        <div className="date-check">
          <div>
            <Checkbox checked={isExpirable} onChange={e => setIsExpirable(e.target.checked)}><span className="check-des">Mandatory claim expiration date (Optional)</span></Checkbox>
          </div>
          <div>
            <div className="date-des">When offering a claim,there will be a requirement to fill the expiration date.</div>
            <div className="date-des">Leaving this unchecked will keep the expiration date as optional.</div>
          </div>
        </div>
        <div className="button-group">
          <div onClick={doCreate}>
            Save template</div>
          <div><Button type="primary" size="large" onClick={doCreateAndOffer}>
            Save & Offer claims
          </Button></div>
        </div>
      </div>
    </div>
  );
}