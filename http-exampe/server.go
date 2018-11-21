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

func check(w http.ResponseWriter, r *http.Request) {
	secret := os.Getenv("SECRET")
	if secret == "" {
		err := errors.New("missing secret")
		log.Printf("%v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get x-encoded-hash header.
	hashHDR := r.Header.Get("x-encoded-hash")
	if hashHDR == "" {
		err := errors.New("missing x-encoded-hash")
		log.Printf("%v", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	log.Printf("x-encoded-hash: %s\n", hashHDR)

	// Get message body.
	var body []byte
	if r.Body != nil {
		body, _ = ioutil.ReadAll(r.Body)
	}

	log.Printf("message body received: %s\n", string(body))
	log.Printf("message body received: %+v bytes\n", len(string(body)))

	// Take a simple hash of the message.
	secretBytes := []byte(secret)
	hash := hmac.New(sha256.New, secretBytes)
	hash.Write(body)
	hexHASH := hex.EncodeToString(hash.Sum(nil))

	// If x-auth-hash does not match the hashed message, the authorization server
	// responds with 401 which will be picked by the filter and returned to the
	// downstream.
	if hexHASH != hashHDR {
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
	http.HandleFunc("/extauth/", check)

	log.Fatal(http.ListenAndServe("0.0.0.0:8088", nil))
}
