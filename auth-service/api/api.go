package api

import (
	genprotos "COMPETITIONS/olympy/auth-service/genprotos"
	"COMPETITIONS/olympy/auth-service/internal/config"
	"log"
	"net"

	"google.golang.org/grpc"
)

type (
	API struct {
		service genprotos.AuthServiceServer
	}
)

func New(service genprotos.AuthServiceServer) *API {
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
	genprotos.RegisterAuthServiceServer(serverRegisterer, a.service)

	log.Println("server has started running on port", config.Server.Port)

	return serverRegisterer.Serve(listener)
}
