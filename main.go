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
	var s1 Skripsi
	fmt.Scanln()
	fmt.Println("+===========================+")
	fmt.Printf("%-5s%-23s%-5s\n", "+", "Tambah Data Skripsi", "+")
	fmt.Println("+===========================+")
	fmt.Printf("Jumlah Skripsi yang akan ditambahkan: ")
	fmt.Scanln(n)
	for i := 0; i < *n; i++ {
		fmt.Printf("Skripsi %d\n", i+1)
		s1.ID = i + 1
		fmt.Printf("Tahun: ")
		fmt.Scanf("%d\n", &s1.Year)
		fmt.Printf("Judul: ")
		fmt.Scanf("%s\n", &s1.Title)
		fmt.Printf("NIM Mahasiswa: ")
		fmt.Scanf("%s\n", &s1.Author.NIM)
		fmt.Printf("Nama Mahasiswa: ")
		fmt.Scanf("%s\n", &s1.Author.Name)
		fmt.Printf("Dosen Pembimbing: ")
		fmt.Scanf("%s\n", &s1.Author.DosBing)
		fmt.Printf("Lulus (true/false): ")
		fmt.Scanf("%t\n", &s1.Author.IsGraduated)
		s[i] = s1
		m[i] = s1.Author
	}
	fmt.Println("➤ Data Berhasil Ditambahkan!")
}

// Fungsi untuk mencari id skripsi berdasarkan nama mahasiswa atau Judul Penelitian menggunakan sequential search
func getSkripsiId(s SkripsiList, n int, keyword string) int {
	for i := 0; i < n; i++ {
		if s[i].Title == keyword || s[i].Author.Name == keyword {
			return s[i].ID
		}
	}
	return -1
}

// Prosedur untuk mengupdate data skripsi berdasarkan ID
func updateSkripsi(s *SkripsiList, m *MahasiswaList, id int) {
	var choice, idx, i int
	idx = -1
	i = 0
	fmt.Println("+===========================+")
	fmt.Printf("%-5s%-23s%-5s\n", "+", "Edit Data Skripsi", "+")
	fmt.Println("+===========================+")

	for i < len(s) {
		if (*s)[i].ID == id {
			idx = i
			break
		}
		i++
	}

	fmt.Println("Pilih Menu")
	fmt.Println("1. Judul Skripsi")
	fmt.Println("2. Tahun Skripsi")
	fmt.Println("3. Penulis Skripsi")
	fmt.Println("4. Dosen Pembimbing Skripsi")
	fmt.Println("5. Status Kelulusan")
	fmt.Println("6. Kembali")
	fmt.Print("Masukkan pilihan: ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Printf("Masukkan Judul Skripsi baru: ")
		fmt.Scan(&s[idx].Title)
	case 2:
		fmt.Printf("Masukkan Tahun Skripsi baru: ")
		fmt.Scan(&s[idx].Year)
	case 3:
		fmt.Printf("Masukkan NIM Mahasiswa baru: ")
		fmt.Scan(&s[idx].Author.NIM)
		fmt.Printf("Masukkan Nama Mahasiswa baru: ")
		fmt.Scan(&s[idx].Author.Name)
	case 4:
		fmt.Printf("Masukkan Dosen Pembimbing baru: ")
		fmt.Scan(&s[idx].Author.DosBing)
	case 5:
		fmt.Printf("Masukkan Status Kelulusan baru (true/false): ")
		fmt.Scan(&s[idx].Author.IsGraduated)
		for j := 0; j < len(m); j++ {
			if m[j].NIM == s[idx].Author.NIM {
				m[j].IsGraduated = s[idx].Author.IsGraduated
			}
		}
	case 6:
		return
	default:
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&choice)
	}
	fmt.Println("➤ Data Berhasil Diperbarui!")
}

// Prosedur untuk menghapus data skripsi berdasarkan ID
func deleteSkripsi(s *SkripsiList, n *int, id int) {
	var i, idx int
	idx = -1
	fmt.Println("+===========================+")
	fmt.Printf("%-5s%-23s%-5s\n", "+", "Hapus Data Skripsi", "+")
	fmt.Println("+===========================+")
	for i < *n {
		if (*s)[i].ID == id {
			(*s)[idx] = (*s)[i]
			idx++
			*n--
		}
		i++
	}
	fmt.Println("➤ Data Berhasil Dihapus!")
}

// Prosedur untuk menampilkan semua data skripsi
func getAllSkripsi(s SkripsiList, n int) {
}

