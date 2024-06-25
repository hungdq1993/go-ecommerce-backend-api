package router

import (
	c "go-ecommerce-backend-api/internal/controller"

	g "github.com/gin-gonic/gin"
)

// MVC pattern
// Router: Controller -> Service -> Repo
func NewRouter() *g.Engine {
	r := g.Default()

	v1 := r.Group("/api/v1")

	{
		v1.GET("/ping/:id", c.NewPingController().Ping)
		v1.GET("/user/:id", c.NewUserController().GetUserById)
	}

	return r
}
