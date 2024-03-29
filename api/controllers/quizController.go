package controllers

import (
	"ajher-server/internal/quiz"
	"ajher-server/internal/user"
	"ajher-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type quizHandler struct {
	quizService quiz.Service
}

func NewQuizHandler(quizService quiz.Service) *quizHandler {
	return &quizHandler{quizService}
}

// Save Quiz  godoc
//
// @Summary  save quiz
// @Description Adding new quiz to the database
// @Tags   Quiz
// @Accept   json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add refresh token here>)
// @Param   createQuizInput body  quiz.CreateQuizInput true "User Data"
// @Success  200   {object} quiz.Quiz
// @Router   /quiz/save [post]
func (h *quizHandler) Save(ctx *gin.Context) {
	var input quiz.CreateQuizInput

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := utils.APIResponse("Save Quiz Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := ctx.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	newQuiz, err := h.quizService.Save(input, userID)

	if err != nil {
		response := utils.APIResponse("Save Quiz Failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Save Quiz Category Success", http.StatusOK, "success", newQuiz)

	ctx.JSON(http.StatusOK, response)

}

// Get Quiz Detail  godoc
//
// @Summary  get quiz detail
// @Description Get quiz from the database
// @Tags   Quiz
// @Accept   json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add refresh token here>)
// @Param id path string true "Quiz Id"
// @Success  200   {object} quiz.Quiz
// @Router   /quiz/{id} [get]
func (h *quizHandler) GetDetailQuiz(ctx *gin.Context) {
	quizId := ctx.Param("id")

	quizDetail, participation, err := h.quizService.GetQuizDetail(quizId)

	if err != nil {
		response := utils.APIResponse("Get Quiz Detail Failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := quiz.FormatQuiz(quizDetail, participation)
	response := utils.APIResponse("Get Quiz Detail Success", http.StatusOK, "success", formatter)
	ctx.JSON(http.StatusOK, response)
}

// Join Quiz  godoc
//
// @Summary  join quiz
// @Description Joining certain quiz
// @Tags   Quiz
// @Accept   json
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add refresh token here>)
// @Param quizCode path string true "Quiz Code"
// @Success  200   {object} quiz.Quiz
// @Router   /quiz/join/{quizCode} [post]
func (h *quizHandler) JoinQuiz(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	quizId := ctx.Param("quizCode")

	quiz, err := h.quizService.JoinQuiz(quizId, userID)

	if err != nil {
		response := utils.APIResponse("Join Quiz Failed", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Join Quiz Success", http.StatusOK, "success", quiz)
	ctx.JSON(http.StatusOK, response)
}
