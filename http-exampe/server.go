package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Server implements a HTTP handler that validates a request body.
type Server struct {
	Secret string
}

// Check request object.
func (s *Server) check(w http.ResponseWriter, r *http.Request) {
	// Get x-encoded-hash header.
	header := r.Header.Get("x-encoded-hash")
	if header == "" {
		err := errors.New("missing x-encoded-hash")
		log.Printf("%v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	log.Printf("x-encoded-hash: %s\n", header)

	// Get message body.
	var body []byte
	if r.Body != nil {
		body, _ = ioutil.ReadAll(r.Body)
	}

	log.Printf("message body received: %s\n", string(body))
	log.Printf("message body received: %+v bytes\n", len(string(body)))

	// Take a simple hash of the message.
	secretBytes := []byte(s.Secret)
	hash := hmac.New(sha256.New, secretBytes)
	hash.Write(body)

	// If x-auth-hash does not match the hashed message, the authorization server
	// responds with 401 which will be picked by the filter and returned to the
	// downstream.
	if hex.EncodeToString(hash.Sum(nil)) != header {
		log.Printf("invalid hmac")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// If x-auth-hash matches the hashed message, the authorization server
	// responds with 200 OK.
	w.WriteHeader(http.StatusOK)
	return
}

func main() {
	log.Print("serving on port 8088")
	server := &Server{
		Secret: os.Getenv("SECRET"),
	}

	if server.Secret == "" {
		err := errors.New("missing secret")
		log.Printf("%v", err)
	}

	http.HandleFunc("/extauth/", server.check)

	log.Fatal(http.ListenAndServe("0.0.0.0:8088", nil))
}
