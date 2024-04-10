FROM node:21 as dev

WORKDIR /usr/src/twin-te/front

COPY package.json yarn.lock ./

RUN yarn install --frozen-lockfile

COPY . .

CMD ["yarn", "vite", "--host", "0.0.0.0", "--port", "5173"]
