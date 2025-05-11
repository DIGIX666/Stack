package dto

import "github.com/DIGIX666/stack/backend/models"

// TODO: changer le nombre de charactères requis pour le mot de passe à 9
// SignupRequest correspond au contrat d’inscription côté frontend
type SignupRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}

// AuthTokenResponse est renvoyé après une inscription réussie (ou un login)
// Il contient le JWT et sa durée de validité
type AuthTokenResponse struct {
	User        models.User `json:"user"`
	AccessToken string      `json:"access_token"`
	TokenType   string      `json:"token_type"` // toujours "bearer"
	ExpiresIn   int64       `json:"expires_in"` // 10min -> en secondes
}
