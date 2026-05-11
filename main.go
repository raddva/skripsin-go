package main

import "fmt"

const nMax = 100

// Struktur data untuk menyimpan informasi tentang mahasiswa
type Mahasiswa struct {
	NIM, Name, DosBing string
	IsGraduated        bool
}

// Struktur data untuk menyimpan informasi tentang skripsi
type Skripsi struct {
	ID, Year int
	Title    string
	Author   Mahasiswa
}

// Tipe data untuk menyimpan daftar skripsi dan daftar mahasiswa
type SkripsiList [nMax]Skripsi
type MahasiswaList [nMax]Mahasiswa

// Prosedur untuk menambah data skripsi beserta data mahasiswa yang terkait
func addSkripsi(s *SkripsiList, n *int, m *MahasiswaList) {
}

// Prosedur untuk mengupdate data skripsi berdasarkan ID
func updateSkripsi(s *SkripsiList, m *MahasiswaList, id int) {
}

// Prosedur untuk menghapus data skripsi berdasarkan ID
func deleteSkripsi(s *SkripsiList, n *int, id int) {
}

// Prosedur untuk menampilkan semua data skripsi
func getAllSkripsi(s SkripsiList, n int, m MahasiswaList) {
}

// Prosedur untuk menampilkan data skripsi berdasarkan ID
func getSkripsi(s SkripsiList, n int, id int) {
}

// Fungsi untuk mencari id skripsi berdasarkan nama mahasiswa atau Judul Penelitian menggunakan sequential search
func findSkripsiSequential(s SkripsiList, n int, keyword string) int {
	return 0
}

// Fungsi untuk mencari id skripsi berdasarkan nama mahasiswa atau Judul Penelitian menggunakan binary search
func findSkripsiBinary(s SkripsiList, n int, keyword string) int {
	return 0
}

// Prosedur untuk mengurutkan skripsi berdasarkan ID menggunakan selection sort
func sortSkripsiSelection(s *SkripsiList, n int, sortType string) {
}

// Prosedur untuk mengurutkan skripsi berdasarkan ID menggunakan insertion sort
func sortSkripsiInsertion(s *SkripsiList, n int, sortType string) {
}

// Prosedur untuk menampilkan statistik skripsi, seperti jumlah skripsi per tahun, jumlah skripsi yang lulus, dll
func statisticSkripsi(s SkripsiList, n int, m MahasiswaList) {
}

// Menu utama untuk menjalankan program
func main() {
	var mainChoice, subChoice, id, n int
	var keyword, sortType string
	var s SkripsiList
	var m MahasiswaList

	for {
		fmt.Println("+================================================================+")
		fmt.Printf("%-5s%-60s%-5s\n", "+", "SkripsIn - Sistem Informasi Inventaris Dokumen Skripsi", "+")
		fmt.Println("+================================================================+")
		fmt.Println("Pilih Menu")
		fmt.Println("1. Kelola Data Skripsi")
		fmt.Println("2. Pencarian Skripsi")
		fmt.Println("3. Pengurutan Skripsi")
		fmt.Println("4. Statistik Skripsi")
		fmt.Println("5. Keluar")

		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&mainChoice)
		if mainChoice == 5 {
			return
		}

		switch mainChoice {
		case 1:
			fmt.Println("+===========================+")
			fmt.Printf("%-5s%-23s%-5s\n", "+", "Kelola Data Skripsi", "+")
			fmt.Println("+===========================+")
			fmt.Println("0. Kembali ke Menu Utama")
			fmt.Println("1. Tambah Skripsi")
			fmt.Println("2. Update Skripsi")
			fmt.Println("3. Hapus Skripsi")
			fmt.Print("Masukkan pilihan: ")
			fmt.Scan(&subChoice)
			if subChoice == 0 {
				continue
			}
			switch subChoice {
			case 1:
				addSkripsi(&s, &n, &m)
			case 2:
				fmt.Print("Masukkan keyword pencarian: ")
				fmt.Scan(&keyword)
				id = findSkripsiSequential(s, n, keyword)
				updateSkripsi(&s, &m, id)
			case 3:
				fmt.Print("Masukkan keyword pencarian: ")
				fmt.Scan(&keyword)
				id = findSkripsiSequential(s, n, keyword)
				deleteSkripsi(&s, &n, id)
			default:
				return
			}
		case 2:
			fmt.Println("+===========================+")
			fmt.Printf("%-5s%-23s%-5s\n", "+", "Cari Skripsi", "+")
			fmt.Println("+===========================+")
			fmt.Println("0. Kembali ke Menu Utama")
			fmt.Println("1. Pencarian Sequential")
			fmt.Println("2. Pencarian Binary")
			fmt.Print("Masukkan pilihan: ")
			fmt.Scan(&subChoice)
			if subChoice == 0 {
				continue
			}
			switch subChoice {
			case 1:
				fmt.Print("Masukkan keyword pencarian: ")
				fmt.Scan(&keyword)
				findSkripsiSequential(s, n, keyword)
			case 2:
				fmt.Print("Masukkan keyword pencarian: ")
				fmt.Scan(&keyword)
				findSkripsiBinary(s, n, keyword)
			default:
				return
			}
		case 3:
			fmt.Println("+===========================+")
			fmt.Printf("%-5s%-23s%-5s\n", "+", "Urutkan Skripsi", "+")
			fmt.Println("+===========================+")
			fmt.Println("0. Kembali ke Menu Utama")
			fmt.Println("1. Selection Sort")
			fmt.Println("2. Insertion Sort")
			fmt.Print("Masukkan pilihan: ")
			fmt.Scan(&subChoice)
			if subChoice == 0 {
				continue
			}
			switch subChoice {
			case 1:
				fmt.Print("Masukkan tipe pengurutan (asc/desc): ")
				fmt.Scan(&sortType)
				sortSkripsiSelection(&s, n, sortType)
			case 2:
				fmt.Print("Masukkan tipe pengurutan (asc/desc): ")
				fmt.Scan(&sortType)
				sortSkripsiInsertion(&s, n, sortType)
			default:
				return
			}
		case 4:
			fmt.Println("+===========================+")
			fmt.Printf("%-5s%-23s%-5s\n", "+", "Statistik Skripsi", "+")
			fmt.Println("+===========================+")
			statisticSkripsi(s, n, m)
			continue
		case 5:
			return
		default:
			fmt.Print("Masukkan pilihan: ")
			fmt.Scan(&mainChoice)
		}
	}
}
