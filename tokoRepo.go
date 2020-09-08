package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

type TokoRepo struct {
	tokoCollection      *[]*Toko
	tokoCollectionInfra *TokoRepoInfra
}

func newTokoRepo(dataPath string, tokoCollection []*Toko) *TokoRepo {
	TokoRepoInfra := newTokoRepoInfra(dataPath)
	return &TokoRepo{tokoCollection: &tokoCollection, tokoCollectionInfra: TokoRepoInfra}
}

func (tr *TokoRepo) AddNewToko(toko *Toko) {
	data := []byte(toko.Barang)
	toko.Barcode = fmt.Sprintf("%x", md5.Sum(data))
	*tr.tokoCollection = append(*tr.tokoCollection, toko)
	tr.tokoCollectionInfra.saveToFile(tr.tokoCollection)
}
func (tr *TokoRepo) AddNewTrans(toko *Toko) {
data := []byte(toko.Barang)
toko.Id = fmt.Sprintf("%x", md5.Sum(data))
tanggal := time.Now()
	toko.Tanggal = fmt.Sprintf("%x",tanggal.Format("2006.01.02"))
*tr.tokoCollection = append(*tr.tokoCollection, toko)
tr.tokoCollectionInfra.saveToFile(tr.tokoCollection)
}

func (tr *TokoRepo) FindAllBarang() []*Toko {
	tr.tokoCollectionInfra.readFile(tr.tokoCollection)
	return *tr.tokoCollection
}

