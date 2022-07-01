package mysql

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/mvrdgs/go-web-server/internal/app/sellers/domain"
	"log"
)

func (m mysqlRepository) GetCountryID(ctx context.Context, countryName string) (uuid.UUID, error) {
	var binaryID []byte

	err := m.db.QueryRowContext(ctx, getCountryID, countryName).
		Scan(&binaryID)
	if err != nil {
		return uuid.Nil, err
	}

	countryID, err := uuid.FromBytes(binaryID)
	if err != nil {
		return uuid.Nil, err
	}

	return countryID, nil
}

func (m mysqlRepository) CreateCountry(ctx context.Context, countryName string, countryID uuid.UUID) (uuid.UUID, error) {
	stmt, err := m.db.PrepareContext(ctx, createCountry)
	if err != nil {
		log.Println(err.Error())
		return uuid.Nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(stmt)

	binaryID, err := uuid.UUID.MarshalBinary(countryID)
	if err != nil {
		return uuid.Nil, err
	}

	_, err = stmt.ExecContext(ctx, binaryID, countryName)
	if err != nil {
		return uuid.Nil, err
	}

	return countryID, nil
}

func (m mysqlRepository) GetProvinceID(ctx context.Context, provinceName string, countryID uuid.UUID) (uuid.UUID, error) {
	var binaryID []byte

	binaryID, err := uuid.UUID.MarshalBinary(countryID)
	if err != nil {
		return uuid.Nil, err
	}

	err = m.db.QueryRowContext(ctx, getProvinceID, provinceName, countryID).
		Scan(&binaryID)
	if err != nil {
		return uuid.Nil, err
	}

	provinceID, err := uuid.FromBytes(binaryID)
	if err != nil {
		return uuid.Nil, err
	}

	return provinceID, nil
}

func (m mysqlRepository) CreateProvince(
	ctx context.Context, provinceName string, countryID, provinceID uuid.UUID,
) (uuid.UUID, error) {
	stmt, err := m.db.PrepareContext(ctx, createProvince)
	if err != nil {
		log.Println(err.Error())
		return uuid.Nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(stmt)

	binaryCountryID, err := uuid.UUID.MarshalBinary(countryID)
	if err != nil {
		return uuid.Nil, err
	}

	binaryProvinceID, err := uuid.UUID.MarshalBinary(countryID)
	if err != nil {
		return uuid.Nil, err
	}

	_, err = stmt.ExecContext(ctx, binaryProvinceID, provinceName, binaryCountryID)
	if err != nil {
		return uuid.Nil, err
	}

	return provinceID, nil
}

func (m mysqlRepository) CheckLocalityExists(ctx context.Context, localityID string) (bool, error) {
	var locality string

	err := m.db.QueryRowContext(ctx, checkLocalityExists, localityID).
		Scan(&locality)
	if err != nil {
		return false, err
	}

	if locality != localityID {
		return false, nil
	}

	return true, nil
}

func (m mysqlRepository) CreateLocality(
	ctx context.Context, locality domain.Locality, provinceID uuid.UUID,
) (domain.Locality, error) {
	stmt, err := m.db.PrepareContext(ctx, createLocality)
	if err != nil {
		log.Println(err.Error())
		return locality, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}(stmt)

	_, err = stmt.ExecContext(ctx, locality.ID, locality.LocalityName, provinceID)
	if err != nil {
		log.Println(err.Error())
		return domain.Locality{}, err
	}

	return locality, nil
}
