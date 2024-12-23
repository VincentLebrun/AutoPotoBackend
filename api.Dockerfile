FROM golang:1.23.4-alpine

WORKDIR /app

RUN apk update && apk add libc-dev gcc make && rm -rf /var/cache/apk/*

COPY go.mod ./

RUN go mod download && go mod verify

RUN go install github.com/githubnemo/CompileDaemon@latest

COPY . .

COPY ./entrypoint.sh /entrypoint.sh

ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for

RUN chmod +x /usr/local/bin/wait-for /entrypoint.sh

ENTRYPOINT [ "sh", "/entrypoint.sh" ]
