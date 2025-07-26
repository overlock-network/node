FROM golang:1.24-alpine AS build-stage

LABEL org.opencontainers.image.source="https://github.com/overlock-network/node"

RUN apk add --no-cache git build-base wget

ADD https://github.com/CosmWasm/wasmvm/releases/download/v2.2.4/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
ADD https://github.com/CosmWasm/wasmvm/releases/download/v2.2.4/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a

RUN sha256sum /lib/libwasmvm_muslc.aarch64.a | grep 27fb13821dbc519119f4f98c30a42cb32429b111b0fdc883686c34a41777488f
RUN sha256sum /lib/libwasmvm_muslc.x86_64.a | grep 70c989684d2b48ca17bbd55bb694bbb136d75c393c067ef3bdbca31d2b23b578

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ /app


RUN CGO_ENABLED=1 go build \
    -tags="muslc" \
    -ldflags='-w -s -linkmode external -extldflags "-static"' \
    -o /overlockd /app/cmd/overlockd/main.go

FROM alpine:3.22

RUN mkdir -p /node/keys

WORKDIR /node

COPY --from=build-stage /overlockd /overlockd

COPY --from=build-stage /app/config /node/config

RUN wget -O /node/config/genesis.json https://raw.githubusercontent.com/overlock-network/net/refs/heads/main/devnet/genesis.json

EXPOSE 26657 1317 9090

COPY cluster/docker/node/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

CMD ["/overlockd", "start", "--home", "/node", "--rpc.laddr", "tcp://0.0.0.0:26657", "--grpc.enable", "--grpc.address", "0.0.0.0:9090", "--api.enable", "--api.enabled-unsafe-cors", "--api.address", "tcp://0.0.0.0:1317", "--minimum-gas-prices", "0stake"]
