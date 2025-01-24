package models

type Sequence struct {
	AccountID                    int64  `json:"account_id"`
	CreatedAt                    int64  `json:"created_at"`
	UpdatedAt                    int64  `json:"updated_at"`
	SequenceID                   int64  `json:"sequence_id"`
	SequenceName                 string `json:"sequence_name"`
	SequenceOpenTrackingEnabled  bool   `json:"sequence_open_tracking_enabled"`
	SequenceClickTrackingEnabled bool   `json:"sequence_click_tracking_enabled"`
}

type Step struct {
	StepID            int64  `json:"step_id"`
	SequenceID        int64  `json:"sequence_id"`
	CreatedAt         int64  `json:"created_at"`
	UpdatedAt         int64  `json:"updated_at"`
	StepEmailSubject  string `json:"step_email_subject"`
	StepEmailBody     string `json:"step_email_body"`
	WaitDays          int    `json:"wait_days"`
	EligibleStartTime int64  `json:"eligible_start_time"`
	EligibleEndTime   int64  `json:"eligible_end_time"`
}

type AddSequenceRequest struct {
	Sequence
	Steps []Step `json:"steps"`
}

func (asr *AddSequenceRequest) Validate() (bool, []string) {
	var invalidFields []string
	var isValid bool = true

	if asr.AccountID <= 0 {
		invalidFields = append(invalidFields, "account_id")
		isValid = false
	}

	if asr.SequenceName == "" {
		invalidFields = append(invalidFields, "sequence_name")
		isValid = false
	}

	return isValid, invalidFields
}

type AddSequenceResponse struct {
	SequenceID int64  `json:"sequence_id"`
	Status     string `json:"status"`
}

type UpdateSequenceRequest struct {
	AccountID                    int64 `json:"account_id"`
	SequenceID                   int64 `json:"sequence_id"`
	SequenceOpenTrackingEnabled  *bool `json:"sequence_open_tracking_enabled"`
	SequenceClickTrackingEnabled *bool `json:"sequence_click_tracking_enabled"`
}

func (usr *UpdateSequenceRequest) Validate() (bool, []string) {
	var invalidFields []string
	var isValid bool = true

	if usr.AccountID <= 0 {
		invalidFields = append(invalidFields, "account_id")
		isValid = false
	}

	if usr.SequenceID <= 0 {
		invalidFields = append(invalidFields, "sequence_id")
		isValid = false
	}

	if usr.SequenceOpenTrackingEnabled == nil {
		invalidFields = append(invalidFields, "sequence_open_tracking_enabled")
		isValid = false
	}

	if usr.SequenceClickTrackingEnabled == nil {
		invalidFields = append(invalidFields, "sequence_click_tracking_enabled")
		isValid = false
	}

	return isValid, invalidFields
}

type UpdateSequenceResponse struct {
	SequenceID int64  `json:"sequence_id"`
	Status     string `json:"status"`
}

type UpdateStepRequest struct {
	AccountID        int64  `json:"account_id"`
	StepID           int64  `json:"step_id"`
	SequenceID       int64  `json:"sequence_id"`
	StepEmailSubject string `json:"step_email_subject"`
	StepEmailBody    string `json:"step_email_body"`
}

func (usr *UpdateStepRequest) Validate() (bool, []string) {
	var invalidFields []string
	var isValid bool = true

	if usr.AccountID <= 0 {
		invalidFields = append(invalidFields, "account_id")
		isValid = false
	}

	if usr.StepID <= 0 {
		invalidFields = append(invalidFields, "step_id")
		isValid = false
	}

	if usr.SequenceID <= 0 {
		invalidFields = append(invalidFields, "sequence_id")
		isValid = false
	}

	if usr.StepEmailSubject == "" {
		invalidFields = append(invalidFields, "step_email_subject")
		isValid = false
	}

	if usr.StepEmailBody == "" {
		invalidFields = append(invalidFields, "step_email_body")
		isValid = false
	}

	return isValid, invalidFields
}

type UpdateStepResponse struct {
	SequenceID int64  `json:"sequence_id"`
	StepID     int64  `json:"step_id"`
	Status     string `json:"status"`
}

type DeleteStepRequest struct {
	AccountID  int64 `json:"account_id"`
	StepID     int64 `json:"step_id"`
	SequenceID int64 `json:"sequence_id"`
}

func (dsr *DeleteStepRequest) Validate() (bool, []string) {
	var invalidFields []string
	var isValid bool = true

	if dsr.AccountID <= 0 {
		invalidFields = append(invalidFields, "account_id")
		isValid = false
	}

	if dsr.StepID <= 0 {
		invalidFields = append(invalidFields, "step_id")
		isValid = false
	}

	if dsr.SequenceID <= 0 {
		invalidFields = append(invalidFields, "sequence_id")
		isValid = false
	}

	return isValid, invalidFields
}

type DeleteStepResponse struct {
	SequenceID int64  `json:"sequence_id"`
	StepID     int64  `json:"step_id"`
	Status     string `json:"status"`
}
