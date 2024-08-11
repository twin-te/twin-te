FROM oven/bun:1.1 AS local
WORKDIR /usr/src/twin-te/front

COPY package.json bun.lockb ./
RUN bun install --frozen-lockfile

COPY . .
CMD ["bunx", "--bun", "vite", "--host", "0.0.0.0", "--port", "5173"]

FROM oven/bun:1.1 AS staging-builder
WORKDIR /app

COPY package.json bun.lockb ./
RUN bun install --production --frozen-lockfile

COPY . .
RUN bun run build:staging

FROM nginx AS staging-deploy
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=staging-builder /app/dist /usr/share/nginx/html

FROM oven/bun:1.1 AS production-builder
WORKDIR /app

COPY package.json bun.lockb ./
RUN bun install --production --frozen-lockfile

COPY . .
RUN --mount=type=secret,id=SENTRY_AUTH_TOKEN \
  SENTRY_AUTH_TOKEN=$(cat /run/secrets/SENTRY_AUTH_TOKEN) \
  bun run build:production

FROM nginx AS production-deploy
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=production-builder /app/dist /usr/share/nginx/html
