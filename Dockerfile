FROM golang

WORKDIR /go/src/app

COPY main.go .

RUN go build

USER 1000

ENTRYPOINT [ "app" ]
