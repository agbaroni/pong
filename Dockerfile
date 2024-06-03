FROM docker.io/golang:1.21

WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./
#RUN go mod download && go mod verify

COPY main.go ./

RUN go build -v -o /usr/local/bin/app ./...

USER 1000

CMD [ "app" ]
