package services

import (
	"fmt"
	. "skripsin/models"
	. "skripsin/utils"
	"strings"
)

// Prosedur untuk menambah data skripsi beserta data mahasiswa yang terkait
func AddSkripsi(s *SkripsiList, n *int) {
	var s1 Skripsi
	var addition, start, end, j int
	var input string
	var found bool

	fmt.Println()
	fmt.Println("✦===========================✦")
	fmt.Printf("%-5s%-23s%-5s\n", "✦", "Tambah Data Skripsi", "✦")
	fmt.Println("✦===========================✦")

	input = ReadInput("Jumlah Skripsi yang akan ditambahkan: ")
	isValidAdd := IsNumber(input)
	for !isValidAdd {
		fmt.Println("⚠︎ Input harus angka!")
		input = ReadInput("Jumlah Skripsi yang akan ditambahkan: ")
		isValidAdd = IsNumber(input)
	}

	fmt.Sscanf(input, "%d", &addition)
	if *n+addition > len(*s) {
		fmt.Println("⚠︎ Kapasitas data penuh!")
		return
	}

	start = *n
	end = start + addition
	for i := start; i < end; i++ {
		fmt.Printf("\nSkripsi %d\n", i+1)

		input = ReadInput("Tahun: ")
		isValidYear := IsValidYear(input)
		for !isValidYear {
			fmt.Println("⚠︎ Tahun harus valid!")
			input = ReadInput("Tahun: ")
			isValidYear = IsValidYear(input)
		}
		fmt.Sscanf(input, "%d", &s1.Year)

		input = ReadInput("Judul: ")
		s1.Title = input

		input = ReadInput("Topik Penelitian: ")
		isValidTopic := IsAlpha(input)
		for !isValidTopic {
			fmt.Println("⚠︎ Topik Penelitian hanya boleh huruf!")
			input = ReadInput("Topik Penelitian: ")
			isValidTopic = IsAlpha(input)
		}
		s1.Topic = input

		input = ReadInput("NIM Mahasiswa: ")
		isValidNIM := IsNumber(input)
		for !isValidNIM {
			fmt.Println("⚠︎ NIM hanya boleh angka!")
			input = ReadInput("NIM Mahasiswa: ")
			isValidNIM = IsNumber(input)
		}
		s1.Author.NIM = input

		input = ReadInput("Nama Mahasiswa: ")
		isValidName := IsAlpha(input)
		for !isValidName {
			fmt.Println("⚠︎ Nama hanya boleh huruf!")
			input = ReadInput("Nama Mahasiswa: ")
			isValidName = IsAlpha(input)
		}
		s1.Author.Name = input

		input = ReadInput("Dosen Pembimbing: ")
		isValidDosbing := IsAlpha(input)
		for !isValidDosbing {
			fmt.Println("⚠︎ Nama dosen hanya boleh huruf!")
			input = ReadInput("Dosen Pembimbing: ")
			isValidDosbing = IsAlpha(input)
		}
		s1.Author.DosBing = input

		s1.Author.IsGraduated = false
		j = 0
		found = false
		for j < *n && !found {
			if s[j].Author.NIM == s1.Author.NIM {
				s1.Author.IsGraduated = s[j].Author.IsGraduated
				found = true
			}
			j++
		}
		s[i] = s1
	}
	*n = end
	fmt.Println("\n⌯⌲⌲ Data Berhasil Ditambahkan!")
	fmt.Println()
}

// Fungsi untuk mencari Index skripsi berdasarkan nama mahasiswa atau Judul Penelitian menggunakan sequential search
func GetSkripsiIdx(s SkripsiList, n int, keyword string) int {
	for i := 0; i < n; i++ {
		// Fungsi ini bukan keyword saja tapi harus exact match, "contains" hanya handling just in case ada spasi di depan/belakang judul yang diketik
		if SequentialContains(s[i].Title, keyword) || SequentialContains(s[i].Author.Name, keyword) {
			return i
		}
	}
	return -1
}

