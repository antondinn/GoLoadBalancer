/*
Runs a simple HTTP server for the testing of the load balancer.
We can run multiple instances by running the program with different INSTANCE_ID values.
*/
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Instance string `json:"instance"`
	Method string `json:"method"`
	Path string `json:"path"`
}

func main() {
	instance := os.Getenv("INSTANCE_ID")

	if instance == "" {
		instance = "unknown"
	}

	//Creating HTTP router
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		//Telling the client the body is JSON
		w.Header().Set("Content-Type", "application/json")

		response := Response{
			Instance:	instance,
			Method:		r.Method,
			Path:		r.URL.Path,
		}

		//Struct converted into JSON and outputs into w
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Printf("failed to encode response: %v", err)
		}
	})

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"status":	"ok",
			"instance":	instance,
		})
	})

	log.Printf("%s listening on :8080", instance)

	//Listening on port 8080 from all interfaces
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}