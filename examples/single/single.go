package main

// This example launches an IPFS-Lite peer and fetches a hello-world
// hash from the IPFS network.

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/marcetin/jdb"
	"github.com/w-ingsolutions/c/model"
)

func main() {
	j := jdb.New("tete")
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//crypto.MinRsaKeyBits = 1024
	//ds, err := jdb.BadgerDatastore("tes33333t")
	//if err != nil {
	//	panic(err)
	//}
	//peer := jdb.GetPeer(ctx, ds)
	//fmt.Println(string(content))
	var materijal model.WingMaterijal

	m := model.WingMaterijal{
		Id:                2,
		Naziv:             "Masa za španski zid",
		Opis:              "Masa za španski zid",
		Obracun:           "Obračun po kilogrammu",
		Proizvodjac:       "evrojug",
		OsobineNamena:     "Španski zid je.",
		NacinRada:         "Masa se meša s.",
		JedinicaPotrosnje: "m2/kg",
		Potrosnja:         2,
		RokUpotrebe:       "12 meseci od datuma proizvodnje istaknutog na ambalaži. Cuvati u originalnoj, dobro zatvorenoj i neoštecenoj ambalaži, pri temperaturi od +5°C do +25",
		Jedinica:          "kg",
		Pakovanje:         25,
		Cena:              0.19,
		Slug:              "masa_za_spanski_zid",
	}

	var bytesBuf bytes.Buffer
	encoder := gob.NewEncoder(&bytesBuf)
	err := encoder.Encode(m)
	if err != nil {

	}
	j.Write(bytesBuf.Bytes())

	j.Read("bafybeihs3p4g232wocqd5ouoakrwm5kocjaexpar6oekkn6qzez5nqj3vu", &materijal)
	fmt.Println("WingMaterijal", materijal)
}
