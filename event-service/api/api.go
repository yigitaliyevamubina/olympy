package api

import (
	"log"
	"net"
	genprotos "olympy/event-service/genproto/event_service"
	"olympy/event-service/internal/config"

	"google.golang.org/grpc"
)

type (
	API struct {
		service genprotos.EventServiceServer
	}
)

func New(service genprotos.EventServiceServer) *API {
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
	genprotos.RegisterEventServiceServer(serverRegisterer, a.service)

	log.Println("server has started running on port", config.Server.Port)

	return serverRegisterer.Serve(listener)
}
