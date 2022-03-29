package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/timych/word-of-wisdom-pow/api"
	"github.com/timych/word-of-wisdom-pow/pow"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	addr = flag.String("addr", "localhost:8888", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}
	defer conn.Close()
	c := api.NewWordOfWisdomClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetChallenge(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("GetChallenge error: %v", err)
	}

	log.Printf("Challenge recived (token=%v, complexity=%v)", r.GetToken(), r.GetComplexity())
	start := time.Now()
	s, err := pow.Compute(ctx, r.GetToken(), r.GetComplexity())
	if err != nil {
		log.Fatalf("PoW compute error: %v", err)
	}

	log.Printf("PoW solution found: %v (%v)", s, time.Since(start))
	w, err := c.InspireMe(ctx, &api.ChallengeSolution{Token: r.GetToken(), Solution: s})
	if err != nil {
		log.Fatalf("InspireMe failed: %v", err)
	}

	log.Printf("Word of Wisdom: %v", w.Value)

	// fmt.Print("Press 'Enter' to exit")
	// bufio.NewReader(os.Stdin).ReadBytes('\n')
}
