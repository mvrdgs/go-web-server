package domain

import "context"

type Seller struct {
	ID          int    `db:"id" json:"id"`
	CID         int    `db:"cid" json:"cid"`
	CompanyName string `db:"company_name" json:"company_name"`
	Address     string `db:"address" json:"address"`
	Telephone   string `db:"telephone" json:"telephone"`
}

type SellerRepository interface {
	GetAllSellers(context.Context) ([]Seller, error)
	GetOneSeller(context.Context, int) (Seller, error)
	CreateSeller(context.Context, Seller) (Seller, error)
	UpdateSeller(context.Context, Seller) (Seller, error)
	DeleteSeller(context.Context, int) error
}

type SellerService interface {
	GetAllSeller(context.Context) ([]Seller, error)
	GetOneSeller(context.Context, int) (Seller, error)
	CreateSeller(ctx context.Context, id, cid int, companyName, address, telephone string) (Seller, error)
	UpdateSeller(ctx context.Context, id, cid int, companyName, address, telephone string) (Seller, error)
	DeleteSeller(context.Context, int) error
}
