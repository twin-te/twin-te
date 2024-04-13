FROM node:21 as dev

WORKDIR /usr/src/twin-te/sponsorship

COPY package.json yarn.lock ./

RUN yarn install --frozen-lockfile

COPY . .

CMD ["yarn", "dev", "--hostname", "0.0.0.0", "--port", "3000"]
