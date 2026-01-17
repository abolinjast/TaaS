package store

import (
	"context"
	"database/sql"

	"github.com/abolinjast/taas/internal/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type PostgresStore struct {
	DB *sql.DB
}

func NewPostgresStore(db *sql.DB) *PostgresStore {
	return &PostgresStore{
		DB: db,
	}
}

func (s *PostgresStore) Create(ctx context.Context, session *models.Session) error {
	query := `
		INSERT INTO study_sessions (
			id, user_id, course_id, module,
			topic, start_time, activity_type,
			status
		) VALUES (
		$1, $2, $3, $4, $5,
		$6, $7, 'running'
		)
		RETURNING id
	`
	if session.ID == uuid.Nil {
		session.ID = uuid.New()
	}
	return s.DB.QueryRowContext(
		ctx, query,
		session.ID,
		session.UserID, session.CourseID,
		session.Module, session.Topic,
		session.StartTime, session.ActivityType,
	).Scan(&session.ID)
}

func (s *PostgresStore) GetActive(ctx context.Context, userID uuid.UUID) (*models.Session, error) {
	query := `
		SELECT id, user_id, course_id, module, topic,
		start_time, activity_type, status
		FROM study_sessions
		WHERE user_id = $1 and status = 'running'
		LIMIT 1
	`
	row := s.DB.QueryRowContext(ctx, query, userID)
	var session models.Session
	err := row.Scan(
		&session.ID, &session.UserID, &session.CourseID,
		&session.Module, &session.Topic, &session.StartTime,
		&session.ActivityType, &session.Status,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (s *PostgresStore) Update(ctx context.Context, session *models.Session) error {
	query := `
		UPDATE study_sessions
		SET
			end_time = $1,
			duration_seconds = EXTRACT(EPOCH FROM ($1 - start_time)),
			status = $2,
			notes = $3,
			quiz_score = $4,
			quiz_passed = $5
		WHERE id = $6
	`
	_, err := s.DB.ExecContext(
		ctx, query,
		session.EndTime,
		session.Status,
		session.Notes,
		session.QuizScore,
		session.QuizPassed,
		session.ID,
	)
	return err
}
