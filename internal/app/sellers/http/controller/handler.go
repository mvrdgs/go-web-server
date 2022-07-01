package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mvrdgs/go-web-server/internal/app/sellers/domain"
)

type SellerHandler struct {
	sellerService domain.SellerService
}

func NewSellerHandler(r *gin.Engine, s domain.SellerService) {
	handler := SellerHandler{sellerService: s}

	sg := r.Group("/api/v1/seller")
	{
		sg.GET("/", handler.GetAllSellers())
		sg.GET("/:id", handler.GetOneSeller())
		sg.POST("/", handler.CreateSeller())
		sg.PATCH("/:id", handler.UpdateSeller())
		sg.DELETE("/:id", handler.DeleteSeller())
	}
}
