FROM node:20-alpine

WORKDIR /usr/src/app

ARG env

COPY package.json ./

RUN apk update
RUN npm install

COPY . .

CMD ["npm", "run", "serve"]
