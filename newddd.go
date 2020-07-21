package db

import (
	"context"
	"fmt"

	"github.com/ipfs/go-cid"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/multiformats/go-multiaddr"
)

type Driver struct {
	ctx context.Context
	p   *Peer
	db  string
}

func DataBase(db string) *Driver {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	crypto.MinRsaKeyBits = 1024
	ds, err := BadgerDatastore("datastore")
	if err != nil {
		panic(err)
	}
	priv, _, err := crypto.GenerateKeyPair(crypto.RSA, 2048)
	if err != nil {
		panic(err)
	}
	listen, _ := multiaddr.NewMultiaddr("/ip4/0.0.0.0/tcp/4005")
	h, dht, err := SetupLibp2p(
		ctx,
		priv,
		nil,
		[]multiaddr.Multiaddr{listen},
		ds,
		Libp2pOptionsExtra...,
	)
	if err != nil {
		panic(err)
	}
	p, err := newPeer(ctx, ds, h, dht, nil)
	if err != nil {
		panic(err)
	}
	p.Bootstrap(DefaultBootstrapPeers())
	ib := &Driver{
		ctx: ctx,
		p:   p,
		db:  db,
	}
	return ib
}

func (ib *Driver) Collection() {
	c, _ := cid.Decode(ib.db)
	node, err := ib.p.Get(ib.ctx, c)
	if err != nil {
		panic(err)
	}
	nnn := node.Links()
	for _, er := range nnn {
		fmt.Println("00000000000000000")
		fmt.Println("Name", er.Name)
		fmt.Println("Cid", er.Cid)
		fmt.Println("Size", er.Size)
		//fmt.Println("terer", er.GetNode())
	}
	//navNode := format.NewNavigableIPLDNode(node, ib.p.DAGService)
	//for i := 0; i < int(navNode.ChildTotal()); i++ {
	//	childNode, err := navNode.FetchChild(ib.ctx, uint(i))
	//	if err != nil {
	//		panic(err)
	//	}
	//	n := format.ExtractIPLDNode(childNode)
	//	childCID := n.Cid().String()
	//	//achildCID,_ := n.Stat()
	//	fmt.Println("graphOut", childCID)
	//	//fmt.Println("graphOuaaaaaat", achildCID)
	//}
}
