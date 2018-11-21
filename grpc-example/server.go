package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"flag"
	"log"
	"net"
	"os"

	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	envoy_type "github.com/envoyproxy/go-control-plane/envoy/type"
	gogo_type "github.com/gogo/protobuf/types"

	pb "github.com/datawire/grpc-auth-server/envoy/service/auth/v2alpha"
	"github.com/gogo/googleapis/google/rpc"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

// Server implements a gRPC service that handles authorization
// and HMAC validation.
type Server struct {
	Secret       string
	Authorized   *pb.CheckResponse
	Unauthorized *pb.CheckResponse
}

// Check request object.
func (s *Server) Check(ctx context.Context, r *pb.CheckRequest) (*pb.CheckResponse, error) {
	// Get request headers.
	headers := r.GetAttributes().GetRequest().GetHttp().GetHeaders()
	if headers["x-encoded-hash"] == "" {
		log.Printf("%v", errors.New("missing x-encoded-hash"))
		return s.Unauthorized, nil
	}
	log.Printf("x-encoded-hash: %s\n", headers["x-encoded-hash"])

	// Get message body.
	body := r.GetAttributes().GetRequest().GetHttp().GetBody().GetInlineBytes()
	log.Printf("message body received: %s\n", string(body))
	log.Printf("message body received: %+v bytes\n", len(string(body)))

	// Take a simple hash of the message.
	hash := hmac.New(sha256.New, []byte(s.Secret))
	hash.Write(body)
	if hex.EncodeToString(hash.Sum(nil)) != headers["x-encoded-hash"] {
		log.Printf("invalid hmac")
		return s.Unauthorized, nil
	}

	return s.Authorized, nil
}

// Run initializes the gRPC service and the server object.
func Run() error {
	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		return err
	}

	instance := &Server{Secret: os.Getenv("SECRET")}
	if instance.Secret == "" {
		return errors.New("missing secret")
	}

	// Authorized returns an response object with status OK and a header
	// that should be sent to the upstream service.
	okHDRs := make([]*core.HeaderValueOption, 1)
	okHDRs[0] = &core.HeaderValueOption{
		Header: &core.HeaderValue{
			Key:   "x-ok-example",
			Value: "this will be sent to the upstream service..",
		},
		Append: &gogo_type.BoolValue{Value: true},
	}
	instance.Authorized = &pb.CheckResponse{
		Status: &rpc.Status{Code: int32(0)},
		HttpResponse: &pb.CheckResponse_OkResponse{
			OkResponse: &pb.OkHttpResponse{
				Headers: okHDRs,
			},
		},
	}

	// Unauthorized will return header, status code and a body to the
	// downstream client.
	deniedHDRs := make([]*core.HeaderValueOption, 1)
	deniedHDRs[0] = &core.HeaderValueOption{
		Header: &core.HeaderValue{
			Key:   "x-failed-example",
			Value: "this will be sent to the client..",
		},
		Append: &gogo_type.BoolValue{Value: false},
	}
	instance.Unauthorized = &pb.CheckResponse{
		Status: &rpc.Status{Code: int32(16)},
		HttpResponse: &pb.CheckResponse_DeniedResponse{
			DeniedResponse: &pb.DeniedHttpResponse{
				Status: &envoy_type.HttpStatus{
					Code: envoy_type.StatusCode_Unauthorized,
				},
				Headers: deniedHDRs,
				Body:    "invalid hmac",
			},
		},
	}

	// Initialize the service.
	server := grpc.NewServer()
	pb.RegisterAuthorizationServer(server, instance)
	log.Printf("serving on port 50051")
	server.Serve(listen)

	return nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := Run(); err != nil {
		log.Printf("%v", err)
	}
}
