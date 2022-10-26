package repo

import (
	"context"

	"github.com/indeedhat/juniper"
	"gorm.io/gorm"
)

type baseRepo struct {
	fallbackDB *gorm.DB
}

// db gets the most appropriate db connection available
func (r baseRepo) db(ctx context.Context) *gorm.DB {
	tx := juniper.TxHandle(ctx)
	if tx == nil {
		tx = r.fallbackDB
	}

	return tx
}
