import { configureStore } from '@reduxjs/toolkit';

import ui from './ui';
import { apiSlice } from './api';

export const store = configureStore({
  middleware: getDefaultMiddleware =>
    getDefaultMiddleware()
      .concat(apiSlice.middleware),
  reducer: {
    ui,
    [apiSlice.reducerPath]: apiSlice.reducer,
  },
});

export type AppStore = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
