import React from 'react';
import axios from '../modules/axios';

interface State {
  kind: string;
  [key: string]: any;
}
export default class SummaryStats extends React.Component<{}, State> {
  constructor(props) {
    super(props);
    this.state = {
      kind: '',
      key: '',
    };

    axios
      .get('/summary')
      .then((response) => {
        console.log(response.data);
        this.setState({
          kind: response.data.kind,
          key: response.data.summary,
        });
      })
      .catch((err) => {
        console.log(err);
        throw new Error('something wrong.');
      });
  }

  render() {
    const style: any = {
      paddingTop: '3%',
      width: '40%',
      color: '#fff',
      background: '#639',
      margin: 'auto',
    };
    return (
      <div style={style}>
        summary stats:
        <div>kind: {this.state.kind}</div>
        <div>summary:</div>
        {Object.keys(this.state.key).map((key) => {
          return (
            <div key={key}>
              <p>{key}:</p>
              <div>{this.state.key[key]}</div>
            </div>
          );
        })}
      </div>
    );
  }
}
