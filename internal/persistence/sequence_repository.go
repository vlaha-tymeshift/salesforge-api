package persistence

import (
	"context"
	"database/sql"
	"salesforge-api/internal/models"
	"time"
)

type SequenceRepository interface {
	AddSequence(ctx context.Context, sequence *models.Sequence, steps *[]models.Step) (sequenceId int64, err error)
	UpdateSequence(ctx context.Context, update *models.UpdateSequenceRequest) (sequenceId int64, err error)
	UpdateStep(ctx context.Context, update *models.UpdateStepRequest) (sequenceId int64, stepId int64, err error)
	DeleteStep(ctx context.Context, delete *models.DeleteStepRequest) (sequenceId int64, stepId int64, err error)
}

type sequenceRepository struct {
	db *sql.DB
}

func NewSequenceRepository(db *sql.DB) SequenceRepository {
	return &sequenceRepository{
		db: db,
	}
}

func (r *sequenceRepository) AddSequence(ctx context.Context, sequence *models.Sequence, steps *[]models.Step) (sequenceId int64, err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	accountId, sequenceId, err := r.addSequence(ctx, tx, sequence)
	if err != nil {
		return 0, err
	}

	err = r.addSteps(ctx, tx, accountId, sequenceId, steps)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return sequenceId, nil
}

func (r *sequenceRepository) addSequence(ctx context.Context, tx *sql.Tx, sequence *models.Sequence) (accountId int64, sequenceId int64, err error) {
	query := `INSERT INTO sequences (account_id, created_at, sequence_name, sequence_open_tracking_enabled, sequence_click_tracking_enabled) VALUES ($1, $2, $3, $4, $5) RETURNING account_id, sequence_id`
	createdAt := time.Now().Unix()
	err = tx.QueryRowContext(ctx, query, sequence.AccountID, createdAt, sequence.SequenceName, sequence.SequenceOpenTrackingEnabled, sequence.SequenceClickTrackingEnabled).Scan(&accountId, &sequenceId)
	if err != nil {
		return 0, 0, err
	}

	return accountId, sequenceId, nil
}

func (r *sequenceRepository) addSteps(ctx context.Context, tx *sql.Tx, accountId int64, sequenceId int64, steps *[]models.Step) error {
	query := `INSERT INTO steps (account_id, sequence_id, created_at, step_email_subject, step_email_body, wait_days, eligible_start_time, eligible_end_time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	createdAt := time.Now().Unix()
	for _, step := range *steps {
		_, err := tx.ExecContext(ctx, query, accountId, sequenceId, createdAt, step.StepEmailSubject, step.StepEmailBody, step.WaitDays, step.EligibleStartTime, step.EligibleEndTime)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *sequenceRepository) UpdateSequence(ctx context.Context, update *models.UpdateSequenceRequest) (sequenceId int64, err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	sequenceId, err = r.updateSequence(ctx, tx, update)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return sequenceId, nil
}

func (r *sequenceRepository) updateSequence(ctx context.Context, tx *sql.Tx, update *models.UpdateSequenceRequest) (sequenceId int64, err error) {
	query := `UPDATE sequences SET sequence_open_tracking_enabled = $1, sequence_click_tracking_enabled = $2, updated_at = $3 WHERE account_id = $4 AND sequence_id = $5 RETURNING sequence_id`
	updatedAt := time.Now().Unix()
	err = tx.QueryRowContext(ctx, query, update.SequenceOpenTrackingEnabled, update.SequenceClickTrackingEnabled, updatedAt, update.AccountID, update.SequenceID).Scan(&sequenceId)
	if err != nil {
		return 0, err
	}

	return sequenceId, nil
}

func (r *sequenceRepository) UpdateStep(ctx context.Context, update *models.UpdateStepRequest) (sequenceId int64, stepId int64, err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, 0, err
	}
	defer tx.Rollback()

	sequenceId, stepId, err = r.updateStep(ctx, tx, update)
	if err != nil {
		return 0, 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, 0, err
	}

	return sequenceId, stepId, nil
}

func (r *sequenceRepository) updateStep(ctx context.Context, tx *sql.Tx, update *models.UpdateStepRequest) (sequenceId int64, stepId int64, err error) {
	query := `UPDATE steps SET step_email_subject = $1, step_email_body = $2, updated_at = $3 WHERE account_id = $4 AND sequence_id = $5 AND step_id = $6 RETURNING sequence_id, step_id`
	updatedAt := time.Now().Unix()
	err = tx.QueryRowContext(ctx, query, update.StepEmailSubject, update.StepEmailBody, updatedAt, update.AccountID, update.SequenceID, update.StepID).Scan(&sequenceId, &stepId)
	if err != nil {
		return 0, 0, err
	}

	return sequenceId, stepId, nil
}

func (r *sequenceRepository) DeleteStep(ctx context.Context, delete *models.DeleteStepRequest) (sequenceId int64, stepId int64, err error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, 0, err
	}
	defer tx.Rollback()

	sequenceId, stepId, err = r.deleteStep(ctx, tx, delete)
	if err != nil {
		return 0, 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, 0, err
	}

	return sequenceId, stepId, nil
}

func (r *sequenceRepository) deleteStep(ctx context.Context, tx *sql.Tx, delete *models.DeleteStepRequest) (sequenceId int64, stepId int64, err error) {
	query := `DELETE FROM steps WHERE account_id = $1 AND sequence_id = $2 AND step_id = $3 RETURNING sequence_id, step_id`
	err = tx.QueryRowContext(ctx, query, delete.AccountID, delete.SequenceID, delete.StepID).Scan(&sequenceId, &stepId)
	if err != nil {
		return 0, 0, err
	}

	return sequenceId, stepId, nil
}
