package mysql

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
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

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(rows)

	for rows.Next() {
		var seller domain.Seller
		err = rows.Scan(&seller.ID, &seller.CID, &seller.CompanyName, &seller.Address, &seller.Telephone)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		sellers = append(sellers, seller)
	}

	return sellers, nil
}

func (m mysqlRepository) GetOneSeller(ctx context.Context, id uuid.UUID) (domain.Seller, error) {
	var seller domain.Seller

	rows, err := m.db.QueryContext(ctx, getOne, id)
	if err != nil {
		log.Println(err)
		return seller, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(rows)

	for rows.Next() {
		err = rows.Scan(&seller.ID, &seller.CID, &seller.CompanyName, &seller.Address, &seller.Telephone)
		if err != nil {
			log.Println(err.Error())
			return seller, err
		}
	}

	return seller, nil
}

func (m mysqlRepository) CreateSeller(ctx context.Context, seller domain.Seller) (domain.Seller, error) {
	stmt, err := m.db.PrepareContext(ctx, create)
	if err != nil {
		log.Println(err.Error())
		return seller, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(stmt)

	_, err = stmt.ExecContext(ctx, seller.ID, seller.CID, seller.CompanyName, seller.Address, seller.Telephone)
	if err != nil {
		log.Println(err.Error())
		return domain.Seller{}, err
	}

	return seller, nil
}

func (m mysqlRepository) UpdateSeller(ctx context.Context, seller domain.Seller) (domain.Seller, error) {
	stmt, err := m.db.PrepareContext(ctx, update)
	if err != nil {
		log.Println(err.Error())
		return seller, err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(stmt)

	_, err = stmt.ExecContext(ctx, seller.ID, seller.CID, seller.CompanyName, seller.Address, seller.Telephone)
	if err != nil {
		return domain.Seller{}, err
	}

	return seller, nil
}

func (m mysqlRepository) DeleteSeller(ctx context.Context, id uuid.UUID) error {
	stmt, err := m.db.PrepareContext(ctx, delete)
	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(stmt)

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func NewMysqlRepository(db *sql.DB) domain.SellerRepository {
	return &mysqlRepository{
		db,
	}
}
