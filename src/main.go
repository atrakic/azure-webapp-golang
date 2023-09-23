package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var port string

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(
		w, `
Hello from azure-webapp-deploy!
`,
	)
}

func main() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handler)
	fmt.Println("Go backend started at port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
