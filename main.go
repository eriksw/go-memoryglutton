package main

import (
	"fmt"
	"math/rand"
)

const sliceSize = 1024 * 1024 * 10

func main() {
	var horde [][]byte

	for ; ; {
		s := make([]byte, sliceSize, sliceSize)
		_, err := rand.Read(s)
		if err != nil {
			panic(err)
		}
		if len(horde) % 10000 == len(horde) {
			fmt.Printf("Used: %d\n", len(horde) * sliceSize)
		}
		horde = append(horde, s)
	}
}
