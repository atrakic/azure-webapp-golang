package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Heartbeat("/healthz"))
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := struct {
			Message string `json:"message"`
		}{
			Message: "Hello from azure-webapp-golang!",
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonResponse)
	})

	r.Get("/sys", sysHandler)

	fmt.Println("Go backend started at port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func sysHandler(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	ip, _ := net.LookupIP(host)

	response := struct {
		IP   string `json:"ip"`
		Host string `json:"host"`
		Date string `json:"date"`
		ClientIp string `json:"clientIp"`
	}{
		IP:   ip[0].String(),
		Host: host,
		Date: time.Now().UTC().Format(time.RFC3339),
		ClientIp: r.RemoteAddr,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
