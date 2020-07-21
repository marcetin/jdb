package jdb

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"github.com/ipfs/go-cid"
	"os"
)

//func Write() {
//	person := model.WingMaterijal{
//		Id:                1,
//		Naziv:             "Glet masa-gletokol",
//		Opis:              "Glet masa-gletokol",
//		Obracun:           "Obračun po kilogrammu",
//		Proizvodjac:       "maxima",
//		OsobineNamena:     "GLETOLIN G je pastozna masa na bazi akrilata, namenjena za no gletovanje unutrašnjih, cvrsto malterisanih zidnih površina, kao i gips kartonskih ploca. GLETOLIN G odlikuje izuzetna belina, lakoca nanošenja i obradivost. Izravnane površine, koje su potpuno suve, mogu se bojiti svim vrstama disperzionih boja.",
//		NacinRada:         "GLETOLIN G je pripremljen za ugradnju. Zidne površine moraju biti nosivo sposobne, ocišcene od prašine, masti i drugih necistoca. Pre nanošenja GLETOLIN-a G, površine grundirati odgovarajucom MAXIKRIL podlogom, u zavisnosti od stanja površine. GLETOLIN G nanositi na zidne površine pomocu celicne gletarice ili nerajucom farbarskom lopaticom.\nMože se nanositi i mašinski, pri cemu ga treba razrediti sa malo vode. Nakon sušenja, svaki sloj nanosa treba dobro izbrusiti brusnim papirom. Vreme sušenja, pre nanošenja narednog sloja, odnosno brušenja, pri optimalnim uslovima, je min. 12 casova. Maksimalna preporucena ukupna debljina sloja je do 3mm. Posle upotrebe, alat odmah oprati vodom",
//		JedinicaPotrosnje: "kg/m², u zavisnosti od hrapavosti podloge",
//		Potrosnja:         2.0,
//		RokUpotrebe:       "12 meseci od datuma proizvodnje istaknutog na ambalaži. Cuvati u originalnoj, dobro zatvorenoj i neoštecenoj ambalaži, pri temperaturi od +5°C do +25",
//		Jedinica:          "kg",
//		Pakovanje:         25,
//		Cena:              0.13,
//		Slug:              "glet_masa_gletokol",
//	}
//	saveGob("glet_masa_gletokol.gob", person)
//}

func Write(ctx context.Context, peer *Peer, fileName string, key interface{}) {
	//outFile, err := os.Create(fileName)
	//checkError(err)
	//encoder := gob.NewEncoder(outFile)
	//err = encoder.Encode(key)
	//checkError(err)
	//outFile.Close()

	content := []byte("hola")

	buf := bytes.NewReader(content)
	n, err := peer.AddFile(context.Background(), buf, nil)
	checkError(err)

	fmt.Println("cii:", n.Cid())

	//encoder := gob.NewEncoder(buf)
	//err = encoder.Encode(content)
	//checkError(err)

}

func Read(ctx context.Context, peer *Peer, fileName string, key interface{}) {
	c, _ := cid.Decode(fileName)
	rsc, err := peer.GetFile(ctx, c)
	if err != nil {
		panic(err)
	}
	decoder := gob.NewDecoder(rsc)
	err = decoder.Decode(key)
	checkError(err)
	defer rsc.Close()
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
