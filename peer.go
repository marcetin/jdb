package jdb

import (
	"context"
	"github.com/ipfs/go-datastore"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/multiformats/go-multiaddr"
)

func GetPeer(ctx context.Context, ds datastore.Batching) *Peer {
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
	p, err := New(ctx, ds, h, dht, nil)
	if err != nil {
		panic(err)
	}
	p.Bootstrap(DefaultBootstrapPeers())
	return p
}
