package main

import "fmt"

type pengguna struct {
	hariKe  int
	durasi  int
	kalori  int
	tanggal string // Format: DDMMYYYY
	jenis   string
}

const Ndays int = 365

type workout [Ndays]pengguna

var P workout
var N int // Menyimpan jumlah data saat ini
var stop bool

func main() {
	N = 0 // Inisialisasi data awal dari 0
	for !stop {
		menu()
	}
}

func menu() {
	var Nomor int

	fmt.Println("-------------")
	fmt.Println("    Menu")
	fmt.Println("-------------")
	fmt.Println("1. Menambahkan data")
	fmt.Println("2. Edit data")
	fmt.Println("3. Hapus data")
	fmt.Println("4. Rekomendasi workout (Berdasarkan Pola Latihan)")
	fmt.Println("5. Cari workout (Sequential Search - Jenis Olahraga)")
	fmt.Println("6. Cari workout (Binary Search - Jenis Olahraga)")
	fmt.Println("7. Urutkan berdasarkan kalori (Selection Sort Descending)")
	fmt.Println("8. Urutkan berdasarkan durasi (Insertion Sort Ascending)")
	fmt.Println("9. Laporan (10 Aktivitas Terakhir & Total Kalori Periode)")
	fmt.Println("10. Tampilkan Semua Data")
	fmt.Println("11. Exit")
	fmt.Print("Pilih menu 1-11: ")
	fmt.Scan(&Nomor)
	switch Nomor {
	case 1:
		addData(&P, &N)
	case 2:
		editData(&P, N)
	case 3:
		hapusData(&P, &N)
	case 4:
		rekomendasi(P, N)
	case 5:
		cariOlahragaSequential(P, N)
	case 6:
		cariOlahragaBinary(P, N)
	case 7:
		sortKalori(P, N)
	case 8:
		sortDurasi(P, N)
	case 9:
		menuLaporan(P, N)
	case 10:
		fmt.Println("\n--- Seluruh Data Workout ---")
		showData(P, N)
	default:
		stop = true
	}
}

func addData(w *workout, n *int) {
	var pilih int = 1

	for pilih == 1 && *n < Ndays {
		fmt.Println("=======================")
		fmt.Println("   Menambahkan Data    ")
		fmt.Println("=======================")
		
		// Hari ke otomatis ditentukan dari jumlah data saat ini + 1
		w[*n].hariKe = *n + 1 
		fmt.Println("Masukkan data hari ke-", w[*n].hariKe)
		
		fmt.Print("Masukkan jadwal latihan (DDMMYYYY): ")
		fmt.Scan(&w[*n].tanggal)
		fmt.Print("Masukkan jenis latihan: ")
		fmt.Scan(&w[*n].jenis)
		fmt.Print("Masukkan durasi latihan (menit): ")
		fmt.Scan(&w[*n].durasi)
		fmt.Print("Masukkan total kalori: ")
		fmt.Scan(&w[*n].kalori)
		fmt.Println()
		
		*n++
		
		fmt.Print("Ketik 1 untuk menambah data lagi, 0 untuk keluar: ")
		fmt.Scan(&pilih)
		fmt.Println()
	}
}

// Poin C: Sequential Search berdasarkan Jenis Olahraga
func cariOlahragaSequential(w workout, n int) {
	var x string
	var found bool = false
	fmt.Print("Jenis olahraga yang ingin dicari (Sequential): ")
	fmt.Scan(&x)
	
	fmt.Printf("| %-10s | %-9s | %-20s | %-7s | %-7s |\n", "Hari ke-", "Tanggal", "Latihan", "Durasi", "Kalori")
	for i := 0; i < n; i++ {
		if w[i].jenis == x {
			fmt.Printf("| %-10d | %-9s | %-20s | %-7d | %-7d |\n", w[i].hariKe, w[i].tanggal, w[i].jenis, w[i].durasi, w[i].kalori)
			found = true
		}
	}
	if !found {
		fmt.Println("Data tidak ditemukan!")
	}
}

// Fungsi bantu untuk mengurutkan data berdasarkan alfabet Jenis Olahraga sebelum Binary Search
func sortAlphabet(w *workout, n int) {
	for pass := 1; pass < n; pass++ {
		i := pass
		temp := w[pass]
		for i > 0 && temp.jenis < w[i-1].jenis {
			w[i] = w[i-1]
			i--
		}
		w[i] = temp
	}
}

