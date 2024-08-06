package routes

import (
	"BE-ecommerce-web-template/controllers"
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

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	controllers.NewCategoryController(r, categoryService)

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	controllers.NewProductController(r, productService)

}
