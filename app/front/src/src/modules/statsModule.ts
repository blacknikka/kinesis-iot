import {createSlice, PayloadAction, createAsyncThunk} from '@reduxjs/toolkit';
import ICurrentStats from '../models/ICurrentStats';
import axios from 'axios';

// const currentStatsReducer = (state = 0, action) => {
//   switch (action.type) {
//     case 'SET_CURRENT_STATS':
//       return state + action.payload;
//     default:
//       return state;
//   }
// };

export const fetchStats = createAsyncThunk<ICurrentStats>(
  'fetchCurrent',
  async (arg, thunk): Promise<ICurrentStats> => {
    const res = await axiosInstance
      .get('/current')
      .then((response) => {
        console.log(response.data);
        return {
          kind: response.data.kind,
          stats: response.data.stats,
        } as ICurrentStats;
      })
      .catch((err) => {
        console.log(err);
        throw new Error('fetch count error');
      });

    return res;
  }
);

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
  //   extraReducers: (builder) => {
  //     builder.addCase(
  //       fetchStats.fulfilled,
  //       (state, action: PayloadAction<ICurrentStats>) => {
  //         return {
  //           ...state,
  //           kind: action.payload.kind,
  //           state: action.payload.stats,
  //         };
  //         // state.kind = action.payload.kind;
  //         // state.stats = action.payload.stats;
  //       }
  //     );
  //   },
});

export const {setCurrentStats} = currentStatsModule.actions;
export default currentStatsModule;

const axiosInstance = axios.create({
  baseURL: process.env.REACT_APP_BACKEND_ENDPOINT,
  headers: {
    'Content-Type': 'application/json',
  },
  responseType: 'json',
});

export const fetchCurrentStats = () => {
  return (dispatch) => {
    const res = axiosInstance
      .get('/current')
      .then((response) => {
        console.log(response);
        return {
          kind: response.data.kind,
          stats: response.data.stats,
        };
      })
      .catch((err) => {
        console.log(err);
        throw new Error('something wrong.');
      });

    dispatch(
        setCurrentStats({
          kind: 'current',
          stats: 10,
        })
      );
  };
};
