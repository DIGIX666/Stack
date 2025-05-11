package authentification

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/DIGIX666/stack/backend/interfaces/dto"
	"github.com/DIGIX666/stack/backend/internal/security"
	"github.com/DIGIX666/stack/backend/models"
)

var ErrEmailAlreadyUsed = models.ErrUserAlreadyExists
var ErrInvalidCredentials = models.ErrUserInvalidCredentials

type AuthService struct {
	db         *gorm.DB
	jwtManager *security.Service
}

func NewAuthService(db *gorm.DB, secretKey string, expiration time.Duration) *AuthService {
	jwtMgr := security.JWTManager(secretKey, db, expiration)
	return &AuthService{db: db, jwtManager: jwtMgr}
}

var GURepo models.GormUserRepo

// Signup crée un utilisateur, hash son mot de passe, enregistre en base,
// puis génère et renvoie un JWT.
func (s *AuthService) Signup(req dto.SignupRequest) (dto.AuthTokenResponse, error) {

	// 1. Vérifier unicité de l’email
	exists, _, err := GURepo.FindByEmail(req.Email)
	if err != nil {
		return dto.AuthTokenResponse{}, err
	}
	if exists {
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
		User:        user,
		AccessToken: token,
		TokenType:   "bearer",
		ExpiresIn:   expiresIn,
	}, nil
}

func (s *AuthService) Login(req dto.LoginRequest) (dto.AuthTokenResponse, error) {
	// 1. Vérifie si l’utilisateur existe
	exists, user, err := GURepo.FindByEmail(req.Email)
	if err != nil {
		return dto.AuthTokenResponse{}, err
	}
	if exists {
		return dto.AuthTokenResponse{}, models.ErrUserNotFound
	}

	// 2. Vérifie le mot de passe
	passwordVerified, err := security.VerifyPassword(req.Password, user.PasswordHash)
	if err != nil {
		return dto.AuthTokenResponse{}, err
	}
	if !passwordVerified {
		return dto.AuthTokenResponse{}, models.ErrInvalidPassword
	}

	// 3. Check if the Session is refresh token is still valid

	token, expiresIn, err := s.jwtManager.GenerateRefreshToken(user)
	if err != nil {
		return dto.AuthTokenResponse{}, err
	}

	// 4. Retourner le DTO
	return dto.AuthTokenResponse{
		User:        *user,
		AccessToken: token,
		TokenType:   "bearer",
		ExpiresIn:   expiresIn,
	}, nil
}

// Logout gère la déconnexion d’un utilisateur
//TODO: implémenter la déconnexion
// func (s *AuthService) Logout(userID string) error {
// 	// 1. Vérifie si l’utilisateur existe
// 	strconvUserID, err := uuid.Parse(userID)
// 	if err != nil {
// 		return models.ErrUserNotFound
// 	}
// 	user, err := GURepo.FindByID(strconvUserID)
// 	if err != nil {
// 		return err
// 	}
// 	if user == nil {
// 		return models.ErrUserNotFound
// 	}
// 	// 2. Invalide le token JWT (si nécessaire)
// 	// 3. Autres opérations de déconnexion (si nécessaire)
// 	return nil
// }
