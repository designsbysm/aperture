FROM golang:1.19.3-alpine

WORKDIR /go/src/go-libs
COPY ./packages/go-libs .

WORKDIR /go/src/server
COPY .env .
COPY ./packages/service-logger .
RUN rm config.yaml
RUN mv config.prod.yaml config.yaml
# COPY ./packages/service-logger/config.prod.yaml config.yaml

CMD ["go", "run", "."]
