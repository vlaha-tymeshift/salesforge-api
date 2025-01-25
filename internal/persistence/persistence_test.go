package persistence_test

import (
	"context"
	"database/sql"
	"log"
	"os"
	"salesforge-api/internal/config"
	"strconv"
	"testing"

	_ "github.com/lib/pq"
	"salesforge-api/internal/models"
	"salesforge-api/internal/persistence"
)

var db *sql.DB

func TestMain(m *testing.M) {
	cfg, err := config.LoadConfig(os.DirFS("../.."))
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	dsn := "user=" + cfg.TestDB.User + " password=" + cfg.TestDB.Pass + " dbname=" + cfg.TestDB.Db + " host=" + cfg.TestDB.Host + " port=" + strconv.Itoa(cfg.TestDB.Port) + " sslmode=disable"
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to connect to test database: %v", err)
	}

	// Run tests
	code := m.Run()

	// Clean up the database and close the connection
	setupTestDB()
	db.Close()

	// Exit with the test result code
	os.Exit(code)
}

func setupTestDB() {
	// Clean up the database before and after each test
	_, err := db.Exec("TRUNCATE TABLE sequences, steps RESTART IDENTITY CASCADE")
	if err != nil {
		log.Fatalf("failed to clean test database: %v", err)
	}
}

func TestAddSequence_Integration(t *testing.T) {
	setupTestDB()
	repo := persistence.NewSequenceRepository(db)

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
			EligibleStartTime: 1706132001,
			EligibleEndTime:   1706304801,
		},
	}

	ctx := context.Background()
	sequenceId, err := repo.AddSequence(ctx, &sequence, &steps)
	if err != nil {
		t.Fatalf("failed to add sequence: %v", err)
	}

	if sequenceId == 0 {
		t.Fatalf("expected sequenceId to be non-zero")
	}
}

func TestUpdateSequence_Integration(t *testing.T) {
	setupTestDB()
	repo := persistence.NewSequenceRepository(db)

	// First, add a sequence to update
	sequence := models.Sequence{
		AccountID:                    1,
		SequenceName:                 "Test Sequence",
		SequenceOpenTrackingEnabled:  false,
		SequenceClickTrackingEnabled: true,
	}
	steps := []models.Step{
		{
			StepEmailSubject:  "Subject 1",
			StepEmailBody:     "Body 1",
			WaitDays:          1,
			EligibleStartTime: 1706132001,
			EligibleEndTime:   1706304801,
		},
	}

	ctx := context.Background()
	sequenceId, err := repo.AddSequence(ctx, &sequence, &steps)
	if err != nil {
		t.Fatalf("failed to add sequence: %v", err)
	}

	enabled := true
	update := models.UpdateSequenceRequest{
		AccountID:                    1,
		SequenceID:                   sequenceId,
		SequenceOpenTrackingEnabled:  &enabled,
		SequenceClickTrackingEnabled: &enabled,
	}

	updatedSequenceId, err := repo.UpdateSequence(ctx, &update)
	if err != nil {
		t.Fatalf("failed to update sequence: %v", err)
	}

	if updatedSequenceId != sequenceId {
		t.Fatalf("expected updatedSequenceId to be %d, got %d", sequenceId, updatedSequenceId)
	}
}

func TestUpdateStep_Integration(t *testing.T) {
	setupTestDB()
	repo := persistence.NewSequenceRepository(db)

	// First, add a sequence and step to update
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
			EligibleStartTime: 1706132001,
			EligibleEndTime:   1706304801,
		},
	}

	ctx := context.Background()
	sequenceId, err := repo.AddSequence(ctx, &sequence, &steps)
	if err != nil {
		t.Fatalf("failed to add sequence: %v", err)
	}

	update := models.UpdateStepRequest{
		AccountID:        1,
		SequenceID:       sequenceId,
		StepID:           1,
		StepEmailSubject: "Updated Subject",
		StepEmailBody:    "Updated Body",
	}

	updatedSequenceId, updatedStepId, err := repo.UpdateStep(ctx, &update)
	if err != nil {
		t.Fatalf("failed to update step: %v", err)
	}

	if updatedSequenceId != sequenceId || updatedStepId != 1 {
		t.Fatalf("expected updatedSequenceId to be %d and updatedStepId to be 1, got %d and %d", sequenceId, updatedSequenceId, updatedStepId)
	}
}

func TestDeleteStep_Integration(t *testing.T) {
	setupTestDB()
	repo := persistence.NewSequenceRepository(db)

	// First, add a sequence and step to delete
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
			EligibleStartTime: 1706132001,
			EligibleEndTime:   1706304801,
		},
	}

	ctx := context.Background()
	sequenceId, err := repo.AddSequence(ctx, &sequence, &steps)
	if err != nil {
		t.Fatalf("failed to add sequence: %v", err)
	}

	delete := models.DeleteStepRequest{
		AccountID:  1,
		SequenceID: sequenceId,
		StepID:     1,
	}

	deletedSequenceId, deletedStepId, err := repo.DeleteStep(ctx, &delete)
	if err != nil {
		t.Fatalf("failed to delete step: %v", err)
	}

	if deletedSequenceId != sequenceId || deletedStepId != 1 {
		t.Fatalf("expected deletedSequenceId to be %d and deletedStepId to be 1, got %d and %d", sequenceId, deletedSequenceId, deletedStepId)
	}
}
