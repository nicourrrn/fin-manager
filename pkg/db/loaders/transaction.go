package loaders

import (
	"github.com/nicourrrn/fin-manager/pkg/db"
	"github.com/nicourrrn/fin-manager/pkg/db/models"
	"time"
)

type TransactionRepo struct {
	cache map[int64]struct {
		Transaction *models.Transaction
		RmTime      time.Time
	}
	cacheLifeTime time.Duration
}

func NewTransactionRepo(cacheLifeTime time.Duration) *TransactionRepo {
	cache := make(map[int64]struct {
		Transaction *models.Transaction
		RmTime      time.Time
	})
	return &TransactionRepo{cache: cache, cacheLifeTime: cacheLifeTime}
}

func (t TransactionRepo) LoadTransaction(c *db.Connection, categoryRepo *CategoryRepo, id int64) (*models.Transaction, error) {
	if cachedTransaction, ok := t.cache[id]; ok {
		return cachedTransaction.Transaction, nil
	}
	row := c.QueryRow("SELECT id, total, UNIX_TIMESTAMP(created_at), category_id "+
		"FROM transactions WHERE id = ?", id)
	transaction := new(models.Transaction)
	var (
		createAt, categoryId int64
	)
	err := row.Scan(&transaction.Id, &transaction.Total, &createAt, &categoryId)
	transaction.Category = categoryRepo.GetCategorie(categoryId)
	transaction.CreatedAd = time.Unix(createAt, 0)
	if err != nil {
		return nil, err
	}
	t.cache[transaction.Id] = struct {
		Transaction *models.Transaction
		RmTime      time.Time
	}{Transaction: transaction, RmTime: time.Now().Add(t.cacheLifeTime)}
	return transaction, nil
}
