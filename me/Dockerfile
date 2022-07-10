FROM node:16

WORKDIR /usr/src/me

COPY package.json ./

ARG SERVER_IP
ENV SERVER_IP $SERVER_IP

RUN npm config set legacy-peer-deps true

RUN npm install

RUN npm ci --only=production

COPY . .

EXPOSE 80

CMD [ "npm", "start" ]