package service

import "github.com/mvrdgs/go-web-server/internal/app/sellers/domain"

type sellerService struct {
	repository domain.SellerRepository
}

func NewSellerService(sr domain.SellerRepository) domain.SellerService {
	return &sellerService{
		repository: sr,
	}
}
