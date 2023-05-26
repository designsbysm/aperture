import { createSlice } from '@reduxjs/toolkit';

import actions from './actions';
import selectors from './selectors';

const initialState: UIStore = {
  loading: false,
};

const slice = createSlice({
  initialState,
  name: 'ui',
  reducers: {
    loading: actions.loading,
  },
});

export interface UIStore {
  loading: boolean;
}

export default slice.reducer;

export const {
  loading,
} = slice.actions;

export const selectLoading = selectors.loading;
