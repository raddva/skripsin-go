package services

import (
	. "skripsin/models"
	. "skripsin/utils"
)

// Prosedur untuk mengurutkan skripsi berdasarkan nama/tahun menggunakan selection sort
func SortSkripsiSelection(s *SkripsiList, n int, sortBy string, sortType string) {
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if CompareSkripsi((*s)[j], (*s)[idx], sortBy, sortType) {
				idx = j
			}
		}
		(*s)[i], (*s)[idx] = (*s)[idx], (*s)[i]
	}
}

// Prosedur untuk mengurutkan skripsi berdasarkan nama/tahun menggunakan insertion sort
func SortSkripsiInsertion(s *SkripsiList, n int, sortBy string, sortType string) {
	for i := 1; i < n; i++ {
		key := (*s)[i]
		j := i - 1
		for j >= 0 && CompareSkripsi(key, (*s)[j], sortBy, sortType) {
			(*s)[j+1] = (*s)[j]
			j--
		}
		(*s)[j+1] = key
	}
}
