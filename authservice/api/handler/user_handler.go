package handler

import (
	"auth_service/api/email"
	"auth_service/api/token"
	pb "auth_service/genproto/user"
	"auth_service/model"
	"auth_service/service"
	"auth_service/storage/redis"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticaionHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	Logout(c *gin.Context)
	ChangePassword(c *gin.Context)
	ForgotPassword(c *gin.Context)
	ResetPassword(c *gin.Context)
}

type AuthenticaionHandlerImpl struct {
	UserManage service.UserManagementService
	Logger     *slog.Logger
	pb.UnimplementedUsersServer
}

func NewAuthenticaionHandler(userManage service.UserManagementService, logger *slog.Logger) AuthenticaionHandler {
	return &AuthenticaionHandlerImpl{UserManage: userManage, Logger: logger}
}



// @Summary      Register a new user
// @Description  Log in a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body     user.LoginRequest  true  "User registration request"
// @Success      200   {object}  model.Tokens
// @Failure      400   {object}  map[string]string    "Bad request"
// @Failure      500   {object}  map[string]string    "Internal server error"
// @Router       /login [post]
func (h *AuthenticaionHandlerImpl) Login(c *gin.Context) {
	req := pb.LoginRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.UserManage.Login(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	tokenString := token.GenerateJWT(resp)

	c.JSON(http.StatusOK, tokenString)
}

// @Summary      Register a new user
// @Description  Register a new user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body     user.RegisterRequest  true  "User registration request"
// @Success      200   {object}  user.RegisterResponse
// @Failure      400   {object}  map[string]string    "Bad request"
// @Failure      500   {object}  map[string]string    "Internal server error"
// @Router       /register [post]
func (h *AuthenticaionHandlerImpl) Register(c *gin.Context) {
	req := pb.RegisterRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.UserManage.Register(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary      Logout a user
// @Description  Logout a user account
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body     user.LogoutRequest  true  "User logout request"
// @Success      200   {object}  user.Message
// @Failure      400   {object}  map[string]string    "Bad request"
// @Failure      500   {object}  map[string]string    "Internal server error"
// @Router       /logout [post]
func (h *AuthenticaionHandlerImpl) Logout(c *gin.Context) {
	req := pb.LogoutRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	resp, err := h.UserManage.Logout(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary      Change user password
// @Description  Change user password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body     user.ChangePasswordRequest  true  "User change password request"
// @Success      200   {object}  user.Message
// @Failure      400   {object}  map[string]string    "Bad request"
// @Failure      500   {object}  map[string]string    "Internal server error"
// @Router       /change-password [post]
func (h *AuthenticaionHandlerImpl) ChangePassword(c *gin.Context) {
	req := pb.ChangePasswordRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if len(req.NewPassword) == 0 {
		h.Logger.Error("New password is empty")
		c.JSON(http.StatusBadRequest, "New password is empty")
		return
	}

	if len(req.CurrentPassword) == 0 {
		h.Logger.Error("Current password is empty")
		c.JSON(http.StatusBadRequest, "Current password is empty")
		return
	}

	resp, err := h.UserManage.ChangePassword(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Summary      Forgot user password
// @Description  Forgot user password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body     user.ForgotPasswordRequest  true  "User forgot password request"
// @Success      200   {object}  user.Message
// @Failure      400   {object}  map[string]string    "Bad request"
// @Failure      500   {object}  map[string]string    "Internal server error"
// @Router       /forgot-password [post]
func (h *AuthenticaionHandlerImpl) ForgotPassword(c *gin.Context) {
	req := pb.ForgotPasswordRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	code, err := email.Email(req.Email)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(code, "++++++")

	c.JSON(http.StatusOK, &pb.Message{Message: "Password reset instructions sent to your email"})
}

// @Summary      Reset user password
// @Description  Reset user password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body     model.ResetPasswordRequest  true  "User reset password request"
// @Success      200   {object}  user.Message
// @Failure      400   {object}  map[string]string    "Bad request"
// @Failure      500   {object}  map[string]string    "Internal server error"
// @Router       /reset-password [post]
func (h *AuthenticaionHandlerImpl) ResetPassword(c *gin.Context) {
	fmt.Println(1)
	req := model.ResetPasswordRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		h.Logger.Error("Error decoding request body")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	e, err := redis.ReadEmail()
	if err != nil {
		h.Logger.Error("Error reading email from redis")
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	pass, err := redis.ReadPassword(e)
	if err != nil {
		h.Logger.Error("Error sending email")
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if pass != req.ResetToken {
		h.Logger.Error("Invalid code")
		c.JSON(http.StatusBadRequest, "Invalid code")
		return
	}
	resp, err := h.UserManage.ResetPassword(c, &pb.ResetPasswordRequest{
		Email:       e,
		NewPassword: req.NewPassword,
	})
	if err != nil {
		h.Logger.Error("Error updating user password")
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	fmt.Println(resp)
	c.JSON(http.StatusOK, resp)
}
