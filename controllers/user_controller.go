package controllers

import (
	"example.com/estudo/database"
	"example.com/estudo/models"
	"example.com/estudo/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(u *gin.Context) {
	db := database.GetDatabase()
	var user models.User

	json := u.ShouldBindJSON(&user)
	if json != nil {
		u.JSON(400, gin.H{
			"error": "Ocorreu um erro ao injetar o JSON: " + json.Error(),
		})
		return
	}

	user.Senha = services.SHA256Encoder(user.Senha)

	create := db.Create(&user).Error
	if create != nil {
		u.JSON(400, gin.H{
			"error": "Ocorreu um erro ao criar o user: " + create.Error(),
		})
		return
	}
	u.Status(204)
}

func ListUser(u *gin.Context) {
	db := database.GetDatabase()
	var users []models.User

	dbError := db.Find(&users).Error
	if dbError != nil {
		u.JSON(400, gin.H{
			"error": "Ocorreu um erro na consulta com o banco de dados" + dbError.Error(),
		})
		return
	}

	u.JSON(200, users)
}
