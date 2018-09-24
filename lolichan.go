package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

const dbFile = "lolichan.db"
const lolisBucket = "lolis"

// Blockchain keeps a sequence of Blocks
type Lolichan struct {
	tip []byte
	db  *bolt.DB
}

// BlockchainIterator is used to iterate over blockchain blocks
type LolichanIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// AddBlock saves provided data as a block in the blockchain
func (lc *Lolichan) AddLoli(data string) {
	var lastHash []byte

	err := lc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(lolisBucket))
		lastHash = b.Get([]byte("l"))

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	newLoli := NewLoli(data, lastHash)

	err = lc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(lolisBucket))
		err := b.Put(newLoli.Hash, newLoli.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("l"), newLoli.Hash)
		if err != nil {
			log.Panic(err)
		}

		lc.tip = newLoli.Hash

		return nil
	})
}

// Iterator ...
func (lc *Lolichan) Iterator() *LolichanIterator {
	lci := &LolichanIterator{lc.tip, lc.db}

	return lci
}

// Next returns next block starting from the tip
func (i *LolichanIterator) Next() *Loli {
	var loli *Loli

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(lolisBucket))
		encodedLoli := b.Get(i.currentHash)
		loli = DeserializeLoli(encodedLoli)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = loli.PrevLoliHash

	return loli
}

// NewBlockchain creates a new Blockchain with genesis Block
func NewLolichan() *Lolichan {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(lolisBucket))

		if b == nil {
			fmt.Println("No existing blockchain found. Creating a new one...")
			genesis := NewGenesisLoli()

			b, err := tx.CreateBucket([]byte(lolisBucket))
			if err != nil {
				log.Panic(err)
			}

			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				log.Panic(err)
			}

			err = b.Put([]byte("l"), genesis.Hash)
			if err != nil {
				log.Panic(err)
			}
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	lc := Lolichan{tip, db}

	return &lc
}
