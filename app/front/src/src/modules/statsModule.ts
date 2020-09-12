import {createSlice, PayloadAction} from '@reduxjs/toolkit';
import ICurrentStats from '../models/ICurrentStats';
import axios from 'axios';

const initialCurrentStats: ICurrentStats = {
  kind: '',
  stats: 0,
};

const currentStatsModule = createSlice({
  name: 'current',
  initialState: initialCurrentStats,
  reducers: {
    setCurrentStats: (
      state: ICurrentStats,
      action: PayloadAction<ICurrentStats>
    ) => {
      state = action.payload;
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

export const fetchCurrentStats = () => {
  return async (dispatch, getState) => {
    const {current} = getState();
    axiosInstance
      .get('/current')
      .then((response) => {
        console.log(response);
        console.log(current);
      })
      .catch((err) => {
        console.log(err);
      });
  };
};

export const {actions: currentActions} = currentStatsModule;
export default currentStatsModule;
