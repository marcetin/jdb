package jdb

import (
	"context"
	"github.com/dgraph-io/badger"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
)

func Open(options *Options) (*JavazacDB, error) {
	if err := ValidateKey(options.PrivateKey); err != nil {
		return nil, err
	}
	j := new(JavazacDB)
	j.encryptKey = options.PrivateKey
	j.principalNode = options.PrincipalNode
	j.options = options
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	crypto.MinRsaKeyBits = 1024
	ds, err := BadgerDatastore("datastore")
	if err != nil {
		panic(err)
	}
	j.ctx = ctx
	j.peer = GetPeer(j.ctx, ds)
	opts := badger.DefaultOptions(options.LocalDBDir).WithLogger(nil)
	ldb, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	j.localDB = ldb
	return j, nil
}
