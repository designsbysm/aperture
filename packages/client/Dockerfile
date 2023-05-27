FROM node:16.14.0-alpine

RUN apk update
RUN apk add git

RUN npm install -g http-server
RUN npm install -g expo-cli

ARG API_URL
ENV API_URL ${API_URL}

ARG FONTAWESOME_NPM_TOKEN
RUN npm config set "@fortawesome:registry" https://npm.fontawesome.com/
RUN npm config set "//npm.fontawesome.com/:_authToken" ${FONTAWESOME_NPM_TOKEN}

WORKDIR /client

COPY ./packages/client/package.json .
RUN npm install --legacy-peer-deps

COPY ./packages/client .
COPY ./packages/client/.env.docker .env
RUN npx expo build:web

CMD http-server web-build -p ${PORT_CLIENT}
