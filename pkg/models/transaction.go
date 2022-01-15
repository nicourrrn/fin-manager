package models

import "time"

type Transaction struct {
	Id          int64
	Total       int
	Category    *string
	Description string
	User        *User
	CreatedAd   time.Time
}

func NewTransaction(total int, category *string, desc string) *Transaction {
	return &Transaction{Total: total, Category: category, Description: desc}
}
