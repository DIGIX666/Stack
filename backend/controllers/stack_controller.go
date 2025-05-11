package controllers

import (
	"github.com/DIGIX666/stack/backend/models"
	"gorm.io/gorm"
)

// AuthController gère l’authentification
type StackController struct {
	repo models.StackRepo
}

// NewAuthController initialise l’AuthController avec ses dépendances
func NewStackController(db *gorm.DB) *StackController {
	// 1. Instancie le repo GORM
	stackRepo := models.NewGormStackRepo(db)
	// 2. Retourne le controller prêt à l’emploi
	return &StackController{repo: stackRepo}
}
