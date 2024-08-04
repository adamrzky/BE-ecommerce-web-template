package main

import (
	"BE-ecommerce-web-template/api"
)

// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	api.App.Run()
}
