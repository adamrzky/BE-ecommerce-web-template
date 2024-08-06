package routes

import (
	"BE-ecommerce-web-template/controllers"
	"BE-ecommerce-web-template/middlewares"
	"BE-ecommerce-web-template/repositories"
	"BE-ecommerce-web-template/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, r *gin.Engine) {

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	r.Use(cors.New(corsConfig))

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// User and Authentication
	userRepo := repositories.NewUserRepository(db)
	authService := &services.AuthService{
		UserRepo: userRepo,
	}
	authController := controllers.NewAuthController(authService)

	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)
	r.GET("/auth/me", authController.Me)
	r.POST("/auth/change-password", authController.ChangePassword)

	// Transaction
	transactionRepo := repositories.NewTransactionRepository(db)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionController := controllers.NewTransactionController(transactionService)

	r.GET("/transactions/:id", transactionController.GetTransactionByID)
	r.POST("/transactions", transactionController.CreateTransaction)
	r.PUT("/transactions/:id", transactionController.UpdateTransaction)
	r.DELETE("/transactions/:id", transactionController.DeleteTransaction)

	// Swagger API Docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Category
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	controllers.NewCategoryController(r, categoryService)

	// Product
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	controllers.NewProductController(r, productService)

	// Reviews
	reviewRepo := repositories.NewReviewRepository(db)
	reviewService := services.NewReviewService(reviewRepo)
	reviewController := controllers.NewReviewController(reviewService)

	r.GET("/my-reviews", middlewares.JwtAuthMiddleware(), reviewController.GetMyReviews)
	r.GET("/reviews/:id", reviewController.GetReviewById)
	r.GET("/reviews-product/:id", reviewController.GetReviewByProductId)
	r.POST("/reviews", middlewares.JwtAuthMiddleware(), reviewController.CreateReview)
	r.PUT("/reviews/:id", middlewares.JwtAuthMiddleware(), reviewController.UpdateReview)
	r.DELETE("/reviews/:id", middlewares.JwtAuthMiddleware(), reviewController.DeleteReview)

	//Profile
	ProfileRepo := repositories.NewProfileRepository(db)
	profileService := &services.ProfileService{
		ProfileRepo: ProfileRepo,
	}
	profileController := controllers.NewProfileController(profileService)
	r.GET("/profiles/:id", profileController.GetByID)
	r.GET("/profiles/:id/user", profileController.GetByUserID)
	r.POST("/profiles", profileController.Create)
	r.PUT("/profiles/:id", profileController.Update)
	r.DELETE("/profiles/:id", profileController.Delete)

}
