package handler

import (
	"net/http"

	"github.com/abolinjast/taas/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SessionHandler struct {
	svc *service.SessionService
}

func NewSessionHandler(svc *service.SessionService) *SessionHandler {
	return &SessionHandler{
		svc: svc,
	}
}

type StartRequest struct {
	UserID       string `json:"user_id" binding:"required,uuid"`
	CourseID     string `json:"course_id" binding:"required"`
	Module       string `json:"module" binding:"required"`
	Topic        string `json:"topic" binding:"required"`
	ActivityType string `json:"activity_type"`
}

type StopRequest struct {
	UserID     string `json:"user_id" binding:"required,uuid"`
	Notes      string `json:"notes"`
	QuizScore  *int   `json:"quiz_score"`
	QuizPassed *bool  `json:"quiz_passed"`
}

func (h *SessionHandler) Start(ctx *gin.Context) {
	var req StartRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userID := uuid.MustParse(req.UserID)
	courseID := uuid.Nil
	if req.CourseID != "" {
		courseID, _ = uuid.Parse(req.CourseID)
	}

	session, err := h.svc.StartSession(
		ctx.Request.Context(),
		userID,
		courseID,
		req.Module,
		req.Topic,
		req.ActivityType,
	)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, session)
}

func (h *SessionHandler) Stop(ctx *gin.Context) {
	var req StopRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userID := uuid.MustParse(req.UserID)
	session, err := h.svc.StopSession(
		ctx.Request.Context(),
		userID,
		req.Notes,
		req.QuizScore,
		req.QuizPassed,
	)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, session)
}
