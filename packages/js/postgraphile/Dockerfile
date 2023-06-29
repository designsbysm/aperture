FROM node:18.16.1-alpine

WORKDIR /app

COPY ./packages/js/postgraphile/package.json .
RUN yarn install

COPY ./packages/js/postgraphile .
COPY .env .

CMD yarn start
