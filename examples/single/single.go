package main

// This example launches an IPFS-Lite peer and fetches a hello-world
// hash from the IPFS network.

import (
	"context"
	"fmt"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/marcetin/jdb"
	"github.com/w-ingsolutions/c/model"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	crypto.MinRsaKeyBits = 1024
	ds, err := jdb.BadgerDatastore("test")
	if err != nil {
		panic(err)
	}
	peer := jdb.GetPeer(ctx, ds)
	//fmt.Println(string(content))
	var materijal model.WingMaterijal

	//m := model.WingMaterijal{
	//	Id:                2,
	//	Naziv:             "Masa za španski zid",
	//	Opis:              "Masa za španski zid",
	//	Obracun:           "Obračun po kilogrammu",
	//	Proizvodjac:       "evrojug",
	//	OsobineNamena:     "\"\nŠpanski zid je namenjen za dekorativnu zaštitu unutrašnjih zidova. Izrađen je od mineralnih agregata i hidrauličnih veziva aditiva. Španski zid se proizvodi u dve veličine zrna: 0mm i 1mm.\"",
	//	NacinRada:         "Masa se meša sa vodom u odnosu 1:3 (3 dela mase i 1 deo vode) ručno ili mešalicom. Zamešana masa treba da odstoji 10-25 minuta, zatim ponovo promeša i nanosi. Španski zid se nanosi na unutrašnje zidove koji su suvi, čisti i bez labavih delova i impregnirani Akrilnom emulzijom podloga. Masa se nanosi u jednom sloju metalnom gletaricom, a za krajnji izgled mogu se upotrebiti različiti valjci, pluta, stiropor, plastična gletarica itd.",
	//	JedinicaPotrosnje: "m2/kg",
	//	Potrosnja:         2,
	//	RokUpotrebe:       "\"12 meseci od datuma proizvodnje istaknutog na ambalaži. Cuvati\nu originalnoj, dobro zatvorenoj i neoštecenoj ambalaži, pri temperaturi od +5°C do +25\"",
	//	Jedinica:          "kg",
	//	Pakovanje:         25,
	//	Cena:              0.19,
	//	Slug:              "masa_za_spanski_zid",
	//}
	//jdb.Write("glet_masa_gletokol.gob", m)

	jdb.Read(ctx, peer, "QmPwhwbzPapHA3sA4UjD1m1Zw78pboQx6m9FyhK3cYN1dB", &materijal)
	fmt.Println("WingMaterijal", materijal)
}
