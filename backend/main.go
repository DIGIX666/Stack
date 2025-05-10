package backend

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	config "github.com/DIGIX666/stack/backend/config"
	"github.com/DIGIX666/stack/backend/routes"
)

func main() {

	config.InitDB()
	// … démarrage du serveur …
	db := config.DB
	router := gin.Default()

	// 5. Définir les routes (controllers associés)
	routes.RegisterAPIRoutes(router, db)

	// 6. Lancer le serveur
	addr := ":" + os.Getenv("SERVER_PORT")
	log.Printf("🚀 serveur démarré sur %s", addr)
	router.Run(addr)
}
