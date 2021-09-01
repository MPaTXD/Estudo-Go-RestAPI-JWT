package routes

import (
	"example.com/estudo/controllers"
	"example.com/estudo/services/middlewares"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("usuario")
		{
			user.POST("/", controllers.CreateUser)
			user.GET("/list", controllers.ListUser, middlewares.Auth())
		}
		produtos := main.Group("produtos", middlewares.Auth())
		{
			produtos.POST("/", controllers.CreateProduto)
			produtos.GET("/:id", controllers.ShowProdutoId)
			produtos.GET("/", controllers.ShowProdutos)
			produtos.PUT("/:id", controllers.EditProdutoId)
			produtos.PUT("/status", controllers.EditProdutosStatus)
			produtos.DELETE("/delete/:id", controllers.DeleteProdutoId)
			produtos.DELETE("/delete/all", controllers.DeleteProdutos)
		}
		main.POST("login", controllers.Login)
	}
	return router
}
