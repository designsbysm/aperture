FROM node:18.16.1-alpine

RUN npm install -g http-server
RUN npm install -g expo-cli

ARG API_URL
ENV API_URL ${API_URL}

ARG FONTAWESOME_NPM_TOKEN
RUN npm config set "@fortawesome:registry" https://npm.fontawesome.com/
RUN npm config set "//npm.fontawesome.com/:_authToken" ${FONTAWESOME_NPM_TOKEN}

WORKDIR /app

COPY ./packages/client/package.json .
RUN yarn install

COPY ./packages/client .
COPY ./packages/client/.env.docker .env
RUN npx expo export:web

CMD http-server web-build -p ${PORT_CLIENT}
