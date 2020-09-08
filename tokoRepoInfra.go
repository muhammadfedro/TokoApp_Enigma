package main

import (
	"encoding/json"
	"io/ioutil"
)

type TokoRepoInfra struct {
	dataPath string
}

func newTokoRepoInfra(dataPath string) *TokoRepoInfra {
	return &TokoRepoInfra{dataPath}
}

func (tri *TokoRepoInfra) saveToFile(tokoCollection *[]*Toko) {
	file, _ := json.MarshalIndent(tokoCollection, "", " ")
	_ = ioutil.WriteFile(tri.dataPath, file, 0644)
}

func (tri *TokoRepoInfra) readFile(tokoCollection *[]*Toko) {
	file, _ := ioutil.ReadFile(tri.dataPath)
	_ = json.Unmarshal(file, tokoCollection)
}
