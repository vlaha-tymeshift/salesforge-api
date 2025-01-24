package sequence

import (
	"github.com/go-chi/render"
	"log"
	"net/http"
	"salesforge-api/internal/models"
	"salesforge-api/internal/service"
)

type SequenceHandler struct {
	sequenceService service.SequenceService
}

func NewSequenceHandler(sequenceService service.SequenceService) *SequenceHandler {
	return &SequenceHandler{
		sequenceService: sequenceService,
	}
}

func (sh *SequenceHandler) AddSequence(w http.ResponseWriter, r *http.Request) {
	addSequenceRequest, err := NewAddSequenceRequestFromHttpRequest(r)
	if err != nil {
		log.Printf("Error decoding request: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	sequence := addSequenceRequest.Sequence
	steps := addSequenceRequest.Steps
	sequenceId, err := sh.sequenceService.AddSequence(r.Context(), &sequence, &steps)
	if err != nil {
		log.Printf("Error processing request: %v", err)
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
	updateSequenceRequest, err := NewUpdateSequenceRequestFromHttpRequest(r)
	if err != nil {
		log.Printf("Error decoding request: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	sequenceId, err := sh.sequenceService.UpdateSequence(r.Context(), updateSequenceRequest)
	if err != nil {
		log.Printf("Error processing request: %v", err)
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
	updateStepRequest, err := NewUpdateStepRequestFromHttpRequest(r)
	if err != nil {
		log.Printf("Error decoding request: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	sequenceId, stepId, err := sh.sequenceService.UpdateStep(r.Context(), updateStepRequest)
	if err != nil {
		log.Printf("Error processing request: %v", err)
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
	deleteStepRequest, err := NewDeleteStepRequestFromHttpRequest(r)
	if err != nil {
		log.Printf("Error decoding request: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	sequenceId, stepId, err := sh.sequenceService.DeleteStep(r.Context(), deleteStepRequest)
	if err != nil {
		log.Printf("Error processing request: %v", err)
		http.Error(w, "An error occurred", http.StatusInternalServerError)
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
