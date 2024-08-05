package grpc

import (
	"context"
	"fmt"
	"time"

	pb "COMPETITIONS/olympy/streaming-service"
	"COMPETITIONS/olympy/streaming-service/websocket"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *server) StreamEvent(ctx context.Context, req *pb.StreamEventRequest) (*pb.StreamEventResponse, error) {
	// Save to MongoDB
	event := bson.D{
		{"_id", uuid.NewString()},
		{"event_id", req.GetEventId()},
		{"text", req.GetText()},
		{"timestamp", time.Now()},
	}
	if err := s.mongoClient.InsertEvent(ctx, event); err != nil {
		return nil, fmt.Errorf("failed to save event to MongoDB: %v", err)
	}

	// Broadcast to clients
	websocket.Broadcast(event)

	return &pb.StreamEventResponse{Message: "Event streamed successfully"}, nil
}
