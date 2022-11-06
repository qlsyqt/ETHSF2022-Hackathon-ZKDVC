import React, { useState } from 'react';
import './index.scss';
import CreatTempalte from './../components/CreatTempalte';
import TemplateList from './../components/TemplateList';
import OfferClaims from './../components/OfferClaims';
import Revocation from './../components/Revocation';
import SetLink from './../components/SetLink';
import { moduleActive } from '../../store/atom';
import { useRecoilState } from 'recoil';

export default function List() {

  const [activeTabStr, setActiveTabStr] = useRecoilState(moduleActive);

  return (
    <div className="template">
      {activeTabStr === 'templateList' && (
        <TemplateList />
      )}
      {activeTabStr === 'creatTempalte' && (
        <CreatTempalte />
      )}
      {activeTabStr === 'offerClaims' && (
        <OfferClaims />
      )}
      {activeTabStr === 'revocation' && (
        <Revocation />
      )}
      {activeTabStr === 'setLink' && (
        <SetLink />
      )}
    </div>
  );
}
