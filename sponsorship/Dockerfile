FROM node:22-bullseye-slim AS local
WORKDIR /usr/src/twin-te/sponsorship

RUN npm install -g bun@v1.1.22

COPY package.json bun.lockb ./
RUN bun install --frozen-lockfile

COPY . .
ENV NODE_ENV development
EXPOSE 3000
CMD ["bun", "run", "dev", "--hostname", "0.0.0.0", "--port", "3000"]

FROM node:22-bullseye-slim AS staging-builder
WORKDIR /app

RUN npm install -g bun@v1.1.22

COPY package.json bun.lockb ./
RUN bun install --frozen-lockfile

COPY . .
COPY .env.staging .env.local
RUN bun run build

FROM node:22-bullseye-slim AS staging-deploy
WORKDIR /app

RUN npm install -g bun@v1.1.22

COPY --from=staging-builder /app/public ./public
COPY --from=staging-builder /app/.next/standalone ./
COPY --from=staging-builder /app/.next/static ./.next/static

ENV NODE_ENV="production"
ENV HOSTNAME="0.0.0.0"
ENV PORT="3000"
EXPOSE 3000
CMD ["bun", "server.js"]

FROM node:22-bullseye-slim AS production-builder
WORKDIR /app

RUN npm install -g bun@v1.1.22

COPY package.json bun.lockb ./
RUN bun install --frozen-lockfile

COPY . .
COPY .env.production .env.local
RUN bun run build

FROM node:22-bullseye-slim AS production-deploy
WORKDIR /app

RUN npm install -g bun@v1.1.22

COPY --from=production-builder /app/public ./public
COPY --from=production-builder /app/.next/standalone ./
COPY --from=production-builder /app/.next/static ./.next/static

ENV NODE_ENV="production"
ENV HOSTNAME="0.0.0.0"
ENV PORT="3000"
EXPOSE 3000
CMD ["bun", "server.js"]
