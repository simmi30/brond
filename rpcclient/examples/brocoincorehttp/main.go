// Copyright (c) 2014-2017 The brsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/brsuite/brond/rpcclient"
)

func main() {
	// Connect to local brocoin core RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host:         "localhost:8332",
		User:         "yourrpcuser",
		Pass:         "yourrpcpass",
		HTTPPostMode: true, // Brocoin core only supports HTTP POST mode
		DisableTLS:   true, // Brocoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	// Get the current block count.
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block count: %d", blockCount)
}
