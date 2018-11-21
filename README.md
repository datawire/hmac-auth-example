***External Authorization***

This project provides two examples of authorization servers that allow validating the request body received in Ambassador.

***HTTP & gRPC examples***

1. Deploy `ambassador.yaml`
2. Get ambassador's external IP. (Referred to as ${EXTERNAL_IP} below)
3. Try the following:

Authorized call:
```
$ curl -v -X POST -H "x-encoded-hash: 4643978965ffcec6e6d73b36a39ae43ceb15f7ef8131b8307862ebc560e7f988" \
  ${EXTERNAL_IP}/httpbin/post -d "the message to hash here"
```

Unauthorized call:
```
$ curl -v -X POST -H "x-encoded-hash: 123" ${EXTERNAL_IP}/httpbin/post -d "the message to hash here"
```

***Development***

Use the provided Dockerfile to build, e.g:
```
$ cd grpc-example && docker build -t auth-service-grpc .
```

Please refer to `server.go` for code example and `ambassador.yaml` on how to configure the filter. 

***Important***

In order to send the message body, Envoy needs to stop the headers and buffer so that the authorization server 
never receives a partial message. It's strongly recommended to buffer the message at the beginning of the filter 
chain and this includes limiting the size of the message as well as the maximum latency allowed. For example:
```
annotations:
    getambassador.io/config: |
      ---
      apiVersion: ambassador/v1
      kind: BufferService
      name: buffer
      max_request_bytes: 16384
      max_request_time: 500
```

In order to send the message body over the authorization filter, the following setting in `getambassador.io/config` must be set to `True`:
```
allow_request_body: true
```
