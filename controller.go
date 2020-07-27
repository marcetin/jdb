package jdb

import (
	"bytes"
	"github.com/dgraph-io/badger"
	"github.com/ipfs/go-cid"
	"io/ioutil"
)

func (j *JavazacDB) writeKeyValuePair(key, value string) error {
	txn := j.localDB.NewTransaction(true)
	defer txn.Discard()

	err := txn.Set([]byte(key), []byte(value))
	if err != nil {
		return err
	}

	return txn.Commit()
}

func (j *JavazacDB) readKeyValuePair(key string) (string, error) {
	txn := j.localDB.NewTransaction(false)
	defer txn.Discard()

	item, err := txn.Get([]byte(key))
	if err != nil {
		return "", err
	}

	var returnedValue []byte
	err = item.Value(func(val []byte) error {
		returnedValue = val
		return nil
	})
	if err != nil {
		return "", err
	}

	return string(returnedValue), nil

}

// Set acts in two stages,
// first encode your <value> and put into distributed IPFS.
// second, save (locally) a relation between your <key> and
// the cid generated from your save performed by ipfs
func (j *JavazacDB) Set(key string, value []byte) error {
	data := encrypt(value, j.encryptKey)

	buf := bytes.NewBuffer(data)

	n, err := j.peer.AddFile(j.ctx, buf, nil)
	if err != nil {
		return err
	}

	err = j.writeKeyValuePair(key, n.Cid().String())
	if err != nil {
		return err
	}

	return nil
}

// Get performs a get action of your JavazacDB,
// it works in two stages (like set function), firs read
// the relation between your key and a one cid, second, collect
// this cid from ipfs and finally tries to decrypt and return
// the value of your key
func (j *JavazacDB) Get(key string) ([]byte, error) {
	hash, err := j.readKeyValuePair(key)
	if err != nil {
		return nil, err
	}
	c, _ := cid.Decode(hash)

	reader, err := j.peer.GetFile(j.ctx, c)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	dData := decrypt(data, j.encryptKey)
	return dData, nil
}

// GetQmFromKey  returns your cid related to your key,
// if exist, of course
func (db *JavazacDB) GetQmFromKey(key string) (string, error) {
	hash, err := db.readKeyValuePair(key)
	if err != nil {
		return "", err
	}
	return hash, nil
}

// Update takes your value, encode and save the relation,
// it is exactly equal to set action
func (j *JavazacDB) Update(key string, value []byte) error {
	data := encrypt(value, j.encryptKey)

	buf := bytes.NewBuffer(data)

	n, err := j.peer.AddFile(j.ctx, buf, nil)
	if err != nil {
		return err
	}

	err = j.writeKeyValuePair(key, n.Cid().String())
	if err != nil {
		return err
	}

	return nil
}

// MetaDB is a db by reference, a way you can to create a very distrubuted DB
type MetaDB struct {
	encryptKey    []byte
	principalNode string

	publicKey  string
	privateKey []byte

	localDB *badger.DB
}

func (j *JavazacDB) snapshotMetaDB(publicKey string, privateKey []byte) (*MetaDB, error) {
	data, err := encodeDBFile(j.options.LocalDBDir, privateKey)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(data)
	n, err := j.peer.AddFile(j.ctx, buf, nil)
	if err != nil {
		return nil, err
	}

	return &MetaDB{
		encryptKey:    j.encryptKey,
		principalNode: j.principalNode,
		publicKey:     n.Cid().String(),
		privateKey:    privateKey,
		localDB:       j.localDB,
	}, nil
}

func (j *JavazacDB) readMetaDB(publicKey string, privateKey []byte) (*MetaDB, error) {
	c, _ := cid.Decode(publicKey)

	reader, err := j.peer.GetFile(j.ctx, c)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	err = decodeDBFile(data, privateKey, j.options.LocalDBDir)
	if err != nil {
		return nil, err
	}

	ldb, err := badger.Open(badger.DefaultOptions(j.options.LocalDBDir))
	if err != nil {
		return nil, err
	}

	return &MetaDB{
		encryptKey:    j.encryptKey,
		principalNode: j.principalNode,
		publicKey:     publicKey,
		privateKey:    privateKey,
		localDB:       ldb,
	}, nil

}
