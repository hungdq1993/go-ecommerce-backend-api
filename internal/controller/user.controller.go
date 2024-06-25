package controller

import (
	"go-ecommerce-backend-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Desc: Tạo 1 struct controller có tên là UserController có thuộc tính userService là con trỏ đến service.UserService
type UserController struct {
	userService *service.UserService
}

// Desc: Tạo mới 1 controller user controller và trả về con trỏ của nó (con trỏ đến struct UserController) với
// userService là con trỏ đến service.UserService

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	p := c.Param("id")
	getId := uc.userService.GetUserById(p)
	c.JSON(http.StatusOK, gin.H{
		"message":  "Get user by id" + "" + p,
		"userInfo": []string{"Alice", "Bob", "Charlie"},
		"id":       getId,
	})
}
