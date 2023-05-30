import { API_URL, NODE_ENV } from '@env';

const environment: Environment = {
  apiURL: API_URL,
  node: NODE_ENV,
};

type Environment = {
  apiURL: string;
  node: string;
};

export default environment;
