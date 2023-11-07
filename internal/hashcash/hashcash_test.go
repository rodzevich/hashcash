package hashcash

import (
	"testing"
	"time"
)

func TestHash(t *testing.T) {
	h := Hashcash{
		Bits:  2,
		Date:  time.Now(),
		Rand:  123,
		Nonce: 0,
	}
	h, err := h.Solve()
	if err != nil {
		t.Fatal(err)
	}
	if h.IsValid() == false {
		t.Fatal("invalid solution")
	}
	if h.Nonce != 296 {
		t.Fatal("invalid solution")
	}
}
