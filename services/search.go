package services

import (
	"fmt"
	. "skripsin/models"
	. "skripsin/utils"
)

// Prosedur untuk mencari id skripsi berdasarkan nama mahasiswa atau Judul Penelitian menggunakan sequential search
func FindSkripsiSequential(s SkripsiList, n int, keyword string) {
	var f SkripsiList
	var length int = 0
	for i := 0; i < n; i++ {
		if SequentialContains(s[i].Title, keyword) || SequentialContains(s[i].Author.Name, keyword) {
			f[length] = s[i]
			length++
		}
	}

	if length == 0 {
		fmt.Println("⚠︎ Data tidak ditemukan.")
	} else {
		PrintSkripsi(f, length)
	}
}

// Prosedur untuk mencari id skripsi berdasarkan nama mahasiswa atau Judul Penelitian menggunakan binary search
func FindSkripsiBinary(s SkripsiList, n int, keyword string) {
	var tempList SkripsiList

	for i := 0; i < n; i++ {
		tempList[i] = s[i]
	}

	SortSkripsiInsertion(&tempList, n, "name", "asc")
	idx := BinaryFind(tempList, n, keyword, "name")
	if idx != -1 {
		SinglePrint(tempList[idx], "DATA DITEMUKAN BERDASARKAN NAMA MAHASISWA")
		return
	}

	SortSkripsiInsertion(&tempList, n, "title", "asc")
	idx = BinaryFind(tempList, n, keyword, "title")
	if idx != -1 {
		SinglePrint(tempList[idx], "DATA DITEMUKAN BERDASARKAN JUDUL SKRIPSI")
		return
	}
	fmt.Println("⚠︎ Data tidak ditemukan. (Pastikan huruf kapital dan ejaan sama persis karena ini Binary Search).")
}
