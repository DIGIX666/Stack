package models

import "gorm.io/gorm"

type PostRepo interface {
	Create(post *Post) error                    // Create crée un nouveau post
	Update(post *Post) error                    // Update met à jour un post existant
	Delete(id int) error                        // Delete supprime un post par son ID
	List() ([]*Post, error)                     // List renvoie tous les posts
	Count() (int64, error)                      // Count renvoie le nombre total de posts
	FindByID(id int) (*Post, error)             // FindByID renvoie un post par son ID
	FindByTitle(title string) (*Post, error)    // FindByTitle renvoie un post par son titre
	FindByStackID(stackID int) ([]*Post, error) // FindByStackID renvoie les posts associés à une stack
	FindByUserID(userID int) ([]*Post, error)   // FindByUserID renvoie les posts associés à un utilisateur
}

type GormPostRepo struct {
	db *gorm.DB
}

// NewGormPostRepo crée un nouveau repo GORM
func NewGormPostRepo(db *gorm.DB) PostRepo {
	return &GormPostRepo{db: db}
}
func (r *GormPostRepo) Create(post *Post) error {
	return r.db.Create(post).Error
}
func (r *GormPostRepo) Update(post *Post) error {
	return r.db.Save(post).Error
}
func (r *GormPostRepo) Delete(id int) error {
	return r.db.Delete(&Post{}, id).Error
}
func (r *GormPostRepo) List() ([]*Post, error) {
	var posts []*Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
func (r *GormPostRepo) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&Post{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
func (r *GormPostRepo) FindByID(id int) (*Post, error) {
	var post Post
	if err := r.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
func (r *GormPostRepo) FindByTitle(title string) (*Post, error) {
	var post Post
	if err := r.db.Where("title = ?", title).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
func (r *GormPostRepo) FindByStackID(stackID int) ([]*Post, error) {
	var posts []*Post
	if err := r.db.Where("stack_id = ?", stackID).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
func (r *GormPostRepo) FindByUserID(userID int) ([]*Post, error) {
	var posts []*Post
	if err := r.db.Where("user_id = ?", userID).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
