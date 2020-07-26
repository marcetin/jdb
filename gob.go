package jdb

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/ipfs/go-cid"
	"os"
)

func (j *JavazacDB) Write(key []byte) {
	buf := bytes.NewReader(key)
	n, err := j.peer.AddFile(j.ctx, buf, nil)
	checkError(err)
	fmt.Println("cii:", n.Cid())
	fmt.Println("cii:", n.Tree("", -1))
}

func (j *JavazacDB) Read(fileName string, key interface{}) {
	c, _ := cid.Decode(fileName)
	rsc, err := j.peer.GetFile(j.ctx, c)
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
