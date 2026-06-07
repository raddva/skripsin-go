package utils

import (
	. "skripsin/models"
	"strings"
)

// Helper Functions
func SequentialContains(text, keyword string) bool {
	text = strings.ToLower(text)
	keyword = strings.ToLower(keyword)
	nText := len(text)
	nKey := len(keyword)
	if nKey > nText {
		return false
	}

	for i := 0; i <= nText-nKey; i++ {
		match := true
		for j := 0; j < nKey; j++ {
			if text[i+j] != keyword[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func BinaryContains(s SkripsiList, n int, keyword string, searchBy string) int {
	keyword = strings.ToLower(keyword)
	left := 0
	right := n - 1

	for left <= right {
		mid := (left + right) / 2
		var midValue string
		switch searchBy {
		case "name":
			midValue = s[mid].Author.Name
		case "title":
			midValue = s[mid].Title
		}

		midValue = strings.ToLower(midValue)
		if midValue == keyword {
			return mid
		}

		if midValue < keyword {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func CompareSkripsi(a, b Skripsi, sortBy, sortType string) bool {
	isAsc := sortType == "asc"
	switch sortBy {
	case "name":
		nameA := strings.ToLower(a.Author.Name)
		nameB := strings.ToLower(b.Author.Name)
		if isAsc {
			return nameA < nameB
		}
		return nameA > nameB
	case "year":
		if isAsc {
			return a.Year < b.Year
		}
		return a.Year > b.Year
	case "title":
		titleA := strings.ToLower(a.Title)
		titleB := strings.ToLower(b.Title)
		if isAsc {
			return titleA < titleB
		}
		return titleA > titleB
	}
	return false
}

func GetStatusString(isGraduated bool) string {
	if isGraduated {
		return "Lulus"
	}
	return "Belum"
}

// End of Helper Functions
