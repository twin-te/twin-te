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

COPY package.json bun.lockb ./
RUN bun install --frozen-lockfile

COPY . .
RUN bun run build:staging

FROM nginx AS staging-deploy
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=staging-builder /app/dist /usr/share/nginx/html

FROM oven/bun:${BUN_VERSION}-slim AS production-builder
WORKDIR /app

COPY package.json bun.lockb ./
RUN bun install --frozen-lockfile

COPY . .
RUN --mount=type=cache,target=/var/lib/apt,sharing=locked \
  --mount=type=cache,target=/var/cache/apt,sharing=locked \
  apt-get update && apt-get install -y ca-certificates
RUN --mount=type=secret,id=SENTRY_AUTH_TOKEN \
  SENTRY_AUTH_TOKEN=$(cat /run/secrets/SENTRY_AUTH_TOKEN) \
  bun run build:production

FROM nginx AS production-deploy
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=production-builder /app/dist /usr/share/nginx/html
