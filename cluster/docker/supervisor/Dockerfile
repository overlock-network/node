ARG BASE_IMAGE=ghcr.io/overlock-network/node:latest
FROM ${BASE_IMAGE}

RUN apk add --no-cache \
    nginx \
    supervisor \
    curl \
    ca-certificates \
    bash \
    libc6-compat \
    gettext

ENV RPC_DOMAIN=rpc.default.local
ENV API_DOMAIN=api.default.local
ENV GRPC_DOMAIN=grpc.default.local

COPY cluster/docker/supervisor/nginx.conf.template /etc/nginx/nginx.conf.template
COPY cluster/docker/supervisor/supervisord.conf.template /etc/supervisor/supervisord.conf.template

COPY cluster/docker/supervisor/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

EXPOSE 80
