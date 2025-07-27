FROM node:22-bullseye-slim AS local
WORKDIR /usr/src/twin-te/front

RUN npm install -g bun@v1.1.22

COPY package.json bun.lockb ./
RUN bun install --frozen-lockfile

COPY . .
ENV NODE_ENV development
EXPOSE 5173
CMD ["bun", "run", "vite", "--host", "0.0.0.0", "--port", "5173"]

FROM node:22-bullseye-slim AS staging-builder
WORKDIR /app

RUN npm install -g bun@v1.1.22

COPY package.json bun.lockb ./
RUN bun install --frozen-lockfile

COPY . .
RUN bun run build:staging

FROM nginx AS staging-deploy
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=staging-builder /app/dist /usr/share/nginx/html

FROM node:22-bullseye-slim AS production-builder
WORKDIR /app

RUN npm install -g bun@v1.1.22

COPY package.json bun.lockb ./
RUN bun install --frozen-lockfile

COPY . .
RUN DEBIAN_FRONTEND=noninteractive apt-get update && apt-get install -y ca-certificates
RUN --mount=type=secret,id=SENTRY_AUTH_TOKEN \
  SENTRY_AUTH_TOKEN=$(cat /run/secrets/SENTRY_AUTH_TOKEN) \
  bun run build:production

FROM nginx AS production-deploy
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=production-builder /app/dist /usr/share/nginx/html
