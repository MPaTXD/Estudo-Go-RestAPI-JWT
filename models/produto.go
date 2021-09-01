package models

import (
	"time"

	"gorm.io/gorm"
)

type Produto struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Nome            string         `json:"nome"`
	Valor           float64        `json:"valor"`
	Marca           string         `json:"marca"`
	Ein             string         `json:"ein"`
	Quantidade      int            `json:"quantidade"`
	Disponibilidade bool           `json:"disponibilidade"`
	CreatedAt       time.Time      `json:"created"`
	UpdatedAt       time.Time      `json:"updated"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted"`
}
