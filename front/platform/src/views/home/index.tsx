import React, { useState, useEffect } from 'react';
import './index.scss';
import Logo from '../../static/img/logo.png';
import Claim from './../claim';
import Template from './../template';
import { BrowserRouter as Router, Switch, Route, Redirect, Link } from 'react-router-dom';
import { RouteComponentProps, useHistory } from 'react-router-dom';
import { moduleActive,routerNm } from '../../store/atom';
import { useRecoilState } from 'recoil';

export default function Home() {

  const [activeTabStr, setActiveTabStr] = useRecoilState(moduleActive);

  const [routerName, setRouterName] = useRecoilState<any>(routerNm);

  const history = useHistory();

  useEffect(() => {
    setComp(history.location.pathname);
  }, [history.location.pathname]);

  useEffect(() => {
    if(routerName === 'claim'){
      history.push(`/home/claim`);
      setRouterName('');
    }
  }, [routerName]);

  const routerTo = (str: string) => {
    history.push(`/home/${str}`);
    setComp(str);
  };

  const setComp = (str: string) => {
    if (str.includes('claim')) {
      setActiveTabStr('claimList')
    } else {
      setActiveTabStr('templateList')
    }
  };

  return (
    <div className="page-home">
      <div className="page-left-content">
        <div className="page-left-content-info">
          <div className="page-left-content-info-head">
            <img src={Logo} alt=""></img>
          </div>
          <div>
            <div className="page-left-content-info-name">
              KNN3 Network
            </div>
            <div className="page-left-content-info-email">
              account@knn3.xyz
            </div>
          </div>
        </div>
        <div className="page-left-router">
          <div onClick={() => routerTo('template')} className={history.location.pathname === '/home/template' || history.location.pathname === '/home' ? 'active' : ''}>
            Template
          </div>
          <div onClick={() => routerTo('claim')} className={history.location.pathname === '/home/claim' ? 'active' : ''}>Claim</div>
        </div>
      </div>
      <div className="page-right-content" key={history.location.key}>
        <Router>
          <Switch>
            <Route path="/home/claim" component={Claim} />
            <Route path="/home/template" component={Template} />
            <Redirect to="/home/template" />
          </Switch>
        </Router>
      </div>
    </div>
  );
}
