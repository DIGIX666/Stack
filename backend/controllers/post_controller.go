package controllers

import (
	"github.com/DIGIX666/stack/backend/models"
	"gorm.io/gorm"
)

type PostController struct {
	repo models.PostRepo // Repositories pour les opérations CRUD sur les posts
}

// NewAuthController initialise l’AuthController avec ses dépendances
func NewPostController(db *gorm.DB) *PostController {
	// 1. Instancie le repo GORM
	postRepo := models.NewGormPostRepo(db)
	// 2. Retourne le controller prêt à l’emploi
	return &PostController{repo: postRepo}
}
