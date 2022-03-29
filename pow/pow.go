package pow

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"

	"github.com/google/uuid"
)

const (
	zero rune = 48
)

func GenerateChallengeToken() string {
	t := uuid.NewString()
	activeTokens.Append(t)
	return t
}

func Compute(ctx context.Context, token string, complexity uint32) (int64, error) {
	var maxIterations int64 = 2 << (complexity * 4) // twice as many as the average required
	var counter int64
	for !checkSolution(token, complexity, counter) {
		counter++
		if counter >= maxIterations {
			return 0, fmt.Errorf("exceeded 2*2^%v iterations, failed to find solution", complexity*4)
		}

		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
			continue
		}
	}
	return counter, nil
}

func Verify(token string, complexity uint32, solution int64) bool {
	if checkSolution(token, complexity, solution) {
		if activeTokens.Delete(token) {
			return true
		}
	}
	return false
}

func checkSolution(token string, complexity uint32, solution int64) bool {
	hash := shaHash(token + ":" + base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(solution, 10))))

	if len(hash) < int(complexity) {
		return false
	}
	for _, val := range hash[:complexity] {
		if val != zero {
			return false
		}
	}
	return true
}

func shaHash(s string) string {
	hash := sha1.New()
	_, err := io.WriteString(hash, s)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(hash.Sum(nil))
}
