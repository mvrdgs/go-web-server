package domain

import "context"

type Seller struct {
	ID          int    `db:"uuid" json:"id"`
	CID         int    `db:"" json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
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
