package security

import (
	"fmt"
	"time"

	"github.com/DIGIX666/stack/backend/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	secretKey  string
	db         *gorm.DB
	expiration time.Duration
}

func JWTManager(secretKey string, db *gorm.DB, expiration time.Duration) *Service {
	return &Service{
		secretKey:  secretKey,
		db:         db,
		expiration: expiration,
	}
}

// CreateToken génère un token JWT pour l'utilisateur donné
// Le token contient l'ID de l'utilisateur et son rôle
// Le token expire après une durée définie (environ 10 minutes)
// Le token est signé avec la clé secrète
// Le token est retourné sous forme de chaîne
// Si une erreur se produit lors de la création du token, elle est retournée
// Le token est utilisé pour authentifier l'utilisateur lors des requêtes ultérieures
// Le token est envoyé au client et doit être stocké de manière sécurisée
// Le token est envoyé dans l'en-tête Authorization des requêtes HTTP
func (s *Service) GenerateToken(user *models.User) (string, int64, error) {
	// Set expiration time
	expirationTime := time.Now().Add(s.expiration).Unix()
	claims := jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  expirationTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", 0, err
	}

	return signedToken, expirationTime, nil
}

// GenerateRefreshToken génère un token JWT pour l'utilisateur donné
// Le token contient l'ID de l'utilisateur et son rôle
// Le token expire après une durée définie (environ 5 jours)
// Le token est signé avec la clé secrète
// Le token est retourné sous forme de chaîne
// Si une erreur se produit lors de la création du token, elle est retournée
// Le token est utilisé pour authentifier l'utilisateur lors des requêtes ultérieures
// Le token est envoyé au client et doit être stocké de manière sécurisée
// Le token est envoyé dans l'en-tête Authorization des requêtes HTTP
// Le token est signé avec la clé secrète
func (s *Service) GenerateRefreshToken(user *models.User) (string, int64, error) {
	// Set expiration time
	expirationTime := time.Now().Add(5 * 24 * time.Hour).Unix() // 5 jours
	claims := jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  expirationTime, // 5 jours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", -1, err
	}
	return signedToken, expirationTime, nil
}

// ValidateToken vérifie la validité du token JWT
// Si le token est valide, il retourne l'ID de l'utilisateur et son rôle
// Si le token est invalide, il retourne une erreur
// Le token est utilisé pour authentifier l'utilisateur lors des requêtes ultérieures
// Le token est envoyé au client et doit être stocké de manière sécurisée
// Le token est envoyé dans l'en-tête Authorization des requêtes HTTP
// Le token est signé avec la clé secrète
func (s *Service) ValidateToken(tokenString string) (uint, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return 0, "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["sub"].(float64))
		role := claims["role"].(string)
		return userID, role, nil
	}

	return 0, "", fmt.Errorf("invalid token")
}

// RefreshToken génère un nouveau token JWT pour l'utilisateur donné
// Le nouveau token contient l'ID de l'utilisateur et son rôle
// Le nouveau token expire après une durée définie (environ 5 jours)
// Le nouveau token est signé avec la clé secrète
// Le nouveau token est retourné sous forme de chaîne
// Si une erreur se produit lors de la création du nouveau token, elle est retournée
// Le nouveau token est utilisé pour authentifier l'utilisateur lors des requêtes ultérieures
// Le nouveau token est envoyé au client et doit être stocké de manière sécurisée
// Le nouveau token est envoyé dans l'en-tête Authorization des requêtes HTTP
func (s *Service) RefreshToken(tokenString string) (string, int64, error) {
	userID, role, err := s.ValidateToken(tokenString)
	if err != nil {
		return "", -1, err
	}

	//TODO: check if returning the expire time is necessary
	newToken, expirTime, err := s.GenerateRefreshToken(&models.User{
		ID:   uuid.MustParse(fmt.Sprintf("%d", userID)),
		Role: role,
	})
	if err != nil {
		return "", expirTime, err
	}

	return newToken, expirTime, nil
}
func (s *Service) InvalidateToken(tokenString string) error {
	// Invalidate the token by adding it to a blacklist or similar mechanism
	// This is a placeholder implementation
	return nil
}
