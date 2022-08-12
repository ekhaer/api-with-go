package main

import (
    "github.com/gin-gonic/gin"
    "github.com/ekhaer/api-with-go/db"
    "github.com/ekhaer/api-with-go/users"
    "github.com/spf13/viper"
	"net/http"

)

func main() {
    viper.SetConfigFile(".env")
    viper.ReadInConfig()

    // port := viper.Get("PORT").(string)
    dbUrl := viper.Get("DATABASE_URL").(string)

    r := gin.Default()
    h := db.Init(dbUrl)

	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"DBuRL": dbUrl})    
	})

	users.RegisterRoutes(r, h)

    r.Run()
}