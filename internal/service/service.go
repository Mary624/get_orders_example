package service

import (
	"strconv"
	"strings"

	"github.com/lib/pq"
)

func GetStringOrders(orders []int64) string {
	var b strings.Builder
	for i, order := range orders {
		b.WriteString(strconv.Itoa(int(order)))
		if i != len(orders)-1 {
			b.WriteString(", ")
		}
	}
	return b.String()
}

func GetStringAdditionalSheves(shelves pq.StringArray) string {
	var b strings.Builder
	for i, shelf := range shelves {
		b.WriteString(shelf)
		if i != len(shelves)-1 {
			b.WriteString(", ")
		}
	}
	return b.String()
}
