package main

import (
	"fmt"
)

func main() {
	bc := NewLolichan()

	bc.AddLoli("Send 1 lolicoin to A")
	bc.AddLoli("Send 2 more lolicoin to A")

	for _, loli := range bc.lolis {
		fmt.Printf("Prev. hash: %x\n", loli.PrevLoliHash)
		fmt.Printf("Data: %s\n", loli.Data)
		fmt.Printf("Hash: %x\n", loli.Hash)
		fmt.Println()
	}

	connect()
	run()
}

