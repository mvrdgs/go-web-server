package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/mvrdgs/go-web-server/internal/app/sellers/http/controller"
	"github.com/mvrdgs/go-web-server/internal/app/sellers/repository/mysql"
	"github.com/mvrdgs/go-web-server/internal/app/sellers/service"
	"github.com/mvrdgs/go-web-server/internal/db"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load("./cmd/server/.env")
	if err != nil {
		log.Fatalln(err.Error())
	}

	db := db.CreateMySqlDB()

	r := gin.Default()

	repo := mysql.NewMysqlRepository(db)

	sellerService := service.NewSellerService(repo)
	controller.NewSellerHandler(r, sellerService)

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Pong")
	})

	err = r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
