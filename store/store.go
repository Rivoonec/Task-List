package store

import (
	"encoding/json"
	"fmt"
	"os"
)

type TaskStatus uint8

const (
	// 0 - not done
	StatusNotDone TaskStatus = iota
	// 1 - in progress
	StatusInProgress
	// 2 - done
	StatusDone

	StorageFileName string = "list.json"
)

type Task struct {
	Task   string     `json:"task"`
	Status TaskStatus `json:"status"` // 0 - not done, 1 - in progress, 2 - done
}

// TaskStore defines contract for data operations
// "Any structure that has Load() and Save() methods is considered a storage"
type TaskStore interface {
	Save(tasks []Task) error
	Load() ([]Task, error)
}

// JSONFileStore implementation for JSON file storage
type JSONFileStore struct {
	filename string
}

// NewJSONFileStore creates new JSON file storage instance
func NewJSONFileStore(filename string) *JSONFileStore {
	return &JSONFileStore{filename: filename}
}

// Load data from JSON file, returns slice of Tasks
func (s *JSONFileStore) Load() ([]Task, error) {
	data, err := os.ReadFile(s.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, fmt.Errorf("file load error: %w", err)
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, fmt.Errorf("JSON decoding error: %w", err)
	}

	return tasks, nil
}

// Saves data to JSON file, returns error
func (s *JSONFileStore) Save(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("JSON encoding error: %w", err)
	}

	return os.WriteFile(s.filename, data, 0644)
}
