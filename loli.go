package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

// Block keeps block headers
type Loli struct {
	Timestamp     int64
	Data          []byte
	PrevLoliHash []byte
	Hash          []byte
	Nonce         int
}

// Serialize serializes the block
func (b *Loli) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// NewBlock creates and returns Block
func NewLoli(data string, prevLoliHash []byte) *Loli {
	loli := &Loli{time.Now().Unix(), []byte(data), prevLoliHash, []byte{}, 0}
	pow := NewProofOfWork(loli)
	nonce, hash := pow.Run()

	loli.Hash = hash[:]
	loli.Nonce = nonce

	return loli
}

// Genesis Block
func NewGenesisLoli() *Loli {
	return NewLoli("Genesis Loli", []byte{})
}

func DeserializeLoli(d []byte) *Loli {
	var loli Loli

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&loli)
	if err != nil {
		log.Panic(err)
	}

	return &loli
}
