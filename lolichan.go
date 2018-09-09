package main

// Blockchain keeps a sequence of Blocks
type Lolichan struct {
	lolis []*Loli
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Lolichan) AddLoli(data string) {
	prevLoli := bc.lolis[len(bc.lolis)-1]
	newLoli := NewLoli(data, prevLoli.Hash)
	bc.lolis = append(bc.lolis, newLoli)
}

// genesis Block
func NewLolichan() *Lolichan {
	return &Lolichan{[]*Loli{NewGenesisLoli()}}
}
