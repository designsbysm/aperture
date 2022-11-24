FROM golang:1.19.3-alpine

WORKDIR /go/src/server
COPY .env .
COPY ./packages/service-authentication .
COPY ./packages/service-authentication/config.prod.yaml config.yaml

CMD ["go", "run", "."]
