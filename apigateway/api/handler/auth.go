package handler

import (
	pb "api_service/genproto/user"
	"api_service/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary      Add a new user
// @Description  Register a new user with the provided details
// @Tags         User
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param        user  body     user.RegisterRequest  true  "Add User Request"
// @Success      201   {object} user.RegisterResponse
// @Failure      400   {object} map[string]interface{}  "Bad request"
// @Failure      500   {object} map[string]interface{}  "Internal server error"
// @Router       /auth/register [post]
func (h *Handler) Register(c *gin.Context) {
	var req pb.RegisterRequest
	err := c.BindJSON(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := h.ClientUser.Register(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary      Get user profile
// @Description  Get the current user's profile
// @Tags         User
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Success      200   {object} user.ProfileResponse
// @Failure      400   {object} map[string]interface{}  "Bad request"
// @Failure      500   {object} map[string]interface{} “ "Internal server error"
// @Router       /auth/profile [get]
func (h *Handler) GetUserProfile(c *gin.Context) {
	fmt.Println("done")
	res, err := h.ClientUser.GetUserProfile(c, &pb.Void{})
	if err != nil {
		fmt.Println("fasdfiuefuhaewthuewgdfdsfgadsfewyry7qwer8ew7r7ewt t6sdf8adsfgads88wety")
		h.Logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(err, "dfsdfidsugeahgtuerwty343rt")
	c.JSON(http.StatusOK, res)
}

// @Summary      Update user profile
// @Description  Update the current user's profile
// @Tags         User
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id    path     string                      true  "User ID"
// @Param        user  body     model.UpdateUser   true  "Update User Request"
// @Success      200   {object} user.ProfileUpdateResponse
// @Failure      400   {object} map[string]interface{}  "Bad request"
// @Failure      500   {object} map[string]interface{}  "Internal server error"
// @Router       /auth/profile/{id} [put]
func (h *Handler) UpdateUserProfile(c *gin.Context) {
	r := model.UpdateUser{}

	// id := c.Param("id")

	err := c.BindJSON(&r)
	if err != nil {
		h.Logger.Error(err.Error())
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	req := pb.ProfileUpdateRequest{
		Fullname:       r.FullName,
		Nativelanguage: r.LeanguageCode,
	}

	res, err := h.ClientUser.UpdateUserProfile(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary      Delete user
// @Description  Delete the user with the provided ID
// @Tags         User
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param        id  path     string  true  "User ID"
// @Success      200   {object} user.DeleteUserResponse
// @Failure      400   {object} map[string]interface{}  "Bad request"
// @Failure      500   {object} map[string]interface{}  "Internal server error"
// @Router       /auth/users/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	req := pb.DeleteUserRequest{}

	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	req.UserId = id

	res, err := h.ClientUser.DeleteUser(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary      Get user by ID
// @Description  Get the user with the provided ID
// @Tags         User
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param        id  path     string  true  "User ID"
// @Success      200   {object} user.GetByUserIdResponse
// @Failure      400   {object} map[string]interface{}  "Bad request"
// @Failure      500   {object} map[string]interface{}  "Internal server error"
// @Router       /auth/users/{id} [get]
func (h *Handler) GetByUserId(c *gin.Context) {
	req := pb.GetByUserIdRequest{}

	id := c.Param("id")
	req.Id = id

	res, err := h.ClientUser.GetByUserId(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
