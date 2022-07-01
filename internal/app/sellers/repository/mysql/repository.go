package mysql

import (
	"database/sql"
	"github.com/mvrdgs/go-web-server/internal/app/sellers/domain"
)

type mysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository(db *sql.DB) domain.SellerRepository {
	return &mysqlRepository{
		db,
	}
}
