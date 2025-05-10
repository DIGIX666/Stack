package config_test

import (
	"testing"

	testconfig "github.com/DIGIX666/stack/backend/config"
	models "github.com/DIGIX666/stack/backend/models"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
)

func TestInitDBWithDialector_Success(t *testing.T) {
	// On utilise SQLite en mémoire pour isoler le test
	db, err := testconfig.InitDBWithDialector(sqlite.Open("file::memory:?cache=shared"))
	require.NoError(t, err)
	require.NotNil(t, db)

	// Vérifie que la table User existe après migration
	hasUser := db.Migrator().HasTable(&models.User{})
	require.True(t, hasUser, "la table \"users\" doit être créée")

	// Vérifie que la table Stack existe après migration
	hasStack := db.Migrator().HasTable(&models.Stack{})
	require.True(t, hasStack, "la table \"stacks\" doit être créée")
	// Vérifie que la table Post existe après migration
	hasPost := db.Migrator().HasTable(&models.Post{})
	require.True(t, hasPost, "la table \"posts\" doit être créée")
	// Vérifie que la table Contributor existe après migration
	hasContributor := db.Migrator().HasTable(&models.Contributor{})
	require.True(t, hasContributor, "la table \"contributors\" doit être créée")
	// Vérifie que la table Notification existe après migration
	hasNotification := db.Migrator().HasTable(&models.Notification{})
	require.True(t, hasNotification, "la table \"notifications\" doit être créée")

}

func TestInitDBWithDialector_Failure(t *testing.T) {
	// Un dialector nil doit retourner une erreur
	_, err := testconfig.InitDBWithDialector(nil)
	require.Error(t, err)
}
