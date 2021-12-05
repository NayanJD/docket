package controllers

import (
	"nayanjd/docket/models"
	"nayanjd/docket/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (ctrl *UserController) Register(c *gin.Context) {
	user := c.MustGet(gin.BindKey).(*models.User)

	err := models.GetDB().Create(&user)

	if err != nil {

	}

	utils.AbortWithGenericJson(c, utils.CreateOKResponse(&gin.H{"message": "success"}, nil), nil)
}
