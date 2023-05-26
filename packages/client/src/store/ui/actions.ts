import { PayloadAction } from '@reduxjs/toolkit';

import { UIStore } from './index';

const loading = (state: UIStore, action: PayloadAction<boolean>) => {
  state.loading = action.payload;
};

export default {
  loading,
};
