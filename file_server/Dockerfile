FROM golang:1.18.3-alpine

WORKDIR /usr/src/file_server

COPY . .

ARG MYSQL_PASSWORD
ENV MYSQL_PASSWORD $MYSQL_PASSWORD

ARG MYSQL_IP
ENV MYSQL_IP $MYSQL_IP

ARG FILE_SERVER_ENV
ENV FILE_SERVER_ENV $FILE_SERVER_ENV

RUN go get

RUN go install

EXPOSE 8282

CMD [ "go", "run", "file_server.go" ]

