package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

// CLI responsible for processing command line arguments
type CLI struct {
	lc *Lolichan
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("  printchain - print all the blocks of the blockchain")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addLoli(data string) {
	cli.lc.AddLoli(data)
	fmt.Println("Success!")
}

func (cli *CLI) printChain() {
	lci := cli.lc.Iterator()

	for {
		loli := lci.Next()

		fmt.Printf("Prev. hash: %x\n", loli.PrevLoliHash)
		fmt.Printf("Data: %s\n", loli.Data)
		fmt.Printf("Hash: %x\n", loli.Hash)
		pow := NewProofOfWork(loli)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(loli.PrevLoliHash) == 0 {
			break
		}
	}
}

// Run parses command line arguments and processes commands
func (cli *CLI) Run() {
	cli.validateArgs()

	addLoliCmd := flag.NewFlagSet("addloli", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addLoliData := addLoliCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addloli":
		err := addLoliCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addLoliCmd.Parsed() {
		if *addLoliData == "" {
			addLoliCmd.Usage()
			os.Exit(1)
		}
		cli.addLoli(*addLoliData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}
