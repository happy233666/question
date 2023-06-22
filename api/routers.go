package api

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	v := r.Group("/users")
	{
		v.POST("/register", Register)
		v.POST("/login", Login)
		v.PUT("/changePassword", ChangePassword)
	}
	m := r.Group("/massage")
	{
		m.POST("/massage", Sendmassage)
		m.PUT("/massage", ChangeMassage)
		m.DELETE("/massage", DeleteMassage)

	}
	c := r.Group("/massage")
	{
		c.POST("/comment", SendComment)
		c.PUT("/comment", UpdateComment)
		c.DELETE("/comment", DeleteComment)
	}
	r.Run()
}
