FROM golang:1.17

ENV GIN_MODE=release
ENV PORT=8000

WORKDIR /go/src

COPY . /go/src

RUN go mod download

RUN go build

CMD ["/go/src/docket"]