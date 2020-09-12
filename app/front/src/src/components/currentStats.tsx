import React from 'react';
import {useDispatch} from 'react-redux';

import ICurrentStats from '../models/ICurrentStats';

const CurrentStats: React.FC<{stats: ICurrentStats}> = ({stats}) => {
  const style: any = {
    width: '40%',
    color: '#fff',
    background: '#639',
    margin:'auto',
  };
  return (
    <div style={style}>
      current stats:
      <div>kind: {stats.kind}</div>
      <div>stats: {stats.stats}</div>
    </div>
  );
};

export default CurrentStats;
