import {combineReducers, createStore, applyMiddleware} from 'redux';
import {configureStore} from '@reduxjs/toolkit';
import thunk from 'redux-thunk';
import statsModule from './modules/statsModule';

const middlewares = [thunk];

const rootReducer = combineReducers({
  currentStats: statsModule.reducer,
});

export const setupStore = () => {
  return createStore(
    rootReducer,
    applyMiddleware(...middlewares)
  );
};
