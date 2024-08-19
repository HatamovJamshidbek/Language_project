package api

import (
	"auth_service/api/handler"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "auth_service/api/docs"
)

type Server struct {
	AuthHandler handler.AuthenticaionHandler
}

func NewServer(authHandler handler.AuthenticaionHandler) *Server {
	return &Server{AuthHandler: authHandler}
}


// @version      1.0
// @description  This is an API for user authentication.
// @termsOfService http://swagger.io/terms/
// @contact.name  API Support
// @contact.email support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath      /auth
func (s *Server) NewRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("auth")

	api.POST("/login", s.AuthHandler.Login)
	api.POST("/register", s.AuthHandler.Register)
	api.POST("/logout", s.AuthHandler.Logout)
	api.POST("/change-password", s.AuthHandler.ChangePassword)
	api.POST("/forgot-password", s.AuthHandler.ForgotPassword)
	api.POST("/reset-password", s.AuthHandler.ResetPassword)

	return router
}
