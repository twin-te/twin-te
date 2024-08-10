FROM oven/bun:1.1 AS local

WORKDIR /usr/src/twin-te/sponsorship

# install nodejs to prevent the following error.
# `Watchpack Error (watcher): EINVAL: Invalid argument`
RUN apt update && apt install -y nodejs

COPY package.json bun.lockb ./

RUN bun install --frozen-lockfile

COPY . .

CMD ["bun", "run", "dev", "--hostname", "0.0.0.0", "--port", "3000"]
