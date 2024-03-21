package storage

import (
	"errors"

	"github.com/lib/pq"
)

var ErrNoResult = errors.New("no result")

type Product struct {
	Order            int64
	NameProduct      string
	IdProduct        int64
	CountProduct     int64
	MainShelf        string
	AdditionalSheves pq.StringArray
}
