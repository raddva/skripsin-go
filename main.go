package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Global Variables
const nMax = 100

var reader = bufio.NewReader(os.Stdin) // temporary use, pengganti fmt kocak gabisa input string spasi euy

// End of Global Variables

// Struktur Data
// Struktur data untuk menyimpan informasi tentang mahasiswa
type Mahasiswa struct {
	NIM, Name, DosBing string
	IsGraduated        bool
}

// Struktur data untuk menyimpan informasi tentang skripsi
type Skripsi struct {
	Year   int
	Title  string
	Author Mahasiswa
}

// End of Struktur Data

// Tipe data untuk menyimpan daftar skripsi dan daftar mahasiswa
type FoundList [nMax]Skripsi
type SkripsiList [nMax]Skripsi
type MahasiswaList [nMax]Mahasiswa

// Helper Functions
func isNumber(s string) bool {
	if s == "" {
		return false
	}

	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func isAlpha(s string) bool {
	if s == "" {
		return false
	}

	for _, c := range s {
		if (c < 'A' || c > 'Z') &&
			(c < 'a' || c > 'z') &&
			c != ' ' {
			return false
		}
	}
	return true
}

func isValidYear(s string) bool {
	var year int
	if !isNumber(s) {
		return false
	}

	fmt.Sscanf(s, "%d", &year)
	if year < 1900 || year > 2026 { // batasi karena sekarang baru 2026 :v
		return false
	}
	return true
}

func contains(text, keyword string) bool {
	nText := len(text)
	nKey := len(keyword)
	if nKey > nText {
		return false
	}

	for i := 0; i <= nText-nKey; i++ {
		match := true
		for j := 0; j < nKey; j++ {
			if text[i+j] != keyword[j] {
				match = false
			}
		}
		if match {
			return true
		}
	}
	return false
}

// End of Helper Functions

// Start Subprograms
// Prosedur untuk menambah data skripsi beserta data mahasiswa yang terkait
func addSkripsi(s *SkripsiList, n *int, m *MahasiswaList) {
	var s1 Skripsi
	var addition int
	var input string
	fmt.Println()
	fmt.Println("✦===========================✦")
	fmt.Printf("%-5s%-23s%-5s\n", "✦", "Tambah Data Skripsi", "✦")
	fmt.Println("✦===========================✦")
	for {
		fmt.Print("Jumlah Skripsi yang akan ditambahkan: ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if isNumber(input) {
			fmt.Sscanf(input, "%d", &addition)
			break
		} else {
			fmt.Println("✗ Input harus angka!")
		}
	}

	for i := *n; i < *n+addition; i++ {
		fmt.Printf("\nSkripsi %d\n", i+1)
		for {
			fmt.Print("Tahun: ")
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if isValidYear(input) {
				fmt.Sscanf(input, "%d", &s1.Year)
				break
			}
			fmt.Println("✗ Tahun harus valid!")
		}
		fmt.Print("Judul: ")
		input, _ = reader.ReadString('\n')
		s1.Title = strings.TrimSpace(input)
		for {
			fmt.Print("NIM Mahasiswa: ")
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if isNumber(input) {
				s1.Author.NIM = input
				break
			}
			fmt.Println("✗ NIM hanya boleh angka!")
		}
		for {
			fmt.Print("Nama Mahasiswa: ")
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if isAlpha(input) {
				s1.Author.Name = input
				break
			}
			fmt.Println("✗ Nama hanya boleh huruf!")
		}

		for {
			fmt.Print("Dosen Pembimbing: ")
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if isAlpha(input) {
				s1.Author.DosBing = input
				break
			}
			fmt.Println("✗ Nama dosen hanya boleh huruf!")
		}
		s1.Author.IsGraduated = false
		s[i] = s1
		m[i] = s1.Author
	}
	*n += addition
	fmt.Println("\n➤ Data Berhasil Ditambahkan!")
}

// Fungsi untuk mencari Index skripsi berdasarkan nama mahasiswa atau Judul Penelitian menggunakan sequential search
func getSkripsiIdx(s SkripsiList, n int, keyword string) int {
	for i := 0; i < n; i++ {
		if contains(s[i].Title, keyword) ||
			contains(s[i].Author.Name, keyword) {
			return i
		}
	}
	return -1
}

// Prosedur untuk mengupdate data skripsi berdasarkan Index
func updateSkripsi(s *SkripsiList, m *MahasiswaList, idx int) {
	var choice int
	var input string
	fmt.Println("✦===========================✦")
	fmt.Printf("%-5s%-23s%-5s\n", "✦", "Edit Data Skripsi", "✦")
	fmt.Println("✦===========================✦")
	fmt.Println("1. Judul Skripsi")
	fmt.Println("2. Tahun Skripsi")
	fmt.Println("3. Penulis Skripsi")
	fmt.Println("4. Dosen Pembimbing Skripsi")
	fmt.Println("5. Status Kelulusan")
	fmt.Println("6. Kembali")
	fmt.Print("⟡ ݁₊ .Masukkan pilihan: ")
	for {
		fmt.Print("⟡ ݁₊ .Masukkan pilihan: ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if isNumber(input) {
			fmt.Sscanf(input, "%d", &choice)
			break
		}

		fmt.Println("✗ Pilihan harus angka!")
	}

	switch choice {
	case 1:
		fmt.Println("⛧ Judul Sebelumnya: ", s[idx].Title)

		fmt.Printf("⏾ Masukkan Judul Skripsi baru: ")
		fmt.Scan(&s[idx].Title)
	case 2:
		fmt.Println("⛧ Tahun Sebelumnya: ", s[idx].Year)
		for {
			fmt.Printf(" Masukkan Tahun Skripsi baru: ")
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if isValidYear(input) {
				fmt.Sscanf(input, "%d", &s[idx].Year)
				break
			}
			fmt.Println("✗ Tahun harus valid!")
		}
	case 3:
		fmt.Printf("⏾ Masukkan NIM Mahasiswa baru: ")
		fmt.Scan(&s[idx].Author.NIM)
		fmt.Printf("⏾ Masukkan Nama Mahasiswa baru: ")
		fmt.Scan(&s[idx].Author.Name)
	case 4:
		fmt.Printf("⏾ Masukkan Dosen Pembimbing baru: ")
		fmt.Scan(&s[idx].Author.DosBing)
	case 5:
		fmt.Printf("⏾ Masukkan Status Kelulusan baru (true/false): ")
		fmt.Scan(&s[idx].Author.IsGraduated)
		for j := 0; j < len(m); j++ {
			if m[j].NIM == s[idx].Author.NIM {
				m[j].IsGraduated = s[idx].Author.IsGraduated
			}
		}
	case 6:
		return
	default:
		for {
			fmt.Print("⟡ ݁₊ .Masukkan pilihan: ")
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if isNumber(input) {
				fmt.Sscanf(input, "%d", &choice)
				break
			}
			fmt.Println("✗ Pilihan harus angka!")
		}
	}
	fmt.Println("➤ Data Berhasil Diperbarui!")
}

// Prosedur untuk menghapus data skripsi berdasarkan Index
func deleteSkripsi(s *SkripsiList, n *int, idx int) {
	if *n == 0 {
		fmt.Println("✦===========================✦")
		fmt.Printf("%-5s%-25s%-5s\n", "✦", "Data Skripsi Masih Kosong!", "✦")
		fmt.Println("✦===========================✦")
		return
	}

	fmt.Println("✦===========================✦")
	fmt.Printf("%-5s%-23s%-5s\n", "✦", "Hapus Data Skripsi", "✦")
	fmt.Println("✦===========================✦")
	for i := idx; i < *n-1; i++ {
		(*s)[i] = (*s)[i+1]
	}
	*n--
	fmt.Println("➤ Data Berhasil Dihapus!")
}

// Prosedur untuk menampilkan semua data skripsi
func getAllSkripsi(s SkripsiList, n int) {
}

// Prosedur untuk mencari id skripsi berdasarkan nama mahasiswa atau Judul Penelitian menggunakan sequential search
func findSkripsiSequential(s SkripsiList, n int, keyword string) {
	var f FoundList
}

// Prosedur untuk mencari id skripsi berdasarkan nama mahasiswa atau Judul Penelitian menggunakan binary search
func findSkripsiBinary(s SkripsiList, n int, keyword string) {
	var f FoundList
}

// Prosedur untuk mengurutkan skripsi berdasarkan nama/tahun menggunakan selection sort
func sortSkripsiSelection(s *SkripsiList, n int, sortBy string, sortType string) {
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if sortType == "asc" {
				if sortBy == "title" {
					if (*s)[j].Title < (*s)[idx].Title {
						idx = j
					}
				} else {
					if (*s)[j].Year < (*s)[idx].Year {
						idx = j
					}
				}
			} else {
				if sortBy == "title" {
					if (*s)[j].Title > (*s)[idx].Title {
						idx = j
					}
				} else {
					if (*s)[j].Year > (*s)[idx].Year {
						idx = j
					}
				}
			}
		}
		(*s)[i], (*s)[idx] = (*s)[idx], (*s)[i]
	}
}

