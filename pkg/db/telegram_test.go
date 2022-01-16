package db_test

import (
	"github.com/nicourrrn/fin-manager/pkg/db"
	"github.com/nicourrrn/fin-manager/pkg/db/telegram"
	"testing"
)

func Test(t *testing.T) {
	conn, err := db.NewConnection()
	if err != nil {
		t.Error(err)
		return
	}
	err = telegram.AddTelegramSession(conn, telegram.TgSession{
		Id:     123,
		Login:  "text",
		Status: "",
	})
	if err != nil {
		t.Error(err)
		return
	}

	session, err := telegram.LoadTelegramSessions(conn, 123)
	if err != nil {
		t.Error(err)
		return
	}
	if session.Login != "text" {
		t.Error("logins not equals")
	}
}
