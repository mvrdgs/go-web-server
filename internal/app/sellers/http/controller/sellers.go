package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mvrdgs/go-web-server/pkg/web"
	"log"
	"net/http"
)

func (s *SellerHandler) GetAllSellers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sellers, code, err := s.sellerService.GetAllSellers(ctx)
		if err != nil {
			ctx.JSON(code, web.NewResponse(code, nil, err.Error()))
			return
		}

		ctx.JSON(code, web.NewResponse(code, sellers, ""))
	}
}

func (s *SellerHandler) GetOneSeller() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "invalid id"))
			return
		}

		seller, code, err := s.sellerService.GetOneSeller(ctx, id)
		if err != nil {
			ctx.JSON(code, web.NewResponse(code, nil, err.Error()))
			return
		}

		ctx.JSON(code, web.NewResponse(code, seller, ""))
	}
}

func (s *SellerHandler) CreateSeller() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req postSellerRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(
				http.StatusUnprocessableEntity, nil, "all fields must be correctly filled",
			))
			return
		}

		seller, code, err := s.sellerService.CreateSeller(ctx, req.CID, req.CompanyName, req.Address, req.Telephone)
		if err != nil {
			ctx.JSON(code, web.NewResponse(code, nil, err.Error()))
			return
		}

		ctx.JSON(code, web.NewResponse(code, seller, ""))
	}
}

func (s *SellerHandler) UpdateSeller() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req updateSellerRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(
				http.StatusUnprocessableEntity, nil, "all fields must be correctly filled",
			))
			return
		}

		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "invalid id"))
			return
		}

		seller, code, err := s.sellerService.UpdateSeller(ctx, id, req.CID, req.CompanyName, req.Address, req.Telephone)
		if err != nil {
			ctx.JSON(code, web.NewResponse(code, nil, err.Error()))
			return
		}

		ctx.JSON(code, web.NewResponse(code, seller, ""))
	}
}

func (s *SellerHandler) DeleteSeller() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "invalid id"))
			return
		}

		code, err := s.sellerService.DeleteSeller(ctx, id)
		if err != nil {
			ctx.JSON(code, web.NewResponse(code, nil, err.Error()))
			return
		}

		ctx.JSON(code, web.NewResponse(code, nil, ""))
	}
}
