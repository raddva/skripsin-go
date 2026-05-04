package main

import "fmt"

type Mahasiswa struct {
	NIM, Name, DosBing string
	IsGraduated        bool
}

type Skripsi struct {
	ID, Year int
	Title    string
	Author   Mahasiswa
}

type SkripsiList []Skripsi
type MahasiswaList []Mahasiswa

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
	var choice int
	var s SkripsiList
	var m MahasiswaList

	fmt.Printf("%-25s\n", "SkripsIn - Sistem Informasi Inventaris Dokumen Skripsi")
	fmt.Printf("%-25s\n", "Pilih Menu")

	fmt.Printf("%-25s\n", "1. Kelola Data Skripsi")
	fmt.Printf("%-25s\n", "2. Pencarian Skripsi")
	fmt.Printf("%-25s\n", "3. Pengurutan Skripsi")
	fmt.Printf("%-25s\n", "4. Statistik Skripsi")
	fmt.Printf("%-25s\n", "5. Keluar")

	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("Kelola Data Skripsi")
		fmt.Println("1. Tambah Skripsi")
		fmt.Println("2. Update Skripsi")
		fmt.Println("3. Hapus Skripsi")
		fmt.Scan(&choice)
		switch choice {
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
		fmt.Println("Cari Skripsi")
		fmt.Println("1. Pencarian Sequential")
		fmt.Println("2. Pencarian Binary")
		fmt.Scan(&choice)
	case 3:
		fmt.Println("Urutkan Skripsi")
		fmt.Println("1. Selection Sort")
		fmt.Println("2. Insertion Sort")
		fmt.Scan(&choice)
	case 4:
		fmt.Println("Statistik Skripsi")
		statisticSkripsi()
	default:
		return
	}
}
