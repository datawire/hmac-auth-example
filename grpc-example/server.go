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

	envoy_type "github.com/envoyproxy/go-control-plane/envoy/type"

	pb "github.com/datawire/grpc-auth-server/envoy/service/auth/v2alpha"
	"github.com/gogo/googleapis/google/rpc"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

// Server is a naive implementation of a gRPC service that handles authorization via
// HMAC validation.
type Server struct {
	Secret string
}

// Check signed message.
func (s *Server) Check(ctx context.Context, r *pb.CheckRequest) (*pb.CheckResponse, error) {
	authorized := pb.CheckResponse{
		Status: &rpc.Status{Code: int32(0)},
		// Denied HTTP response object that will be returned to the downstream.
		HttpResponse: &pb.CheckResponse_OkResponse{},
	}

	unauthorized := pb.CheckResponse{
		Status: &rpc.Status{Code: int32(16)},
		// Denied HTTP response object that will be returned to the downstream.
		HttpResponse: &pb.CheckResponse_DeniedResponse{
			DeniedResponse: &pb.DeniedHttpResponse{
				Status: &envoy_type.HttpStatus{
					Code: envoy_type.StatusCode_Unauthorized,
				},
			},
		},
	}

	headers := r.GetAttributes().GetRequest().GetHttp().GetHeaders()
	if headers["x-encoded-hash"] == "" {
		log.Printf("%v", errors.New("missing x-encoded-hash"))

		return &unauthorized, nil
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

		return &unauthorized, nil
	}

	return &authorized, nil
}

// Run server.
func Run() error {
	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		return err
	}

	instance := &Server{Secret: os.Getenv("SECRET")}
	if instance.Secret == "" {
		return errors.New("missing secret")
	}

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
