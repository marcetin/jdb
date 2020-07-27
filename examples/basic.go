package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/marcetin/jdb"
)

func main() {
	// initialize db options
	opts := jdb.DefaultOptions
	// Set PrivateKey. This should come from an ENV or a secret store in the real world
	opts.PrivateKey, _ = hex.DecodeString("44667768254d593b7ea48c3327c18a651f6031554ca4f5e3e641f6ff1ea72e98")
	db, err := jdb.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Set("language", []byte("Go"))
	if err != nil {
		log.Fatal(err)
	}
	answer, err := db.Get("language")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("This software is made with Go: ", string(answer))
	// The meaning of life is:  42
}
