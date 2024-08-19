package handler

import (
	"api_service/config"
	pbuLearning "api_service/genproto/learning"
	pbuAuth "api_service/genproto/user"
	"api_service/pkg"
	"log/slog"

	"github.com/casbin/casbin/v2"
)

type Handler struct {
	ClientUser     pbuAuth.UsersClient
	ClientLearning pbuLearning.LearningServiceClient
	Logger         *slog.Logger
	Enforse        *casbin.Enforcer
}

func NewHandler(cfg config.Config, logger *slog.Logger, enforcer *casbin.Enforcer) *Handler {
	return &Handler{
		ClientUser:     pkg.NewAuthenticationClient(&cfg),
		ClientLearning: pkg.NewLearningClient(&cfg),
		Logger:         logger,
		Enforse:        enforcer,
	}
}
