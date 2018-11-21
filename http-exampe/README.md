***External Authorization - HTTP example

1. Deploy `ambassador.yaml`
2. Get ambassador's external ip.
3. Try the following:

Authorized call:
```
$ curl -v -X POST -H "x-encoded-hash: 4643978965ffcec6e6d73b36a39ae43ceb15f7ef8131b8307862ebc560e7f988" {EXTERNAL_IP}/httpbin/post -d "the message to hash here"
```

Unauthorized call:
```
$ curl -v -X POST -H "x-encoded-hash: 123" EXTERNAL_IP/httpbin/post -d "the message to hash here"
```

Please refer to `server.go` for code example and `ambassador.yaml` comments on how to configure the filter. 
