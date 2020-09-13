import {createSlice, PayloadAction, createAsyncThunk} from '@reduxjs/toolkit';
import ICurrentStats from '../models/ICurrentStats';
import axios from 'axios';

const currentStatsModule = createSlice({
  name: 'current',
  initialState: {
    kind: '',
    stats: 0,
  },
  reducers: {
    setCurrentStats: (state, action: PayloadAction<ICurrentStats>) => ({
      ...state,
      stats: action.payload.stats,
      kind: action.payload.kind,
    }),
  },
});

export const {setCurrentStats} = currentStatsModule.actions;
export default currentStatsModule;
