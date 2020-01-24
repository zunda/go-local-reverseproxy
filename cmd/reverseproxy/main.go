package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	bport := os.Getenv("BPORT")
	if bport == "" {
		bport = "5000"
	}
	bURL, err := url.Parse("http://localhost:" + bport)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Backend URL: " + bURL.String())
	h := http.NewServeMux()
	h.HandleFunc("/", httputil.NewSingleHostReverseProxy(bURL).ServeHTTP)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Println("Listening at port " + port)
	err = http.ListenAndServe(":"+port, h)
	log.Fatal(err)
}
