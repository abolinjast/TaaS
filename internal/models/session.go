package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID       uuid.UUID `json:"id"`
	UserID   uuid.UUID `json:"user_id"`
	CourseID uuid.UUID `json:"course_id"`

	Module string `json:"module"`
	Topic  string `json:"topic"`

	StartTime time.Time  `json:"start_time"`
	EndTime   *time.Time `json:"end_time"`
	Duration  *int64     `json:"duration,omitempty"`

	ActivityType string `json:"activity_type"`
	QuizScore    *int   `json:"quiz_score,omitempty"`
	QuizPassed   *bool  `json:"quiz_passed,omitempty"`

	Status string `json:"status"`
	Notes  string `json:"notes"`
}
