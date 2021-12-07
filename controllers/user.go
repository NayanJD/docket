package controllers

import (
	"nayanjd/docket/models"
	"nayanjd/docket/utils"

	"github.com/gin-gonic/gin"
)

type UserInputForm struct {
	ID           *string `form:"id"`
	First_name   *string `form:"first_name" binding:"required"`
	Last_name    *string `form:"last_name"  binding:"required"`
	Username     *string `form:"username"   binding:"required,email"`
	Password     *string `form:"password"   binding:"required"`
	Is_superuser *bool   `form:"-"`
	Is_staff     *bool   `form:"-"`
}

type UserController struct{}

// Register Create user
// @Summary	Create user
// @Accept	json
// @Produce json
// @Param	newUser	body	UserInputForm	true	"Create user"
// @Success	200		{object}	models.User	"Success"
// @Router	/user/register	[post]
func (ctrl *UserController) Register(c *gin.Context) {
	userForm := c.MustGet(gin.BindKey).(*UserInputForm)

	newUser := models.User{
		First_name: userForm.First_name,
		Last_name:  userForm.Last_name,
		Username:   userForm.Username,
		Password:   userForm.Password,
	}

	err := models.GetDB().Create(&newUser).Error

	if err != nil {
		c.Error(err).SetType(utils.ErrorTypeDB)
		return
	}

	utils.AbortWithGenericJson(c, utils.CreateOKResponse(newUser, nil), nil)
}

// Register Get a user
// @Summary	Get a user
// @Produce json
// @Param	id	path	int64	true	"Create user"
// @Success	200		{object}	models.User	"Success"
// @Success	204		{}	nil	"Not found"
// @Router	/user/:id	[get]
func (ctrl *UserController) GetUser(c *gin.Context) {
	userId := c.Param("id")

	user := models.User{}

	if err := models.GetDB().First(&user, "id = ?", userId).Error; err != nil {
		c.Error(err).SetType(utils.ErrorTypeDB)
		return
	}

	utils.AbortWithGenericJson(c, utils.CreateOKResponse(user, nil), nil)
}

// Register Get all users
// @Summary	Get all users
// @Produce json
// @Param	id	path	int64	true	"Create user"
// @Success	200		{object}	models.User	"Success"
// @Success	204		{}	nil	"Not found"
// @Router	/user/:id	[get]
func (ctrl *UserController) GetUsers(c *gin.Context) {

	users := []models.User{}

	if err := models.GetDB().Find(&users).Error; err != nil {
		c.Error(err).SetType(utils.ErrorTypeDB)
		return
	}

	utils.AbortWithGenericJson(c, utils.CreateOKResponse(users, nil), nil)
}
