package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Desc: Tạo 1 struct controller có tên là PingController không có thuộc tính nào cả
type PingController struct{}

// Desc: Tạo mới 1 controller ping controller và trả về con trỏ của nó (con trỏ đến struct PingController)
func NewPingController() *PingController {
	return &PingController{}
}

// http://localhost:8080/api/v1/ping?name=DoQuocHung
// curl -X GET "http://localhost:8080/api/v1/ping?name=DoQuocHung" -H "accept: application/json"
func (pc *PingController) Ping(c *gin.Context) {
	// Lấy giá trị từ query string default là "Đỗ Quốc Hưng"
	p := c.DefaultQuery("name", "Nguyen A")

	// Lấy giá trị từ path param
	idParam := c.Param("id")

	// Lấy giá trị từ query string
	q := c.Query("uid")

	// Trả về response
	c.JSON(http.StatusOK, gin.H{
		"message": "Xin chào Đỗ Quốc Hưng " + p + " " + idParam,
		"uid":     q,
		"user": []string{
			"Đỗ Quốc Hưng",
			"Đỗ Quốc Hưng",
			"Đỗ Quốc Hưng",
		},
	})
}
