import { AppStore } from '../index';

const loading = (state: AppStore) => state.ui?.loading;

export default {
  loading,
};