// Prosedur untuk mengupdate data skripsi berdasarkan Index
func UpdateSkripsi(s *SkripsiList, n int, idx int) {
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

	input = ReadInput("╰┈➤Masukkan pilihan: ")
	isValidMenu := IsValidMenuChoice(input, 0, 7)
	for !isValidMenu {
		fmt.Println("⚠︎ Pilihan harus berupa angka pada rentang 0 sampai 7!")
		input = ReadInput("╰┈➤Masukkan pilihan: ")
		isValidMenu = IsValidMenuChoice(input, 0, 7)
	}
	fmt.Sscanf(input, "%d", &choice)

	switch choice {
	case 0:
		return
	case 1:
		fmt.Println("⛧ Judul Sebelumnya: ", s[idx].Title)
		input = ReadInput("╰•➤ Masukkan Judul Skripsi baru: ")
		s[idx].Title = strings.TrimSpace(input)
	case 2:
		input = ReadInput("╰•➤ Masukkan Topik Penelitian baru: ")
		isValidTopic := IsAlpha(input)
		for !isValidTopic {
			fmt.Println("⚠︎ Topik Penelitian hanya boleh huruf!")
			input = ReadInput("╰•➤ Masukkan Topik Penelitian baru: ")
			isValidTopic = IsAlpha(input)
		}
		s[idx].Topic = input
	case 3:
		fmt.Println("⛧ Tahun Sebelumnya: ", s[idx].Year)
		input = ReadInput("╰•➤ Masukkan Tahun Skripsi baru: ")
		isValidYear := IsValidYear(input)
		for !isValidYear {
			fmt.Println("⚠︎ Tahun harus valid!")
			input = ReadInput("╰•➤ Masukkan Tahun Skripsi baru: ")
			isValidYear = IsValidYear(input)
		}
		fmt.Sscanf(input, "%d", &s[idx].Year)
	case 4:
		fmt.Println("⛧ NIM Sebelumnya: ", s[idx].Author.NIM)
		fmt.Println("⛧ Nama Penulis Sebelumnya: ", s[idx].Author.Name)

		input = ReadInput("╰•➤ Masukkan NIM Mahasiswa baru: ")
		isValidNIM := IsNumber(input)
		for !isValidNIM {
			fmt.Println("⚠︎ NIM hanya boleh angka!")
			input = ReadInput("╰•➤ Masukkan NIM Mahasiswa baru: ")
			isValidNIM = IsNumber(input)
		}
		s[idx].Author.NIM = input

		input = ReadInput("╰•➤ Masukkan Nama Mahasiswa baru: ")
		isValidName := IsAlpha(input)
		for !isValidName {
			fmt.Println("⚠︎ Nama hanya boleh huruf!")
			input = ReadInput("╰•➤ Masukkan Nama Mahasiswa baru: ")
			isValidName = IsAlpha(input)
		}
		s[idx].Author.Name = input
	case 5:
		input = ReadInput("╰•➤ Masukkan Dosen Pembimbing baru: ")
		isValidDosbing := IsAlpha(input)
		for !isValidDosbing {
			fmt.Println("⚠︎ Nama Dosen hanya boleh huruf!")
			input = ReadInput("╰•➤ Masukkan Dosen Pembimbing baru: ")
			isValidDosbing = IsAlpha(input)
		}
		s[idx].Author.DosBing = input
	case 6:
		input = ReadInput("⏾ Update Status Kelulusan (1=lulus, 0=belum): ")
		isValidStatus := (input == "1" || input == "0")
		for !isValidStatus {
			fmt.Println("⚠︎ Hanya boleh input 1 atau 0!")
			input = ReadInput("⏾ Update Status Kelulusan (1=lulus, 0=belum): ")
			isValidStatus = (input == "1" || input == "0")
		}

		if input == "1" {
			s[idx].Author.IsGraduated = true
		} else {
			s[idx].Author.IsGraduated = false
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
	SinglePrint(s[idx])
}

// Prosedur untuk menghapus data skripsi berdasarkan Index
func DeleteSkripsi(s *SkripsiList, n *int, idx int) {
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
func PrintSkripsi(s SkripsiList, n int) {
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
		status = GetStatusString(s[i].Author.IsGraduated)
		fmt.Printf("✦ %-5d | %-52s | %-23s | %-7s | %-15s | %-23s | %-5s ✦\n", s[i].Year, s[i].Title, s[i].Topic, s[i].Author.NIM, s[i].Author.Name, s[i].Author.DosBing, status)
	}
	fmt.Println("✦======================================================================================================================================================✦")
	fmt.Println()
}

// Prosedur untuk menampilkan (cetak) satu data skripsi
func SinglePrint(s Skripsi) {
	fmt.Println()
	var status string
	status = GetStatusString(s.Author.IsGraduated)
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
