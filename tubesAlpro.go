package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NMAX int = 100

type workout struct {
	namaLatihan string
	jenis       string
	durasi      int
	kalori      int
	jadwal      string
}

type tabWorkout [NMAX]workout

var data tabWorkout

func main() {
	var n, pilih int

	dummyData(&n)

	for {
		menu()
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahData(&data, &n)
		case 2:
			tampilData(data, n)
		case 3:
			editData(&data, n)
		case 4:
			hapusData(&data, &n)
		case 5:
			menuSearch(&data, n)
		case 6:
			menuSort(&data, n)
		case 7:
			laporan(data, n)
		case 8:
			rekomendasi(data, n)
		case 9:
			fmt.Println("Program selesai")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func menu() {
	fmt.Println("\n===== APLIKASI WORKOUT =====")
	fmt.Println("1. Tambah Workout")
	fmt.Println("2. Tampilkan Workout")
	fmt.Println("3. Edit Workout")
	fmt.Println("4. Hapus Workout")
	fmt.Println("5. Cari Workout")
	fmt.Println("6. Sorting Workout")
	fmt.Println("7. Laporan")
	fmt.Println("8. Rekomendasi Workout")
	fmt.Println("9. Keluar")
	fmt.Print("Pilih menu: ")
}

func tambahData(A *tabWorkout, n *int) {
	var latihan, jenis, jadwal string
	var durasi, kalori int

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nama latihan: ")
	reader.ReadString('\n')
	latihan, _ = reader.ReadString('\n')
	latihan = strings.TrimSpace(latihan)

	fmt.Print("Jenis olahraga: ")
	jenis, _ = reader.ReadString('\n')
	jenis = strings.TrimSpace(jenis)

	fmt.Print("Durasi (menit): ")
	fmt.Scan(&durasi)

	fmt.Print("Kalori terbakar: ")
	fmt.Scan(&kalori)

	reader.ReadString('\n')

	fmt.Print("Jadwal latihan: ")
	jadwal, _ = reader.ReadString('\n')
	jadwal = strings.TrimSpace(jadwal)

	A[*n].namaLatihan = latihan
	A[*n].jenis = jenis
	A[*n].durasi = durasi
	A[*n].kalori = kalori
	A[*n].jadwal = jadwal

	*n++

	fmt.Println("Data berhasil ditambahkan")
}

func tampilData(A tabWorkout, n int) {
	var i int

	fmt.Println("\n===== DATA WORKOUT =====")
	fmt.Printf("%-3s %-15s %-15s %-10s %-10s %-10s\n",
		"No", "Latihan", "Jenis", "Durasi", "Kalori", "Jadwal")

	for i = 0; i < n; i++ {
		fmt.Printf("%-3d %-15s %-15s %-10d %-10d %-10s\n",
			i+1,
			A[i].namaLatihan,
			A[i].jenis,
			A[i].durasi,
			A[i].kalori,
			A[i].jadwal)
	}
}

func editData(A *tabWorkout, n int) {
	var idx int
	var latihan, jenis, jadwal string
	var durasi, kalori int

	reader := bufio.NewReader(os.Stdin)

	tampilData(*A, n)

	fmt.Print("Edit data ke: ")
	fmt.Scan(&idx)

	if idx >= 1 && idx <= n {

		reader.ReadString('\n')

		fmt.Print("Nama latihan baru: ")
		latihan, _ = reader.ReadString('\n')
		latihan = strings.TrimSpace(latihan)

		fmt.Print("Jenis olahraga baru: ")
		jenis, _ = reader.ReadString('\n')
		jenis = strings.TrimSpace(jenis)

		fmt.Print("Durasi baru: ")
		fmt.Scan(&durasi)

		fmt.Print("Kalori baru: ")
		fmt.Scan(&kalori)

		reader.ReadString('\n')

		fmt.Print("Jadwal baru: ")
		jadwal, _ = reader.ReadString('\n')
		jadwal = strings.TrimSpace(jadwal)

		A[idx-1].namaLatihan = latihan
		A[idx-1].jenis = jenis
		A[idx-1].durasi = durasi
		A[idx-1].kalori = kalori
		A[idx-1].jadwal = jadwal

		fmt.Println("Data berhasil diubah")

	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func hapusData(A *tabWorkout, n *int) {
	var idx, i int

	tampilData(*A, *n)

	fmt.Print("Hapus data ke: ")
	fmt.Scan(&idx)

	if idx >= 1 && idx <= *n {

		for i = idx - 1; i < *n-1; i++ {
			A[i] = A[i+1]
		}

		*n--

		fmt.Println("Data berhasil dihapus")

	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

func sequentialSearch(A tabWorkout, n int, key string) {
	var i int
	var found bool

	found = false
	key = strings.ToLower(key)

	for i = 0; i < n; i++ {
		if strings.ToLower(A[i].jenis) == key {
			fmt.Println("Data ditemukan:")
			fmt.Println(A[i].namaLatihan, "-", A[i].jenis)
			found = true
		}
	}

	if !found {
		fmt.Println("Data tidak ditemukan")
	}
}

func selectionSortAsc(A *tabWorkout, n int) {
	var pass, i, idx int
	var temp workout

	pass = 1

	for pass < n {
		idx = pass - 1
		i = pass

		for i < n {
			if A[i].durasi < A[idx].durasi {
				idx = i
			}
			i++
		}

		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp

		pass++
	}
}

func selectionSortDesc(A *tabWorkout, n int) {
	var pass, i, idx int
	var temp workout

	pass = 1

	for pass < n {
		idx = pass - 1
		i = pass

		for i < n {
			if A[i].durasi > A[idx].durasi {
				idx = i
			}
			i++
		}

		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp

		pass++
	}
}

func insertionSortAsc(A *tabWorkout, n int) {
	var pass, i int
	var temp workout

	pass = 1

	for pass <= n-1 {
		i = pass
		temp = A[pass]

		for i > 0 && temp.kalori < A[i-1].kalori {
			A[i] = A[i-1]
			i--
		}

		A[i] = temp
		pass++
	}
}

func insertionSortDesc(A *tabWorkout, n int) {
	var pass, i int
	var temp workout

	pass = 1

	for pass <= n-1 {
		i = pass
		temp = A[pass]

		for i > 0 && temp.kalori > A[i-1].kalori {
			A[i] = A[i-1]
			i--
		}

		A[i] = temp
		pass++
	}
}

func binarySearch(A tabWorkout, n int, x int) {
	var left, right, mid int
	var found bool

	left = 0
	right = n - 1
	found = false

	for left <= right {
		mid = (left + right) / 2

		if A[mid].kalori == x {
			found = true
			fmt.Println("Workout ditemukan:")
			fmt.Println(A[mid].namaLatihan, "-", A[mid].kalori, "kalori")
			left = right + 1
		} else if A[mid].kalori < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if !found {
		fmt.Println("Data tidak ditemukan")
	}
}

func menuSearch(A *tabWorkout, n int) {
	var pilih int
	var key string
	var kalori int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n===== MENU SEARCH =====")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	if pilih == 1 {

		reader.ReadString('\n')

		fmt.Print("Cari jenis olahraga: ")
		key, _ = reader.ReadString('\n')
		key = strings.TrimSpace(key)

		sequentialSearch(*A, n, key)

	} else if pilih == 2 {

		insertionSortAsc(A, n)

		fmt.Print("Cari jumlah kalori: ")
		fmt.Scan(&kalori)

		binarySearch(*A, n, kalori)

	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func menuSort(A *tabWorkout, n int) {
	var pilih int

	fmt.Println("\n===== MENU SORT =====")
	fmt.Println("1. Selection Sort Ascending")
	fmt.Println("2. Selection Sort Descending")
	fmt.Println("3. Insertion Sort Ascending")
	fmt.Println("4. Insertion Sort Descending")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		selectionSortAsc(A, n)
		tampilData(*A, n)

	} else if pilih == 2 {
		selectionSortDesc(A, n)
		tampilData(*A, n)

	} else if pilih == 3 {
		insertionSortAsc(A, n)
		tampilData(*A, n)

	} else if pilih == 4 {
		insertionSortDesc(A, n)
		tampilData(*A, n)

	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func laporan(A tabWorkout, n int) {
	var i, total int

	fmt.Println("\n===== 10 WORKOUT TERAKHIR =====")

	if n < 10 {
		for i = 0; i < n; i++ {
			fmt.Println(i+1, A[i].namaLatihan)
		}
	} else {
		for i = n - 10; i < n; i++ {
			fmt.Println(i+1, A[i].namaLatihan)
		}
	}

	for i = 0; i < n; i++ {
		total = total + A[i].kalori
	}

	fmt.Println("Total kalori terbakar:", total)
}

func rekomendasi(A tabWorkout, n int) {
	var i int
	var total int
	var rata int

	for i = 0; i < n; i++ {
		total = total + A[i].durasi
	}

	if n > 0 {
		rata = total / n

		if rata < 30 {
			fmt.Println("Rekomendasi: Cardio Ringan")
		} else if rata < 60 {
			fmt.Println("Rekomendasi: Jogging atau Cycling")
		} else {
			fmt.Println("Rekomendasi: HIIT Workout")
		}
	} else {
		fmt.Println("Belum ada data workout")
	}
}

func dummyData(n *int) {

	data[0].namaLatihan = "Push Up"
	data[0].jenis = "Strength"
	data[0].durasi = 30
	data[0].kalori = 200
	data[0].jadwal = "Senin"

	data[1].namaLatihan = "Running"
	data[1].jenis = "Cardio"
	data[1].durasi = 45
	data[1].kalori = 350
	data[1].jadwal = "Selasa"

	data[2].namaLatihan = "Cycling"
	data[2].jenis = "Cardio"
	data[2].durasi = 60
	data[2].kalori = 500
	data[2].jadwal = "Rabu"

	*n = 3
}