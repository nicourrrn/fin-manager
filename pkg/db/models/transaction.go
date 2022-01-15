package models

import "time"

type Transaction struct {
	Id          int64
	Total       int
	Category    *string
	Description string
	CreatedAd   time.Time
}
