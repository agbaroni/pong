FROM golang

WORKDIR /go/src/app

COPY main.go .

RUN go build

ENTRYPOINT [ "app" ]
