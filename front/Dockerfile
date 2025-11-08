ARG BUN_VERSION=1.1.22
FROM oven/bun:${BUN_VERSION}-slim AS local
WORKDIR /usr/src/twin-te/front

COPY package.json bun.lockb ./
RUN bun install --frozen-lockfile

COPY . .
ENV NODE_ENV development
EXPOSE 5173
CMD ["bun", "run", "vite", "--host", "0.0.0.0", "--port", "5173"]

FROM oven/bun:${BUN_VERSION}-slim AS staging-builder
WORKDIR /app

RUN --mount=type=cache,target=~/.bun/install/cache \
  --mount=type=bind,source=package.json,target=package.json \
  --mount=type=bind,source=bun.lockb,target=bun.lockb \
  bun install --frozen-lockfile

COPY . .
RUN bun run build:staging

FROM caddy:2 AS staging-deploy
COPY Caddyfile /etc/caddy/Caddyfile
COPY --from=staging-builder /app/dist /usr/share/caddy

FROM oven/bun:${BUN_VERSION}-slim AS production-builder
WORKDIR /app

RUN --mount=type=cache,target=~/.bun/install/cache \
  --mount=type=bind,source=package.json,target=package.json \
  --mount=type=bind,source=bun.lockb,target=bun.lockb \
  bun install --frozen-lockfile

COPY . .
RUN --mount=type=cache,target=/var/lib/apt,sharing=locked \
  --mount=type=cache,target=/var/cache/apt,sharing=locked \
  DEBIAN_FRONTEND=noninteractive apt-get update && apt-get install -y ca-certificates
RUN --mount=type=secret,id=SENTRY_AUTH_TOKEN \
  SENTRY_AUTH_TOKEN=$(cat /run/secrets/SENTRY_AUTH_TOKEN) \
  bun run build:production

FROM caddy:2 AS production-deploy
COPY Caddyfile /etc/caddy/Caddyfile
COPY --from=production-builder /app/dist /usr/share/caddy
