package api

import (
	"log"
	"net"

	athleteservice "olympy/athlete-service/genproto/athlete_service"
	"olympy/athlete-service/internal/config"

	"google.golang.org/grpc"
)

type (
	API struct {
		service athleteservice.AthleteServiceServer
	}
)

func New(service athleteservice.AthleteServiceServer) *API {
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
	athleteservice.RegisterAthleteServiceServer(serverRegisterer, a.service)

	log.Println("server has started running on port", config.Server.Port)

	return serverRegisterer.Serve(listener)
}
