#!/bin/sh
set -e

echo "[entrypoint] Starting Western Skies server..."
/usr/local/bin/western-skies &

echo "[entrypoint] Waiting for app to be ready..."
i=0
app_ready=false
while [ $i -lt 10 ]; do
  if wget -q -O /dev/null "http://127.0.0.1:${APP_PORT:-8080}/" 2>/dev/null; then
    echo "[entrypoint] App is ready."
    app_ready=true
    break
  fi
  i=$((i + 1))
  sleep 1
done

if [ "$app_ready" = false ]; then
  echo "[entrypoint] ERROR: App failed to start after 10 attempts. Exiting."
  exit 1
fi

echo "[entrypoint] Starting Caddy..."
exec caddy run --config /etc/caddy/Caddyfile --adapter caddyfile
