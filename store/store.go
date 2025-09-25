package store

import (
	"encoding/json"
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

// Определеяет контракт для работы с данными
// "Любая структура, у которой 
// есть методы Load() и Save(), считается хранилищем"
type TaskStore interface {
	Save(tasks []Task) error
	Load() ([]Task, error)
}
// Реализация для JSON файла
type JSONFileStore struct {
	filename string
}

// Новый экземпляр хранилища JSON (имя файла)
func NewJSONFileStore(filename string) *JSONFileStore {
	return &JSONFileStore{filename: filename}
}

// Loads data from JSON file, returns slice of Tasks
func (s *JSONFileStore) Load() ([]Task, error) {
	data, err := os.ReadFile(s.filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// Saves data to JSON file, returns error
func (s *JSONFileStore) Save(taskList []Task) error {
	data, err := json.Marshal(taskList)
	if err != nil {
		return err
	}

	return os.WriteFile(s.filename, data, 0644)
}