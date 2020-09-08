package main

type TokoService struct {
	r *TokoRepo
}

func newTokoService(repo *TokoRepo) *TokoService {
	return &TokoService{r: repo}
}

func (ts *TokoService) registerNewBarang(b *Toko) {
	ts.r.AddNewToko(b)
}
func (ts *TokoService) registerNewTrans(b *Toko) {
	ts.r.AddNewTrans(b)
}

func (ts *TokoService) getAllTokoCollection() []*Toko {
	return ts.r.FindAllBarang()
}

