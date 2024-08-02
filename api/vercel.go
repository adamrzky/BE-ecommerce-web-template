package api

import (
	"BE-ecommerce-web-template/config"
	// "BE-ecommerce-web-template/docs"
	"BE-ecommerce-web-template/routes"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	App *gin.Engine
)

func init() {
	App = gin.New()

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "development"
	}

	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// docs.SwaggerInfo.Title = "Movie REST API"
	// docs.SwaggerInfo.Description = "This is REST API Movie."
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = os.Getenv("HOST")
	// if docs.SwaggerInfo.Host == "" {
	// 	docs.SwaggerInfo.Host = "localhost:8080"
	// }
	// if environment == "development" {
	// 	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	// } else {
	// 	docs.SwaggerInfo.Schemes = []string{"https"}
	// }

	db := config.ConnectDataBase()
	routes.SetupRouter(db, App)
}

// Entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	App.ServeHTTP(w, r)
}
