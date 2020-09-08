package main

type Toko struct {
	Barcode string
	Barang string
	Kategori string
	Harga float64
	Stok float64
	Deskripsi string
	Ukuran string
	Fungsi string
	Jumlah float64
	Id string
	Tanggal string
	Nama string
}

func newToko(barcode string, barang string, kategori string, harga float64, stok float64, deskrpsi string, ukuran string, fungsi string) Toko {
	return Toko{
		Barcode: barcode,
		Barang: barang,
		Kategori: kategori,
		Harga: harga,
		Stok: stok,
		Deskripsi: deskrpsi,
		Ukuran: ukuran,
		Fungsi: fungsi,
	}
}

func newTransaksi(barcode string, id string, nama string, jumlah float64, tanggal string) Toko{
	return Toko{
		Barcode: barcode,
		Id:        id,
		Nama:    nama,
		Jumlah:    jumlah,
		Tanggal:   tanggal,
	}
}