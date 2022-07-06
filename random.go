package main

import (
	"math/rand"
	"time"
)

func RandomInit() {
	rand.Seed(time.Now().UnixNano())
}

func RandomRange(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func RandomEntry[V comparable](arr []V) V {
	index := rand.Intn(len(arr))
	return arr[index]
}

func FlipACoin() string {
	value := RandomRange(0, 1)
	if value == 0 {
		return "heads"
	}

	return "tails"
}
