FROM golang:1.10
ADD . /go/src/github.com/furikuri/food-machine
WORKDIR /go/src/github.com/furikuri/food-machine/worker
RUN go get ./
RUN go build

FROM alpine:3.8
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /root/
COPY --from=0 /go/bin/worker .

ENTRYPOINT ["/root/worker"]