// backend/models/models.go
package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	ID                   uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Username             string    `gorm:"type:varchar(50);not null;uniqueIndex"`
	Email                string    `gorm:"type:varchar(100);not null;uniqueIndex"`
	PasswordHash         string    `gorm:"type:char(60);not null"` // ou varchar(100) si vous variez
	AvatarURL            *string   `gorm:"type:varchar(255)"`
	CreatedAt            time.Time `gorm:"autoCreateTime"`
	UpdatedAt            time.Time `gorm:"autoUpdateTime"`
	Role                 string    `gorm:"type:varchar(20);not null;default:'user'"`
	Session              string    `gorm:"type:varchar(255);not null;default:gen_random_uuid()"`
	IsVerifiedEmail      bool      `gorm:"type:boolean;default:false"`
	PasswordResetToken   string    `gorm:"type:varchar(255)"`
	PasswordResetExpires time.Time `gorm:"type:timestamptz"` // durée environ 15minutes
	// Role peut être 'user', 'admin', etc.
	// Vous pouvez également utiliser un enum si vous préférez
	// ou une table de référence pour les rôles

}

type Stack struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Title     string    `gorm:"type:varchar(255);not null"`
	OwnerID   uuid.UUID `gorm:"not null;index"`
	Owner     User      `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// Post représente la table posts
type Post struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	StackID     uuid.UUID      `gorm:"type:uuid;not null;index"`
	Stack       Stack          `gorm:"foreignKey:StackID;constraint:OnDelete:CASCADE"`
	AuthorID    *uuid.UUID     `gorm:"type:uuid;index"`
	Author      *User          `gorm:"foreignKey:AuthorID;constraint:OnDelete:SET NULL"`
	ContentType string         `gorm:"type:varchar(30);not null;check:content_type IN ('text','markdown','image','video')"`
	Content     string         `gorm:"type:text;not null"`
	Media       datatypes.JSON `gorm:"type:jsonb;not null;default:'[]'"`
	Checked     bool           `gorm:"type:boolean;not null;default:false"`
	OrderIndex  int            `gorm:"column:order_index;not null;default:0"`
	CreatedAt   time.Time      `gorm:"type:timestamptz;not null;default:now()"`
	UpdatedAt   time.Time      `gorm:"type:timestamptz;not null;default:now()"`
}

// Contributor représente la table contributors (relation N–N)
type Contributor struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	StackID   uuid.UUID `gorm:"type:uuid;not null;index:idx_contrib_stack_user"`
	Stack     Stack     `gorm:"foreignKey:StackID;constraint:OnDelete:CASCADE"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index:idx_contrib_stack_user"`
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `gorm:"type:timestamptz;not null;default:now()"`
	// contrainte d'unicité (stack_id, user_id)
	// GORM crée automatiquement l'index composite idx_contrib_stack_user
}

// Notification représente la table notifications
type Notification struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;index:idx_notifications_user_read,priority:1"`
	User      User       `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	StackID   *uuid.UUID `gorm:"type:uuid;index"`
	Stack     *Stack     `gorm:"foreignKey:StackID;constraint:OnDelete:SET NULL"`
	Type      string     `gorm:"type:varchar(50);not null"`
	Message   string     `gorm:"type:text;not null"`
	IsRead    bool       `gorm:"type:boolean;not null;default:false;index:idx_notifications_user_read,priority:2"`
	CreatedAt time.Time  `gorm:"type:timestamptz;not null;default:now()"`
}

// AutoMigrateModels exécute la migration GORM pour créer/mettre à jour les tables
func AutoMigrateModels(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&Stack{},
		&Post{},
		&Contributor{},
		&Notification{},
	)
}
