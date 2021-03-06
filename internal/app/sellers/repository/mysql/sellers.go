package mysql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/mvrdgs/go-web-server/internal/app/sellers/domain"
	"log"
)

func (m mysqlRepository) GetAllSellers(ctx context.Context) ([]domain.Seller, error) {
	var sellers []domain.Seller

	rows, err := m.db.QueryContext(ctx, getAllSellers)
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
		var id []byte
		err = rows.Scan(&id, &seller.CID, &seller.CompanyName, &seller.Address, &seller.Telephone)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		seller.ID, err = uuid.FromBytes(id)
		if err != nil {
			return sellers, err
		}

		sellers = append(sellers, seller)
	}

	return sellers, nil
}

func (m mysqlRepository) GetOneSeller(ctx context.Context, id uuid.UUID) (domain.Seller, error) {
	var seller domain.Seller

	binaryId, err := uuid.UUID.MarshalBinary(id)
	if err != nil {
		return seller, err
	}

	err = m.db.QueryRowContext(ctx, getOneSeller, binaryId).
		Scan(&binaryId, &seller.CID, &seller.CompanyName, &seller.Address, &seller.Telephone)
	if err != nil {
		return seller, err
	}

	seller.ID = id

	return seller, nil
}

func (m mysqlRepository) CreateSeller(ctx context.Context, seller domain.Seller) (domain.Seller, error) {
	stmt, err := m.db.PrepareContext(ctx, createSeller)
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

	binaryId, err := uuid.UUID.MarshalBinary(seller.ID)
	if err != nil {
		return seller, err
	}

	_, err = stmt.ExecContext(ctx, binaryId, seller.CID, seller.CompanyName, seller.Address, seller.Telephone)
	if err != nil {
		log.Println(err.Error())
		return domain.Seller{}, err
	}

	return seller, nil
}

func (m mysqlRepository) UpdateSeller(ctx context.Context, seller domain.Seller) (domain.Seller, error) {
	stmt, err := m.db.PrepareContext(ctx, updateSeller)
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

	binaryId, err := uuid.UUID.MarshalBinary(seller.ID)
	if err != nil {
		return seller, err
	}

	_, err = stmt.ExecContext(ctx, seller.CID, seller.CompanyName, seller.Address, seller.Telephone, binaryId)
	if err != nil {
		return domain.Seller{}, err
	}

	return seller, nil
}

func (m mysqlRepository) DeleteSeller(ctx context.Context, id uuid.UUID) error {
	binaryId, err := uuid.UUID.MarshalBinary(id)
	if err != nil {
		return err
	}

	res, err := m.db.ExecContext(ctx, deleteSeller, binaryId)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("0 rows affected")
	}

	return nil
}
