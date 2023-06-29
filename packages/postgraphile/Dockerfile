FROM node:18.16.1-alpine

WORKDIR /app

COPY ./packages/postgraphile/package.json .
RUN yarn install

COPY ./packages/postgraphile .

CMD yarn start
