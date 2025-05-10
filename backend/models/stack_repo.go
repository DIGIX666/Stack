package models

import "gorm.io/gorm"

// StackRepo définit les opérations de persistence de la stack
type StackRepo interface {
	Create(stack *Stack) error                               // Create crée une nouvelle stack
	Update(stack *Stack) error                               // Update met à jour les informations de la stack
	Delete(id int) error                                     // Delete supprime la stack par son ID
	List() ([]*Stack, error)                                 // List renvoie toutes les stacks
	Count() (int64, error)                                   // Count renvoie le nombre total de stacks
	FindByID(id int) (*Stack, error)                         // FindByID renvoie la stack par son ID
	FindByTitle(title string) (*Stack, error)                // FindByTitle renvoie la stack par son titre
	FindByOwnerID(ownerID int) ([]*Stack, error)             // FindByOwnerID renvoie les stacks associées à un propriétaire
	FindByContributorID(contributorID int) ([]*Stack, error) // FindByContributorID renvoie les stacks associées à un contributeur
	FindByPostID(postID int) ([]*Stack, error)               // FindByPostID renvoie les stacks associées à un post
	FindByStackID(stackID int) ([]*Stack, error)             // FindByStackID renvoie les stacks associées à une stack
	FindByStackOwnerID(stackID int) (*Stack, error)          // FindByStackOwnerID renvoie le propriétaire d'une stack
}

type GormStackRepo struct {
	db *gorm.DB
}

// NewGormStackRepo crée un nouveau repo GORM
func NewGormStackRepo(db *gorm.DB) StackRepo {
	return &GormStackRepo{db: db}
}
func (r *GormStackRepo) Create(stack *Stack) error {
	return r.db.Create(stack).Error
}
func (r *GormStackRepo) Update(stack *Stack) error {
	return r.db.Save(stack).Error
}
func (r *GormStackRepo) Delete(id int) error {
	return r.db.Delete(&Stack{}, id).Error
}
func (r *GormStackRepo) List() ([]*Stack, error) {
	var stacks []*Stack
	if err := r.db.Find(&stacks).Error; err != nil {
		return nil, err
	}
	return stacks, nil
}

func (r *GormStackRepo) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&Stack{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
func (r *GormStackRepo) FindByID(id int) (*Stack, error) {
	var stack Stack
	if err := r.db.First(&stack, id).Error; err != nil {
		return nil, err
	}
	return &stack, nil
}
func (r *GormStackRepo) FindByTitle(title string) (*Stack, error) {
	var stack Stack
	if err := r.db.Where("title = ?", title).First(&stack).Error; err != nil {
		return nil, err
	}
	return &stack, nil
}
func (r *GormStackRepo) FindByOwnerID(ownerID int) ([]*Stack, error) {
	var stacks []*Stack
	if err := r.db.Where("owner_id = ?", ownerID).Find(&stacks).Error; err != nil {
		return nil, err
	}
	return stacks, nil
}
func (r *GormStackRepo) FindByContributorID(contributorID int) ([]*Stack, error) {
	var stacks []*Stack
	if err := r.db.Where("contributor_id = ?", contributorID).Find(&stacks).Error; err != nil {
		return nil, err
	}
	return stacks, nil
}
func (r *GormStackRepo) FindByPostID(postID int) ([]*Stack, error) {
	var stacks []*Stack
	if err := r.db.Where("post_id = ?", postID).Find(&stacks).Error; err != nil {
		return nil, err
	}
	return stacks, nil
}
func (r *GormStackRepo) FindByStackID(stackID int) ([]*Stack, error) {
	var stacks []*Stack
	if err := r.db.Where("stack_id = ?", stackID).Find(&stacks).Error; err != nil {
		return nil, err
	}
	return stacks, nil
}
func (r *GormStackRepo) FindByStackOwnerID(stackID int) (*Stack, error) {
	var stack Stack
	if err := r.db.Where("stack_id = ?", stackID).First(&stack).Error; err != nil {
		return nil, err
	}
	return &stack, nil
}
