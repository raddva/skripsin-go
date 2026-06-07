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
	s, n = GetDummyData() // jika testing pakai dummy data
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
			input = ReadInput("╰┈➤Masukkan pilihan: ")
			if IsValidMenuChoice(input, 1, 6) {
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
				input = ReadInput("╰┈➤Masukkan pilihan: ")
				if IsValidMenuChoice(input, 0, 3) {
					fmt.Sscanf(input, "%d", &subChoice)
					break
				}
				fmt.Println("⚠︎ Pilihan harus berupa angka pada rentang 0 sampai 3!")
			}

			switch subChoice {
			case 0:
				fmt.Println()
				continue
			case 1:
				AddSkripsi(&s, &n)
			case 2:
				keyword = ReadInput("⪼---➢ Masukkan judul penelitian/nama mahasiswa yang akan di-update: ")
				idx = GetSkripsiIdx(s, n, keyword)
				if idx == -1 {
					fmt.Println("Skripsi tidak ditemukan")
					continue
				}
				UpdateSkripsi(&s, n, idx)
			case 3:
				keyword = ReadInput("⪼---➢ Masukkan judul penelitian/nama mahasiswa yang akan dihapus: ")
				idx = GetSkripsiIdx(s, n, keyword)
				if idx == -1 {
					fmt.Println("Skripsi tidak ditemukan")
					continue
				}
				DeleteSkripsi(&s, &n, idx)
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
				input = ReadInput("╰┈➤Masukkan pilihan: ")
				if IsValidMenuChoice(input, 0, 2) {
					fmt.Sscanf(input, "%d", &subChoice)
					break
				}
				fmt.Println("⚠︎ Pilihan harus berupa angka 0, 1, atau 2!")
			}
			switch subChoice {
			case 0:
				fmt.Println()
				continue
			case 1:
				keyword = ReadInput("⪼---➢ Masukkan keyword pencarian: ")
				FindSkripsiSequential(s, n, keyword)
			case 2:
				keyword = ReadInput("⪼---➢ Masukkan Judul/Nama Mahasiswa yang akan dicari: ")
				FindSkripsiBinary(s, n, keyword)
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
				input = ReadInput("╰┈➤Masukkan pilihan: ")
				if IsValidMenuChoice(input, 0, 2) {
					fmt.Sscanf(input, "%d", &subChoice)
					break
				}
				fmt.Println("⚠︎ Pilihan harus berupa angka 0, 1, atau 2!")
			}

			if subChoice == 0 {
				fmt.Println()
				continue
			}

			for {
				sortBy = ReadInput("⪼---➢ Masukkan tipe pengurutan (name/year/title): ")
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
				sortType = ReadInput("⪼---➢ Masukkan tipe pengurutan (asc/desc): ")
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
				SortSkripsiSelection(&s, n, sortBy, sortType)
				fmt.Printf("\nPengurutan Skripsi Menggunakan Selection Sort Berdasarkan %s secara %s:\n", by, tipe)
				PrintSkripsi(s, n)
			case 2:
				SortSkripsiInsertion(&s, n, sortBy, sortType)
				fmt.Printf("\nPengurutan Skripsi Menggunakan Insertion Sort Berdasarkan %s secara %s:\n", by, tipe)
				PrintSkripsi(s, n)
			}
		case 4:
			fmt.Println()
			StatisticSkripsi(s, n)
			continue
		case 5:
			PrintSkripsi(s, n)
			continue
		case 6:
			return
		}
	}
}
