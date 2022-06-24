package mysql

import (
	"context"
	"database/sql"
	"github.com/mvrdgs/go-web-server/internal/app/sellers/domain"
	"log"
)

type mysqlRepository struct {
	db *sql.DB
}

func (m mysqlRepository) GetAllSellers(ctx context.Context) ([]domain.Seller, error) {
	var sellers []domain.Seller

	rows, err := m.db.Query(getAll)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		var seller domain.Seller
		err = rows.Scan(&seller.ID, &seller.CID, &seller.CompanyName, &seller.Address, &seller.Telephone)
		if err != nil {
			log.Fatalln(err.Error())
			return nil, err
		}

		sellers = append(sellers, seller)
	}

	return sellers, nil
}

func (m mysqlRepository) GetOneSeller(ctx context.Context, i int) (domain.Seller, error) {
	//TODO implement me
	panic("implement me")
}

func (m mysqlRepository) CreateSeller(ctx context.Context, seller domain.Seller) (domain.Seller, error) {
	//TODO implement me
	panic("implement me")
}

func (m mysqlRepository) UpdateSeller(ctx context.Context, seller domain.Seller) (domain.Seller, error) {
	//TODO implement me
	panic("implement me")
}

func (m mysqlRepository) DeleteSeller(ctx context.Context, i int) error {
	//TODO implement me
	panic("implement me")
}

func NewMysqlRepository(db *sql.DB) domain.SellerRepository {
	return &mysqlRepository{
		db,
	}
}
