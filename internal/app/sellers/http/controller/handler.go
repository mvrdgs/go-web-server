package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mvrdgs/go-web-server/internal/app/sellers/domain"
	"github.com/mvrdgs/go-web-server/pkg/web"
	"net/http"
)

type SellerHandler struct {
	sellerService domain.SellerService
}

//func NewSellerHandler(r *gin.Engine, s domain.SellerService) {
//	handler := SellerHandler{sellerService: s}
//
//	sg := r.Group("/api/v1/seller")
//}

func (s *SellerHandler) GetAllSellers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sellers, code, err := s.sellerService.GetAllSeller(ctx)
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
		var req request
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(
				http.StatusUnprocessableEntity, nil, "all fields must be correctly filled",
				return
			))
		}

		id, err := uuid.Parse(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "invalid id"))
			return
		}

		cid, err := uuid.Parse(req.CID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "invalid cid"))
			return
		}

		seller, code, err := s.sellerService.CreateSeller(ctx, id, cid, req.CompanyName, req.Address, req.Telephone)
		if err != nil {
			ctx.JSON(code, web.NewResponse(code, nil, err.Error()))
			return
		}

		ctx.JSON(code, web.NewResponse(code, seller, ""))
	}
}
