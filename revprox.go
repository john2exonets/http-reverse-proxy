//
//  revprox.go  -- Simple HTTP Reverse Proxy
//  - based on: https://gist.github.com/thurt/2ae1be5fd12a3501e7f049d96dc68bb9
//
//  John D. Allen
//  September, 2020
//

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var port string
var phost string
var rhost string
var exists bool

func main() {
	port, exists = os.LookupEnv("REVPROX_PORT")
	if !exists {
		port = "8080"
	}
	phost, exists = os.LookupEnv("REVPROX_LOCAL_IP")
	if !exists {
		phost = "0.0.0.0"
	}
	rhost, exists = os.LookupEnv("REVPROX_REMOTE")
	if !exists {
		rhost = "google.com"
	}
	proxy := httputil.NewSingleHostReverseProxy(&url.URL{Scheme: "http", Host: rhost})
	// use http.Handle instead of http.HandleFunc when your struct implements http.Handler interface
	http.Handle("/", &ProxyHandler{proxy})
	host := phost + ":" + port
	fmt.Println("HTTP Reverse Proxy started on port", port)
	err := http.ListenAndServe(host, nil)
	if err != nil {
		panic(err)
	}
}

// ProxyHandler ... comment to satisfy stupid go-lint
type ProxyHandler struct {
	p *httputil.ReverseProxy
}

func (ph *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	// w.Header().Set("X-Ben", "Rad")
	ph.p.ServeHTTP(w, r)
}

