package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
)

func fib(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}
	return new(big.Int).Add(fib(n-1), fib(n-2))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["n"]
	if !ok || len(keys[0]) < 1 {
		http.Error(w, "Missing 'n' query param", http.StatusBadRequest)
		return
	}
	n, err := strconv.Atoi(keys[0])
	if err != nil || n < 0 {
		http.Error(w, "'n' must be a positive integer", http.StatusBadRequest)
		return
	}
	result := fib(n)
	fmt.Fprintf(w, "Fibonacci(%d) = %s\n", n, result.String())
}

func main() {
	http.HandleFunc("/sum", sumHandler)
	http.HandleFunc("/health", healthHandler)
	port := "8080"
	log.Printf("Server starting on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}