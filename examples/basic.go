package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"github.com/marcetin/jdb"
	"log"
)

type jdbTest struct {
	Name   string
	Number int
	Float  float64
}

func main() {
	//jt := jdbTest{}
	// initialize db options
	opts := jdb.DefaultOptions
	//Set PrivateKey. This should come from an ENV or a secret store in the real world
	opts.PrivateKey, _ = hex.DecodeString("44667768254d593b7ea48c3327c18a651f6031554ca4f5e3e641f6ff1ea72e98")
	db, err := jdb.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//err = db.Set("language", []byte("Go"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//answer, err := db.Get("language")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("This software is made with programming language: ", string(answer))
	// The meaning of life is:  42

	jtNew := jdbTest{
		Name:   "test",
		Number: 1,
		Float:  6.18,
	}

	//var bytesBuf bytes.Buffer
	//encoder := gob.NewEncoder(&bytesBuf)
	//err = encoder.Encode(jtNew)
	//err = db.Set("test", bytesBuf.Bytes())
	//if err != nil {
	//}
	//
	//
	//
	//p, err := db.Get("test")
	//if err != nil {
	//}
	//
	//buf := bytes.NewReader(p)
	//decoder := gob.NewDecoder(buf)
	//err = decoder.Decode(jt)
	//if err != nil {
	//}
	//fmt.Println("Ima tameblu da da",jt)
	var write bytes.Buffer
	//var read bytes.Buffer
	enc := gob.NewEncoder(&write) // Will write to network.
	// Encode (send) the value.
	err = enc.Encode(jtNew)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	err = db.Set("test", write.Bytes())
	if err != nil {
	}

	// Decode (receive) the value.
	p, err := db.Get("test")
	if err != nil {
	}
	read := bytes.NewReader(p)
	dec := gob.NewDecoder(read) // Will read from network.

	var q jdbTest
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Println("test", q)
}
