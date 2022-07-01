package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mvrdgs/go-web-server/pkg/web"
	"log"
	"net/http"
)

func (s *SellerHandler) CreateLocality() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req postLocalityRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(
				http.StatusUnprocessableEntity, nil, "all fields must be correctly filled",
			))
			return
		}

		seller, code, err := s.sellerService.CreateLocality(ctx, req.ID, req.LocalityName, req.ProvinceName, req.CountryName)
		if err != nil {
			ctx.JSON(code, web.NewResponse(code, nil, err.Error()))
			return
		}

		ctx.JSON(code, web.NewResponse(code, seller, ""))
	}
}
