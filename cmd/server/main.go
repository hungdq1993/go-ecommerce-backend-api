// Desc: Tên 1 file main.go trong thư mục cmd/server
package main

// Thêm các thư viện cần thiết
import (
	"go-ecommerce-backend-api/internal/router"
)

func main() {
	// Khởi tạo router
	r := router.NewRouter()

	// Chạy server
	r.Run()
}
