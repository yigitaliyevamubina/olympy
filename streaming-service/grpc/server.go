package grpc

import (
	"log"
	"net"

	pb "olympy/streaming-service/genproto/stream_service"
	"olympy/streaming-service/storage"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStreamingServiceServer
	mongoClient *storage.MongoClient
}

func NewServer(mongoClient *storage.MongoClient) *grpc.Server {
	grpcServer := grpc.NewServer()
	pb.RegisterStreamingServiceServer(grpcServer, &server{mongoClient: mongoClient})
	return grpcServer
}

func StartGRPCServer(mongoClient *storage.MongoClient, address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer := NewServer(mongoClient)
	log.Printf("gRPC server listening on %s", address)
	return grpcServer.Serve(lis)
}
