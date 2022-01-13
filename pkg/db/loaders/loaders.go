package loaders

import (
	"github.com/nicourrrn/fin-manager/pkg/db"
	"github.com/nicourrrn/fin-manager/pkg/db/models"
	"time"
)

// TODO Добавить сборщик мусора для простроченых велечин
type UserRepo struct {
	cache map[int64]struct{
		User *models.User
		RmTime time.Time
	}
	cacheLifeTime time.Duration
}

func NewUserRepo(cacheLifeTime time.Duration) *UserRepo {
	cache := make(map[int64]struct{User *models.User
		RmTime time.Time})
	return &UserRepo{cache: cache, cacheLifeTime: cacheLifeTime}
}

func (u UserRepo)LoadUser(c db.Connection, id int64) (*models.User, error){
	if cachedUser, ok := u.cache[id]; ok {
		return cachedUser.User, nil
	}
	row := c.QueryRow("SELECT id, login, money, for_days, money_per_day, now_money, created_at" +
		" FROM users WHERE id = ?", id)
	var user *models.User
	err := row.Scan(user)
	if err != nil {
		return nil, err
	}
	u.cache[user.Id] = struct {
		User   *models.User
		RmTime time.Time
	}{User: user, RmTime: time.Now().Add(u.cacheLifeTime) }
	return user, nil
}
func (u UserRepo) AddUser(c db.Connection, user *models.User) (int64, error) {
	result, err := c.Exec("INSERT INTO users (login, money, for_days, money_per_day, now_money) VALUE (?, ?, ?, ?, ?)",
		user.Login, user.Money, user.ForDays, user.PerDayMoney, user.NowMoney)
	if err != nil {
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	user.Id = lastId
	u.cache[lastId] = struct {
		User   *models.User
		RmTime time.Time
	}{User: user, RmTime: time.Now().Add(u.cacheLifeTime) }
	return lastId, nil
}
