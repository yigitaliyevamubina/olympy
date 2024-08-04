package api

import (
	"log"
	"net"

	"olympy/medal-service/genprotos"
	"olympy/medal-service/internal/config"

	"google.golang.org/grpc"
)

type (
	API struct {
		service genprotos.ProductServiceServer
	}
)

func New(service genprotos.ProductServiceServer) *API {
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
	genprotos.RegisterProductServiceServer(serverRegisterer, a.service)

	log.Println("server has started running on port", config.Server.Port)

	return serverRegisterer.Serve(listener)
}
