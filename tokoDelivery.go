package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type TokoDelivery struct {
	tokoService *TokoService
}

func newTokoDelivery(c *AppConfig) *TokoDelivery {
	repo := newTokoRepo(c.getDbPath(), c.getDB())
	tokoService := newTokoService(repo)
	return &TokoDelivery{tokoService}
}

func (td *TokoDelivery) Create() {
	var isExist = false
	var userChoice string

	td.mainMenuForm()
	for isExist == false {
		fmt.Printf("\n%s", "Your Choice: ")
		fmt.Scan(&userChoice)
		switch {
		case userChoice == "01":
			td.registrationTokoForm()
		case userChoice == "02":
			td.viewTokoCollectionForm()
		case userChoice == "03":
			td.transaksiForm()
		case userChoice == "04":
			td.viewTransaksi()
		case userChoice == "q":
			isExist = true
		default:
			fmt.Println("Unknown Menu Code")

		}
	}
}

func (td *TokoDelivery) menuChoiceOrdered(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func (td *TokoDelivery) mainMenuForm() {
	var appMenu = map[string]string{
		"01": "Tambah Barang",
		"02": "Daftar Barang",
		"03": "Tambah Transaksi",
		"04": "Daftar Transaksi",
		"q":  "Quit aplication",
	}
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	fmt.Printf("%26s\n", "Daftar Barang")
	fmt.Printf("%s\n", strings.Repeat("*", 30))
	for _, menuCode := range td.menuChoiceOrdered(appMenu) {
		fmt.Printf("%s. %s\n", menuCode, appMenu[menuCode])
	}
}

func (td *TokoDelivery) registrationTokoForm() {
	consoleClear()
	var barcode string
	var barang string
	var kategori string
	var harga float64
	var stok float64
	var deskripsi string
	var ukuran string
	var fungsi string
	var confirmation string
	fmt.Println()
	fmt.Printf("%s\n", "Registrasi Form")
	fmt.Printf("%s\n", strings.Repeat("-", 30))
	scanner := bufio.NewReader(os.Stdin)

	sbarcode, _ := scanner.ReadString('\n')
	barcode = strings.TrimSpace(sbarcode)

	fmt.Print("Barang : ")
	sbarang, _ := scanner.ReadString('\n')
	barang = strings.TrimSpace(sbarang)

	fmt.Print("Kategori : ")
	skategori, _ := scanner.ReadString('\n')
	kategori = strings.TrimSpace(skategori)

	fmt.Print("Harga : ")
	sharga, _ := scanner.ReadString('\n')
	harga, _ = strconv.ParseFloat(strings.TrimSpace(sharga), 64)

	fmt.Print("Stok : ")
	sstok, _ := scanner.ReadString('\n')
	stok, _ = strconv.ParseFloat(strings.TrimSpace(sstok),64)

	fmt.Print("Deskriosi : ")
	sdeskrpisi, _ := scanner.ReadString('\n')
	deskripsi = strings.TrimSpace(sdeskrpisi)

	fmt.Print("Ukuran : ")
	sukuran, _ := scanner.ReadString('\n')
	ukuran = strings.TrimSpace(sukuran)
	fmt.Print("Fungsi : ")
	sfungsi, _ := scanner.ReadString('\n')
	fungsi = strings.TrimSpace(sfungsi)

	fmt.Println("Save to collection? :Y/n")
	fmt.Scanln(&confirmation)

	if confirmation == "Y" {
		newToko := newToko(barcode, barang, kategori, harga, stok, deskripsi,ukuran,fungsi)
		td.tokoService.registerNewBarang(&newToko)
	}
	consoleClear()
	td.mainMenuForm()

}

func (td *TokoDelivery) viewTokoCollectionForm() {
	consoleClear()
	Tokos := td.tokoService.getAllTokoCollection()
	fmt.Println("")
	fmt.Println("Daftar Toko Seadanya")
	fmt.Printf("%s\n", strings.Repeat("=", 150))
	fmt.Printf("%-40s%-15s%-15s%-12s%-10s%-15s%-15s%-15s\n", "Barcode", "Barang", "Kategori", "Harga", "Stok","Deskripsi","Ukuran","Fungsi")
	fmt.Printf("%s\n", strings.Repeat("-", 150))
	for _, b := range Tokos {
		if b.Stok != 0 {
			fmt.Printf("%-40s%-15s%-15s%-12.2f%-10.0f%-15s%-15s%-15s\n", b.Barcode, b.Barang, b.Kategori, b.Harga, b.Stok, b.Deskripsi, b.Ukuran, b.Fungsi)
		}
		}
	var confirmation string
	fmt.Print("Back To Menu ? : (Y)")
	fmt.Scanln(&confirmation)

	if confirmation == "Y" {
		consoleClear()
		td.mainMenuForm()
	}
}



func (td *TokoDelivery) transaksiForm() {
	consoleClear()
	var barcode string
	var id string
	var nama string
	var jumlah float64
	var tanggal string
	var confirmation string
	fmt.Println()
	fmt.Printf("%s\n", "Formulir Transaksi Seadanya")
	fmt.Printf("%s\n", strings.Repeat("-", 30))
	scanner := bufio.NewReader(os.Stdin)


	sbarcode, _ := scanner.ReadString('\n')
	barcode = strings.TrimSpace(sbarcode)
	fmt.Print("Id Transaction : ")
	sid, _ := scanner.ReadString('\n')
	id = strings.TrimSpace(sid)

	fmt.Print("Barang : ")
	snama, _ := scanner.ReadString('\n')
	nama = strings.TrimSpace(snama)

	fmt.Print("Jumlah : ")
	sjumlah, _ := scanner.ReadString('\n')
	jumlah, _ = strconv.ParseFloat(strings.TrimSpace(sjumlah), 64)

	fmt.Print("Tanggal : ")
	stanggal, _ := scanner.ReadString('\n')
	tanggal = strings.TrimSpace(stanggal)


	fmt.Println("Save to collection? :Y/n")
	fmt.Scanln(&confirmation)

	if confirmation == "Y" {
		newTransaksi := newTransaksi(barcode, id, nama, jumlah, tanggal)
		td.tokoService.registerNewBarang(&newTransaksi)
	}
	consoleClear()
	td.mainMenuForm()

}

func (td *TokoDelivery) viewTransaksi() {
	consoleClear()
	Tokos := td.tokoService.getAllTokoCollection()
	fmt.Println("")
	fmt.Println("Transaksi Toko Seadanya")
	fmt.Printf("%s\n", strings.Repeat("=", 150))
	fmt.Printf("%-40s%-15s%-15s%-12s\n", "ID", "Barang", "Jumlah", "Tanggal")
	fmt.Printf("%s\n", strings.Repeat("-", 150))


	for _, b := range Tokos {
		if b.Jumlah != 0{
			fmt.Printf("%-40s%-15s%-12.0f%-15s\n", b.Id,b.Nama,b.Jumlah,b.Tanggal)
		}

	}

	var confirmation string
	fmt.Print("Back To Menu ? : (Y)")
	fmt.Scanln(&confirmation)

	if confirmation == "Y" {
		consoleClear()
		td.mainMenuForm()
	}
}