// Poin C: Binary Search berdasarkan Jenis Olahraga
func cariOlahragaBinary(w workout, n int) {
	var x string
	fmt.Print("Jenis olahraga yang ingin dicari (Binary): ")
	fmt.Scan(&x)

	// Binary search membutuhkan data terurut. Kita urutkan salinan datanya terlebih dahulu.
	sortAlphabet(&w, n)

	left := 0
	right := n - 1
	idx := -1

	for left <= right && idx == -1 {
		mid := (left + right) / 2
		if x < w[mid].jenis {
			right = mid - 1
		} else if x > w[mid].jenis {
			left = mid + 1
		} else {
			idx = mid
		}
	}

	if idx == -1 {
		fmt.Println("Data tidak ditemukan!")
	} else {
		fmt.Println("Data ditemukan (Hasil setelah data diurutkan berdasarkan alfabet):")
		fmt.Printf("| %-10s | %-9s | %-20s | %-7s | %-7s |\n", "Hari ke-", "Tanggal", "Latihan", "Durasi", "Kalori")
		
		// Menampilkan semua data yang sama di sekitar indeks yang ditemukan
		// Melangkah ke kiri untuk mencari batas awal olahraga yang sama
		start := idx
		for start > 0 && w[start-1].jenis == x {
			start--
		}
		// Cetak dari batas awal sampai namanya tidak sama lagi
		for i := start; i < n && w[i].jenis == x; i++ {
			fmt.Printf("| %-10d | %-9s | %-20s | %-7d | %-7d |\n", w[i].hariKe, w[i].tanggal, w[i].jenis, w[i].durasi, w[i].kalori)
		}
	}
}

// Fungsi standar menampilkan seluruh data saat ini
func showData(w workout, n int) {
	if n == 0 {
		fmt.Println("Belum ada data workout.")
		return
	}
	fmt.Printf("| %-10s | %-9s | %-20s | %-7s | %-7s |\n", "Hari ke-", "Tanggal", "Latihan", "Durasi", "Kalori")
	for i := 0; i < n; i++ {
		fmt.Printf("| %-10d | %-9s | %-20s | %-7d | %-7d |\n", w[i].hariKe, w[i].tanggal, w[i].jenis, w[i].durasi, w[i].kalori)
	}
}

func editData(w *workout, n int) {
	var targetHari, pilihan, edit int
	var idx int = -1
	
	if n == 0 {
		fmt.Println("Belum ada data untuk diedit.")
		return
	}
	
	showData(*w, n)
	fmt.Print("Pilih data Hari ke- berapa yang ingin diedit: ")
	fmt.Scan(&targetHari)
	
	// Cari indeks yang memiliki hariKe sesuai input
	for i := 0; i < n; i++ {
		if w[i].hariKe == targetHari {
			idx = i
			break
		}
	}
	
	if idx == -1 {
		fmt.Println("Hari tidak ditemukan!")
		return
	}

	fmt.Printf("| %-10s | %-9s | %-20s | %-7s | %-7s |\n", "Hari ke-", "Tanggal", "Latihan", "Durasi", "Kalori")
	fmt.Printf("| %-10d | %-9s | %-20s | %-7d | %-7d |\n", w[idx].hariKe, w[idx].tanggal, w[idx].jenis, w[idx].durasi, w[idx].kalori)
	
	fmt.Print("Pilih 1 untuk edit keseluruhan, 0 untuk edit sebagian: ")
	fmt.Scan(&edit)
	if edit == 1 {
		fmt.Print("Tanggal baru (DDMMYYYY): ")
		fmt.Scan(&w[idx].tanggal)
		fmt.Print("Jenis olahraga baru: ")
		fmt.Scan(&w[idx].jenis)
		fmt.Print("Durasi baru: ")
		fmt.Scan(&w[idx].durasi)
		fmt.Print("Kalori baru: ")
		fmt.Scan(&w[idx].kalori)
	} else if edit == 0 {
		fmt.Println("Pilih bagian yang ingin diedit: ")
		fmt.Println("1. Ubah tanggal")
		fmt.Println("2. Ubah jenis olahraga")
		fmt.Println("3. Ubah durasi")
		fmt.Println("4. Ubah total kalori")
		fmt.Println("5. Batalkan edit")
		fmt.Print("Pilih menu 1-5: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			fmt.Print("Masukkan tanggal yang baru (DDMMYYYY): ")
			fmt.Scan(&w[idx].tanggal)
		case 2:
			fmt.Print("Masukkan jenis olahraga yang baru: ")
			fmt.Scan(&w[idx].jenis)
		case 3:
			fmt.Print("Masukkan durasi yang baru: ")
			fmt.Scan(&w[idx].durasi)
		case 4:
			fmt.Print("Masukkan jumlah kalori yang baru: ")
			fmt.Scan(&w[idx].kalori)
		}
	}
}

