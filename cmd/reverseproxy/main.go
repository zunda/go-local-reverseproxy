package main

import (
	"fmt"
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
	bURL, err := url.Parse(fmt.Sprintf("http://localhost:%s", bport))
	if err != nil {
		panic(err)
	}
	log.Println("Backend URL: ", bURL)
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
