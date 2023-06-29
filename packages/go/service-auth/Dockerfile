FROM golang:1.19.3-alpine

WORKDIR /go/src/libs
COPY ./packages/go/libs .

WORKDIR /go/src/types
COPY ./packages/go/types .

WORKDIR /go/src/server
COPY .env .
COPY ./packages/go/service-auth .
RUN rm config.yaml
RUN mv config.prod.yaml config.yaml

CMD ["go", "run", "."]
