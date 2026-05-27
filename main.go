// disclaimer: kodingan masih berantakan dikit, nanti final dirapihin (insyaAllah)
// ini komen bner komen manusia (saya)(nadya), bingung jelasinnyh ya gitulah pokonya
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Global Variables
const nMax = 999
const nYear = 2026 // Batasi tahun: 2026

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
	Topic  string
	Author Mahasiswa
}

// End of Struktur Data

// Tipe data untuk menyimpan daftar skripsi dan daftar mahasiswa
type SkripsiList [nMax]Skripsi

// Helper Functions
func readInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func isNumber(s string) bool {
	if s == "" {
		return false
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
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

	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c < 'A' || c > 'Z') &&
			(c < 'a' || c > 'z') &&
			c != ' ' &&
			c != '.' &&
			c != ',' {
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
	if year < 1900 || year > nYear { // batasi karena sekarang baru 2026 :v
		return false
	}
	return true
}

// Fungsi untuk mengecek apakah input adalah angka dan berada di range menu valid
func isValidMenuChoice(input string, min int, max int) bool {
	if !isNumber(input) {
		return false
	}

	var choice int
	fmt.Sscanf(input, "%d", &choice)
	if choice >= min && choice <= max {
		return true
	}
	return false
}

func sequentialContains(text, keyword string) bool {
	text = strings.ToLower(text)
	keyword = strings.ToLower(keyword)
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
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func binaryContains(s SkripsiList, n int, keyword string, searchBy string) int {
	keyword = strings.ToLower(keyword)
	left := 0
	right := n - 1

	for left <= right {
		mid := (left + right) / 2
		var midValue string
		switch searchBy {
		case "name":
			midValue = s[mid].Author.Name
		case "title":
			midValue = s[mid].Title
		}

		midValue = strings.ToLower(midValue)
		if midValue == keyword {
			return mid
		}

		if midValue < keyword {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func compareSkripsi(a, b Skripsi, sortBy, sortType string) bool {
	isAsc := sortType == "asc"
	switch sortBy {
	case "name":
		nameA := strings.ToLower(a.Author.Name)
		nameB := strings.ToLower(b.Author.Name)
		if isAsc {
			return nameA < nameB
		}
		return nameA > nameB
	case "year":
		if isAsc {
			return a.Year < b.Year
		}
		return a.Year > b.Year
	case "title":
		titleA := strings.ToLower(a.Title)
		titleB := strings.ToLower(b.Title)
		if isAsc {
			return titleA < titleB
		}
		return titleA > titleB
	}
	return false
}

func getStatusString(isGraduated bool) string {
	if isGraduated {
		return "Lulus"
	}
	return "Belum"
}

// End of Helper Functions

// Start Subprograms
// Prosedur untuk menambah data skripsi beserta data mahasiswa yang terkait
func addSkripsi(s *SkripsiList, n *int) {
	var s1 Skripsi
	var addition, start, end, j int
	var input string
	var found bool = false

	fmt.Println()
	fmt.Println("✦===========================✦")
	fmt.Printf("%-5s%-23s%-5s\n", "✦", "Tambah Data Skripsi", "✦")
	fmt.Println("✦===========================✦")
	for {
		input = readInput("Jumlah Skripsi yang akan ditambahkan: ")
		if isNumber(input) {
			fmt.Sscanf(input, "%d", &addition)
			if *n+addition > len(s) {
				fmt.Println("⚠︎ Kapasitas data penuh!")
				return
			}
			break
		} else {
			fmt.Println("⚠︎ Input harus angka!")
		}
	}

	start = *n
	end = start + addition
	for i := start; i < end; i++ {
		fmt.Printf("\nSkripsi %d\n", i+1)
		for {
			input = readInput("Tahun: ")
			if isValidYear(input) {
				fmt.Sscanf(input, "%d", &s1.Year)
				break
			}
			fmt.Println("⚠︎ Tahun harus valid!")
		}
		input = readInput("Judul: ")
		s1.Title = input
		for {
			input = readInput("Topik Penelitian: ")
			if isAlpha(input) {
				s1.Topic = input
				break
			}
			fmt.Println("⚠︎ Topik Penelitian hanya boleh huruf!")
		}
		for {
			input = readInput("NIM Mahasiswa: ")
			if isNumber(input) {
				s1.Author.NIM = input
				break
			}
			fmt.Println("⚠︎ NIM hanya boleh angka!")
		}

		for {
			input = readInput("Nama Mahasiswa: ")
			if isAlpha(input) {
				s1.Author.Name = input
				break
			}
			fmt.Println("⚠︎ Nama hanya boleh huruf!")
		}

		for {
			input = readInput("Dosen Pembimbing: ")
			if isAlpha(input) {
				s1.Author.DosBing = input
				break
			}
			fmt.Println("⚠︎ Nama dosen hanya boleh huruf!")
		}
		s1.Author.IsGraduated = false
		for j < *n && !found {
			if s[j].Author.NIM == s1.Author.NIM {
				s1.Author.IsGraduated = s[j].Author.IsGraduated
				found = true
			}
			j++
		}
	}
	*n = end
	fmt.Println("\n⌯⌲⌲ Data Berhasil Ditambahkan!")
	fmt.Println()
}

// Fungsi untuk mencari Index skripsi berdasarkan nama mahasiswa atau Judul Penelitian menggunakan sequential search
func getSkripsiIdx(s SkripsiList, n int, keyword string) int {
	for i := 0; i < n; i++ {
		// Fungsi ini bukan keyword saja tapi harus exact match, "contains" hanya handling just in case ada spasi di depan/belakang judul yang diketik
		if sequentialContains(s[i].Title, keyword) || sequentialContains(s[i].Author.Name, keyword) {
			return i
		}
	}
	return -1
}

// Prosedur untuk mengupdate data skripsi berdasarkan Index
func updateSkripsi(s *SkripsiList, n int, idx int) {
	var choice int
	var input string
	fmt.Println()
	fmt.Println("✦===========================✦")
	fmt.Printf("%-5s%-23s%-5s\n", "✦", "Edit Data Skripsi", "✦")
	fmt.Println("✦===========================✦")
	fmt.Println("0. Kembali ke Menu Utama")
	fmt.Println("1. Judul Skripsi")
	fmt.Println("2. Topik Skripsi")
	fmt.Println("3. Tahun Skripsi")
	fmt.Println("4. Penulis Skripsi")
	fmt.Println("5. Dosen Pembimbing Skripsi")
	fmt.Println("6. Status Kelulusan")
	fmt.Println("7. Kembali")
	for {
		input = readInput("╰┈➤ˎˊ˗Masukkan pilihan: ")
		if isValidMenuChoice(input, 0, 7) {
			fmt.Sscanf(input, "%d", &choice)
			break
		}
		fmt.Println("⚠︎ Pilihan harus berupa angka pada rentang 0 sampai 7!")
	}

	switch choice {
	case 0:
		return
	case 1:
		fmt.Println("⛧ Judul Sebelumnya: ", s[idx].Title)
		input = readInput("╰•➤ Masukkan Judul Skripsi baru: ")
		s[idx].Title = strings.TrimSpace(input)
	case 2:
		for {
			input = readInput("╰•➤ Masukkan Topik Penelitian baru: ")
			if isAlpha(input) {
				s[idx].Topic = input
				break
			}
			fmt.Println("⚠︎ Topik Penelitian hanya boleh huruf!")
		}
	case 3:
		fmt.Println("⛧ Tahun Sebelumnya: ", s[idx].Year)
		for {
			input = readInput("╰•➤ Masukkan Tahun Skripsi baru: ")
			if isValidYear(input) {
				fmt.Sscanf(input, "%d", &s[idx].Year)
				break
			}
			fmt.Println("⚠︎ Tahun harus valid!")
		}
	case 4:
		fmt.Println("⛧ NIM Sebelumnya: ", s[idx].Author.NIM)
		fmt.Println("⛧ Nama Penulis Sebelumnya: ", s[idx].Author.Name)
		for {
			input = readInput("╰•➤ Masukkan NIM Mahasiswa baru: ")
			if isNumber(input) {
				s[idx].Author.NIM = input
				break
			}
			fmt.Println("⚠︎ NIM hanya boleh angka!")
		}

		for {
			input = readInput("╰•➤ Masukkan Nama Mahasiswa baru: ")
			if isAlpha(input) {
				s[idx].Author.Name = input
				break
			}
			fmt.Println("⚠︎ Nama hanya boleh huruf!")
		}
	case 5:
		for {
			input = readInput("╰•➤ Masukkan Dosen Pembimbing baru: ")
			if isAlpha(input) {
				s[idx].Author.DosBing = input
				break
			}
			fmt.Println("⚠︎ Nama Dosen hanya boleh huruf!")
		}
	case 6:
		for {
			input = readInput("⏾ Update Status Kelulusan (1=lulus, 0=belum): ")
			if input == "1" {
				s[idx].Author.IsGraduated = true
				break
			} else if input == "0" {
				s[idx].Author.IsGraduated = false
				break
			}
			fmt.Println("⚠︎ Hanya boleh input 1 atau 0!")
		}

		for j := 0; j < n; j++ {
			if s[j].Author.NIM == s[idx].Author.NIM {
				s[j].Author.IsGraduated = s[idx].Author.IsGraduated
			}
		}
	case 7:
		return
	}
	fmt.Println("⌯⌲⌲ Data Berhasil Diperbarui!")
	fmt.Println("Data Sekarang:")
	singlePrint(s[idx])
}

// Prosedur untuk menghapus data skripsi berdasarkan Index
func deleteSkripsi(s *SkripsiList, n *int, idx int) {
	fmt.Println()
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
	fmt.Println("⌯⌲⌲ Data Berhasil Dihapus!")
	fmt.Println()
}

// Prosedur untuk menampilkan (cetak) semua data skripsi
func printSkripsi(s SkripsiList, n int) {
	fmt.Println()
	if n == 0 {
		fmt.Println("✦===========================✦")
		fmt.Println("    Data Skripsi Masih Kosong!   ")
		fmt.Println("✦===========================✦")
		return
	}

	var i int
	var status string
	fmt.Println("✦======================================================================================================================================================✦")
	fmt.Printf("✦ %-5s | %-52s | %-23s | %-7s | %-15s | %-23s | %-5s✦\n", "TAHUN", "JUDUL SKRIPSI", "TOPIK", "NIM", "NAMA", "DOSBING", "STATUS")
	fmt.Println("✦======================================================================================================================================================✦")
	for i = 0; i < n; i++ {
		status = getStatusString(s[i].Author.IsGraduated)
		fmt.Printf("✦ %-5d | %-52s | %-23s | %-7s | %-15s | %-23s | %-5s ✦\n", s[i].Year, s[i].Title, s[i].Topic, s[i].Author.NIM, s[i].Author.Name, s[i].Author.DosBing, status)
	}
	fmt.Println("✦======================================================================================================================================================✦")
	fmt.Println()
}

// Prosedur untuk menampilkan (cetak) satu data skripsi
func singlePrint(s Skripsi) {
	fmt.Println()
	var status string
	status = getStatusString(s.Author.IsGraduated)
	fmt.Println("✦==================================================✦")
	fmt.Printf("✦ TAHUN: %d\n", s.Year)
	fmt.Printf("✦ JUDUL: %s\n", s.Title)
	fmt.Printf("✦ TOPIK: %s\n", s.Topic)
	fmt.Printf("✦ NIM: %s\n", s.Author.NIM)
	fmt.Printf("✦ NAMA: %s\n", s.Author.Name)
	fmt.Printf("✦ DOSEN PEMBIMBING: %s\n", s.Author.DosBing)
	fmt.Printf("✦ STATUS KELULUSAN: %s\n", status)
	fmt.Println("✦==================================================✦")
	fmt.Println()
}

// Prosedur untuk mencari id skripsi berdasarkan nama mahasiswa atau Judul Penelitian menggunakan sequential search
func findSkripsiSequential(s SkripsiList, n int, keyword string) {
	var f SkripsiList
	var length int = 0
	for i := 0; i < n; i++ {
		if sequentialContains(s[i].Title, keyword) || sequentialContains(s[i].Author.Name, keyword) {
			f[length] = s[i]
			length++
		}
	}
	printSkripsi(f, length)
}

// Prosedur untuk mencari id skripsi berdasarkan nama mahasiswa atau Judul Penelitian menggunakan binary search
func findSkripsiBinary(s SkripsiList, n int, keyword string) {
	var tempList SkripsiList

	for i := 0; i < n; i++ {
		tempList[i] = s[i]
	}

	sortSkripsiInsertion(&tempList, n, "name", "asc")
	idx := binaryContains(tempList, n, keyword, "name")
	if idx != -1 {
		singlePrint(tempList[idx])
		return
	}

	sortSkripsiInsertion(&tempList, n, "title", "asc")
	idx = binaryContains(tempList, n, keyword, "title")
	if idx != -1 {
		singlePrint(tempList[idx])
		return
	}
	fmt.Println("⚠︎ Data tidak ditemukan. (Pastikan huruf kapital dan ejaan sama persis karena ini Binary Search).")
}

// Prosedur untuk mengurutkan skripsi berdasarkan nama/tahun menggunakan selection sort
func sortSkripsiSelection(s *SkripsiList, n int, sortBy string, sortType string) {
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if compareSkripsi((*s)[j], (*s)[idx], sortBy, sortType) {
				idx = j
			}
		}
		(*s)[i], (*s)[idx] = (*s)[idx], (*s)[i]
	}
}

// Prosedur untuk mengurutkan skripsi berdasarkan nama/tahun menggunakan insertion sort
func sortSkripsiInsertion(s *SkripsiList, n int, sortBy string, sortType string) {
	for i := 1; i < n; i++ {
		key := (*s)[i]
		j := i - 1
		for j >= 0 && compareSkripsi(key, (*s)[j], sortBy, sortType) {
			(*s)[j+1] = (*s)[j]
			j--
		}
		(*s)[j+1] = key
	}
}

// Prosedur untuk menampilkan statistik skripsi, seperti jumlah skripsi per tahun, jumlah skripsi yang lulus, dll
func statisticSkripsi(s SkripsiList, n int) {
	var CountPerYear [nYear + 1]int
	var CountGraduated int
	for i := 0; i < n; i++ {
		CountPerYear[s[i].Year]++
		if s[i].Author.IsGraduated {
			CountGraduated++
		}
	}
	fmt.Println("✦===============================✦")
	fmt.Printf("%-7s %-23s %-5s\n", "✦", "STATISTIK SKRIPSI", "✦")
	fmt.Println("✦===============================✦")
	fmt.Println("Jumlah Skripsi per tahun:")
	for year := 0; year <= 2026; year++ {
		if CountPerYear[year] > 0 {
			fmt.Printf("➯ %d : %d Skripsi\n", year, CountPerYear[year])
		}
	}
	fmt.Printf("Jumlah Skripsi Lulus: %d\n", CountGraduated)
	fmt.Printf("Total Skripsi: %d\n", n)
	fmt.Println()
}

// End Subprograms

// Menu utama untuk menjalankan program
func main() {
	var mainChoice, subChoice, idx, n int
	var keyword, sortBy, sortType, input, by, tipe string
	var s SkripsiList

	// dummy data testing (temp)
	copy(s[:], dummySkripsi)
	n = len(dummySkripsi) // jika testing pakai dummy data
	// n = 0 jika tidak testing pakai dummy data

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
			input = readInput("╰┈➤ˎˊ˗Masukkan pilihan: ")
			if isValidMenuChoice(input, 1, 6) {
				fmt.Sscanf(input, "%d", &mainChoice)
				break
			}
			fmt.Println("⚠︎ Pilihan harus berupa angka pada rentang 1 sampai 6!")
		}

		switch mainChoice {
		case 1:
			fmt.Println()
			fmt.Println("✦===========================✦")
			fmt.Printf("%-5s%-23s%-5s\n", "✦", "Kelola Data Skripsi", "✦")
			fmt.Println("✦===========================✦")
			fmt.Println("0. Kembali ke Menu Utama")
			fmt.Println("1. Tambah Skripsi")
			fmt.Println("2. Update Skripsi")
			fmt.Println("3. Hapus Skripsi")
			for {
				input = readInput("╰┈➤ˎˊ˗Masukkan pilihan: ")
				if isValidMenuChoice(input, 0, 3) {
					fmt.Sscanf(input, "%d", &subChoice)
					break
				}
				fmt.Println("⚠︎ Pilihan harus berupa angka pada rentang 0 sampai 3!")
			}

			switch subChoice {
			case 0:
				continue
			case 1:
				addSkripsi(&s, &n)
			case 2:
				keyword = readInput("⪼---➢ Masukkan judul penelitian/nama mahasiswa yang akan di-update: ")
				idx = getSkripsiIdx(s, n, keyword)
				if idx == -1 {
					fmt.Println("Skripsi tidak ditemukan")
					continue
				}
				updateSkripsi(&s, n, idx)
			case 3:
				keyword = readInput("⪼---➢ Masukkan judul penelitian/nama mahasiswa yang akan dihapus: ")
				idx = getSkripsiIdx(s, n, keyword)
				if idx == -1 {
					fmt.Println("Skripsi tidak ditemukan")
					continue
				}
				deleteSkripsi(&s, &n, idx)
			}
		case 2:
			fmt.Println()
			fmt.Println("✦===========================✦")
			fmt.Printf("%-5s%-23s%-5s\n", "✦", "Cari Skripsi", "✦")
			fmt.Println("✦===========================✦")
			fmt.Println("0. Kembali ke Menu Utama")
			fmt.Println("1. Pencarian Sequential")
			fmt.Println("2. Pencarian Binary")
			for {
				input = readInput("╰┈➤ˎˊ˗Masukkan pilihan: ")
				if isValidMenuChoice(input, 0, 2) {
					fmt.Sscanf(input, "%d", &subChoice)
					break
				}
				fmt.Println("⚠︎ Pilihan harus berupa angka 0, 1, atau 2!")
			}
			switch subChoice {
			case 0:
				continue
			case 1:
				keyword = readInput("⪼---➢ Masukkan keyword pencarian: ")
				findSkripsiSequential(s, n, keyword)
			case 2:
				keyword = readInput("⪼---➢ Masukkan Judul/Nama Mahasiswa yang akan dicari: ")
				findSkripsiBinary(s, n, keyword)
			}
		case 3:
			fmt.Println()
			fmt.Println("✦===========================✦")
			fmt.Printf("%-5s%-23s%-5s\n", "✦", "Urutkan Skripsi", "✦")
			fmt.Println("✦===========================✦")
			fmt.Println("0. Kembali ke Menu Utama")
			fmt.Println("1. Selection Sort")
			fmt.Println("2. Insertion Sort")
			for {
				input = readInput("╰┈➤ˎˊ˗Masukkan pilihan: ")
				if isValidMenuChoice(input, 0, 2) {
					fmt.Sscanf(input, "%d", &subChoice)
					break
				}
				fmt.Println("⚠︎ Pilihan harus berupa angka 0, 1, atau 2!")
			}

			if subChoice == 0 {
				continue
			}

			for {
				sortBy = readInput("⪼---➢ Masukkan tipe pengurutan (name/year/title): ")
				// based on title sekalian biar ga nganggur
				if sortBy == "name" || sortBy == "year" || sortBy == "title" {
					break
				}
				fmt.Println("⚠︎ Pilihan hanya boleh 'name', 'year', atau 'title'!")
			}
			switch sortBy {
			case "name":
				by = "Nama Penulis"
			case "year":
				by = "Tahun"
			case "title":
				by = "Judul Penelitian"
			}

			for {
				sortType = readInput("⪼---➢ Masukkan tipe pengurutan (asc/desc): ")
				if sortType == "asc" || sortType == "desc" {
					break
				}
				fmt.Println("⚠︎ Pilihan hanya boleh 'asc' atau 'desc'!")
			}

			if sortType == "asc" {
				tipe = "Ascending"
			} else {
				tipe = "Descending"
			}

			switch subChoice {
			case 1:
				sortSkripsiSelection(&s, n, sortBy, sortType)
				fmt.Printf("\nPengurutan Skripsi Menggunakan Selection Sort Berdasarkan %s secara %s:\n", by, tipe)
				printSkripsi(s, n)
			case 2:
				sortSkripsiInsertion(&s, n, sortBy, sortType)
				fmt.Printf("\nPengurutan Skripsi Menggunakan Insertion Sort Berdasarkan %s secara %s:\n", by, tipe)
				printSkripsi(s, n)
			}
		case 4:
			fmt.Println()
			statisticSkripsi(s, n)
			continue
		case 5:
			printSkripsi(s, n)
			continue
		case 6:
			return
		}
	}
}
