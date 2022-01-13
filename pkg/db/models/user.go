package models

import "time"

type User struct {
	Id        int
	Login     string
	Money     int
	ForDays   int `db: "for_days"`
	NowMoney  int `db: "now_money"`
	CreatedAd time.Time
}
