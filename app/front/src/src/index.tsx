import React from 'react';
import ReactDOM from 'react-dom';
import {Provider} from 'react-redux';
import {useDispatch, useSelector} from 'react-redux';
import statsModule, {
    fetchCurrentStats,
  } from './modules/statsModule';
  
import ICurrentStats from './models/ICurrentStats';
import ISummaryStats from './models/ISummaryStats';
import {setupStore} from './store';
import {IRootState} from './modules';
import CurrentStats from './components/currentStats';
import SummaryStats from './components/summaryStats';

const store = setupStore();

const App: React.FC = () => {
  const dispatch = useDispatch();
  dispatch(fetchCurrentStats());

  const summary: ISummaryStats = {
    kind: 'current',
    summary: {
      total: 500,
      event: 200,
    },
  };
  return (
    <>
      <CurrentStats />
      <SummaryStats stats={summary} />
    </>
  );
};

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
);
