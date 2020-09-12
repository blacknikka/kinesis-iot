import React from 'react';
import {useDispatch} from 'react-redux';

import ISummaryStats from '../models/ISummaryStats';

const SummaryStats: React.FC<{stats: ISummaryStats}> = ({stats}) => {
  const style: any = {
    paddingTop: '3%',
    width: '40%',
    color: '#fff',
    background: '#639',
    margin: 'auto',
  };
  console.log(stats.summary);
  return (
    <div style={style}>
      summary stats:
      <div>kind: {stats.kind}</div>
      <div>summary:</div>
      {Object.keys(stats.summary).map((key) => {
        return (
          <div key={key}>
            <p>{key}:</p>
            <div>{stats.summary[key]}</div>
          </div>
        );
      })}
    </div>
  );
};

export default SummaryStats;
