package main

import (
	"COMPETITIONS/olympy/streaming-service/storage"
	"COMPETITIONS/olympy/streaming-service/websocket"
	"fmt"
	"log"
	"net/http"

	grpcServer "COMPETITIONS/olympy/streaming-service/grpc"
)

func main() {
	// Initialize MongoDB connection
	mongoClient, err := storage.NewMongoClient("mongodb://localhost:27017")
	fmt.Println("mongoooooooo")
	if err != nil {
		log.Fatalf("failed to connect to MongoDB: %v", err)
	}

	// Start WebSocket server
	go func() {
		if err := http.ListenAndServe(":8080", http.HandlerFunc(websocket.HandleConnections)); err != nil {
			log.Fatalf("failed to start WebSocket server: %v", err)
		}
	}()

	// Start gRPC server
	log.Println("Starting gRPC server on port 8777...")
	if err := grpcServer.StartGRPCServer(mongoClient, ":8777"); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}
