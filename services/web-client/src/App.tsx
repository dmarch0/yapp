import React from 'react';
import { RecoilRoot } from 'recoil'
import Router from './components/Router/Router';

function App() {
  return (
    <RecoilRoot>
      <Router />
    </RecoilRoot>
  );
}

export default App;
