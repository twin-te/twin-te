FROM oven/bun:1.1 as dev

WORKDIR /usr/src/twin-te/front

COPY package.json bun.lockb ./

RUN bun install --frozen-lockfile

COPY . .

CMD ["bunx", "--bun", "vite", "--host", "0.0.0.0", "--port", "5173"]
