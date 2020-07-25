package jdb

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"github.com/ipfs/go-cid"
	"os"
)

func Write(ctx context.Context, peer *Peer, key interface{}) {
	var bytesBuf bytes.Buffer

	buf := bytes.NewReader(bytesBuf.Bytes())
	n, err := peer.AddFile(ctx, buf, nil)
	checkError(err)
	encoder := gob.NewEncoder(&bytesBuf)
	err = encoder.Encode(key)
	fmt.Println("cii:", n.Cid())
}

func Read(ctx context.Context, peer *Peer, fileName string, key interface{}) {
	c, _ := cid.Decode(fileName)
	rsc, err := peer.GetFile(ctx, c)
	if err != nil {
		panic(err)
	}
	decoder := gob.NewDecoder(rsc)
	err = decoder.Decode(key)
	checkError(err)
	defer rsc.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
