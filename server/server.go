package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/timych/word-of-wisdom-pow/api"
	"github.com/timych/word-of-wisdom-pow/pow"
	"github.com/timych/word-of-wisdom-pow/server/wisdom"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	port          = flag.Int("port", 8888, "Server port to listen")
	powComplexity = flag.Int("pow_complexity", 5, "Number of wanted zeros in solution")
)

type server struct {
	api.UnimplementedWordOfWisdomServer
}

func (s *server) GetChallenge(ctx context.Context, in *emptypb.Empty) (*api.Challenge, error) {
	c := &api.Challenge{
		Token:      pow.GenerateChallengeToken(),
		Complexity: uint32(*powComplexity),
	}
	log.Printf("Challenge requested (%+v)", c)
	return c, nil
}

func (s *server) InspireMe(ctx context.Context, in *api.ChallengeSolution) (*wrapperspb.StringValue, error) {
	if !pow.Verify(in.Token, uint32(*powComplexity), in.Solution) {
		log.Printf("PoW verification failed (%+v)", in)
		return nil, status.Error(codes.PermissionDenied, "PoW verification failed")
	}
	log.Printf("PoW verified (%+v)", in)
	return &wrapperspb.StringValue{Value: wisdom.GetWordOfWisdom()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterWordOfWisdomServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
