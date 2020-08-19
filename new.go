package jdb

import (
	"context"
	"github.com/gioapp/cms/pkg/jdb/cfg"
	"github.com/gioapp/cms/pkg/jdb/repo"
	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	"os"
)

// JavazacDb Structure
type JavazacDB struct {
	ctx   context.Context
	peer  *Peer
	index map[string]string
	store string
}

func New(ctx context.Context, store string) *JavazacDB {
	j := &JavazacDB{
		ctx:   ctx,
		index: make(map[string]string),
		store: store,
	}
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//c, _ := cid.Decode(hash)
	//j.cid = c
	root := j.store + string(os.PathSeparator) + repo.Root
	conf, err := cfg.ConfigInit(2048)
	checkError(err)
	err = repo.Init(root, conf)
	checkError(err)

	r, err := repo.Open(root)
	checkError(err)
	peer, err := NewPeer(j.ctx, r)
	checkError(err)
	peer.Bootstrap(DefaultBootstrapPeers())
	j.peer = peer
	return j
}

func (j *JavazacDB) ReadList(c cid.Cid) []*format.Link {
	rsc, err := j.peer.Get(j.ctx, c)
	checkError(err)
	return rsc.Links()
}
