package service

import (
	"context"
	"net/http"
	"salesforge-api/internal/errors"
	"salesforge-api/internal/models"
	"salesforge-api/internal/persistence"
)

type sequenceService struct {
	sequenceRepo persistence.SequenceRepository
}

type SequenceService interface {
	AddSequence(ctx context.Context, sequence *models.Sequence, steps *[]models.Step) (sequenceId int64, err error)
	UpdateSequence(ctx context.Context, update *models.UpdateSequenceRequest) (sequenceId int64, err error)
	UpdateStep(ctx context.Context, update *models.UpdateStepRequest) (sequenceId int64, stepId int64, err error)
	DeleteStep(ctx context.Context, delete *models.DeleteStepRequest) (sequenceId int64, stepId int64, err error)
}

func NewSequenceService(
	sequenceRepo persistence.SequenceRepository,
) SequenceService {
	return &sequenceService{
		sequenceRepo: sequenceRepo,
	}
}

func (s *sequenceService) AddSequence(ctx context.Context, sequence *models.Sequence, steps *[]models.Step) (sequenceId int64, err error) {
	sequenceId, err = s.sequenceRepo.AddSequence(ctx, sequence, steps)
	if err != nil {
		return 0, errors.NewAppError(http.StatusInternalServerError, "failed to add sequence", err)
	}
	return sequenceId, nil
}

func (s *sequenceService) UpdateSequence(ctx context.Context, update *models.UpdateSequenceRequest) (sequenceId int64, err error) {
	sequenceId, err = s.sequenceRepo.UpdateSequence(ctx, update)
	if err != nil {
		return 0, errors.NewAppError(http.StatusInternalServerError, "failed to update sequence", err)
	}
	return sequenceId, nil
}

func (s *sequenceService) UpdateStep(ctx context.Context, update *models.UpdateStepRequest) (sequenceId int64, stepId int64, err error) {
	sequenceId, stepId, err = s.sequenceRepo.UpdateStep(ctx, update)
	if err != nil {
		return 0, 0, errors.NewAppError(http.StatusInternalServerError, "failed to update step", err)
	}
	return sequenceId, stepId, nil
}

func (s *sequenceService) DeleteStep(ctx context.Context, delete *models.DeleteStepRequest) (sequenceId int64, stepId int64, err error) {
	sequenceId, stepId, err = s.sequenceRepo.DeleteStep(ctx, delete)
	if err != nil {
		return 0, 0, errors.NewAppError(http.StatusInternalServerError, "failed to delete step", err)
	}
	return sequenceId, stepId, nil
}
