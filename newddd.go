package db

import (
	"context"
	"fmt"

	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/multiformats/go-multiaddr"
)

func cPeer() (context.Context, *Peer, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//log.SetLogLevel("*", "warn")
	// Bootstrappers are using 1024 keys. See:
	// https://github.com/ipfs/infra/issues/378
	crypto.MinRsaKeyBits = 1024
	ds, err := BadgerDatastore("test")
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
	return ctx, p, err
}

func (d *Driver) DataBase() {
	c, _ := cid.Decode("QmS4ustL54uo8FzR9455qaxZwuMiUhyvMcX9Ba8nUH4uVv")
	node, err := d.p.Get(d.ctx, c)
	if err != nil {
		panic(err)
	}
	navNode := format.NewNavigableIPLDNode(node, d.p.DAGService)
	for i := 0; i < int(navNode.ChildTotal()); i++ {
		childNode, err := navNode.FetchChild(d.ctx, uint(i))
		if err != nil {
			panic(err)
		}
		n := format.ExtractIPLDNode(childNode)
		schildCID := n.Cid().String()
		//childCID := n.Tree("",-1)
		//fmt.Println("graphOut", childCID)
		fmt.Println("ssss", schildCID)
	}
}