// Prosedur untuk mengurutkan skripsi berdasarkan nama/tahun menggunakan insertion sort
func sortSkripsiInsertion(s *SkripsiList, n int, sortBy string, sortType string) {
	for i := 0; i < n; i++ {
		key := (*s)[i]
		j := i - 1
		if sortType == "asc" {
			if sortBy == "title" {
				for j >= 0 && (*s)[j].Title > key.Title {
					(*s)[j+1] = (*s)[j]
					j--
				}
			} else {
				for j >= 0 && (*s)[j].Year > key.Year {
					(*s)[j+1] = (*s)[j]
					j--
				}
			}
		} else {
			if sortBy == "title" {
				for j >= 0 && (*s)[j].Title < key.Title {
					(*s)[j+1] = (*s)[j]
					j--
				}
			} else {
				for j >= 0 && (*s)[j].Year < key.Year {
					(*s)[j+1] = (*s)[j]
					j--
				}
			}
		}
	}
}

// Prosedur untuk menampilkan statistik skripsi, seperti jumlah skripsi per tahun, jumlah skripsi yang lulus, dll
func statisticSkripsi(s SkripsiList, n int) {
}

// End Subprograms

// Menu utama untuk menjalankan program
func main() {
	var mainChoice, subChoice, idx, n int
	var keyword, sortBy, sortType, input string
	var s SkripsiList
	var m MahasiswaList

	for {
		fmt.Println("✦================================================================✦")
		fmt.Printf("%-5s%-60s%-5s\n", "✦", "SkripsIn - Sistem Informasi Inventaris Dokumen Skripsi", "✦")
		fmt.Println("✦================================================================✦")
		fmt.Println("1. Kelola Data Skripsi")
		fmt.Println("2. Pencarian Skripsi (Berdasarkan Nama Mahasiswa atau Judul Penelitian)")
		fmt.Println("3. Pengurutan Skripsi")
		fmt.Println("4. Statistik Skripsi")
		fmt.Println("5. Tampilkan Semua Data Skripsi")
		fmt.Println("6. Keluar")
		for {
			fmt.Print("⟡ ݁₊ .Masukkan pilihan: ")
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if isNumber(input) {
				fmt.Sscanf(input, "%d", &mainChoice)
				break
			}

			fmt.Println("✗ Pilihan harus angka!")
		}
		if mainChoice == 6 {
			return
		}

		switch mainChoice {
		case 1:
			fmt.Println("✦===========================✦")
			fmt.Printf("%-5s%-23s%-5s\n", "✦", "Kelola Data Skripsi", "✦")
			fmt.Println("✦===========================✦")
			fmt.Println("0. Kembali ke Menu Utama")
			fmt.Println("1. Tambah Skripsi")
			fmt.Println("2. Update Skripsi")
			fmt.Println("3. Hapus Skripsi")
			for {
				fmt.Print("⟡ ݁₊ .Masukkan pilihan: ")
				input, _ = reader.ReadString('\n')
				input = strings.TrimSpace(input)
				if isNumber(input) {
					fmt.Sscanf(input, "%d", &subChoice)
					break
				}

				fmt.Println("✗ Pilihan harus angka!")
			}
			if subChoice == 0 {
				continue
			}
			switch subChoice {
			case 1:
				addSkripsi(&s, &n, &m)
			case 2:
				fmt.Print("⟡ ݁₊ .Masukkan judul penelitian/nama mahasiswa yang akan di-update: ")
				keyword, _ = reader.ReadString('\n')
				keyword = strings.TrimSpace(keyword)

				idx = getSkripsiIdx(s, n, keyword)

				if idx == -1 {
					fmt.Println("Skripsi tidak ditemukan")
					continue
				}

				idx = getSkripsiIdx(s, n, keyword)
				if idx == -1 {
					fmt.Println("Skripsi tidak ditemukan")
					continue
				}
				updateSkripsi(&s, &m, idx)
			case 3:
				fmt.Print("⟡ ݁₊ .Masukkan judul penelitian/nama mahasiswa yang akan dihapus: ")
				fmt.Print("⟡ ݁₊ .Masukkan keyword: ")
				keyword, _ = reader.ReadString('\n')
				keyword = strings.TrimSpace(keyword)
				idx = getSkripsiIdx(s, n, keyword)
				if idx == -1 {
					fmt.Println("Skripsi tidak ditemukan")
					continue
				}
				deleteSkripsi(&s, &n, idx)
			default:
				return
			}
		case 2:
			fmt.Println("✦===========================✦")
			fmt.Printf("%-5s%-23s%-5s\n", "✦", "Cari Skripsi", "✦")
			fmt.Println("✦===========================✦")
			fmt.Println("0. Kembali ke Menu Utama")
			fmt.Println("1. Pencarian Sequential")
			fmt.Println("2. Pencarian Binary")
			fmt.Print("⟡ ݁₊ .Masukkan pilihan: ")
			for {
				fmt.Print("⟡ ݁₊ .Masukkan pilihan: ")
				input, _ = reader.ReadString('\n')
				input = strings.TrimSpace(input)
				if isNumber(input) {
					fmt.Sscanf(input, "%d", &subChoice)
					break
				}

				fmt.Println("✗ Pilihan harus angka!")
			}
			if subChoice == 0 {
				continue
			}
			switch subChoice {
			case 1:
				fmt.Print("⟡ ݁₊ .Masukkan keyword pencarian: ")
				keyword, _ = reader.ReadString('\n')
				keyword = strings.TrimSpace(keyword)
				findSkripsiSequential(s, n, keyword)
			case 2:
				fmt.Print("⟡ ݁₊ .Masukkan keyword pencarian: ")
				keyword, _ = reader.ReadString('\n')
				keyword = strings.TrimSpace(keyword)
				sortSkripsiInsertion(&s, n, "year", "asc")
				findSkripsiBinary(s, n, keyword)
			default:
				return
			}
		case 3:
			fmt.Println("✦===========================✦")
			fmt.Printf("%-5s%-23s%-5s\n", "✦", "Urutkan Skripsi", "✦")
			fmt.Println("✦===========================✦")
			fmt.Println("0. Kembali ke Menu Utama")
			fmt.Println("1. Selection Sort")
			fmt.Println("2. Insertion Sort")
			fmt.Print("⟡ ݁₊ .Masukkan pilihan: ")
			for {
				fmt.Print("⟡ ݁₊ .Masukkan pilihan: ")
				input, _ = reader.ReadString('\n')
				input = strings.TrimSpace(input)
				if isNumber(input) {
					fmt.Sscanf(input, "%d", &subChoice)
					break
				}

				fmt.Println("✗ Pilihan harus angka!")
			}
			if subChoice == 0 {
				continue
			}
			switch subChoice {
			case 1:
				fmt.Print("⟡ ݁₊ .Masukkan tipe pengurutan (title/year): ")
				sortBy, _ = reader.ReadString('\n')
				sortBy = strings.TrimSpace(sortBy)
				fmt.Print("⟡ ݁₊ .Masukkan tipe pengurutan (asc/desc): ")
				sortType, _ = reader.ReadString('\n')
				sortType = strings.TrimSpace(sortType)
				sortSkripsiSelection(&s, n, sortBy, sortType)
			case 2:
				fmt.Print("⟡ ݁₊ .Masukkan tipe pengurutan (title/year): ")
				sortBy, _ = reader.ReadString('\n')
				sortBy = strings.TrimSpace(sortBy)
				fmt.Print("⟡ ݁₊ .Masukkan tipe pengurutan (asc/desc): ")
				sortType, _ = reader.ReadString('\n')
				sortType = strings.TrimSpace(sortType)
				sortSkripsiInsertion(&s, n, sortBy, sortType)
			default:
				return
			}
		case 4:
			fmt.Println("✦===========================✦")
			fmt.Printf("%-5s%-23s%-5s\n", "✦", "Statistik Skripsi", "✦")
			fmt.Println("✦===========================✦")
			statisticSkripsi(s, n)
			continue
		case 5:
			getAllSkripsi(s, n)
			continue
		case 6:
			return
		default:
			for {
				fmt.Print("⟡ ݁₊ .Masukkan pilihan: ")
				input, _ = reader.ReadString('\n')
				input = strings.TrimSpace(input)
				if isNumber(input) {
					fmt.Sscanf(input, "%d", &subChoice)
					break
				}
				fmt.Println("✗ Pilihan harus angka!")
			}
		}
	}
}
