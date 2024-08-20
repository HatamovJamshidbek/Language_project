package api

import (
	"api_service/api/handler"
	"api_service/config"
	"log/slog"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title        LinguaLeap API
// @version      1.0
// @description  This is an API for LinguaLeap.
// @termsOfService http://swagger.io/terms/
// @contact.name  API Support
// @contact.email support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath      /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @security [ApiKeyAuth]
func NewRouter(cfg config.Config, logger *slog.Logger, enforcer *casbin.Enforcer) *gin.Engine {
	router := gin.Default()
	h := handler.NewHandler(cfg, logger, enforcer)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//router.Use(middlewere.JWTMiddleware(), middlewere.CasbinMiddleware(enforcer))
	auth := router.Group("/auth")

	{
		auth.POST("/register", h.Register)
		auth.GET(("/profile"), h.GetUserProfile)
		auth.PUT("/profile/:id", h.UpdateUserProfile)
		auth.DELETE("/users/:id", h.DeleteUser)
		auth.GET("/users/:id", h.GetByUserId)
	}

	learning := router.Group("learning")

	{
		learning.GET("/languages", h.GetLanguages)
		learning.POST("/languages", h.StartLearnLanguage)
		learning.GET("/lessons", h.GetLessonsList)
		learning.GET("/lessons/:id", h.GetInfoLessons)
		learning.POST("/complate/:lesson_id", h.ComplateLesson)
		learning.GET("/exercises/:code", h.GetLanguageExercises)
		learning.POST("/exercises/:id", h.SubmitExerciseAnswer)
		learning.GET("/vocabulary", h.GetVocabularyList)
		learning.POST("/vocabulary", h.AddVocabulary)

	}

	return router

}
