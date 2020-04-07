package main

import (
	"flag"

	"github.com/rs/zerolog/log"

	grpcsetup "github.com/Chans321/client-server-app/server/internal"
)

func main() {
	var addressPtr = flag.String("address", ":50051", "address to connect SayHello service")
	flag.Parse()
	s := grpcsetup.NewServer(*addressPtr)
	err := s.ListenAndServe()
	if err != nil {
		log.Info().Msg("Failed to start grpc server for SayHello service")
	}
}
