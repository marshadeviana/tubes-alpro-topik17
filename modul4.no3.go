package main

import "fmt"

var jenisKendaraan string
var jam1, menit1, detik1 int
var jam2, menit2, detik2 int
var totalUang int

func tampil_menu() {
	fmt.Println("--------------------------")
	fmt.Println("        M E N U           ")
	fmt.Println("--------------------------")
	fmt.Println("1. Input Kendaraan Masuk")
	fmt.Println("2. Input Kendaraan Keluar")
	fmt.Println("3. Hitung Biaya Parkir")
	fmt.Println("4. Cetak Total Uang")
	fmt.Println("5. Exit")
	fmt.Println("--------------------------")
}

func inputKendaraanMasuk() {
	fmt.Print("Masukkan jenis kendaraan (mobil/motor): ")
	fmt.Scan(&jenisKendaraan)
	fmt.Print("Masukkan jam, menit, detik kendaraan masuk: ")
	fmt.Scan(&jam1, &menit1, &detik1)
}

func inputKendaraanKeluar() {
	fmt.Print("Masukkan jam, menit, detik kendaraan keluar: ")
	fmt.Scan(&jam2, &menit2, &detik2)
}

func reset() {
	jenisKendaraan = ""
	jam1, menit1, detik1 = 0, 0, 0
	jam2, menit2, detik2 = 0, 0, 0
}

func hitungBiayaParkir() {
	t1 := jam1*3600 + menit1*60 + detik1
	t2 := jam2*3600 + menit2*60 + detik2
	selisih := t2 - t1

	durasiJam := selisih / 3600
	if selisih%3600 > 0 {
		durasiJam++
	}

	biaya := 0
	if jenisKendaraan == "mobil" {
		if durasiJam > 0 {
			biaya = 5000 + (durasiJam-1)*3000
		}
	} else if jenisKendaraan == "motor" {
		if durasiJam > 0 {
			biaya = 2000 + (durasiJam-1)*1000
		}
	}

	totalUang += biaya
	fmt.Printf("Biaya parkir %s selama %d jam: Rp %d,-\n", jenisKendaraan, durasiJam, biaya)
	reset()
}

func cetakTotalUang() {
	fmt.Printf("Total uang: Rp %d,-\n", totalUang)
}

func main() {
	var pilih int
	for {
		tampil_menu()
		fmt.Print("Pilih (1/2/3/4/5)? ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			inputKendaraanMasuk()
		} else if pilih == 2 {
			inputKendaraanKeluar()
		} else if pilih == 3 {
			hitungBiayaParkir()
		} else if pilih == 4 {
			cetakTotalUang()
		} else if pilih == 5 {
			break
		}
	}
}