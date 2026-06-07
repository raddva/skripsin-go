package models

// Global Variables
const NMax = 999
const NYear = 2026

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
type SkripsiList [NMax]Skripsi
