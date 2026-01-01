package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Jacksondprice/atlas-recess-etl/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/payments", handlers.ListPayments)
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Fatal(http.ListenAndServe(addr, mux))
}
