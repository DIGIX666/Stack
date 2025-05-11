// backend/controllers/auth_controller.go
package controllers

import (
	"net/http"
	"strconv"

	"gorm.io/gorm"

	// tes DTO de requête/réponse
	// hash, JWT…

	"github.com/DIGIX666/stack/backend/interfaces/dto"
	"github.com/DIGIX666/stack/backend/internal/authentification"
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
	var authService authentification.AuthService
	var signupReq dto.SignupRequest

	if err := c.ShouldBindJSON(&signupReq); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}
	signupResponse, err := authService.Signup(signupReq)
	if err != nil {
		if err == authentification.ErrEmailAlreadyUsed {
			c.JSON(400, gin.H{"error": "Email already used"})
			return
		}
		c.JSON(500, gin.H{"error": "Internal server error: " + err.Error()})
		return
	}

	// 5. Retourne le token et les informations de l’utilisateur

	c.Header("Authorization", signupResponse.AccessToken)
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Expose-Headers", "Authorization")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

	c.JSON(http.StatusCreated, signupResponse)
}

// Login gère la connexion d’un utilisateur
func (ctrl *AuthController) Login(c *gin.Context) {
	var authService authentification.AuthService
	var loginReq dto.LoginRequest

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	loginResponse, err := authService.Login(loginReq)
	if err != nil {
		if err == authentification.ErrInvalidCredentials {
			c.JSON(401, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(500, gin.H{"error": "Internal server error: " + err.Error()})
		return
	}
	// 5. Retourne le token et les informations de l’utilisateur
	c.Header("Authorization", loginResponse.AccessToken)
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Expose-Headers", "Authorization")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

	c.JSON(http.StatusOK, loginResponse)

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
