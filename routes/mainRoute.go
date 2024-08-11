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
	roleRepo := repositories.NewRoleRepository(db)
	authService := &services.AuthService{
		UserRepo: userRepo,
		RoleRepo: roleRepo,
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

	r.GET("/transactions", transactionController.GetAllTransactions)
	r.GET("/transactions/:id", transactionController.GetTransactionByID)
	r.POST("/transactions", transactionController.CreateTransaction)
	r.PUT("/transactions/:id", transactionController.UpdateTransaction)
	r.DELETE("/transactions/:id", transactionController.DeleteTransaction)
	r.GET("/mytransactions", transactionController.GetMyTransactions)

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
	profileService := services.NewProfileService(ProfileRepo)
	profileController := controllers.NewProfileController(profileService)
	r.GET("/profiles/:id", profileController.GetByID)
	r.POST("/profiles", profileController.Create)
	r.GET("/my-profiles", profileController.GetMyProfiles)
	r.PUT("/profiles/:id", profileController.Update)
	r.DELETE("/profiles/:id", profileController.DeleteProfile)

	// User
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	r.GET("/users/:id", middlewares.JwtAuthMiddleware(), userController.GetUserByID)
	r.POST("/users", middlewares.JwtAuthMiddleware(), userController.CreateUser)
	r.PUT("/users/:id", middlewares.JwtAuthMiddleware(), userController.UpdateUser)
	r.DELETE("/users/:id", middlewares.JwtAuthMiddleware(), userController.DeleteUser)
	r.GET("/users", userController.GetAllUsers)

	// Role
	roleService := services.NewRoleService(roleRepo)
	roleController := controllers.NewRoleController(roleService)

	r.GET("/roles/:id", middlewares.JwtAuthMiddleware(), roleController.GetRoleByID)
	r.POST("/roles", middlewares.JwtAuthMiddleware(), roleController.CreateRole)
	r.PUT("/roles/:id", middlewares.JwtAuthMiddleware(), roleController.UpdateRole)
	r.DELETE("/roles/:id", middlewares.JwtAuthMiddleware(), roleController.DeleteRole)
	r.GET("/roles", roleController.GetAllRoles)

	// Dummy (temporary controller, will be deleted once merge into main)
	dummyController := controllers.NewDummyController()

	r.GET("/my-claims", middlewares.JwtAuthMiddleware(), dummyController.MyClaims)                            // Sample to get claims from jwt
	r.GET("/admin-and-dev", middlewares.JwtAuthMiddleware("Admin", "Developer"), dummyController.AdminAndDev) // Sample to protect endpoint by multiple roles
}
