package models

import "time"

type User struct {
	Id          int64
	Login       string
	Money       int
	ForDays     int `db: "for_days"`
	NowMoney    int `db: "now_money"`
	PerDayMoney int `db: "per_day_money"`
	CreatedAd   time.Time
}
