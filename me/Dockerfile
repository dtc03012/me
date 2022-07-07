FROM node:16

WORKDIR /usr/src/me

COPY package.json ./

RUN npm config set legacy-peer-deps true

RUN npm ci --only=production

COPY . .

EXPOSE 80

CMD [ "npm", "start" ]