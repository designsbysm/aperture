import * as dotenv from 'dotenv';
import * as dotenvExpand from 'dotenv-expand';
import * as fs from 'fs';
import * as http from 'http';
import { postgraphile } from 'postgraphile';

const envFiles = [
  '.env',
  '../.env',
  '../../.env',
  '../../../.env',
];
const found = envFiles.find(item => fs.existsSync(item));

if (!!found) {
  const envConfig = dotenv.config({ path: found });
  dotenvExpand.expand(envConfig);
}

http
  .createServer(postgraphile(
    process.env.DATABASE_URL,
    'northwind',
    {
      enhanceGraphiql: true,
      graphiql: true,
      watchPg: true,
    },
  ))
  .listen(process.env.PORT_GRAPHQL);
