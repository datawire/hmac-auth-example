**External Authorization**

***HTTP & gRPC examples***

1. Deploy `ambassador.yaml`
2. Get ambassador's external ip.
3. Try the following:

Authorized call:
```
$ curl -v -X POST -H "x-encoded-hash: 4643978965ffcec6e6d73b36a39ae43ceb15f7ef8131b8307862ebc560e7f988" \
  {EXTERNAL_IP}/httpbin/post -d "the message to hash here"
```

Unauthorized call:
```
$ curl -v -X POST -H "x-encoded-hash: 123" EXTERNAL_IP/httpbin/post -d "the message to hash here"
```

***Developement***

Use the provided Dockerfiles to build, e.g:
```
$ cd grpc-example && docker build -t auth-service-grpc .
```

Please refer to `server.go` for code example and `ambassador.yaml` on how to configure the filter. 

***Important***

In order to send the message body, Envoy needs to stop the headers and buffer so that the authorization server 
never receives a partial message. It's strongly recomended to buffer the message at the beguining of the filter 
chain and this includes limitting the size of the message as well as the maximum latency allowed. For example:
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
