FROM docker.io/golang:1.21

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod tidy && go mod download && go mod verify

COPY main.go .

RUN go build -v -o /usr/local/bin/app ./...

USER 1000

CMD [ "app" ]
