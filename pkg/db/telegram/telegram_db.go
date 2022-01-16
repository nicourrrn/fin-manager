package telegram

import "github.com/nicourrrn/fin-manager/pkg/db"

type TgSession struct {
	Id            int64
	Status, Login string
}

func LoadTelegramSessions(db *db.Connection, telegramId int64) (TgSession, error) {
	row := db.QueryRow("SELECT status, user_login FROM telegram_sessions WHERE telegram_id = ?", telegramId)
	session := new(TgSession)
	err := row.Scan(&session.Status, &session.Login)
	if err != nil {
		return TgSession{}, err
	}
	session.Id = telegramId
	return *session, nil
}

func AddTelegramSession(db *db.Connection, session TgSession) error {
	_, err := db.Exec("INSERT INTO telegram_sessions(telegram_id, status, user_login) VALUE (?, ?, ?)",
		session.Id, session.Status, session.Login)
	if err != nil {
		return err
	}
	return nil
}
