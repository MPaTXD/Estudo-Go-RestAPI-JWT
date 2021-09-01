package middlewares

import (
	"example.com/estudo/services"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(a *gin.Context) {
		const Bearer_schema = "Bearer "
		header := a.GetHeader("Authorization")
		if header == "" {
			a.AbortWithStatus(401)
		}
		
		token := header[len(Bearer_schema):]

		if !services.NewJWTService().ValidarToken(token) {
			a.AbortWithStatus(401)
		}
	}
}
