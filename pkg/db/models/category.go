package models

import "github.com/nicourrrn/fin-manager/pkg/db"

type CategoryRepo struct {
	cache map[int64]string
}

func LoadCategories(c *db.Connection) (*CategoryRepo, error) {
	rows, err := c.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		id   int64
		name string
	)
	cr := CategoryRepo{
		cache: make(map[int64]string),
	}
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		cr.cache[id] = name
	}
	return &cr, nil
}
func (c CategoryRepo) GetCategorie(id int64) *string {
	name, _ := c.cache[id]
	return &name
}
