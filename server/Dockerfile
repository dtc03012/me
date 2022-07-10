FROM golang:1.18.3-alpine

WORKDIR /usr/src/server

COPY . .

RUN go get

RUN go install

EXPOSE 9000 9001

CMD [ "go", "run", "server.go" ]

