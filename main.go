package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/A292460/CRUDAPI/apis"
)

func main() {
	// Create a server log file
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	// Setup a JSON handler for the logger
	jsonHandler := slog.NewJSONHandler(logFile, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	// Set the default logger to use the JSON handler
	slog.SetDefault(slog.New(jsonHandler))

	// Define routes here
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/datastore/valkey", apis.CreateValkeyInstance) //Calling the ProductGet function from apis.go
	slog.Info("Server started at port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("Server failed", "error", err)
	}

}
