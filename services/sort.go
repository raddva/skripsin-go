package services

import (
	. "skripsin/models"
	. "skripsin/utils"
)

// Prosedur untuk mengurutkan skripsi berdasarkan nama/tahun menggunakan selection sort
func SortSkripsiSelection(s *SkripsiList, n int, sortBy string, sortType string) {
	var temp Skripsi
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if CompareSkripsi((*s)[j], (*s)[idx], sortBy, sortType) {
				idx = j
			}
		}
		temp = (*s)[i]
		(*s)[i] = (*s)[idx]
		(*s)[idx] = temp
	}
}

// Prosedur untuk mengurutkan skripsi berdasarkan nama/tahun menggunakan insertion sort
func SortSkripsiInsertion(s *SkripsiList, n int, sortBy string, sortType string) {
	var key Skripsi
	var j int
	for i := 1; i < n; i++ {
		key = (*s)[i]
		j = i - 1
		for j >= 0 && CompareSkripsi(key, (*s)[j], sortBy, sortType) {
			(*s)[j+1] = (*s)[j]
			j--
		}
		(*s)[j+1] = key
	}
}
