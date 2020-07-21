package main

import (
	"context"
	"fmt"
	"github.com/marcetin/jdb"

	"github.com/ipfs/go-cid"
	format "github.com/ipfs/go-ipld-format"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/multiformats/go-multiaddr"
)

var (
	testCID = "QmWC7wxt6z69ndm2JPYvkbp6xY6ZJ4ShaKfa98So82wzMy"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//log.SetLogLevel("*", "warn")

	// Bootstrappers are using 1024 keys. See:
	// https://github.com/ipfs/infra/issues/378
	crypto.MinRsaKeyBits = 1024

	ds, err := jdb.BadgerDatastore("test")
	if err != nil {
		panic(err)
	}
	priv, _, err := crypto.GenerateKeyPair(crypto.RSA, 2048)
	if err != nil {
		panic(err)
	}

	listen, _ := multiaddr.NewMultiaddr("/ip4/0.0.0.0/tcp/4005")

	h, dht, err := jdb.SetupLibp2p(
		ctx,
		priv,
		nil,
		[]multiaddr.Multiaddr{listen},
		ds,
		jdb.Libp2pOptionsExtra...,
	)

	if err != nil {
		panic(err)
	}

	lite, err := jdb.New(ctx, ds, h, dht, nil)
	if err != nil {
		panic(err)
	}

	lite.Bootstrap(jdb.DefaultBootstrapPeers())

	c, _ := cid.Decode(testCID)
	node, err := lite.Get(ctx, c)
	if err != nil {
		panic(err)
	}
	navNode := format.NewNavigableIPLDNode(node, lite.DAGService)
	//for i := 0; i < int(navNode.ChildTotal()); i++ {
	//	childNode, err := navNode.FetchChild(ctx, uint(i))
	//	if err != nil {
	//		panic(err)
	//	}
	//	n := format.ExtractIPLDNode(childNode)
	//	childCID := n.Cid().String()
	//	sschildCID := n.Links()
	//	fmt.Println("graphOut", childCID)
	//	fmt.Println("graphOut", sschildCID)
	//}
	for _, link := range navNode.GetIPLDNode().Links() {
		fmt.Println("Cid:", link.Cid)
		fmt.Println("Name:", link.Name)
		fmt.Println("Size:", link.Size)
		fmt.Println("----------")
	}
	//jdb.Write()
	//jdb.Read()
}
