package dto

// TODO: changer le nombre de charactères requis pour le mot de passe minimum à 9
type LoginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=5"`
}

// // AuthTokenResponse est renvoyé après une inscription réussie (ou un login)
// // Il contient le JWT et sa durée de validité
// type RefreshTokenResponse struct {
// 	RefreshTokenToken string `json:"refresh_token"`
// 	TokenType         string `json:"token_type"` // toujours "bearer"
// 	ExpiresIn         int64  `json:"expires_in"` // 5 jours -> en secondes
// }
