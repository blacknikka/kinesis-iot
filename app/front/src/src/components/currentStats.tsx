import React from 'react';
import axios from '../modules/axios';

interface State {
  kind: string;
  stats: number;
}
class CurrentStats extends React.Component<{}, State> {
  constructor(props) {
    super(props);
    this.state = {
      kind: '',
      stats: 0,
    };

    axios
      .get('/current')
      .then((response) => {
        this.setState({
          kind: response.data.kind,
          stats: response.data.stats,
        });
      })
      .catch((err) => {
        console.log(err);
        throw new Error('something wrong.');
      });
  }

  render() {
    const style: any = {
      width: '40%',
      color: '#fff',
      background: '#639',
      margin: 'auto',
    };

    return (
      <div style={style}>
        current stats:
        <div>kind: {this.state.kind}</div>
        <div>stats: {this.state.stats}</div>
      </div>
    );
  }
}

export default CurrentStats;
