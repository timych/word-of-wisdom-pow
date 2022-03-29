package pow

import "sync"

var activeTokens ActiveTokens = ActiveTokens{}

type ActiveTokens struct {
	tokens []string
	sync.Mutex
}

func (at *ActiveTokens) Append(t string) {
	at.Lock()
	at.tokens = append(at.tokens, t)
	at.Unlock()
}

func (at *ActiveTokens) Delete(t string) bool {
	at.Lock()
	defer at.Unlock()
	for i, v := range at.tokens {
		if v == t {
			at.tokens[i] = at.tokens[len(at.tokens)-1]
			at.tokens = at.tokens[:len(at.tokens)-1]
			return true
		}
	}
	return false
}
