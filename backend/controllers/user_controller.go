package controllers

import (
	"gorm.io/gorm"

	// tes DTO de requête/réponse
	// hash, JWT…

	"github.com/DIGIX666/stack/backend/models"
)

// AuthController gère l’authentification
type UserController struct {
	repo models.UserRepo
}

// NewAuthController initialise l’AuthController avec ses dépendances
func NewUserController(db *gorm.DB) *UserController {
	// 1. Instancie le repo GORM
	userRepo := models.NewGormUserRepo(db)
	// 2. Retourne le controller prêt à l’emploi
	return &UserController{repo: userRepo}
}

// Me gère la récupération du profil de l'utilisateur courant
func (ctrl *UserController) Me(userID int) (models.User, error) {
	// 1. Vérifie si l’utilisateur existe
	user, err := ctrl.repo.FindByID(userID)
	if err != nil {
		return models.User{}, err
	}
	if user == nil {
		return models.User{}, models.ErrUserNotFound
	}
	// 2. Retourne l'utilisateur
	return *user, nil
}

// Update gère la mise à jour du profil de l'utilisateur
func (ctrl *UserController) Update(userID int, user models.User) (models.User, error) {
	// 1. Vérifie si l’utilisateur existe
	existingUser, err := ctrl.repo.FindByID(userID)
	if err != nil {
		return models.User{}, err
	}
	if existingUser == nil {
		return models.User{}, models.ErrUserNotFound
	}

	// 2. Met à jour les informations de l'utilisateur
	existingUser.Username = user.Username
	existingUser.Email = user.Email

	err = ctrl.repo.Update(existingUser)
	if err != nil {
		return models.User{}, err
	}

	return *existingUser, nil
}

// Delete gère la suppression d'un utilisateur
func (ctrl *UserController) Delete(userID int) error {
	// 1. Vérifie si l’utilisateur existe
	existingUser, err := ctrl.repo.FindByID(userID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return models.ErrUserNotFound
	}

	// 2. Supprime l'utilisateur
	err = ctrl.repo.Delete(userID)
	if err != nil {
		return err
	}

	return nil
}

// List gère la récupération de tous les utilisateurs
func (ctrl *UserController) List() ([]*models.User, error) {
	// 1. Récupère tous les utilisateurs
	users, err := ctrl.repo.List()
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Count gère le comptage des utilisateurs
func (ctrl *UserController) Count() (int64, error) {
	// 1. Compte le nombre d'utilisateurs
	count, err := ctrl.repo.Count()
	if err != nil {
		return 0, err
	}

	return count, nil
}

// FindByEmail gère la recherche d'un utilisateur par son email
func (ctrl *UserController) FindByEmail(email string) (models.User, error) {
	// 1. Vérifie si l’utilisateur existe
	user, err := ctrl.repo.FindByEmail(email)
	if err != nil {
		return models.User{}, err
	}
	if user == nil {
		return models.User{}, models.ErrUserNotFound
	}

	return *user, nil
}

// FindByID gère la recherche d'un utilisateur par son ID
func (ctrl *UserController) FindByID(userID int) (models.User, error) {
	// 1. Vérifie si l’utilisateur existe
	user, err := ctrl.repo.FindByID(userID)
	if err != nil {
		return models.User{}, err
	}
	if user == nil {
		return models.User{}, models.ErrUserNotFound
	}
	return *user, nil
}

// FindByUsername gère la recherche d'un utilisateur par son nom d'utilisateur
func (ctrl *UserController) FindByUsername(username string) (models.User, error) {
	// 1. Vérifie si l’utilisateur existe
	user, err := ctrl.repo.FindByUsername(username)
	if err != nil {
		return models.User{}, err
	}
	if user == nil {
		return models.User{}, models.ErrUserNotFound
	}

	return *user, nil
}

// FindByEmailOrUsername gère la recherche d'un utilisateur par son email ou son nom d'utilisateur
func (ctrl *UserController) FindByEmailOrUsername(email, username string) (models.User, error) {
	// 1. Vérifie si l’utilisateur existe
	user, err := ctrl.repo.FindByEmailOrUsername(email, username)
	if err != nil {
		return models.User{}, err
	}
	if user == nil {
		return models.User{}, models.ErrUserNotFound
	}

	return *user, nil
}
