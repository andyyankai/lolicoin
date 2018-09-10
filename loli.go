package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block keeps block headers
type Loli struct {
	Timestamp     int64
	Data          []byte
	PrevLoliHash []byte
	Hash          []byte
	Nonce		int
}

// SetHash calculates and sets block hash
func (b *Loli) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevLoliHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// NewBlock creates and returns Block
func NewLoli(data string, prevLoliHash []byte) *Loli {
	loli := &Loli{time.Now().Unix(), []byte(data), prevLoliHash, []byte{}, 0}
	loli.SetHash()
	return loli
}

// Genesis Block
func NewGenesisLoli() *Loli {
	return NewLoli("Genesis Loli", []byte{})
}
