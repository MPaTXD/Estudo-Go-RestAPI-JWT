package controllers

import (
	"example.com/estudo/database"
	"example.com/estudo/models"
	"example.com/estudo/services"
	"github.com/gin-gonic/gin"
)

func Login(l *gin.Context) {
	db := database.GetDatabase()

	var login models.Login

	json := l.ShouldBindJSON(&login)
	if json != nil {
		l.JSON(400, gin.H{
			"error": "Ocorreu um erro ao injetar o JSON: " + json.Error(),
		})
		return
	}

	var user models.User
	dbError := db.Where("email = ?", login.Email).First(&user).Error
	if dbError != nil {
		l.JSON(400, gin.H{
			"error": "User não existe",
		})
		return
	}

	if user.Senha != services.SHA256Encoder(login.Senha) {
		l.JSON(400, gin.H{
			"error": "Credenciais Incorretas",
		})
		return
	}

	token, err := services.NewJWTService().GerarToken(user.ID)
	if err != nil {
		l.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	l.JSON(200, gin.H{
		"token": token,
	})
}
