FROM golang:1.17

ENV GIN_MODE=release
ENV PORT=8000

RUN apt-get update -y \
    && apt-get install -y openssh-server \
    && apt-get install -y iproute2 \
    && apt-get install -y python

RUN rm /bin/sh && ln -s /bin/bash /bin/sh

WORKDIR /go/src

ADD ./.profile.d /app/.profile.d

COPY . /go/src

RUN go mod download

RUN go build

CMD ["/go/src/docket"]