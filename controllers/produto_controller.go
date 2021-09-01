package controllers

import (
	"fmt"
	"strconv"

	"example.com/estudo/database"
	"example.com/estudo/models"
	"github.com/gin-gonic/gin"
)

func CreateProduto(p *gin.Context) {
	db := database.GetDatabase()

	var existenteProduto []models.Produto
	var produto []models.Produto

	err := p.ShouldBindJSON(&produto)
	if err != nil {
		p.JSON(400, gin.H{
			"error": "JSON não inserido: " + err.Error(),
		})
		return
	}

	for i := range produto {
		result := db.Where("nome = ?", produto[i].Nome).Find(&existenteProduto)
		if result.Error != nil {
			p.JSON(400, gin.H{
				"error": "Ocorreu um erro na consulta " + result.Error.Error(),
			})
			return
		}
	}

	if len(existenteProduto) == 0 {
		err = db.Create(&produto).Error
		if err != nil {
			p.JSON(400, gin.H{
				"error": "Ocorreu um erro na criação do produto: " + err.Error(),
			})
			return
		}
		p.JSON(200, produto)
	} else {
		mensagem := fmt.Sprintf("Já existe um produto com o nome: [%s]", produto[0].Nome)
		p.JSON(200, gin.H{
			"Retorno": mensagem,
		})
		return
	}
}

func ShowProdutoId(p *gin.Context) {
	var produto []models.Produto
	db := database.GetDatabase()
	id := p.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		p.JSON(400, gin.H{
			"error": "ID não convertido para Int: " + err.Error(),
		})
		return
	}

	result := db.Where("id = ?", newid).Find(&produto)
	if result.Error != nil {
		p.JSON(400, gin.H{
			"error": "Ocorreu um erro na pesquisa: " + result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		mensagem := fmt.Sprintf("Não existem produto cadastrado com o id = [%d]", newid)
		p.JSON(200, gin.H{
			"Retorno": mensagem,
		})
		return
	} else {
		p.JSON(200, produto)
	}
}

func ShowProdutos(p *gin.Context) {
	db := database.GetDatabase()
	var produtos []models.Produto

	result := db.Order("id").Find(&produtos)
	if result.Error != nil {
		p.JSON(400, gin.H{
			"error": "Ocorreu um error na pesquisa " + result.Error.Error(),
		})
		return
	}

	if result.RowsAffected > 0 {
		p.JSON(200, produtos)
	} else {
		p.JSON(200, gin.H{
			"Retorno": "Nenhum produto encontrado",
		})
		return
	}
}

func EditProdutoId(p *gin.Context) {
	db := database.GetDatabase()
	var newProduto models.Produto
	var existeProduto models.Produto

	id := p.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		p.JSON(400, gin.H{
			"error": "Erro ao converter o ID para Int " + err.Error(),
		})
		return
	}

	edit := db.Where("id = ?", newId).Find(&newProduto)
	if edit.Error != nil {
		p.JSON(400, gin.H{
			"error": "Ocorreu um erro na pesquisa " + edit.Error.Error(),
		})
		return
	}

	if edit.RowsAffected > 0 {
		err := p.ShouldBindJSON(&newProduto)
		if err != nil {
			p.JSON(400, gin.H{
				"error": "JSON não inserido: " + err.Error(),
			})
			return
		}
		err = db.Where("nome = ?", newProduto.Nome).Find(&existeProduto).Error
		if err != nil {
			p.JSON(400, gin.H{
				"error": "Ocorreu um erro na pesquisa " + err.Error(),
			})
			return
		}
		if newProduto.ID == existeProduto.ID {
			err = db.Save(&newProduto).Error
			if err != nil {
				p.JSON(400, gin.H{
					"error": "Ocorreu um error em atualizar as modificações do produto: " + err.Error(),
				})
				return
			}
			p.JSON(200, newProduto)
		} else if newProduto.Nome != existeProduto.Nome {
			err = db.Save(&newProduto).Error
			if err != nil {
				p.JSON(400, gin.H{
					"error": "Ocorreu um error em atualizar as modificações do produto: " + err.Error(),
				})
				return
			}
			p.JSON(200, newProduto)
		} else {
			mensagem := fmt.Sprintf("Já existe um produto com o nome = [%s]", existeProduto.Nome)
			p.JSON(200, gin.H{
				"Retorno": mensagem,
			})
			return
		}
	} else {
		mensagem := fmt.Sprintf("Não existe produto com o id = [%d]", newId)
		p.JSON(200, gin.H{
			"Retorno": mensagem,
		})
		return
	}
}

func EditProdutosStatus(p *gin.Context) {
	db := database.GetDatabase()
	var produtos []models.Produto

	err := db.Exec("UPDATE produtos SET disponibilidade = ? WHERE quantidade > ?", true, 1).Error
	if err != nil {
		p.JSON(400, gin.H{
			"error": "Ocorreu um erro no UPDATE das disponibilidade dos produtos: " + err.Error(),
		})
		return
	}
	err = db.Exec("UPDATE produtos SET disponibilidade = ? WHERE quantidade < ?", false, 1).Error
	if err != nil {
		p.JSON(400, gin.H{
			"error": "Ocorreu um erro no UPDATE das disponibilidade dos produtos: " + err.Error(),
		})
		return
	}
	err = db.Order("id").Find(&produtos).Error
	if err != nil {
		p.JSON(400, gin.H{
			"error": "Ocorreu um erro na pesquisa: " + err.Error(),
		})
		return
	}
	p.JSON(200, produtos)
}

func DeleteProdutoId(p *gin.Context) {
	db := database.GetDatabase()
	id := p.Param("id")
	var produto models.Produto

	newId, err := strconv.Atoi(id)
	if err != nil {
		p.JSON(400, gin.H{
			"error": "ID não convertido para Int " + err.Error(),
		})
		return
	}

	result := db.Where("id = ?", newId).First(&produto)
	if result.Error != nil {
		p.JSON(400, gin.H{
			"error": "Ocorreu um erro na pesquisa: " + result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		mensagem := fmt.Sprintf("Não existe produto cadastrado com id = [%d]", newId)
		p.JSON(200, gin.H{
			"Resultado": mensagem,
		})
		return
	} else {
		err = db.Delete(&produto).Error
		if err != nil {
			p.JSON(400, gin.H{
				"error": "Error ao deletar o Produto " + err.Error(),
			})
			return
		}
		mensagem := fmt.Sprintf("O produto com id = [%d] foi deletado com sucesso", newId)
		p.JSON(200, gin.H{
			"Resultado": mensagem,
		})
		return
	}
}

func DeleteProdutos(p *gin.Context) {
	db := database.GetDatabase()
	var produtos []models.Produto

	result := db.Find(&produtos)
	if result.Error != nil {
		p.JSON(400, gin.H{
			"error": "Ocorreu um erro na consulta: " + result.Error.Error(),
		})
		return
	}
	if result.RowsAffected > 0 {
		err := db.Delete(&produtos).Error
		if err != nil {
			p.JSON(400, gin.H{
				"error": "Ocorreu um erro ao deletar os produtos: " + err.Error(),
			})
			return
		}
		p.Status(204)
		return
	} else {
		p.JSON(200, produtos)
		return
	}
}
