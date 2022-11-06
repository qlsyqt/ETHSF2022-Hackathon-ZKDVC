import React, { useState, useEffect } from 'react';
import './index.scss';
import { RouteComponentProps } from 'react-router-dom';
import QrImg from "./../../static/img/QR_ready.png";
import { Button } from 'antd';
import QRCode  from 'qrcode.react';
import { copyToClipboard } from "./../../utils/tools";
import api from '../../api';

export default function QrCode(props: RouteComponentProps) {

  const [jsonData, setJsonData] = useState('');

  const getQrCode = async () => {
    const res: any = await api.offer.getQrCode(props?.match?.params);
    setJsonData(JSON.stringify(res.result))
  };

  useEffect(() => {
    document.title = 'Claim Offer'
    getQrCode();
  }, []);

  return (
    <div className='qr'>
      <div>
        <img
          alt=""
          src={QrImg}
        />
      </div>
      <div className='title1'>You received a claim offer from account@knn3</div>
      <div>
        {
          jsonData && <QRCode value={jsonData} size={256} />
        }
        
      </div>
      <div className='btn'>
        <Button type="primary" size="large" onClick={() => copyToClipboard(jsonData)}>
          Copy
        </Button>
      </div>
      <div className='title2'>Scan the QR code to add the chaim to your wallet</div>
    </div>
  );
}
