package main

import (
	"context"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/timych/word-of-wisdom-pow/pow"
)

var (
	addr    = flag.String("addr", "localhost:8888", "the address to connect to")
	timeout = flag.Int("timeout", 1000, "Solution timeout in milliseconds")
)

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		log.Fatalf("Connection error: %v", err)
		return
	}
	defer conn.Close()

	// Getting challenge
	buf := make([]byte, 17)
	conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		log.Fatalf("Getting challenge error: %v", err)
	}
	seed := buf[:len(buf)-1]
	complexity := uint8(buf[len(buf)-1])
	log.Printf("Challenge recived (seed=%x, complexity=%v)", seed, complexity)

	// PoW solving
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout)*time.Millisecond)
	defer cancel()
	start := time.Now()
	s, err := pow.Compute(ctx, seed, complexity)
	if err != nil {
		log.Fatalf("PoW compute error: %v", err)
	}
	log.Printf("PoW solution found: %d (%v)", binary.LittleEndian.Uint64(s), time.Since(start))

	// Solution sending
	_, err = conn.Write(s)
	if err != nil {
		log.Fatalf("Solution sending error: %v", err)
	}

	// Wisdom reading
	buf = make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
	_, err = conn.Read(buf)
	if err != nil {
		log.Fatalf("Wisdom reading error: %v", err)
	}
	fmt.Printf("\n%s\n", buf)
}
