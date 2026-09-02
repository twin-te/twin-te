ARG BUN_VERSION=1.4.0

FROM oven/bun:${BUN_VERSION}-slim AS local
WORKDIR /usr/src/twin-te/sponsorship

COPY package.json bun.lock ./
RUN bun install --frozen-lockfile

COPY . .
ENV NODE_ENV development
EXPOSE 3000
CMD ["bun", "run", "dev", "--hostname", "0.0.0.0", "--port", "3000"]

FROM oven/bun:${BUN_VERSION}-slim AS staging-builder
WORKDIR /app

COPY package.json bun.lock ./
RUN bun install --frozen-lockfile

COPY . .
COPY .env.staging .env.local
RUN bun run build

FROM oven/bun:${BUN_VERSION}-slim AS staging-deploy
WORKDIR /app

COPY --from=staging-builder /app/public ./public
COPY --from=staging-builder /app/.next/standalone ./
COPY --from=staging-builder /app/.next/static ./.next/static

ENV NODE_ENV="production"
ENV HOSTNAME="0.0.0.0"
ENV PORT="3000"
EXPOSE 3000
CMD ["bun", "server.js"]

FROM oven/bun:${BUN_VERSION}-slim AS production-builder
WORKDIR /app

COPY package.json bun.lock ./
RUN bun install --frozen-lockfile

COPY . .
COPY .env.production .env.local
RUN bun run build

FROM oven/bun:${BUN_VERSION}-slim AS production-deploy
WORKDIR /app

COPY --from=production-builder /app/public ./public
COPY --from=production-builder /app/.next/standalone ./
COPY --from=production-builder /app/.next/static ./.next/static

ENV NODE_ENV="production"
ENV HOSTNAME="0.0.0.0"
ENV PORT="3000"
EXPOSE 3000
CMD ["bun", "server.js"]