func hapusData(w *workout, n *int) {
	var targetHari int
	var idx int = -1
	
	if *n == 0 {
		fmt.Println("Belum ada data untuk dihapus.")
		return
	}
	
	showData(*w, *n)
	fmt.Print("Pilih data Hari ke- berapa yang ingin dihapus: ")
	fmt.Scan(&targetHari)
	
	for i := 0; i < *n; i++ {
		if w[i].hariKe == targetHari {
			idx = i
			break
		}
	}
	
	if idx == -1 {
		fmt.Println("Hari tidak ditemukan!")
		return
	}

	for i := idx; i < *n-1; i++ {
		w[i] = w[i+1]
	}
	*n -= 1
	fmt.Println("Data berhasil dihapus!")
}

// Poin B: Memberi rekomendasi workout berdasarkan pola latihan sebelumnya
func rekomendasi(w workout, n int) {
	var count int = 1
	if n == 0 {
		fmt.Println("Belum ada riwayat olahraga. Silakan tambahkan data terlebih dahulu.")
		return
	}
	
	fmt.Println("Rekomendasi workout berdasarkan pola 3 latihan terakhir Anda: ")
	if n >= 3 {
		for i := n - 3; i < n; i++ {
			fmt.Printf("%d. %s\n", count, w[i].jenis)
			count++
		}
	} else {
		for i := 0; i < n; i++ {
			fmt.Printf("%d. %s\n", count, w[i].jenis)
			count++
		}
	}
}

// Poin D: Selection Sort secara Descending (Besar ke Kecil) berdasarkan Kalori
func sortKalori(A workout, n int) {
	if n == 0 {
		fmt.Println("Belum ada data untuk diurutkan.")
		return
	}
	
	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if A[j].kalori > A[idxMax].kalori {
				idxMax = j
			}
		}
		// Tukar posisi
		temp := A[i]
		A[i] = A[idxMax]
		A[idxMax] = temp
	}
	fmt.Println("\nData terurut dari kalori terbesar sampai terendah:")
	showData(A, n)
}

// Poin D: Insertion Sort secara Ascending (Kecil ke Besar) berdasarkan Durasi
func sortDurasi(A workout, n int) {
	if n == 0 {
		fmt.Println("Belum ada data untuk diurutkan.")
		return
	}
	
	for pass := 1; pass < n; pass++ {
		i := pass
		temp := A[pass]
		for i > 0 && temp.durasi < A[i-1].durasi {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
	}
	fmt.Println("\nData terurut dari durasi terpendek hingga terlama:")
	showData(A, n)
}

// Poin E: Fitur Menu Laporan Spesifik
func menuLaporan(w workout, n int) {
	var opsi int
	if n == 0 {
		fmt.Println("Belum ada data aktivitas untuk membuat laporan.")
		return
	}
	
	fmt.Println("\n==========================")
	fmt.Println("      MENU LAPORAN        ")
	fmt.Println("==========================")
	fmt.Println("1. Tampilkan 10 Aktivitas Terakhir")
	fmt.Println("2. Total Kalori dalam Periode Tertentu")
	fmt.Print("Pilih opsi laporan (1-2): ")
	fmt.Scan(&opsi)
	
	switch opsi {
	case 1:
		// Menampilkan maksimal 10 aktivitas terakhir yang dilakukan pengguna
		fmt.Println("\n--- 10 Aktivitas Terakhir ---")
		fmt.Printf("| %-10s | %-9s | %-20s | %-7s | %-7s |\n", "Hari ke-", "Tanggal", "Latihan", "Durasi", "Kalori")
		start := 0
		if n > 10 {
			start = n - 10
		}
		for i := start; i < n; i++ {
			fmt.Printf("| %-10d | %-9s | %-20s | %-7d | %-7d |\n", w[i].hariKe, w[i].tanggal, w[i].jenis, w[i].durasi, w[i].kalori)
		}
		
	case 2:
		// Menghitung total kalori pada rentang tanggal/periode tertentu
		var tglMulai, tglSelesai string
		var totalKalori int = 0
		var found bool = false
		
		fmt.Print("Masukkan tanggal mulai (DDMMYYYY): ")
		fmt.Scan(&tglMulai)
		fmt.Print("Masukkan tanggal selesai (DDMMYYYY): ")
		fmt.Scan(&tglSelesai)
		
		for i := 0; i < n; i++ {
			// Melakukan pengecekan sederhana apakah tanggal data berada di dalam rentang input pengguna
			// Catatan: Karena format string DDMMYYYY, pencocokan idealnya dilakukan berurutan berdasarkan indeks data yang terinput kronologis
			if w[i].tanggal >= tglMulai && w[i].tanggal <= tglSelesai {
				totalKalori += w[i].kalori
				found = true
			}
		}
		
		if found {
			fmt.Printf("Total kalori yang terbakar dari periode %s s/d %s adalah: %d kalori\n", tglMulai, tglSelesai, totalKalori)
		} else {
			fmt.Println("Tidak ada aktivitas ditemukan pada periode tanggal tersebut.")
		}
	}
	fmt.Println()
}