package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/mvrdgs/go-web-server/internal/app/sellers/domain"
	"log"
	"net/http"
)

type sellerService struct {
	repository domain.SellerRepository
}

func (s sellerService) GetAllSeller(ctx context.Context) ([]domain.Seller, int, error) {
	sellers, err := s.repository.GetAllSellers(ctx)
	if err != nil {
		log.Println(err.Error())
		return sellers, http.StatusServiceUnavailable, errors.New("service unavailable, try again later")
	}

	return sellers, http.StatusOK, nil
}

func (s sellerService) GetOneSeller(ctx context.Context, uuid uuid.UUID) (domain.Seller, int, error) {
	seller, err := s.repository.GetOneSeller(ctx, uuid)
	if err != nil {
		log.Println(err.Error())
		return seller, http.StatusNotFound, errors.New("seller not found")
	}

	return seller, http.StatusOK, nil
}

func (s sellerService) CreateSeller(ctx context.Context, id, cid uuid.UUID, companyName, address, telephone string) (domain.Seller, int, error) {
	//TODO implement me
	panic("implement me")
}

func (s sellerService) UpdateSeller(ctx context.Context, id, cid uuid.UUID, companyName, address, telephone string) (domain.Seller, int, error) {
	//TODO implement me
	panic("implement me")
}

func (s sellerService) DeleteSeller(ctx context.Context, uuid uuid.UUID) (int, error) {
	//TODO implement me
	panic("implement me")
}

func NewSellerService(sr domain.SellerRepository) domain.SellerService {
	return &sellerService{
		repository: sr,
	}
}
