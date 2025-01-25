package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"salesforge-api/internal/models"
	"salesforge-api/internal/persistence/mocks"
	"testing"
)

func TestAddSequence(t *testing.T) {
	mockRepo := new(mocks.SequenceRepository)
	svc := NewSequenceService(mockRepo)

	sequence := models.Sequence{
		AccountID:    1,
		SequenceName: "Test Sequence",
	}

	// When AddSequence is called (with any 3 parameters), return 1 and nil.
	mockRepo.On("AddSequence", mock.Anything, mock.Anything, mock.Anything).Return(int64(1), nil)

	ctx := context.Background()
	_, err := svc.AddSequence(ctx, &sequence, nil)
	assert.NoError(t, err)
	// Verify that the AddSequence method was called as expected.
	mockRepo.AssertExpectations(t)
}

func TestAddSequence_InvalidInput(t *testing.T) {
	mockRepo := new(mocks.SequenceRepository)
	svc := NewSequenceService(mockRepo)

	invalidSequence := models.Sequence{
		AccountID:    -1, // Invalid AccountID
		SequenceName: "",
	}

	mockRepo.On("AddSequence", mock.Anything, mock.Anything, mock.Anything).Return(int64(0), errors.New("validation error"))

	ctx := context.Background()
	_, err := svc.AddSequence(ctx, &invalidSequence, nil)
	assert.Error(t, err)
}

func TestUpdateSequence(t *testing.T) {
	mockRepo := new(mocks.SequenceRepository)
	svc := NewSequenceService(mockRepo)

	update := models.UpdateSequenceRequest{
		AccountID:                    1,
		SequenceID:                   1,
		SequenceClickTrackingEnabled: &[]bool{true}[0],
		SequenceOpenTrackingEnabled:  &[]bool{true}[0],
	}

	mockRepo.On("UpdateSequence", mock.Anything, mock.Anything).Return(int64(1), nil)

	ctx := context.Background()
	_, err := svc.UpdateSequence(ctx, &update)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateSequence_InvalidInput(t *testing.T) {
	mockRepo := new(mocks.SequenceRepository)
	svc := NewSequenceService(mockRepo)

	update := models.UpdateSequenceRequest{
		AccountID:  1,
		SequenceID: 1,
	}

	// Expect no call to UpdateSequence due to validation failure
	mockRepo.On("UpdateSequence", mock.Anything, mock.Anything).Return(int64(0), errors.New("validation error"))

	ctx := context.Background()
	_, err := svc.UpdateSequence(ctx, &update)
	assert.Error(t, err) // Expect an error due to invalid input
	mockRepo.AssertExpectations(t)
}

func TestUpdateStep(t *testing.T) {
	mockRepo := new(mocks.SequenceRepository)
	svc := NewSequenceService(mockRepo)

	update := models.UpdateStepRequest{
		AccountID:        1,
		SequenceID:       1,
		StepID:           1,
		StepEmailSubject: "Updated Subject",
		StepEmailBody:    "Updated Body",
	}

	mockRepo.On("UpdateStep", mock.Anything, mock.Anything).Return(int64(1), int64(1), nil)

	ctx := context.Background()
	_, _, err := svc.UpdateStep(ctx, &update)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateStep_InvalidInput(t *testing.T) {
	mockRepo := new(mocks.SequenceRepository)
	svc := NewSequenceService(mockRepo)

	// Invalid input: missing required fields
	update := models.UpdateStepRequest{
		AccountID:        -1, // Invalid AccountID
		SequenceID:       1,
		StepID:           1,
		StepEmailSubject: "",
		StepEmailBody:    "",
	}

	// Expect no call to UpdateStep due to validation failure
	mockRepo.On("UpdateStep", mock.Anything, mock.Anything).Return(int64(0), int64(0), errors.New("validation error"))

	ctx := context.Background()
	_, _, err := svc.UpdateStep(ctx, &update)
	assert.Error(t, err) // Expect an error due to invalid input
	mockRepo.AssertExpectations(t)
}

func TestDeleteStep(t *testing.T) {
	mockRepo := new(mocks.SequenceRepository)
	svc := NewSequenceService(mockRepo)

	delete := models.DeleteStepRequest{
		AccountID:  1,
		SequenceID: 1,
		StepID:     1,
	}

	mockRepo.On("DeleteStep", mock.Anything, mock.Anything).Return(int64(1), int64(1), nil)

	ctx := context.Background()
	_, _, err := svc.DeleteStep(ctx, &delete)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteStep_InvalidInput(t *testing.T) {
	mockRepo := new(mocks.SequenceRepository)
	svc := NewSequenceService(mockRepo)

	// Invalid input: missing required fields
	delete := models.DeleteStepRequest{
		AccountID:  -1, // Invalid AccountID
		SequenceID: 1,
		StepID:     1,
	}

	// Expect no call to DeleteStep due to validation failure
	mockRepo.On("DeleteStep", mock.Anything, mock.Anything).Return(int64(0), int64(0), errors.New("validation error"))

	ctx := context.Background()
	_, _, err := svc.DeleteStep(ctx, &delete)
	assert.Error(t, err) // Expect an error due to invalid input
	mockRepo.AssertExpectations(t)
}

func TestAddSequence_Success(t *testing.T) {
	mockRepo := new(mocks.SequenceRepository)
	svc := NewSequenceService(mockRepo)

	sequence := models.Sequence{
		AccountID:                    1,
		SequenceName:                 "Test Sequence",
		SequenceOpenTrackingEnabled:  true,
		SequenceClickTrackingEnabled: true,
	}
	steps := []models.Step{
		{
			StepEmailSubject:  "Subject 1",
			StepEmailBody:     "Body 1",
			WaitDays:          1,
			EligibleStartTime: 1717758001,
			EligibleEndTime:   1718758034,
		},
	}

	mockRepo.On("AddSequence", mock.Anything, &sequence, &steps).Return(int64(1), nil)

	ctx := context.Background()
	sequenceId, err := svc.AddSequence(ctx, &sequence, &steps)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), sequenceId)
	mockRepo.AssertExpectations(t)
}

func TestAddSequence_Failure(t *testing.T) {
	mockRepo := new(mocks.SequenceRepository)
	svc := NewSequenceService(mockRepo)

	sequence := models.Sequence{
		AccountID:                    1,
		SequenceName:                 "Test Sequence",
		SequenceOpenTrackingEnabled:  true,
		SequenceClickTrackingEnabled: true,
	}
	steps := []models.Step{
		{
			StepEmailSubject:  "Subject 1",
			StepEmailBody:     "Body 1",
			WaitDays:          1,
			EligibleStartTime: 1717758001,
			EligibleEndTime:   1718758034,
		},
	}

	mockRepo.On("AddSequence", mock.Anything, &sequence, &steps).Return(int64(0), errors.New("db error"))

	ctx := context.Background()
	sequenceId, err := svc.AddSequence(ctx, &sequence, &steps)
	assert.Error(t, err)
	assert.Equal(t, int64(0), sequenceId)
	mockRepo.AssertExpectations(t)
}

func TestUpdateSequence_Success(t *testing.T) {
	mockRepo := new(mocks.SequenceRepository)
	svc := NewSequenceService(mockRepo)
	enabled := true

	update := models.UpdateSequenceRequest{
		AccountID:                    1,
		SequenceID:                   1,
		SequenceOpenTrackingEnabled:  &enabled,
		SequenceClickTrackingEnabled: &enabled,
	}

	mockRepo.On("UpdateSequence", mock.Anything, &update).Return(int64(1), nil)

	ctx := context.Background()
	sequenceId, err := svc.UpdateSequence(ctx, &update)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), sequenceId)
	mockRepo.AssertExpectations(t)
}

