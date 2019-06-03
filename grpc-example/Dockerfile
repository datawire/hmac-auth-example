FROM golang:1.10.3-alpine3.8

# git
RUN apk add make git openssl && rm /var/cache/apk/*

# envoy control-plane dependencies.
RUN go get github.com/envoyproxy/protoc-gen-validate
RUN go get github.com/envoyproxy/go-control-plane/envoy/type
RUN go get github.com/envoyproxy/go-control-plane/envoy/api/v2/core
RUN go get github.com/gogo/googleapis/google/rpc
RUN go get github.com/gogo/protobuf/types
RUN go get github.com/golang/glog
RUN go get google.golang.org/grpc

WORKDIR $GOPATH/src/github.com/datawire/grpc-auth-server
COPY . ./

CMD ["go", "run", "server.go"]

EXPOSE 50051