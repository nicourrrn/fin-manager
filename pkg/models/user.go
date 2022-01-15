package models

import (
	"errors"
	"time"
)

type User struct {
	Id           int64
	Login        string
	Money        uint
	ForDays      uint
	NowMoney     uint
	PerDayMoney  uint
	Transactions []*Transaction
	CreatedAd    time.Time
}

func NewUser(login string, money uint) *User {
	return &User{Login: login, Money: money}
}

func (u *User) Calc(forDays uint) error {
	if forDays == 0 {
		return errors.New("you can't live 0 days")
	}
	u.ForDays = forDays
	u.PerDayMoney = u.Money / u.ForDays
	u.NowMoney = u.PerDayMoney
	return nil
}
func (u *User) AddTransaction(trans *Transaction) *Transaction {
	trans.User = u
	u.Transactions = append(u.Transactions, trans)
	return trans
}
