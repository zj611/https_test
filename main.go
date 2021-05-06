package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of https service in golang!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":8081", "server.crt",
		"server.key", nil)
}

/*
 1043   openssl genrsa -des3 -out server.key 1024
 1044  openssl req -new -key server.key -out server.csr
 1045  openssl rsa -in server.key -out server.key
 1046  openssl x509 -req -days 365 -in server.csr -signkey server.key -out server.crt

*/
