import {combineReducers} from 'redux';
import {configureStore, getDefaultMiddleware} from '@reduxjs/toolkit';
import statsModule from './modules/statsModule';

const rootReducer = combineReducers({
  task: statsModule.reducer,
});

export const setupStore = () => {
  const middleware = getDefaultMiddleware();
  return configureStore({
    reducer: rootReducer,
    middleware,
  });
};
