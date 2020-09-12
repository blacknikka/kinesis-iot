import React from 'react';
import {useDispatch} from 'react-redux';

import ICurrentStats from '../models/ICurrentStats';

const CurrentStats: React.FC<{stats: ICurrentStats}> = ({stats}) => {
  const dispatch = useDispatch();

  return (
    <div>
      <div>kind: {stats.kind}</div>
      <div>stats: {stats.stats}</div>
    </div>
  );
};

export default CurrentStats;
