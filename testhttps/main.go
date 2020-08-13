package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of https service in golang!")
}

func main_https() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":8081", "server.pem",
		"server-key.pem", nil)
}

const idleTimeout = 5 * time.Minute
const activeTimeout = 10 * time.Minute
func maihttps() {
	var srv http.Server
	//http2.VerboseLogs = true
	srv.Addr = ":8081"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w,"http 协议%s\n",r.Proto)
	})
	//http2.ConfigureServer(&srv, &http2.Server{})
	go func() {
		log.Fatal(http.ListenAndServeTLS(":8000","server.pem", "server-key.pem",nil))
	}()
	select {}
}

func myhttp2() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "http proto %s:%s",r.Proto,r.RemoteAddr)
	})
	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	http2.ConfigureServer(s, &http2.Server{})
	go func() {
		log.Fatal(s.ListenAndServeTLS("server.pem","server-key.pem"))

	}()
	select {}
}

func http2c(){
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
		fmt.Fprintf(w, "Protocol: %s\n", r.Proto)


	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w,"http 协议%s  %s\n",r.Proto,r.RemoteAddr)
	})
	h2s := &http2.Server{
		// ...
	}
	h1s := &http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(handler, h2s),
	}
	log.Fatal(h1s.ListenAndServe())
}

func myhttp() {
	//http2.VerboseLogs = true
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w,"http 协议%s  %s\n",r.Proto,r.RemoteAddr)
	})
	//http2.ConfigureServer(&srv, &http2.Server{})
	go func() {
		log.Fatal(http.ListenAndServe(":8080",nil))
	}()
	select {}
}

func main(){
	//http2c()
	myhttp2()
//	myhttp()
}