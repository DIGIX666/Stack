// backend/controllers/auth_controller.go
package controllers

import (
	"strconv"

	"gorm.io/gorm"

	// tes DTO de requête/réponse
	// hash, JWT…
	"github.com/DIGIX666/stack/backend/internal/security"
	"github.com/DIGIX666/stack/backend/models"
	"github.com/gin-gonic/gin"
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

// Signup gère l’inscription d’un nouvel utilisateur
func (ctrl *AuthController) Signup(c *gin.Context) {
	// 1. Vérifie si l’utilisateur existe déjà
	var user models.User
	existingUser, err := ctrl.repo.FindByEmail(user.Email)
	if err != nil {
		return models.User{}, err
	}
	if existingUser != nil {
		return models.User{}, models.ErrUserAlreadyExists
	}

	// 2. Hash le mot de passe
	hashedPassword, err := security.HashPassword(user.PasswordHash)
	if err != nil {
		return models.User{}, err
	}
	user.PasswordHash = hashedPassword

	// 3. Enregistre l’utilisateur dans la base de données
	err = ctrl.repo.Create(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// Login gère la connexion d’un utilisateur
func (ctrl *AuthController) Login(email, password string) (models.User, error) {
	// 1. Vérifie si l’utilisateur existe
	user, err := ctrl.repo.FindByEmail(email)
	if err != nil {
		return models.User{}, err
	}
	if user == nil {
		return models.User{}, models.ErrUserNotFound
	}

	// 2. Vérifie le mot de passe
	passwordCheck, err := security.VerifyPassword(user.PasswordHash, password)
	if !passwordCheck {
		return models.User{}, models.ErrInvalidPassword
	} else if err != nil {
		return models.User{}, err
	}

	return *user, nil
}

// Logout gère la déconnexion d’un utilisateur
func (ctrl *AuthController) Logout(userID string) error {
	// 1. Vérifie si l’utilisateur existe
	strconvUserID, err := strconv.Atoi(userID)
	user, err := ctrl.repo.FindByID(strconvUserID)
	if err != nil {
		return err
	}
	if user == nil {
		return models.ErrUserNotFound
	}
	// 2. Invalide le token JWT (si nécessaire)
	// 3. Autres opérations de déconnexion (si nécessaire)
	return nil
}

// Me retourne les informations de l’utilisateur courant
