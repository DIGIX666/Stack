// backend/controllers/auth_controller.go
package controllers

import (
	"gorm.io/gorm"

	// tes DTO de requête/réponse
	// hash, JWT…
	"github.com/DIGIX666/stack/backend/models"
)

// AuthController gère l’authentification
type AuthController struct {
	repo models.UserRepo
}

// NewAuthController initialise l’AuthController avec ses dépendances
func NewAuthController(db *gorm.DB) *AuthController {
	// 1. Instancie le repo GORM
	userRepo := models.NewGormUserRepo(db)
	// 2. Retourne le controller prêt à l’emploi
	return &AuthController{repo: userRepo}
}
