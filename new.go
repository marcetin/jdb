package jdb

import (
	"context"
	"gioui.org/widget"
	"github.com/gioapp/cms/pkg/items"
	"github.com/gioapp/cms/pkg/jdb/cfg"
	"github.com/gioapp/cms/pkg/jdb/repo"
	"github.com/ipfs/go-cid"
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

func (j *JavazacDB) ReadList(hash string) (itms items.I) {
	c, _ := cid.Decode(hash)
	rsc, err := j.peer.Get(j.ctx, c)
	checkError(err)
	for _, item := range rsc.Links() {
		//pss, err := rsc.Stat()
		//checkError(err)
		//nonono, err := item.GetNode(j.ctx, j.peer)
		//checkError(err)
		//nns, err := nonono.Stat()
		//checkError(err)

		itms = append(itms, &items.FolderListItem{
			Name: item.Name,
			Cid:  item.Cid,
			Size: item.Size,
			//Type:  uint8,
			Btn:   new(widget.Clickable),
			Check: new(widget.Bool),
		})
	}
	return
}
