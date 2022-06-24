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

	rows, err := m.db.QueryContext(ctx, getAll)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

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

func (m mysqlRepository) GetOneSeller(ctx context.Context, id int) (domain.Seller, error) {
	var seller domain.Seller

	rows, err := m.db.QueryContext(ctx, getOne, id)
	if err != nil {
		log.Println(err)
		return seller, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&seller.ID, &seller.CID, &seller.CompanyName, &seller.Address, &seller.Telephone)
		if err != nil {
			log.Fatalln(err.Error())
			return seller, err
		}
	}

	return seller, nil
}

func (m mysqlRepository) CreateSeller(ctx context.Context, seller domain.Seller) (domain.Seller, error) {
	stmt, err := m.db.Prepare(create)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, seller.ID, seller.CID, seller.CompanyName, seller.Address, seller.Telephone)
	if err != nil {
		log.Println(err.Error())
		return domain.Seller{}, err
	}

	return seller, nil
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
