package models

import "time"

var Categories []string

type Transaction struct {
	Id          int
	UserId      int
	Category    *string
	Total       int
	Description string
	CreatedAd   time.Time
}
