package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/iMil/goxplorer/blockchain"
)

func TestBlockCount(t *testing.T) {
	// Contains the dump of bitcoin's level DB.
	blockFile := "./data"

	db, err := blockchain.OpenDB(blockFile)
	require.NoError(t, err)

	// Load entire Block data.
	blockMap := make(map[uint64]*blockchain.Block)

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		data := iter.Value()
		if data == nil {
			continue
		}

		height, err := blockchain.ReadBlockRecord(data, "nHeight")
		require.NoError(t, err)

		var b blockchain.Block
		b.ReadHeader(data[8:88])
		blockMap[height] = &b
	}

	var preBlock *blockchain.Block
	var count uint64
	for blockHeight := uint64(0); ; blockHeight++ {
		block, ok := blockMap[blockHeight]
		if !ok {
			fmt.Printf("block %d not found\n", blockHeight)
			break
		}

		if preBlock != nil && (int64(block.Header.TimeStamp)-int64(preBlock.Header.TimeStamp) > 20*60) {
			count++
		}
		preBlock = block
	}
	fmt.Println("count of block produced after 20 minutes: ", count)
}
