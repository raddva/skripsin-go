// disclaimer: kodingan masih berantakan dikit, nanti final dirapihin (insyaAllah)
// ini komen bner komen manusia (saya)(nadya), bingung jelasinnyh ya gitulah pokonya
package main

import (
	"fmt"
	. "skripsin/models"
	. "skripsin/services"
	. "skripsin/utils"
)

// Menu utama untuk menjalankan program
func main() {
	var mainChoice, subChoice, idx, n int
	var keyword, sortBy, sortType, input, by, tipe string
	var s SkripsiList

	// dummy data testing (temp)
	s, n = GetDummyData() // komen jika tidak pakai dummy data
	// n = 0 // uncoment jika tidak testing pakai dummy data

	isRunning := true // Flag loop

	for isRunning {
		fmt.Println("╔══════════════════════════════════════════════════════════════════════╗")
		fmt.Printf("║ %-68s ║\n", "SkripsIn - Sistem Informasi Inventaris Dokumen Skripsi")
		fmt.Println("╠══════════════════════════════════════════════════════════════════════╣")
		fmt.Println("║ 1. Kelola Data Skripsi                                               ║")
		fmt.Println("║ 2. Pencarian Skripsi                                                 ║")
		fmt.Println("║ 3. Pengurutan Skripsi                                                ║")
		fmt.Println("║ 4. Statistik Skripsi                                                 ║")
		fmt.Println("║ 5. Tampilkan Semua Data Skripsi                                      ║")
		fmt.Println("║ 6. Keluar                                                            ║")
		fmt.Println("╚══════════════════════════════════════════════════════════════════════╝")

		input = ReadInput("➜ Masukkan pilihan: ")
		isValidMain := IsValidMenuChoice(input, 1, 6)
		for !isValidMain {
			fmt.Println("⚠︎ Pilihan harus berupa angka pada rentang 1 sampai 6!")
			input = ReadInput("➜ Masukkan pilihan: ")
			isValidMain = IsValidMenuChoice(input, 1, 6)
		}
		fmt.Sscanf(input, "%d", &mainChoice)

		switch mainChoice {
		case 1:
			fmt.Println()
			fmt.Println("╔════════════════════════════════════╗")
			fmt.Printf("║ %-34s ║\n", "Kelola Data Skripsi")
			fmt.Println("╠════════════════════════════════════╣")
			fmt.Println("║ 0. Kembali ke Menu Utama           ║")
			fmt.Println("║ 1. Tambah Skripsi                  ║")
			fmt.Println("║ 2. Update Skripsi                  ║")
			fmt.Println("║ 3. Hapus Skripsi                   ║")
			fmt.Println("╚════════════════════════════════════╝")

			input = ReadInput("➜ Masukkan pilihan: ")
			isValidSub := IsValidMenuChoice(input, 0, 3)
			for !isValidSub {
				fmt.Println("⚠︎ Pilihan harus berupa angka pada rentang 0 sampai 3!")
				input = ReadInput("➜ Masukkan pilihan: ")
				isValidSub = IsValidMenuChoice(input, 0, 3)
			}
			fmt.Sscanf(input, "%d", &subChoice)

			switch subChoice {
			case 0:
				fmt.Println()
			case 1:
				AddSkripsi(&s, &n)
			case 2:
				keyword = ReadInput("⪼---➢ Masukkan judul penelitian/nama mahasiswa yang akan di-update: ")
				idx = GetSkripsiIdx(s, n, keyword)
				if idx == -1 {
					fmt.Println("Skripsi tidak ditemukan")
				} else {
					UpdateSkripsi(&s, n, idx)
					fmt.Println()
				}
			case 3:
				keyword = ReadInput("⪼---➢ Masukkan judul penelitian/nama mahasiswa yang akan dihapus: ")
				idx = GetSkripsiIdx(s, n, keyword)
				if idx == -1 {
					fmt.Println("Skripsi tidak ditemukan")
				} else {
					DeleteSkripsi(&s, &n, idx)
				}
			}

		case 2:
			fmt.Println()
			fmt.Println("╔════════════════════════════════════╗")
			fmt.Printf("║ %-34s ║\n", "Cari Skripsi")
			fmt.Println("╠════════════════════════════════════╣")
			fmt.Println("║ 0. Kembali ke Menu Utama           ║")
			fmt.Println("║ 1. Pencarian Sequential            ║")
			fmt.Println("║ 2. Pencarian Binary                ║")
			fmt.Println("╚════════════════════════════════════╝")

			input = ReadInput("➜ Masukkan pilihan: ")
			isValidSubSearch := IsValidMenuChoice(input, 0, 2)
			for !isValidSubSearch {
				fmt.Println("⚠︎ Pilihan harus berupa angka 0, 1, atau 2!")
				input = ReadInput("➜ Masukkan pilihan: ")
				isValidSubSearch = IsValidMenuChoice(input, 0, 2)
			}
			fmt.Sscanf(input, "%d", &subChoice)

			switch subChoice {
			case 0:
				fmt.Println()
			case 1:
				keyword = ReadInput("⪼---➢ Masukkan keyword pencarian: ")
				FindSkripsiSequential(s, n, keyword)
			case 2:
				keyword = ReadInput("⪼---➢ Masukkan Judul/Nama Mahasiswa yang akan dicari: ")
				FindSkripsiBinary(s, n, keyword)
			}

		case 3:
			fmt.Println()
			fmt.Println("╔════════════════════════════════════╗")
			fmt.Printf("║ %-34s ║\n", "Urutkan Skripsi")
			fmt.Println("╠════════════════════════════════════╣")
			fmt.Println("║ 0. Kembali ke Menu Utama           ║")
			fmt.Println("║ 1. Selection Sort                  ║")
			fmt.Println("║ 2. Insertion Sort                  ║")
			fmt.Println("╚════════════════════════════════════╝")

			input = ReadInput("➜ Masukkan pilihan: ")
			isValidSubSort := IsValidMenuChoice(input, 0, 2)
			for !isValidSubSort {
				fmt.Println("⚠︎ Pilihan harus berupa angka 0, 1, atau 2!")
				input = ReadInput("➜ Masukkan pilihan: ")
				isValidSubSort = IsValidMenuChoice(input, 0, 2)
			}
			fmt.Sscanf(input, "%d", &subChoice)

			if subChoice != 0 {
				sortBy = ReadInput("⪼---➢ Masukkan tipe pengurutan (name/year/title): ")
				isValidSortBy := (sortBy == "name" || sortBy == "year" || sortBy == "title")
				for !isValidSortBy {
					fmt.Println("⚠︎ Pilihan hanya boleh 'name', 'year', atau 'title'!")
					sortBy = ReadInput("⪼---➢ Masukkan tipe pengurutan (name/year/title): ")
					isValidSortBy = (sortBy == "name" || sortBy == "year" || sortBy == "title")
				}

				switch sortBy {
				case "name":
					by = "Nama Penulis"
				case "year":
					by = "Tahun"
				case "title":
					by = "Judul Penelitian"
				}

				sortType = ReadInput("⪼---➢ Masukkan tipe pengurutan (asc/desc): ")
				isValidSortType := (sortType == "asc" || sortType == "desc")
				for !isValidSortType {
					fmt.Println("⚠︎ Pilihan hanya boleh 'asc' atau 'desc'!")
					sortType = ReadInput("⪼---➢ Masukkan tipe pengurutan (asc/desc): ")
					isValidSortType = (sortType == "asc" || sortType == "desc")
				}

				if sortType == "asc" {
					tipe = "Ascending"
				} else {
					tipe = "Descending"
				}

				switch subChoice {
				case 1:
					SortSkripsiSelection(&s, n, sortBy, sortType)
					fmt.Printf("\nPengurutan Skripsi Menggunakan Selection Sort Berdasarkan %s secara %s:\n", by, tipe)
					PrintSkripsi(s, n)
				case 2:
					SortSkripsiInsertion(&s, n, sortBy, sortType)
					fmt.Printf("\nPengurutan Skripsi Menggunakan Insertion Sort Berdasarkan %s secara %s:\n", by, tipe)
					PrintSkripsi(s, n)
				}
			}
		case 4:
			fmt.Println()
			StatisticSkripsi(s, n)
		case 5:
			PrintSkripsi(s, n)
		case 6:
			isRunning = false
		}
	}
}

// kelar kah min? belom sih bentar laporan blm
