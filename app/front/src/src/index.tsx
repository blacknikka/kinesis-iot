import React from 'react';
import ReactDOM from 'react-dom';
import {Provider} from 'react-redux';

import {setupStore} from './store';
import CurrentStats from './components/currentStats';
import SummaryStats from './components/summaryStats';

const store = setupStore();

const App: React.FC = () => {
  return (
    <>
      <CurrentStats />
      <SummaryStats />
    </>
  );
};

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
);
