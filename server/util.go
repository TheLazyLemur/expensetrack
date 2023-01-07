package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateUuid() string {
	rand.Seed(time.Now().UnixNano())

	id := fmt.Sprintf("%x-%x-%x-%x-%x",
		rand.Int31n(0x10000),
		rand.Int31n(0x10000),
		rand.Int31n(0x10000),
		rand.Int31n(0x10000),
		rand.Int31n(0x10000),
	)

	return id
}
