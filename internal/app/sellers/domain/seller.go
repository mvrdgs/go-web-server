package domain

import (
	"context"
	"github.com/google/uuid"
)

type Seller struct {
	ID          uuid.UUID `db:"id" json:"id"`
	CID         string    `db:"cid" json:"cid"`
	CompanyName string    `db:"company_name" json:"company_name"`
	Address     string    `db:"address" json:"address"`
	Telephone   string    `db:"telephone" json:"telephone"`
}

type Locality struct {
	ID           string `json:"id"`
	LocalityName string `json:"locality_name"`
	ProvinceName string `json:"province_name"`
	CountryName  string `json:"country_name"`
}

type SellerRepository interface {
	GetAllSellers(context.Context) ([]Seller, error)
	GetOneSeller(context.Context, uuid.UUID) (Seller, error)
	CreateSeller(context.Context, Seller) (Seller, error)
	UpdateSeller(context.Context, Seller) (Seller, error)
	DeleteSeller(context.Context, uuid.UUID) error
	CreateLocality(ctx context.Context, locality Locality, provinceID uuid.UUID) (Locality, error)
}

type SellerService interface {
	GetAllSellers(context.Context) ([]Seller, int, error)
	GetOneSeller(context.Context, uuid.UUID) (Seller, int, error)
	CreateSeller(ctx context.Context, cid, companyName, address, telephone string) (Seller, int, error)
	UpdateSeller(ctx context.Context, id uuid.UUID, cid, companyName, address, telephone string) (Seller, int, error)
	DeleteSeller(context.Context, uuid.UUID) (int, error)
	CreateLocality(ctx context.Context, id, localityName, provinceName, countryName string) (Locality, int, error)
}
