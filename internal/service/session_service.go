package service

import (
	"context"
	"errors"
	"time"

	"github.com/abolinjast/taas/internal/models"
	"github.com/abolinjast/taas/internal/store"
	"github.com/google/uuid"
)

type SessionService struct {
	repo store.SessionRepository
}

func NewSessionService(repo store.SessionRepository) *SessionService {
	return &SessionService{
		repo: repo,
	}
}

func (s *SessionService) StartSession(
	ctx context.Context,
	userID, courseID uuid.UUID,
	module, topic, activityType string,
) (*models.Session, error) {
	active, err := s.repo.GetActive(ctx, userID)
	if err != nil {
		return nil, err
	}
	if active != nil {
		return nil, errors.New("you already have a running session, stop it first")
	}

	if activityType == "" {
		activityType = "study"
	}

	newSession := &models.Session{
		ID:           uuid.New(),
		UserID:       userID,
		CourseID:     courseID,
		Module:       module,
		Topic:        topic,
		ActivityType: activityType,
		StartTime:    time.Now(),
		Status:       "running",
	}
	if err := s.repo.Create(ctx, newSession); err != nil {
		return nil, err
	}
	return newSession, err
}

func (s *SessionService) StopSession(
	ctx context.Context,
	userID uuid.UUID,
	notes string,
	quizScore *int,
	quizPassed *bool,
) (*models.Session, error) {
	active, err := s.repo.GetActive(ctx, userID)
	if err != nil {
		return nil, err
	}
	if active == nil {
		return nil, errors.New("you don't have a running session, start one first")
	}

	if active.ActivityType == "study" {
		if active.QuizScore != nil || active.QuizPassed != nil {
			return nil, errors.New("cannot save quiz results: this is a 'study' session, not a 'quiz'")
		}
	}

	now := time.Now()
	active.EndTime = &now
	active.Status = "completed"
	active.Notes = notes

	active.QuizScore = quizScore
	active.QuizPassed = quizPassed

	// No calculation done here, postgres does it for us
	if err := s.repo.Update(ctx, active); err != nil {
		return nil, err
	}
	return active, nil

}
