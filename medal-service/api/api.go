package api

import (
	"log"
	"net"

	countryservice "olympy/medal-service/genproto/country_service"
	modelservice "olympy/medal-service/genproto/medal_service"
	"olympy/medal-service/internal/config"

	"google.golang.org/grpc"
)

type (
	API struct {
		medalservice   modelservice.MedalServiceServer
		countryservice countryservice.CountryServiceServer
	}
)

func New(medalservice modelservice.MedalServiceServer, countryservice countryservice.CountryServiceServer) *API {
	return &API{
		medalservice:   medalservice,
		countryservice: countryservice,
	}
}

func (a *API) RUN(config *config.Config) error {
	listener, err := net.Listen("tcp", config.Server.Port)
	if err != nil {
		return err
	}

	serverRegisterer := grpc.NewServer()
	modelservice.RegisterMedalServiceServer(serverRegisterer, a.medalservice)
	countryservice.RegisterCountryServiceServer(serverRegisterer, a.countryservice)

	log.Println("server has started running on port", config.Server.Port)

	return serverRegisterer.Serve(listener)
}
