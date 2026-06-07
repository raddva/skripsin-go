package main

import (
	. "skripsin/models"
)

func GetDummyData() (SkripsiList, int) {
	var s SkripsiList
	var dummySkripsi = []Skripsi{
		{
			Year:  2021,
			Title: "Implementasi Sistem Informasi Akademik Berbasis Web",
			Topic: "Sistem Informasi",
			Author: Mahasiswa{
				NIM:         "1901001",
				Name:        "Andi Saputra",
				DosBing:     "Dr. Budi Santoso",
				IsGraduated: true,
			},
		},
		{
			Year:  2022,
			Title: "Analisis Sentimen Twitter Menggunakan Naive Bayes",
			Topic: "Machine Learning",
			Author: Mahasiswa{
				NIM:         "1901002",
				Name:        "Siti Rahma",
				DosBing:     "Dr. Dewi Lestari",
				IsGraduated: true,
			},
		},
		{
			Year:  2020,
			Title: "Aplikasi E-Commerce UMKM Berbasis Mobile",
			Topic: "Mobile Development",
			Author: Mahasiswa{
				NIM:         "1801003",
				Name:        "Rizky Hidayat",
				DosBing:     "Ir. Ahmad Fauzi",
				IsGraduated: true,
			},
		},
		{
			Year:  2023,
			Title: "Deteksi Wajah Menggunakan Deep Learning",
			Topic: "Computer Vision",
			Author: Mahasiswa{
				NIM:         "2001004",
				Name:        "Nabila Putri",
				DosBing:     "Dr. Sari Wulandari",
				IsGraduated: false,
			},
		},
		{
			Year:  2021,
			Title: "Perancangan UI/UX Aplikasi Perpustakaan Digital",
			Topic: "UI/UX",
			Author: Mahasiswa{
				NIM:         "1901005",
				Name:        "Fajar Nugroho",
				DosBing:     "M. Ilham Pratama, M.Kom",
				IsGraduated: true,
			},
		},
		{
			Year:  2022,
			Title: "Sistem Prediksi Penjualan Menggunakan Regresi Linear",
			Topic: "Data Science",
			Author: Mahasiswa{
				NIM:         "1901006",
				Name:        "Dewi Anggraini",
				DosBing:     "Dr. Rina Marlina",
				IsGraduated: true,
			},
		},
		{
			Year:  2020,
			Title: "Keamanan Jaringan Wireless Menggunakan WPA3",
			Topic: "Cyber Security",
			Author: Mahasiswa{
				NIM:         "1801007",
				Name:        "Yoga Prasetyo",
				DosBing:     "Ir. Hendra Wijaya",
				IsGraduated: true,
			},
		},
		{
			Year:  2023,
			Title: "Chatbot Pelayanan Akademik Berbasis NLP",
			Topic: "Artificial Intelligence",
			Author: Mahasiswa{
				NIM:         "2001008",
				Name:        "Clara Oktaviani",
				DosBing:     "Dr. Yuni Kartika",
				IsGraduated: false,
			},
		},
		{
			Year:  2021,
			Title: "Monitoring IoT untuk Smart Farming",
			Topic: "Internet of Things",
			Author: Mahasiswa{
				NIM:         "1901009",
				Name:        "Bagas Maulana",
				DosBing:     "Dr. Taufik Hidayat",
				IsGraduated: true,
			},
		},
		{
			Year:  2022,
			Title: "Optimasi Algoritma Pencarian pada Aplikasi Navigasi",
			Topic: "Algoritma",
			Author: Mahasiswa{
				NIM:         "1901010",
				Name:        "Intan Permata",
				DosBing:     "Prof. Agus Setiawan",
				IsGraduated: true,
			},
		},
	}
	for i := 0; i < len(dummySkripsi); i++ {
		s[i] = dummySkripsi[i]
	}

	return s, len(dummySkripsi)
}
