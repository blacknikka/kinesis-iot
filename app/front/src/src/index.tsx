import React from 'react';
import ReactDOM from 'react-dom';
import {Provider} from 'react-redux';
import {useDispatch, useSelector} from 'react-redux';

import ICurrentStats from './models/ICurrentStats';
import {setupStore} from './store';
import {IRootState} from './modules';
import CurrentStats from './components/currentStats';
import {currentActions} from './modules/statsModule';

const store = setupStore();

const App: React.FC = () => {
  const dispatch = useDispatch();
  const taskState = useSelector((state: IRootState) => state.currentStats);
  const cur: ICurrentStats = {
    kind: 'current',
    stats: 100,
  };
  return (
    <>
      <CurrentStats stats={cur} />
    </>
  );
};

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
);
