package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var port string

func main() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Heartbeat("/healthz"))
	r.Use(middleware.Logger)

	r.Get("/ip", helloHandler)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from azure-webapp-deploy!"))
	})

	fmt.Println("Go backend started at port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	ip, _ := net.LookupIP(host)
	fmt.Fprintf(w, "ip: %s - host: %s", ip, host)
}
