package main

import (
    "github.com/gin-gonic/gin"
    "github.com/ekhaer/api-with-go/db"
    "github.com/ekhaer/api-with-go/users"
    "github.com/spf13/viper"

)

func main() {
    viper.SetConfigFile(".env")
    viper.ReadInConfig()

    port := viper.Get("PORT").(string)
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

    r.Run(port)
}