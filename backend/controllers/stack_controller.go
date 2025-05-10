package controllers

import (
	"github.com/DIGIX666/stack/backend/models"
	"gorm.io/gorm"
)

// AuthController gère l’authentification
type StackController struct {
	repo models.UserRepo
}

// NewAuthController initialise l’AuthController avec ses dépendances
func NewStackController(db *gorm.DB) *UserController {
	// 1. Instancie le repo GORM
	userRepo := models.NewGormUserRepo(db)
	// 2. Retourne le controller prêt à l’emploi
	return &UserController{repo: userRepo}
}
