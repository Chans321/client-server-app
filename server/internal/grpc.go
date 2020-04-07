package grpc

import (
	"context"
	"net"

	pbclientserver "github.com/Chans321/client-server-app/proto-files"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

func (g *Grpc) SayHello(ctx context.Context, input *pbclientserver.SayHelloRequest) (*pbclientserver.SayHelloResponse, error) {
	log.Info().Msg("In SayHello function of server")
	name := input.Name
	newGreeting := "hello " + name
	return &pbclientserver.SayHelloResponse{
		Greeting: newGreeting,
	}, nil

}

func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

func (g *Grpc) ListenAndServe() error {
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "failed to open tcp port")
	}
	g.srv = grpc.NewServer()
	pbclientserver.RegisterSayHelloServiceServer(g.srv, g)
	log.Info().Str("address", g.address).Msg("starting hello service")
	err = g.srv.Serve(lis)
	if err != nil {
		return errors.Wrap(err, "failed to start grpc server")
	}
	return nil
}
