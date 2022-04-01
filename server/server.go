package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/timych/word-of-wisdom-pow/pow"
	"github.com/timych/word-of-wisdom-pow/server/wisdom"
)

var (
	port          = flag.Int("port", 8888, "Server port to listen")
	timeout       = flag.Int("timeout", 5000, "Writing/Reading timeout in milliseconds")
	powComplexity = flag.Int("pow_complexity", 20, "Number of wanted zero bits in solution hash")
)

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Server listening at %v", l.Addr())
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatalf("Failed to accept: %v", err)
		}
		go handleConnection(c)
	}

}

func handleConnection(c net.Conn) {
	defer c.Close()
	c.SetDeadline(time.Now().Add(time.Duration(*timeout) * time.Millisecond))

	// Challenge sending
	seed := pow.NewChallenge()
	_, err := c.Write(append(seed, uint8(*powComplexity))) // 16 bytes of seed + 1 byte of complexity
	if err != nil {
		log.Printf("Writing error: %v", err)
		return
	}

	// Solution reading
	buf := make([]byte, 8)
	n, err := io.ReadFull(c, buf)
	if err != nil {
		log.Printf("Solution reading error: %v", err)
		return
	}

	// Solution verification
	if n < 8 || !pow.Verify(seed, uint8(*powComplexity), buf) {
		log.Print("PoW verification failed")
		return
	}

	c.Write(append([]byte(wisdom.GetWordOfWisdom()), '\n'))
}
