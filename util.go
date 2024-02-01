package main

import (
	"context"
	"math/rand"
	"time"
)

var rng *rand.Rand = rand.New(rand.NewSource(time.Now().Unix()))

// MustRandomIsAlive returns true or false depends on the percentage
func MustRandomIsAlive(ctx context.Context, percentage int) bool {
	if percentage < 0 || percentage > 100 {
		panic("percentage must be between 0-100")
	}

	return rng.Intn(101) < percentage
}
