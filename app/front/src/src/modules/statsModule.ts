import {createSlice, PayloadAction} from '@reduxjs/toolkit';
import ICurrentStats from '../models/ICurrentStats';
import axios from 'axios';

const currentStatsModule = createSlice({
  name: 'current',
  initialState: {
    kind: '',
    stats: 0,
  },
  reducers: {
    setCurrentStats: {
      reducer: (state: ICurrentStats, action: PayloadAction<ICurrentStats>) => {
        state = action.payload;
      },
      prepare: (current) => {
        return {payload: {kind: current.kind, stats: current.stats}};
      },
    },
  },
});

const axiosInstance = axios.create({
  baseURL: process.env.REACT_APP_BACKEND_ENDPOINT,
  headers: {
    'Content-Type': 'application/json',
  },
  responseType: 'json',
});

export const {actions: currentActions} = currentStatsModule;
export default currentStatsModule;

export const fetchCurrentStats = () => {
  return async (dispatch, getState) => {
    axiosInstance
      .get('/current')
      .then((response) => {
        dispatch(
          currentStatsModule.actions.setCurrentStats({
            kind: response.data.kind,
            stats: response.data.stats,
          })
        );
      })
      .catch((err) => {
        console.log(err);
      });
  };
};
