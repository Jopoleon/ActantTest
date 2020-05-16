
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Egor Miloserdov"

RUN apk update && apk add --no-cache git

RUN mkdir /actant
WORKDIR /actant
COPY . .
#RUN echo $PATH
#/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin


#RUN apk update
#RUN apk add git unzip build-essential autoconf libtool
#RUN git clone https://github.com/google/protobuf.git && \
#    cd protobuf && \
#    ./autogen.sh && \
#    ./configure && \
#    make && \
#    make install && \
#    ldconfig && \
#    make clean && \
#    cd .. && \
#    rm -r protobuf
#
## Get the source from GitHub
#RUN go get google.golang.org/grpc
## Install protoc-gen-go
#RUN go get github.com/golang/protobuf/protoc-gen-go
#
#RUN go install google.golang.org/protobuf/cmd/protoc-gen-go

RUN go mod download


COPY . .

#RUN protoc -I=proto --go_out=plugins=grpc:. proto/*.proto


RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /actant/main .

EXPOSE 8899

ENTRYPOINT ./main -send_port 8899
