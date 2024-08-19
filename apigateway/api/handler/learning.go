package handler

import (
	pb "api_service/genproto/learning"
	"api_service/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetLanguages @Summary      Get languages
// @Description  Get languages list
// @Tags         Learning
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        limit    query     int  false  "Limit"   default(10)
// @Param        page     query     int  false  "Page"    default(0)
// @Success      200      {object}  learning.GetLanguageResponse
// @Failure      400      {object}  map[string]interface{}  "Bad request"
// @Failure      500      {object}  map[string]interface{}  "Internal server error"
// @Router       /learning/languages [get]
func (h *Handler) GetLanguages(c *gin.Context) {
	req := pb.GetLanguageRequest{}

	limitStr := c.Query("limit")
	pageStr := c.Query("page")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 0
	}

	req.Limit = int32(limit)
	req.Page = int32(page)

	res, err := h.ClientLearning.GetLanguages(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		fmt.Println("+++++++", err)
		c.JSON(500, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// StartLearnLanguage       Start learn language
// @Description  Start learn language
// @Tags         Learning
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        language  body     learning.StartLanguageRequest  true  "Start Language Request"
// @Success      200      {object}  learning.StartLanguageResponse
// @Failure      400      {object}  map[string]interface{}  "Bad request"
// @Failure      500      {object}  map[string]interface{}  "Internal server error"
// @Router       /learning/languages [post]
func (h *Handler) StartLearnLanguage(c *gin.Context) {
	req := pb.StartLanguageRequest{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	res, err := h.ClientLearning.StartLearnLanguage(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		fmt.Println("+++++++++", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetLessonsList       Get lessons list
// @Description  Get lessons list
// @Tags         Learning
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        code  query     string  true  "Language Request"
// @Success      200      {object}  learning.LessonsListResponse
// @Failure      400      {object}  map[string]interface{}  "Bad request"
// @Failure      500      {object}  map[string]interface{}  "Internal server error"
// @Router       /learning/lessons [get]
func (h *Handler) GetLessonsList(c *gin.Context) {
	code := c.Query("code")
	fmt.Println(code)
	req := pb.Language{
		Code: code,
	}

	res, err := h.ClientLearning.GetLessonsList(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		fmt.Println("++++++++", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetInfoLessons       Get info lessons
// @Description  Get info lessons
// @Tags         Learning
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id  path     string  true  "Lesson ID"
// @Success      200      {object}  learning.GetInfoLessonsResponse
// @Failure      400      {object}  map[string]interface{}  "Bad request"
// @Failure      500      {object}  map[string]interface{}  "Internal server error"
// @Router       /learning/lessons/{id} [get]
func (h *Handler) GetInfoLessons(c *gin.Context) {

	id := c.Param("id")
	req := pb.GetInfoLessonsResponse{
		LessonId: id,
	}

	fmt.Println(id, req.LessonId)
	res, err := h.ClientLearning.GetInfoLessons(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		fmt.Println("+++++++", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// ComplateLesson       Complate lesson
// @Description  Complate lesson
// @Tags         Learning
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        lesson_id  path string true  "Lesson Request"
// @Success      200      {object}  learning.ComplateLessonResponse
// @Failure      400      {object}  map[string]interface{}  "Bad request"
// @Failure      500      {object}  map[string]interface{}  "Internal server error"
// @Router       /learning/complate/{lesson_id}	 [post]
func (h *Handler) ComplateLesson(c *gin.Context) {
	id := c.Param("lesson_id")
	fmt.Println(id)
	req := pb.Lesson{
		Id: id,
	}

	res, err := h.ClientLearning.ComplateLesson(c, &req)
	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
		h.Logger.Error(err.Error())
		fmt.Println("+++++++++++", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetLanguageExercises       Get language exercises
// @Description  Get language exercises
// @Tags         Learning
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        code  query     string  true  "Language Request"
// @Success      200      {object}  learning.LanguageExerciseResponse
// @Failure      400      {object}  map[string]interface{}  "Bad request"
// @Failure      500      {object}  map[string]interface{}  "Internal server error"
// @Router       /learning/exercises/{code} [get]
func (h *Handler) GetLanguageExercises(c *gin.Context) {
	code := c.Query("code")
	fmt.Println(code)

	req := pb.Language{
		Code: code,
	}
	fmt.Println(code)
	res, err := h.ClientLearning.GetLanguageExercises(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		fmt.Println("+++++++++++", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// SubmitExerciseAnswer       Submit exercise answer
// @Description  Submit exercise answer
// @Tags         Learning
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id  path     string  true  "Exercise ID"
// @Param        answer  body     model.ExerciseAnswerRequest  true  "Exercise Answer Request"
// @Success      200      {object}  learning.ExerciseAnswerResponse
// @Failure      400      {object}  map[string]interface{}  "Bad request"
// @Failure      500      {object}  map[string]interface{}  "Internal server error"
// @Router       /learning/exercises/{id} [post]
func (h *Handler) SubmitExerciseAnswer(c *gin.Context) {
	r := model.ExerciseAnswerRequest{}

	err := c.ShouldBindJSON(&r)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	id := c.Param("id")

	req := pb.ExerciseAnswerRequest{
		ExerciseId: id,
		Answer:     r.Answer,
	}

	res, err := h.ClientLearning.SubmitExerciseAnswer(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		fmt.Println("++++++++++", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetVocabularyList       Get vocabulary list
// @Description  Get vocabulary list
// @Tags         Learning
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        code     query    string  true  "Language Code"
// @Success      200      {object}  learning.VocabularyListResponse
// @Failure      400      {object}  map[string]interface{}  "Bad request"
// @Failure      500      {object}  map[string]interface{}  "Internal server error"
// @Router       /learning/vocabulary [get]
func (h *Handler) GetVocabularyList(c *gin.Context) {

	code := c.Query("code")
	fmt.Println(code)

	req := pb.Language{
		Code: code,
	}

	res, err := h.ClientLearning.GetVocabularyList(c, &req)
	fmt.Println(res)
	if err != nil {
		h.Logger.Error(err.Error())
		fmt.Println("+++++++++", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

// AddVocabulary       Add vocabulary
// @Description  Add vocabulary
// @Tags         Learning
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        vocabulary  body     learning.AddVocabularyRequest  true  "Vocabulary Request"
// @Param        language_code  query     string  true  "Language Code"
// @Success      200      {object}  learning.AddVocabularyResponse
// @Failure      400      {object}  map[string]interface{}  "Bad request"
// @Failure      500      {object}  map[string]interface{}  "Internal server error"
// @Router       /learning/vocabulary [post]
func (h *Handler) AddVocabulary(c *gin.Context) {
	req := pb.AddVocabularyRequest{}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	code := c.Query("language_code")

	req.Languagecode = code

	res, err := h.ClientLearning.AddVocabulary(c, &req)
	if err != nil {
		h.Logger.Error(err.Error())
		fmt.Println("++++++++", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
