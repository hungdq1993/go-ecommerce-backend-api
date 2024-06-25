package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	{
		v1.GET("/ping/:param", ping)
	}

	return r
}

func ping(c *gin.Context) {
	// Lấy giá trị từ query string default là "Đỗ Quốc Hưng"
	p := c.DefaultQuery("name", "Đỗ Quốc Hung")

	// Lấy giá trị từ path param
	param := c.Param("param")

	// Lấy giá trị từ query string
	q := c.Query("uid")

	// Trả về response
	c.JSON(http.StatusOK, gin.H{
		"message": "Xin chào Đỗ Quốc Hưng " + p + " " + param,
		"uid":     q,
		"user": []string{
			"Đỗ Quốc Hưng",
			"Đỗ Quốc Hưng",
			"Đỗ Quốc Hưng",
		},
	})
}
