package users

import (
    "github.com/gin-gonic/gin"

    "gorm.io/gorm"
)

type handler struct {
    DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
    h := &handler{
        DB: db,
    }

    routes := r.Group("/users")
    routes.GET("/", h.GetUsers)
    routes.GET("/:userid", h.GetUser)
    routes.POST("/", h.AddUser)
}