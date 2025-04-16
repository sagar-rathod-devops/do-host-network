package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sagar-rathod-devops/do-host-network/config"
	"github.com/sagar-rathod-devops/do-host-network/internal/controllers"
	"github.com/sagar-rathod-devops/do-host-network/internal/repositories"
	"github.com/sagar-rathod-devops/do-host-network/internal/services"
	"github.com/sagar-rathod-devops/do-host-network/middlewares"
	"github.com/sagar-rathod-devops/do-host-network/migrations"
)

func SetupServer() *gin.Engine {
	// Load configuration
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize DB connection
	db, err := config.ConnectDB(&cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Run migrations
	if err := migrations.Migrate(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// Initialize repositories
	userRepo := repositories.UserRepository{DB: db}
	otpRepo := repositories.OTPRepository{DB: db}
	userProfileRepo := repositories.NewUserProfileRepository(db)
	postRepo := repositories.NewPostRepository(db)
	commentRepo := repositories.NewCommentRepository(db)
	reactionRepo := repositories.NewReactionRepository(db)
	notificationRepo := repositories.NewNotificationRepository(db)
	connectionRepo := repositories.NewConnectionRepository(db)
	educationRepo := &repositories.EducationRepository{DB: db}
	experienceRepo := repositories.NewExperienceRepository(db)
	messageRepo := repositories.NewMessageRepository(db)
	groupRepo := repositories.NewGroupRepository(db)
	mediaRepo := repositories.NewMediaRepository(db)
	searchRepo := repositories.NewSearchRepository(db)
	adminRepo := repositories.NewAdminRepository(db)
	analyticsRepo := repositories.NewAnalyticsRepository(db)

	// Initialize services
	authService := services.AuthService{DB: db, UserRepository: userRepo, OTPRepository: otpRepo, TokenExpiration: 3600, OTPLifespan: 300}
	userProfileService := services.NewUserProfileService(*userProfileRepo)
	postService := services.NewPostService(postRepo)
	commentService := services.NewCommentService(commentRepo)
	reactionService := services.NewReactionService(reactionRepo)
	notificationService := services.NewNotificationService(notificationRepo)
	connectionService := services.NewConnectionService(connectionRepo)
	educationService := &services.EducationService{Repo: educationRepo}
	experienceService := services.NewExperienceService(experienceRepo)
	messageService := services.NewMessageService(messageRepo)
	groupService := services.NewGroupService(groupRepo)
	mediaService := services.NewMediaService(mediaRepo)
	feedService := services.NewFeedService(*postRepo)
	searchService := services.NewSearchService(searchRepo)
	adminService := services.NewAdminService(adminRepo)
	analyticsService := services.NewAnalyticsService(analyticsRepo)

	// Initialize controllers
	authController := controllers.AuthController{AuthService: &authService}
	userProfileController := controllers.NewUserProfileController(*userProfileService)
	postController := controllers.NewPostController(postService)
	commentController := controllers.NewCommentController(commentService)
	reactionController := controllers.NewReactionController(reactionService)
	notificationController := controllers.NewNotificationController(notificationService)
	connectionController := controllers.NewConnectionController(connectionService)
	educationController := controllers.NewEducationController(educationService)
	experienceController := controllers.NewExperienceController(experienceService)
	messageController := controllers.NewMessageController(messageService)
	groupController := controllers.NewGroupController(groupService)
	mediaController := controllers.NewMediaController(mediaService)
	feedController := controllers.NewFeedController(feedService)
	searchController := controllers.NewSearchController(searchService)
	adminController := controllers.NewAdminController(adminService)
	analyticsController := controllers.NewAnalyticsController(analyticsService)

	// Set up Gin router
	router := gin.Default()

	// Public API
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", gin.WrapF(authController.Register))
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/verify-otp", gin.WrapF(authController.VerifyOTP))
		authGroup.POST("/forgot-password", gin.WrapF(authController.ForgotPassword))
		authGroup.POST("/reset-password", gin.WrapF(authController.ResetPassword))
	}

	// Private API (Protected by middleware)
	authProtected := router.Group("/auth")
	authProtected.Use(middlewares.DeserializeUser(db))
	{
		authProtected.GET("/logout", authController.LogoutUser)
	}

	userProfileProtected := router.Group("/users")
	userProfileProtected.Use(middlewares.DeserializeUser(db))
	{
		userProfileProtected.GET("/:id", userProfileController.GetUserProfile)
		userProfileProtected.POST("/", userProfileController.CreateOrUpdateUserProfile)
		userProfileProtected.DELETE("/:id", userProfileController.DeleteUserProfile)
		userProfileProtected.POST("/:id/follow", connectionController.FollowUser)
		userProfileProtected.POST("/:id/unfollow", connectionController.UnfollowUser)
		userProfileProtected.POST("/:id/friend-request", connectionController.SendFriendRequest)
	}

	postProtected := router.Group("/posts")
	postProtected.Use(middlewares.DeserializeUser(db))
	{
		postProtected.POST("/", postController.CreatePost)
		postProtected.GET("/:id", postController.GetPost)
		postProtected.DELETE("/:id", postController.DeletePost)
		postProtected.POST("/:id/comment", commentController.CreateComment)
		postProtected.POST("/:id/:reaction_type", reactionController.AddReaction)
		postProtected.DELETE("/:id/:reaction_type/:user_id", reactionController.RemoveReaction)
	}

	feedProtected := router.Group("/feed")
	feedProtected.Use(middlewares.DeserializeUser(db))
	{
		feedProtected.GET("", feedController.GetFeedChronological)
		feedProtected.GET("/trending", feedController.GetFeedTrending)
	}

	searchProtected := router.Group("/search")
	searchProtected.Use(middlewares.DeserializeUser(db))
	{
		searchProtected.GET("/", searchController.Search)
	}

	notificationsProtected := router.Group("/notifications")
	notificationsProtected.Use(middlewares.DeserializeUser(db))
	{
		notificationsProtected.GET("/:userID", notificationController.GetNotifications)
		notificationsProtected.PUT("/:userID/mark-as-read", notificationController.MarkAsRead)
	}

	educationProtected := router.Group("/education")
	educationProtected.Use(middlewares.DeserializeUser(db))
	{
		educationProtected.GET("/:id", educationController.GetEducation)
		educationProtected.POST("/", educationController.CreateEducation)
		educationProtected.DELETE("/:id", educationController.DeleteEducation)
		educationProtected.GET("/user/:user_id", educationController.ListEducationByUser)
	}

	experienceProtected := router.Group("/experience")
	experienceProtected.Use(middlewares.DeserializeUser(db))
	{
		experienceProtected.GET("/:id", experienceController.GetExperienceByID)
		experienceProtected.POST("/", experienceController.CreateExperience)
		experienceProtected.DELETE("/:id", experienceController.DeleteExperience)
		experienceProtected.GET("/all/:user_id", experienceController.GetAllExperiencesByUserID)
	}

	messageProtected := router.Group("/messages")
	messageProtected.Use(middlewares.DeserializeUser(db))
	{
		messageProtected.POST("/send", messageController.SendMessage)
		messageProtected.GET("/conversation/:id", messageController.GetConversation)
	}

	groupProtected := router.Group("/groups")
	groupProtected.Use(middlewares.DeserializeUser(db))
	{
		groupProtected.POST("/", groupController.CreateGroup)
		groupProtected.GET("/:id", groupController.GetGroup)
		groupProtected.POST("/:id/join", groupController.JoinGroup)
	}

	mediaProtected := router.Group("/media")
	mediaProtected.Use(middlewares.DeserializeUser(db))
	{
		mediaProtected.POST("/", mediaController.CreateMedia)
		mediaProtected.GET("/:user_id", mediaController.GetMediaByUserID)
	}

	moderationProtected := router.Group("/moderation")
	moderationProtected.Use(middlewares.DeserializeUser(db))
	{
		moderationProtected.POST("/posts/:id", adminController.ModeratePost)
		moderationProtected.POST("/users/:id", adminController.BanUser)
	}

	analyticsProtected := router.Group("/analytics")
	analyticsProtected.Use(middlewares.DeserializeUser(db))
	{
		analyticsProtected.GET("/posts", analyticsController.GetPostInteractions)
		analyticsProtected.GET("/users", analyticsController.GetUserAnalytics)
	}

	return router
}

func RunServer() {
	router := SetupServer()
	if err := router.Run(":8000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
