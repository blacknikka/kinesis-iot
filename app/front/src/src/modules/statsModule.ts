import {createSlice, PayloadAction} from '@reduxjs/toolkit';
import ICurrentStats from '../models/ICurrentStats';

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

export const {actions: currentActions} = currentStatsModule;
export default currentStatsModule;
