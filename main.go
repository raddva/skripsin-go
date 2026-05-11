package main

import "fmt"

const nMax = 100

type Mahasiswa struct {
	NIM, Name, DosBing string
	IsGraduated        bool
}

type Skripsi struct {
	ID, Year int
	Title    string
	Author   Mahasiswa
}

type SkripsiList [nMax]Skripsi
type MahasiswaList [nMax]Mahasiswa

func addSkripsi() {
}

func updateSkripsi() {
}

func deleteSkripsi() {
}

func getSkripsi() {
}

func findSkripsiSequential() {
}

func findSkripsiBinary() {
}

func sortSkripsiSelection() {
}

func sortSkripsiInsertion() {
}

func statisticSkripsi() {
}

func main() {
	var mainChoice, subChoice int
	// var s SkripsiList
	// var m MahasiswaList

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
				addSkripsi()
			case 2:
				updateSkripsi()
			case 3:
				deleteSkripsi()
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
				findSkripsiSequential()
			case 2:
				findSkripsiBinary()
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
				sortSkripsiSelection()
			case 2:
				sortSkripsiInsertion()
			default:
				return
			}
		case 4:
			fmt.Println("+===========================+")
			fmt.Printf("%-5s%-23s%-5s\n", "+", "Statistik Skripsi", "+")
			fmt.Println("+===========================+")
			statisticSkripsi()
			continue
		case 5:
			return
		default:
			fmt.Print("Masukkan pilihan: ")
			fmt.Scan(&mainChoice)
		}
	}
}
