package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "olympy/streaming-service/genproto/stream_service"
	"olympy/streaming-service/storage"
	pb "COMPETITIONS/olympy/streaming-service"
	"COMPETITIONS/olympy/streaming-service/storage"
	"COMPETITIONS/olympy/streaming-service/websocket"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
)

type StreamServiceServer struct {
	pb.UnimplementedStreamingServiceServer
	mongoClient *storage.MongoClient
}

func NewStreamServiceServer(mongoClient *storage.MongoClient) *grpc.Server {
	grpcServer := grpc.NewServer()
	pb.RegisterStreamingServiceServer(grpcServer, &StreamServiceServer{mongoClient: mongoClient})
	return grpcServer
}

func StartGRPCServer(mongoClient *storage.MongoClient, address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer := NewStreamServiceServer(mongoClient)
	log.Printf("gRPC server listening on %s", address)
	return grpcServer.Serve(lis)
}

func (s *StreamServiceServer) StreamEvent(ctx context.Context, req *pb.StreamEventRequest) (*pb.StreamEventResponse, error) {
	event := bson.D{
		{"_id", uuid.NewString()},
		{"event_id", req.GetEventId()},
		{"text", req.GetText()},
		{"timestamp", time.Now()},
	}
	if err := s.mongoClient.InsertEvent(ctx, event); err != nil {
		return nil, fmt.Errorf("failed to save event to MongoDB: %v", err)
	}

	websocket.Broadcast(event)

	return &pb.StreamEventResponse{Message: "Event streamed successfully"}, nil
}
