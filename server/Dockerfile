FROM golang:1.18.3-alpine

WORKDIR /usr/src/server

COPY . .

ENV MYSQL_PASSWORD $MYSQL_PASSWORD
ENV MYSQL_IP $MYSQL_IP
ENV ME_ENV $ME_ENV
ENV ME_ADMIN_PASSWORD $ME_ADMIN_PASSWORD
ENV weatherAPI $weatherAPI

RUN go get

RUN go install

EXPOSE 9000 9001

CMD [ "go", "run", "server.go" ]

