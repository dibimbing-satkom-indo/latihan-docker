package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	app := gin.Default()
	dsn := "root:password@tcp(db:3306)/app?parseTime=True"
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		log.Fatalln(err)
	}

	app.GET("/", func(context *gin.Context) {
		version := ""
		err := db.Raw("select version()").Scan(&version).Error
		if err != nil {
			log.Println(err)
			context.Status(http.StatusInternalServerError)
			return
		}
		context.String(http.StatusOK, version)
	})

	err = app.Run(":8080")

	if err != nil {
		log.Println(err)
	}
}
