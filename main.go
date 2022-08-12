package main

import (
    "github.com/gin-gonic/gin"
    "github.com/ekhaer/api-with-go/db"
    "github.com/ekhaer/api-with-go/users"
    // "github.com/spf13/viper"
	"net/http"
	"os"
	"fmt"

)

func main() {
    // viper.SetConfigFile(".env")
    // viper.ReadInConfig()

    // port := viper.Get("PORT").(string)
    // dbUrl := viper.Get("DATABASE_URL").(string)

	// dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbport := os.Getenv("DBPORT")
	username := os.Getenv("USER")
	name := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	//database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=require password=%s port=%s", host, username, name, password, dbport)


    r := gin.Default()
    h := db.Init(dbURI)

	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"DBuRL": dbURI})    
	})

	users.RegisterRoutes(r, h)

    r.Run()
}