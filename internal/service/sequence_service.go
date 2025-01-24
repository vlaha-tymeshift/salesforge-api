package service

import (
	"context"
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
	return s.sequenceRepo.AddSequence(ctx, sequence, steps)
}

func (s *sequenceService) UpdateSequence(ctx context.Context, update *models.UpdateSequenceRequest) (sequenceId int64, err error) {
	return s.sequenceRepo.UpdateSequence(ctx, update)
}

func (s *sequenceService) UpdateStep(ctx context.Context, update *models.UpdateStepRequest) (sequenceId int64, stepId int64, err error) {
	return s.sequenceRepo.UpdateStep(ctx, update)
}

func (s *sequenceService) DeleteStep(ctx context.Context, delete *models.DeleteStepRequest) (sequenceId int64, stepId int64, err error) {
	return s.sequenceRepo.DeleteStep(ctx, delete)
}
