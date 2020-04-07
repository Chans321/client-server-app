package main

import (
	"context"
	"flag"
	"time"

	pbclientserver "github.com/Chans321/client-server-app/proto-files"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	var addressPtr = flag.String("address", "localhost:50051", "address to connect")
	flag.Parse()
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	con, err := grpc.Dial(*addressPtr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to dial server")
	}
	c := pbclientserver.NewSayHelloServiceClient(con)
	if c == nil {
		log.Info().Msg("Client Nil")
	}
	r, err := c.SayHello(timeoutCtx, &pbclientserver.SayHelloRequest{
		Name: "Chandu",
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to fetch greetings from server")
	}
	if r != nil {
		log.Info().Interface("Response:", r.GetGreeting()).Msg("Greeting Recieved")
	} else {
		log.Fatal().Err(err).Msg("Failed to fetch greetings")
	}
	defer func() {
		err := con.Close()
		if err != nil {
			log.Info().Msg("Failed to close connection")
		}
	}()
}
