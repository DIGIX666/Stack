// backend/models/user_repo.go
package models

import (
	"log"

	"gorm.io/gorm"
)

// UserRepo définit les opérations de persistence de l’utilisateur
type UserRepo interface {
	Create(user *User) error // Create crée un nouvel utilisateur
	Update(user *User) error // Update met à jour les informations de l'utilisateur
	Delete(id int) error     // Delete supprime l'utilisateur par son ID
	List() ([]*User, error)  // List renvoie tous les utilisateurs
	Count() (int64, error)   // Count renvoie le nombre total d'utilisateurs

	FindByEmail(email string) (bool, *User, error)
	FindByID(id int) (*User, error)                              // FindByID renvoie l'utilisateur par son ID
	FindByUsername(username string) (*User, error)               // FindByUsername renvoie l'utilisateur par son nom d'utilisateur
	FindByEmailOrUsername(email, username string) (*User, error) // FindByEmailOrUsername renvoie l'utilisateur par son email ou son nom d'utilisateur

	FindByStackOwnerID(stackID int) (*User, error)                  // FindByStackOwnerID renvoie le propriétaire d'une stack
	FindByStackOwnerIDAndPostID(stackID, postID int) (*User, error) // FindByStackOwnerIDAndPostID renvoie le propriétaire d'une stack et d'un post
	CountByStackID(stackID int) (int64, error)                      // CountByStackID renvoie le nombre d'utilisateurs associés à une stack
	FindByStackID(stackID int) ([]*User, error)                     // FindByStackID renvoie les utilisateurs associés à une stack
	FindByPostID(postID int) (*User, error)                         // FindByPostID renvoie les utilisateurs associés à un post
	FindByContributorID(contributorID int) ([]*User, error)         // FindByContributorID renvoie les utilisateurs associés à un contributeur
}

type GormUserRepo struct {
	db *gorm.DB
}

// NewGormUserRepo crée un nouveau repo GORM
func NewGormUserRepo(db *gorm.DB) UserRepo {
	return &GormUserRepo{db: db}
}

func (r *GormUserRepo) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *GormUserRepo) FindByEmail(email string) (bool, *User, error) {
	var u User
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {

		log.Printf("Error finding user by email: %v", err)
		return true, nil, err
	}
	return false, &u, nil
}

// Count implements UserRepo.
func (r *GormUserRepo) Count() (int64, error) {
	panic("unimplemented")
}

// CountByStackID implements UserRepo.
func (r *GormUserRepo) CountByStackID(stackID int) (int64, error) {
	panic("unimplemented")
}

// Delete implements UserRepo.
func (r *GormUserRepo) Delete(id int) error {
	panic("unimplemented")
}

// FindByContributorID implements UserRepo.
func (r *GormUserRepo) FindByContributorID(contributorID int) ([]*User, error) {
	panic("unimplemented")
}

// FindByEmailOrUsername implements UserRepo.
func (r *GormUserRepo) FindByEmailOrUsername(email string, username string) (*User, error) {
	panic("unimplemented")
}

// FindByID implements UserRepo.
func (r *GormUserRepo) FindByID(id int) (*User, error) {
	panic("unimplemented")
}

// FindByPostID implements UserRepo.
func (r *GormUserRepo) FindByPostID(postID int) (*User, error) {
	panic("unimplemented")
}

// FindByStackID implements UserRepo.
func (r *GormUserRepo) FindByStackID(stackID int) ([]*User, error) {
	panic("unimplemented")
}

// FindByStackOwnerID implements UserRepo.
func (r *GormUserRepo) FindByStackOwnerID(stackID int) (*User, error) {
	panic("unimplemented")
}

// FindByStackOwnerIDAndPostID implements UserRepo.
func (r *GormUserRepo) FindByStackOwnerIDAndPostID(stackID int, postID int) (*User, error) {
	panic("unimplemented")
}

// FindByUsername implements UserRepo.
func (r *GormUserRepo) FindByUsername(username string) (*User, error) {
	panic("unimplemented")
}

// List implements UserRepo.
func (r *GormUserRepo) List() ([]*User, error) {
	panic("unimplemented")
}

// Update implements UserRepo.
func (r *GormUserRepo) Update(user *User) error {
	panic("unimplemented")
}
