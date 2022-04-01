package pow

import (
	"context"
	"crypto/sha1"
	"encoding/binary"
	"math/bits"

	"github.com/google/uuid"
)

func NewChallenge() []byte {
	t, _ := uuid.New().MarshalBinary()
	return t
}

func Compute(ctx context.Context, seed []byte, complexity uint8) ([]byte, error) {
	var counter uint64
	solution := make([]byte, 8)

	for {
		binary.LittleEndian.PutUint64(solution, counter)
		if Verify(seed, complexity, solution) {
			break
		}

		counter++

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			continue
		}
	}

	return solution, nil
}

func Verify(seed []byte, complexity uint8, solution []byte) bool {
	hash := shaHash(append(seed, solution...))

	if len(hash)*8 < int(complexity) {
		return false
	}

	zeroBytesWanted := complexity / 8
	zeroBitsWanted := int(complexity % 8)
	for _, b := range hash[:zeroBytesWanted] {
		if b != 0 {
			return false
		}
	}
	return zeroBitsWanted == 0 || bits.LeadingZeros8(hash[zeroBytesWanted]) >= zeroBitsWanted
}

func shaHash(in []byte) []byte {
	hash := sha1.New()
	_, err := hash.Write(in)
	if err != nil {
		return []byte{}
	}
	return hash.Sum(nil)
}
