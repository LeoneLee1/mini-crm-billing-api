package productcategoryutils

import "strings"

var ValidCategories = []string{
	"Web Development",
	"Mobile Development",
	"Infrastructure",
	"Digital Marketing",
	"Consultation",
	"Maintenance",
	"Desgin",
	"Other",
}

func IsValidCategory(category string) bool {
	if category == "" {
		return true
	}
	for _, c := range ValidCategories {
		if strings.EqualFold(c, category) {
			return true
		}
	}
	return false
}

func NormalizeCategory(category string) string {
	if category == "" {
		return ""
	}
	for _, c := range ValidCategories {
		if strings.EqualFold(c, category) {
			return c
		}
	}
	return category
}
