package api

import (
	"log"
	"net"

	modelservice "olympy/medal-service/genproto/medal_service"
	"olympy/medal-service/internal/config"

	"google.golang.org/grpc"
)

type (
	API struct {
		service modelservice.MedalServiceServer
	}
)

func New(service modelservice.MedalServiceServer) *API {
	return &API{
		service: service,
	}
}

func (a *API) RUN(config *config.Config) error {
	listener, err := net.Listen("tcp", config.Server.Port)
	if err != nil {
		return err
	}

	serverRegisterer := grpc.NewServer()
	modelservice.RegisterMedalServiceServer(serverRegisterer, a.service)

	log.Println("server has started running on port", config.Server.Port)

	return serverRegisterer.Serve(listener)
}
