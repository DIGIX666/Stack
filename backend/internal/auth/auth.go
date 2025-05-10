package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/DIGIX666/stack/backend/interfaces/dto"
	"github.com/DIGIX666/stack/backend/internal/security"
	"github.com/DIGIX666/stack/backend/models"
)

var ErrEmailAlreadyUsed = errors.New("email déjà utilisé")

type AuthService struct {
	db         *gorm.DB
	jwtManager *security.Service
}

func NewAuthService(db *gorm.DB, secretKey string, expiration time.Duration) *AuthService {
	jwtMgr := security.JWTManager(secretKey, db, expiration)
	return &AuthService{db: db, jwtManager: jwtMgr}
}

// Signup crée un utilisateur, hash son mot de passe, enregistre en base,
// puis génère et renvoie un JWT.
func (s *AuthService) Signup(req dto.SignupRequest) (dto.AuthTokenResponse, error) {
	// 1. Vérifier unicité de l’email
	var count int64
	if err := s.db.Model(&models.User{}).
		Where("email = ?", req.Email).
		Count(&count).Error; err != nil {
		return dto.AuthTokenResponse{}, err
	}
	if count > 0 {
		return dto.AuthTokenResponse{}, ErrEmailAlreadyUsed
	}

	// 2. Hash du mot de passe
	hash, err := security.HashPassword(req.Password)
	if err != nil {
		return dto.AuthTokenResponse{}, err
	}

	// 3. Création du modèle utilisateur
	user := models.User{
		ID:           uuid.New(),
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: hash,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	if err := s.db.Create(&user).Error; err != nil {
		return dto.AuthTokenResponse{}, err
	}

	// 4. Génération du JWT
	token, expiresIn, err := s.jwtManager.GenerateToken(&user)
	if err != nil {
		return dto.AuthTokenResponse{}, err
	}

	// 5. Retourner le DTO
	return dto.AuthTokenResponse{
		AccessToken: token,
		TokenType:   "bearer",
		ExpiresIn:   expiresIn,
	}, nil
}
