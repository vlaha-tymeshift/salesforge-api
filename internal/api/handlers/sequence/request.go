package sequence

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"salesforge-api/internal/models"
)

const (
	RequestDecodeError     = "requestDecodeError"
	InvalidParametersError = "invalidParametersError"
)

func NewAddSequenceRequestFromHttpRequest(r *http.Request) (*models.AddSequenceRequest, error) {
	addSequenceRequest := &models.AddSequenceRequest{}
	err := json.NewDecoder(r.Body).Decode(addSequenceRequest)
	if err != nil {
		return nil, errors.New(RequestDecodeError)
	}

	isValid, invalidFields := addSequenceRequest.Validate()
	if !isValid {
		return nil, fmt.Errorf("%s: %v", InvalidParametersError, invalidFields)
	}

	return addSequenceRequest, nil
}

func NewUpdateSequenceRequestFromHttpRequest(r *http.Request) (*models.UpdateSequenceRequest, error) {
	updateSequenceRequest := &models.UpdateSequenceRequest{}
	err := json.NewDecoder(r.Body).Decode(updateSequenceRequest)
	if err != nil {
		return nil, errors.New(RequestDecodeError)
	}

	isValid, invalidFields := updateSequenceRequest.Validate()
	if !isValid {
		return nil, fmt.Errorf("%s: %v", InvalidParametersError, invalidFields)
	}

	return updateSequenceRequest, nil
}

func NewUpdateStepRequestFromHttpRequest(r *http.Request) (*models.UpdateStepRequest, error) {
	updateStepRequest := &models.UpdateStepRequest{}
	err := json.NewDecoder(r.Body).Decode(updateStepRequest)
	if err != nil {
		return nil, errors.New(RequestDecodeError)
	}

	isValid, invalidFields := updateStepRequest.Validate()
	if !isValid {
		return nil, fmt.Errorf("%s: %v", InvalidParametersError, invalidFields)
	}

	return updateStepRequest, nil
}

func NewDeleteStepRequestFromHttpRequest(r *http.Request) (*models.DeleteStepRequest, error) {
	deleteStepRequest := &models.DeleteStepRequest{}
	err := json.NewDecoder(r.Body).Decode(deleteStepRequest)
	if err != nil {
		return nil, errors.New(RequestDecodeError)
	}

	isValid, invalidFields := deleteStepRequest.Validate()
	if !isValid {
		return nil, fmt.Errorf("%s: %v", InvalidParametersError, invalidFields)
	}

	return deleteStepRequest, nil
}
