package hashcash

import (
	"crypto/sha1"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

const (
	bits        = 3
	maxAttempts = 1000000
)

type Hashcash struct {
	Bits  int
	Date  time.Time
	Rand  int
	Nonce int64
}

func New() Hashcash {
	return Hashcash{
		Bits:  bits,
		Date:  time.Now(),
		Rand:  rand.Intn(math.MaxInt32),
		Nonce: 0,
	}
}

func (h *Hashcash) Stringify() string {
	return fmt.Sprintf("%d:%s:%d:%d", h.Bits, h.Date.Format(time.DateOnly), h.Rand, h.Nonce)
}

func (h *Hashcash) sha1() string {
	s := h.Stringify()
	hash := sha1.New()
	hash.Write([]byte(s))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (h *Hashcash) IsValid() bool {
	hash := h.sha1()
	prefix := strings.Repeat("0", h.Bits)
	return strings.HasPrefix(hash, prefix)
}

func (h *Hashcash) Solve() (Hashcash, error) {
	for h.Nonce <= maxAttempts {
		if h.IsValid() {
			return *h, nil
		}
		h.Nonce++
	}
	return *h, fmt.Errorf("can't solve the challenge")
}
