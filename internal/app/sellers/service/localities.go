package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/mvrdgs/go-web-server/internal/app/sellers/domain"
	"log"
	"net/http"
)

func (s *sellerService) CreateLocality(
	ctx context.Context, id, localityName, provinceName, countryName string,
) (domain.Locality, int, error) {
	locality := domain.Locality{
		ID:           id,
		LocalityName: localityName,
		ProvinceName: provinceName,
		CountryName:  countryName,
	}

	provinceID := uuid.New()

	locality, err := s.repository.CreateLocality(ctx, locality, provinceID)
	if err != nil {
		log.Println(err.Error())
		code, err := errorHandler(err)
		return domain.Locality{}, code, err
	}

	return locality, http.StatusCreated, nil
}
