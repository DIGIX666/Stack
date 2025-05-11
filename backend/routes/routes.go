package routes

import (
	"github.com/DIGIX666/stack/backend/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterAPIRoutes configure l'ensemble des routes de l'API REST et les associe à leurs handlers
func RegisterAPIRoutes(r *gin.Engine, db *gorm.DB) {
	// Initialisation des controllers avec leurs dépendances (DB, services...)
	authCtrl := controllers.NewAuthController(db)

	//TODO: décommenter les autres controllers quand ils seront prêts
	// userCtrl := controllers.NewUserController(db)
	// stackCtrl := controllers.NewStackController(db)
	// postCtrl := controllers.NewPostController(db)

	// Groupe principal /api
	api := r.Group("/api")

	// Routes d'authentification (non protégées)
	auth := api.Group("/auth")
	{
		auth.POST("/signup", authCtrl.Signup)
		auth.POST("/login", authCtrl.Login)
		// Pour le logout, on vérifie d'abord le JWT
		//TODO: décommenter les routes ci-dessous si nécessaire quand elles seront prêtes
		//auth.POST("/logout", middleware.Auth(), authCtrl.Logout)
		// auth.POST("/refresh", middleware.Auth(), authCtrl.Refresh)                // Rafraîchir le token
		// auth.POST("/forgot-password", middleware.Auth(), authCtrl.ForgotPassword) // Mot de passe oublié
		// auth.POST("/reset-password", middleware.Auth(), authCtrl.ResetPassword)   // Réinitialiser le mot de passe
		// auth.POST("/verify-email", middleware.Auth(), authCtrl.VerifyEmail)       // Vérifier l'email
	}

	// Routes utilisateurs (protégées)
	// 	users := api.Group("/users", middleware.Auth())
	// 	{
	// 		users.GET("/me", userCtrl.Me)      // Profil courant
	// 		users.PUT("/:id", userCtrl.Update) // Mise à jour profil
	// 	}

	// 	// Routes Stacks (protégées)
	// 	stacks := api.Group("/stacks", middleware.Auth())
	// 	{
	// 		stacks.GET("", stackCtrl.List)    // Lister toutes les stacks
	// 		stacks.POST("", stackCtrl.Create) // Créer une nouvelle stack
	// 		stacks.GET("/:id", stackCtrl.Get) // Détails d'une stack
	// 		stacks.PUT("/:id", stackCtrl.Update)
	// 		stacks.DELETE("/:id", stackCtrl.Delete)

	// 		// Sous-groupe Posts d'une stack
	// 		posts := stacks.Group("/:id/posts")
	// 		{
	// 			posts.GET("", postCtrl.ListByStack)    // Lister les posts
	// 			posts.POST("", postCtrl.Create)        // Créer un post
	// 			posts.GET("/:postId", postCtrl.Get)    // Détails d'un post
	// 			posts.PUT("/:postId", postCtrl.Update) // Mettre à jour un post
	// 			posts.DELETE("/:postId", postCtrl.Delete)
	// 		}

	// 		// Routes additionnelles pour les stacks
	// 		// Sous-groupe pour les contributeurs
	// 		contributors := stacks.Group("/:id/contributors")
	// 		{
	// 			contributors.GET("", stackCtrl.ListContributors)             // Lister les contributeurs
	// 			contributors.POST("", stackCtrl.AddContributor)              // Ajouter un contributeur
	// 			contributors.GET("/:userId", stackCtrl.GetContributor)       // Détails d'un contributeur
	// 			contributors.DELETE("/:userId", stackCtrl.RemoveContributor) // Supprimer un contributeur

	// 			// Sous-groupe pour les posts des contributeurs
	// 			contributorPosts := contributors.Group("/:userId/posts")
	// 			{
	// 				contributorPosts.GET("", stackCtrl.ListContributorPosts)       // Lister les posts d'un contributeur
	// 				contributorPosts.GET("/:postId", stackCtrl.GetContributorPost) // Détails d'un post d'un contributeur

	// 				// Sous-groupe pour les notifications des posts des contributeurs
	// 				contributorPostNotifications := contributorPosts.Group("/:postId/notifications")
	// 				{
	// 					contributorPostNotifications.GET("", stackCtrl.ListContributorPostNotifications)                     // Lister les notifications d'un post d'un contributeur
	// 					contributorPostNotifications.POST("", stackCtrl.CreateContributorPostNotification)                   // Créer une notification pour un post d'un contributeur
	// 					contributorPostNotifications.PUT("/:notificationId", stackCtrl.UpdateContributorPostNotification)    // Mettre à jour une notification de post d'un contributeur
	// 					contributorPostNotifications.DELETE("/:notificationId", stackCtrl.DeleteContributorPostNotification) // Supprimer une notification de post d'un contributeur
	// 					contributorPostNotifications.GET("/:notificationId", stackCtrl.GetContributorPostNotification)       // Détails d'une notification de post d'un contributeur
	// 					contributorPostNotifications.GET("/:notificationId/summary", stackCtrl.Summary)                      // Résumé automatique d'une notification de post d'un contributeur
	// 				}
	// 			}
	// 		}

	// 		// Sous-groupe pour les notifications
	// 		notifications := stacks.Group("/:id/notifications")
	// 		{
	// 			notifications.GET("", stackCtrl.ListNotifications)                     // Lister les notifications
	// 			notifications.POST("", stackCtrl.CreateNotification)                   // Créer une notification
	// 			notifications.PUT("/:notificationId", stackCtrl.UpdateNotification)    // Mettre à jour une notification
	// 			notifications.DELETE("/:notificationId", stackCtrl.DeleteNotification) // Supprimer une notification
	// 			notifications.GET("/:notificationId", stackCtrl.GetNotification)       // Détails d'une notification
	// 			notifications.GET("/:notificationId/summary", stackCtrl.Summary)       // Résumé automatique
	// 		}

	// 		// Sous-groupe pour les notifications des posts
	// 		postNotifications := stacks.Group("/:id/posts/:postId/notifications")
	// 		{
	// 			postNotifications.GET("", stackCtrl.ListPostNotifications)                     // Lister les notifications d'un post
	// 			postNotifications.POST("", stackCtrl.CreatePostNotification)                   // Créer une notification pour un post
	// 			postNotifications.PUT("/:notificationId", stackCtrl.UpdatePostNotification)    // Mettre à jour une notification de post
	// 			postNotifications.DELETE("/:notificationId", stackCtrl.DeletePostNotification) // Supprimer une notification de post
	// 			postNotifications.GET("/:notificationId", stackCtrl.GetPostNotification)       // Détails d'une notification de post
	// 			postNotifications.GET("/:notificationId/summary", stackCtrl.Summary)           // Résumé automatique
	// 		}

	// 		// Route pour partager une stack
	// 		stacks.POST("/:id/share", stackCtrl.Share) // Partager une stack

	// 		// Résumé automatique pour une stack
	// 		stacks.GET("/:id/summary", stackCtrl.Summary) // Résumé automatique
	// 	}
}