// Prosedur untuk menampilkan data skripsi berdasarkan ID
func getSkripsiById(s SkripsiList, n int, id int) {
}

// Prosedur untuk mencari id skripsi berdasarkan nama mahasiswa atau Judul Penelitian menggunakan sequential search
func findSkripsiSequential(s SkripsiList, n int, keyword string) {
}

// Prosedur untuk mencari id skripsi berdasarkan nama mahasiswa atau Judul Penelitian menggunakan binary search
func findSkripsiBinary(s SkripsiList, n int, keyword string) {
}

// Prosedur untuk mengurutkan skripsi berdasarkan nama/tahun menggunakan selection sort
func sortSkripsiSelection(s *SkripsiList, n int, sortBy string, sortType string) {
}

// Prosedur untuk mengurutkan skripsi berdasarkan nama/tahun menggunakan insertion sort
func sortSkripsiInsertion(s *SkripsiList, n int, sortBy string, sortType string) {
}

// Prosedur untuk menampilkan statistik skripsi, seperti jumlah skripsi per tahun, jumlah skripsi yang lulus, dll
func statisticSkripsi(s SkripsiList, n int) {
}

// Menu utama untuk menjalankan program
func main() {
	var mainChoice, subChoice, id, n int
	var keyword, sortBy, sortType string
	var s SkripsiList
	var m MahasiswaList

	for {
		fmt.Println("+================================================================+")
		fmt.Printf("%-5s%-60s%-5s\n", "+", "SkripsIn - Sistem Informasi Inventaris Dokumen Skripsi", "+")
		fmt.Println("+================================================================+")
		fmt.Println("Pilih Menu")
		fmt.Println("1. Kelola Data Skripsi")
		fmt.Println("2. Pencarian Skripsi (Berdasarkan Nama Mahasiswa atau Judul Penelitian)")
		fmt.Println("3. Ambil data Skripsi Berdasarkan ID")
		fmt.Println("4. Pengurutan Skripsi")
		fmt.Println("5. Statistik Skripsi")
		fmt.Println("6. Tampilkan Semua Data Skripsi")
		fmt.Println("7. Keluar")
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&mainChoice)
		if mainChoice == 7 {
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
				fmt.Print("Masukkan judul penelitian/nama mahasiswa yang akan di-update: ")
				fmt.Scan(&keyword)
				id = getSkripsiId(s, n, keyword)
				if id == -1 {
					fmt.Println("Skripsi tidak ditemukan")
					continue
				}
				updateSkripsi(&s, &m, id)
			case 3:
				fmt.Print("Masukkan judul penelitian/nama mahasiswa yang akan dihapus: ")
				fmt.Scan(&keyword)
				id = getSkripsiId(s, n, keyword)
				if id == -1 {
					fmt.Println("Skripsi tidak ditemukan")
					continue
				}
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
				sortSkripsiInsertion(&s, n, "year", "asc")
				findSkripsiBinary(s, n, keyword)
			default:
				return
			}
		case 3:
			fmt.Println("+===========================+")
			fmt.Printf("%-5s%-23s%-5s\n", "+", "Ambil Skripsi Berdasarkan ID", "+")
			fmt.Println("+===========================+")
			fmt.Print("Masukkan id skripsi yang akan diambil: ")
			fmt.Scan(&id)
			getSkripsiById(s, n, id)
		case 4:
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
				fmt.Print("Masukkan tipe pengurutan (title/year): ")
				fmt.Scan(&sortBy)
				fmt.Print("Masukkan tipe pengurutan (asc/desc): ")
				fmt.Scan(&sortType)
				sortSkripsiSelection(&s, n, sortBy, sortType)
			case 2:
				fmt.Print("Masukkan tipe pengurutan (title/year): ")
				fmt.Scan(&sortBy)
				fmt.Print("Masukkan tipe pengurutan (asc/desc): ")
				fmt.Scan(&sortType)
				sortSkripsiInsertion(&s, n, sortBy, sortType)
			default:
				return
			}
		case 5:
			fmt.Println("+===========================+")
			fmt.Printf("%-5s%-23s%-5s\n", "+", "Statistik Skripsi", "+")
			fmt.Println("+===========================+")
			statisticSkripsi(s, n)
			continue
		case 6:
			getAllSkripsi(s, n)
			continue
		case 7:
			return
		default:
			fmt.Print("Masukkan pilihan: ")
			fmt.Scan(&mainChoice)
		}
	}
}
