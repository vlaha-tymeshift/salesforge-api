package sequence

import (
	"github.com/go-chi/render"
	"go.uber.org/zap"
	"net/http"
	"salesforge-api/internal/errors"
	"salesforge-api/internal/models"
	"salesforge-api/internal/service"
)

type SequenceHandler struct {
	sequenceService service.SequenceService
	logger          *zap.Logger
}

func NewSequenceHandler(sequenceService service.SequenceService, logger *zap.Logger) *SequenceHandler {
	return &SequenceHandler{
		sequenceService: sequenceService,
		logger:          logger,
	}
}

func (sh *SequenceHandler) AddSequence(w http.ResponseWriter, r *http.Request) {
	sh.logger.Info("AddSequence request received")
	addSequenceRequest, err := NewAddSequenceRequestFromHttpRequest(r)
	if err != nil {
		appErr := errors.NewAppError(http.StatusBadRequest, "invalid request payload", err)
		sh.logger.Error("error decoding request", zap.Error(appErr))
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	sequence := addSequenceRequest.Sequence
	steps := addSequenceRequest.Steps
	sequenceId, err := sh.sequenceService.AddSequence(r.Context(), &sequence, &steps)
	if err != nil {
		appErr := errors.NewAppError(http.StatusInternalServerError, "failed to add sequence", err)
		sh.logger.Error("error processing request", zap.Error(appErr))
		http.Error(w, "An error occurred", http.StatusInternalServerError)
		return
	}

	res := models.AddSequenceResponse{
		SequenceID: sequenceId,
		Status:     "ok",
	}

	render.Status(r, 200)
	render.JSON(w, r, res)
	return
}

func (sh *SequenceHandler) UpdateSequence(w http.ResponseWriter, r *http.Request) {
	sh.logger.Info("UpdateSequence request received")
	updateSequenceRequest, err := NewUpdateSequenceRequestFromHttpRequest(r)
	if err != nil {
		appErr := errors.NewAppError(http.StatusBadRequest, "invalid request payload", err)
		sh.logger.Error("error decoding request", zap.Error(appErr))
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	sequenceId, err := sh.sequenceService.UpdateSequence(r.Context(), updateSequenceRequest)
	if err != nil {
		appErr := errors.NewAppError(http.StatusInternalServerError, "failed to update sequence", err)
		sh.logger.Error("error processing request", zap.Error(appErr))
		http.Error(w, "An error occurred", http.StatusInternalServerError)
		return
	}

	res := models.UpdateSequenceResponse{
		SequenceID: sequenceId,
		Status:     "ok",
	}

	render.Status(r, 200)
	render.JSON(w, r, res)
	return
}

func (sh *SequenceHandler) UpdateStep(w http.ResponseWriter, r *http.Request) {
	sh.logger.Info("UpdateStep request received")
	updateStepRequest, err := NewUpdateStepRequestFromHttpRequest(r)
	if err != nil {
		appErr := errors.NewAppError(http.StatusBadRequest, "invalid request payload", err)
		sh.logger.Error("error decoding request", zap.Error(appErr))
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	sequenceId, stepId, err := sh.sequenceService.UpdateStep(r.Context(), updateStepRequest)
	if err != nil {
		appErr := errors.NewAppError(http.StatusInternalServerError, "failed to update step", err)
		sh.logger.Error("error processing request", zap.Error(appErr))
		http.Error(w, "An error occurred", http.StatusInternalServerError)
		return
	}

	res := models.UpdateStepResponse{
		SequenceID: sequenceId,
		StepID:     stepId,
		Status:     "ok",
	}

	render.Status(r, 200)
	render.JSON(w, r, res)
	return
}

func (sh *SequenceHandler) DeleteStep(w http.ResponseWriter, r *http.Request) {
	sh.logger.Info("DeleteStep request received")
	deleteStepRequest, err := NewDeleteStepRequestFromHttpRequest(r)
	if err != nil {
		appErr := errors.NewAppError(http.StatusBadRequest, "invalid request payload", err)
		sh.logger.Error("error decoding request", zap.Error(appErr))
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	sequenceId, stepId, err := sh.sequenceService.DeleteStep(r.Context(), deleteStepRequest)
	if err != nil {
		appErr := errors.NewAppError(http.StatusInternalServerError, "failed to delete step", err)
		sh.logger.Error("error processing request", zap.Error(appErr))
		http.Error(w, "an error occurred", http.StatusInternalServerError)
		return
	}

	res := models.DeleteStepResponse{
		SequenceID: sequenceId,
		StepID:     stepId,
		Status:     "ok",
	}

	render.Status(r, 200)
	render.JSON(w, r, res)
	return
}
