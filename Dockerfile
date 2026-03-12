# =============================================================================
# Stage 1: Generate templ files and build Go binary
# =============================================================================
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git

# Install templ CLI
RUN go install github.com/a-h/templ/cmd/templ@v0.3.1001

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN templ generate
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o western-skies ./cmd/server

# =============================================================================
# Stage 2: Final image — Caddy + Go binary + static assets
# =============================================================================
FROM caddy:2-alpine

# Copy Go binary
COPY --from=builder /build/western-skies /usr/local/bin/western-skies

# Copy content and static files needed at runtime
COPY --from=builder /build/content /app/content
COPY --from=builder /build/static /app/static

# Copy Caddy configuration
COPY Caddyfile /etc/caddy/Caddyfile

# Copy and set entrypoint
COPY docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh
RUN chmod +x /usr/local/bin/docker-entrypoint.sh /usr/local/bin/western-skies

WORKDIR /app

ENV PORT=80 \
    APP_PORT=8080 \
    SITE_CONFIG=content/site.toml

EXPOSE 80

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD wget -qO- http://127.0.0.1:${PORT:-80}/health || exit 1

ENTRYPOINT ["/usr/local/bin/docker-entrypoint.sh"]
