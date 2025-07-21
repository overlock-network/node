#!/bin/bash
set -e

envsubst < /etc/nginx/nginx.conf.template > /etc/nginx/nginx.conf

echo "[debug] RPC_DOMAIN: $RPC_DOMAIN"
echo "[debug] API_DOMAIN: $API_DOMAIN"
echo "[debug] GRPC_DOMAIN: $GRPC_DOMAIN"
echo "[debug] overlockd args: $@"

export OVERLOCKD_ARGS="$*"

envsubst < /etc/supervisor/supervisord.conf.template > /etc/supervisor/supervisord.conf

exec /usr/bin/supervisord -c /etc/supervisor/supervisord.conf