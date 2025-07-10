FROM golang:1.24 AS build-stage

LABEL org.opencontainers.image.source: "https://github.com/overlock-network/node"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ /app

RUN CGO_ENABLED=0 GOOS=linux go build -o /overlockd /app/cmd/overlockd/main.go 

FROM alpine:3.22

WORKDIR /

COPY --from=build-stage /overlockd /overlockd

COPY --from=build-stage /app/config /app/.overlockd/config

RUN wget -O /app/.overlockd/config/genesis.json https://raw.githubusercontent.com/overlock-network/net/refs/heads/main/devnet/genesis.json

EXPOSE 26657 1317 9090

ENTRYPOINT ["/overlockd", "start", "--rpc.laddr", "tcp://0.0.0.0:26657", "--grpc.enable", "--grpc.address", "0.0.0.0:9090", "--api.enable", "--api.enabled-unsafe-cors", "--api.address", "tcp://0.0.0.0:1317", "--minimum-gas-prices", "0stake", "--home", "/app/.overlockd"]
