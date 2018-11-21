FROM golang:1.10.3-alpine3.8

# git
RUN apk add make git openssl curl && rm /var/cache/apk/*

# envoy control-plane dependencies.
RUN go get github.com/alexellis/hmac

WORKDIR $GOPATH/src/github.com/datawire/http-auth-server
COPY . ./

CMD ["go", "run", "server.go"]

EXPOSE 8088
