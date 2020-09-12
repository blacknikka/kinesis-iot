import React from 'react';
import ReactDOM from 'react-dom';
import {Provider} from 'react-redux';
import {useDispatch, useSelector} from 'react-redux';

import ICurrentStats from './models/ICurrentStats';
import ISummaryStats from './models/ISummaryStats';
import {setupStore} from './store';
import {IRootState} from './modules';
import CurrentStats from './components/currentStats';
import SummaryStats from './components/summaryStats';
import {currentActions} from './modules/statsModule';

const store = setupStore();

const App: React.FC = () => {
  const dispatch = useDispatch();
  const taskState = useSelector((state: IRootState) => state.currentStats);
  const cur: ICurrentStats = {
    kind: 'current',
    stats: 100,
  };
  const summary: ISummaryStats = {
    kind: 'current',
    summary: {
      total: 500,
      event: 200,
    },
  };
  return (
    <>
      <CurrentStats stats={cur} />
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
