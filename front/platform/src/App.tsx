import React, { Suspense } from 'react';
import { Spin } from 'antd';
import { BrowserRouter as Router, Switch, Route, Redirect } from 'react-router-dom';
import Home from './views/home';
import QrCode from './views/qrcode';
import './App.css';
import './style/resetAntd.scss';

function App() {
  return (
    <Suspense>
      <Router>
        <Switch>
            <Route path="/home" component={Home} />
            <Route path="/page/:offerId" component={QrCode} />
            <Redirect to="/home" />
        </Switch>
      </Router>
    </Suspense>
  );
}

export default App;
