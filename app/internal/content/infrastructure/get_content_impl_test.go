package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/sky0621/my-test-project/app/internal/content/application/query"
	"github.com/sky0621/my-test-project/app/internal/shared/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupTestData(t *testing.T, db *sql.DB) (string, error) {
	t.Helper()

	tx, err := db.Begin()
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	contentID := uuid.New().String()

	_, err = tx.Exec("INSERT INTO contents (id, name) VALUES (UUID_TO_BIN(?), ?)", contentID, "Test Content")
	if err != nil {
		return "", fmt.Errorf("failed to insert content: %w", err)
	}

	program1ID := uuid.New().String()
	program2ID := uuid.New().String()

	_, err = tx.Exec("INSERT INTO programs (id, question, answer, content_id) VALUES (UUID_TO_BIN(?), ?, ?, UUID_TO_BIN(?))",
		program1ID, "Test Question 1", "Test Answer 1", contentID)
	if err != nil {
		return "", fmt.Errorf("failed to insert program 1: %w", err)
	}

	_, err = tx.Exec("INSERT INTO programs (id, question, answer, content_id) VALUES (UUID_TO_BIN(?), ?, ?, UUID_TO_BIN(?))",
		program2ID, "Test Question 2", "Test Answer 2", contentID)
	if err != nil {
		return "", fmt.Errorf("failed to insert program 2: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return contentID, nil
}

func cleanupTestData(t *testing.T, db *sql.DB, contentID string) {
	t.Helper()

	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	_, err = tx.Exec("DELETE FROM programs WHERE content_id = UUID_TO_BIN(?)", contentID)
	if err != nil {
		t.Fatalf("Failed to delete programs: %v", err)
	}

	_, err = tx.Exec("DELETE FROM contents WHERE id = UUID_TO_BIN(?)", contentID)
	if err != nil {
		t.Fatalf("Failed to delete content: %v", err)
	}

	if err = tx.Commit(); err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}
}

func TestGetContentImpl_Exec(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	db := test.ConnectTestDB(t)
	defer db.Close()

	contentID, err := setupTestData(t, db)
	if err != nil {
		t.Fatalf("Failed to setup test data: %v", err)
	}
	defer cleanupTestData(t, db, contentID)

	getContent := NewGetContent(db)

	tests := []struct {
		name    string
		id      string
		want    func(t *testing.T, result query.GetContentReadModel, err error)
		wantErr bool
	}{
		{
			name: "Valid content ID",
			id:   contentID,
			want: func(t *testing.T, result query.GetContentReadModel, err error) {
				assert.NoError(t, err)
				assert.Equal(t, "Test Content", result.Name)
				assert.Len(t, result.Programs, 2)

				questions := []string{result.Programs[0].Question, result.Programs[1].Question}
				answers := []string{result.Programs[0].Answer, result.Programs[1].Answer}

				assert.Contains(t, questions, "Test Question 1")
				assert.Contains(t, questions, "Test Question 2")
				assert.Contains(t, answers, "Test Answer 1")
				assert.Contains(t, answers, "Test Answer 2")
			},
			wantErr: false,
		},
		{
			name: "Invalid content ID",
			id:   uuid.New().String(),
			want: func(t *testing.T, result query.GetContentReadModel, err error) {
				assert.NoError(t, err)
				assert.True(t, result.IsEmpty())
			},
			wantErr: false,
		},
		{
			name: "Empty content ID",
			id:   "",
			want: func(t *testing.T, result query.GetContentReadModel, err error) {
				assert.Error(t, err)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := getContent.Exec(context.Background(), tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetContentImpl.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				tt.want(t, result, err)
			}
		})
	}
}
