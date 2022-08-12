package users

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/ekhaer/api-with-go/models"
)

type AddReq struct {
	Userid         string
	Name          string
}

func (h handler) AddUser(c *gin.Context) {
    body := AddReq{}

    // getting request's body
    if err := c.BindJSON(&body); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var user models.Users

    user.Userid = body.Userid
    user.Name = body.Name

    if result := h.DB.Create(&user); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    c.JSON(http.StatusCreated, &user)
}