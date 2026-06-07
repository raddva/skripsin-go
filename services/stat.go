package services

import (
	"fmt"
	. "skripsin/models"
)

// Prosedur untuk menampilkan statistik skripsi, seperti jumlah skripsi per tahun, jumlah skripsi yang lulus, dll
func StatisticSkripsi(s SkripsiList, n int) {
	var CountPerYear [NYear + 1]int
	var CountGraduated int
	for i := 0; i < n; i++ {
		CountPerYear[s[i].Year]++
		if s[i].Author.IsGraduated {
			CountGraduated++
		}
	}
	fmt.Println("✦===============================✦")
	fmt.Printf("%-7s %-23s %-5s\n", "✦", "STATISTIK SKRIPSI", "✦")
	fmt.Println("✦===============================✦")
	fmt.Println("Jumlah Skripsi per tahun:")
	for year := 0; year <= 2026; year++ {
		if CountPerYear[year] > 0 {
			fmt.Printf("➯ %d : %d Skripsi\n", year, CountPerYear[year])
		}
	}
	fmt.Printf("Jumlah Skripsi Lulus: %d\n", CountGraduated)
	fmt.Printf("Total Skripsi: %d\n", n)
	fmt.Println()
}
