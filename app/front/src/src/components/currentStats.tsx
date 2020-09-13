import React from 'react';
import axios from 'axios';

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
    const axiosInstance = axios.create({
      baseURL: process.env.REACT_APP_BACKEND_ENDPOINT,
      headers: {
        'Content-Type': 'application/json',
      },
      responseType: 'json',
    });

    axiosInstance
      .get('/current')
      .then((response) => {
        console.log(response);
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
