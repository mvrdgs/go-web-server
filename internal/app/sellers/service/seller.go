package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/mvrdgs/go-web-server/internal/app/sellers/domain"
	"log"
	"net/http"
	"strings"
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
	seller := domain.Seller{
		ID:          id,
		CID:         cid,
		CompanyName: companyName,
		Address:     address,
		Telephone:   telephone,
	}

	seller, err := s.repository.CreateSeller(ctx, seller)
	if err != nil {
		log.Println(err.Error())
		return domain.Seller{}, http.StatusConflict, errors.New("cid already registered")
	}

	return seller, http.StatusCreated, nil
}

func (s sellerService) UpdateSeller(ctx context.Context, id, cid uuid.UUID, companyName, address, telephone string) (domain.Seller, int, error) {
	seller, err := s.repository.GetOneSeller(ctx, id)
	if err != nil {
		log.Println(err.Error())
		return seller, http.StatusNotFound, errors.New("seller not found")
	}

	if cid != uuid.Nil {
		seller.CID = cid
	}

	if strings.Trim(companyName, " ") != "" {
		seller.CompanyName = companyName
	}

	if strings.Trim(address, " ") != "" {
		seller.Address = address
	}

	if strings.Trim(telephone, " ") != "" {
		seller.Telephone = telephone
	}

	seller, err = s.repository.UpdateSeller(ctx, seller)
	if err != nil {
		log.Println(err.Error())
		return seller, http.StatusBadRequest, errors.New("could not update seller")
	}

	return seller, http.StatusOK, nil
}

func (s sellerService) DeleteSeller(ctx context.Context, uuid uuid.UUID) (int, error) {
	err := s.repository.DeleteSeller(ctx, uuid)
	if err != nil {
		log.Println(err.Error())
		return http.StatusNotFound, errors.New("seller not found")
	}

	return http.StatusNoContent, nil
}

func NewSellerService(sr domain.SellerRepository) domain.SellerService {
	return &sellerService{
		repository: sr,
	}
}
