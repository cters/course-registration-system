package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/cmd/swag/docs"
	"gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/internal/initialize"
)

// @title Go Course Registration Backend API by DDD
// @version 1.0
// @description This is a server for a Go Course Registration Backend API, demonstrating DDD principles.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8001
// @BasePath /v1/api

// @externalDocs.description OpenAPI
// @externalDocs.url https://swagger.io/resources/open-api/
func main() {
	r, port := initialize.Run()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":" + port) 
}