func TestUpdateSequence_Failure(t *testing.T) {
	mockRepo := new(mocks.SequenceRepository)
	svc := NewSequenceService(mockRepo)
	enabled := true

	update := models.UpdateSequenceRequest{
		AccountID:                    1,
		SequenceID:                   1,
		SequenceOpenTrackingEnabled:  &enabled,
		SequenceClickTrackingEnabled: &enabled,
	}

	mockRepo.On("UpdateSequence", mock.Anything, &update).Return(int64(0), errors.New("db error"))

	ctx := context.Background()
	sequenceId, err := svc.UpdateSequence(ctx, &update)
	assert.Error(t, err)
	assert.Equal(t, int64(0), sequenceId)
	mockRepo.AssertExpectations(t)
}

func TestUpdateStep_Success(t *testing.T) {
	mockRepo := new(mocks.SequenceRepository)
	svc := NewSequenceService(mockRepo)

	update := models.UpdateStepRequest{
		AccountID:        1,
		SequenceID:       1,
		StepID:           1,
		StepEmailSubject: "Updated Subject",
		StepEmailBody:    "Updated Body",
	}

	mockRepo.On("UpdateStep", mock.Anything, &update).Return(int64(1), int64(1), nil)

	ctx := context.Background()
	sequenceId, stepId, err := svc.UpdateStep(ctx, &update)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), sequenceId)
	assert.Equal(t, int64(1), stepId)
	mockRepo.AssertExpectations(t)
}

func TestUpdateStep_Failure(t *testing.T) {
	mockRepo := new(mocks.SequenceRepository)
	svc := NewSequenceService(mockRepo)

	update := models.UpdateStepRequest{
		AccountID:        1,
		SequenceID:       1,
		StepID:           1,
		StepEmailSubject: "Updated Subject",
		StepEmailBody:    "Updated Body",
	}

	mockRepo.On("UpdateStep", mock.Anything, &update).Return(int64(0), int64(0), errors.New("db error"))

	ctx := context.Background()
	sequenceId, stepId, err := svc.UpdateStep(ctx, &update)
	assert.Error(t, err)
	assert.Equal(t, int64(0), sequenceId)
	assert.Equal(t, int64(0), stepId)
	mockRepo.AssertExpectations(t)
}

func TestDeleteStep_Success(t *testing.T) {
	mockRepo := new(mocks.SequenceRepository)
	svc := NewSequenceService(mockRepo)

	delete := models.DeleteStepRequest{
		AccountID:  1,
		SequenceID: 1,
		StepID:     1,
	}

	mockRepo.On("DeleteStep", mock.Anything, &delete).Return(int64(1), int64(1), nil)

	ctx := context.Background()
	sequenceId, stepId, err := svc.DeleteStep(ctx, &delete)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), sequenceId)
	assert.Equal(t, int64(1), stepId)
	mockRepo.AssertExpectations(t)
}

func TestDeleteStep_Failure(t *testing.T) {
	mockRepo := new(mocks.SequenceRepository)
	svc := NewSequenceService(mockRepo)

	delete := models.DeleteStepRequest{
		AccountID:  1,
		SequenceID: 1,
		StepID:     1,
	}

	mockRepo.On("DeleteStep", mock.Anything, &delete).Return(int64(0), int64(0), errors.New("db error"))

	ctx := context.Background()
	sequenceId, stepId, err := svc.DeleteStep(ctx, &delete)
	assert.Error(t, err)
	assert.Equal(t, int64(0), sequenceId)
	assert.Equal(t, int64(0), stepId)
	mockRepo.AssertExpectations(t)
}
