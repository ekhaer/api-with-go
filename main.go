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
    dbUrl := viper.Get("DBURL").(string)

    r := gin.Default()
    h := db.Init(dbUrl)

    // r.GET("/", func(c *gin.Context) {
    //     c.JSON(200, gin.H{
    //         "port": port,
    //         "dbUrl": dbUrl,
    //     })
    // })

	users.RegisterRoutes(r, h)
	// http.ListenAndServe(":8080", router)

    r.Run()
}