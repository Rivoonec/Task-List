package service

import (
	"ToDoList/locale"
	"ToDoList/store"
	"testing"
)

// InMemoryStore — реализация store.TaskStore в памяти (для тестов)
type InMemoryStore struct {
	tasks []store.Task
}

func (s *InMemoryStore) Save(tasks []store.Task) error {
	s.tasks = tasks
	return nil
}

func (s *InMemoryStore) Load() ([]store.Task, error) {
	return s.tasks, nil
}

// Вспомогательная функция для создания TaskService в тестах
func newTestService() (*TaskService, error) {
	lm := locale.NewManager()
	lm.SetLocale("en") // используем английский — он точно есть

	store := &InMemoryStore{}
	return NewTaskService(store, lm)
}

// Тест: добавление задачи с валидным описанием (≥3 символов)
func TestAddTask_Valid(t *testing.T) {
	service, _ := newTestService()

	err := service.CreateTask("Buy groceries")
	if err != nil {
		t.Fatalf("Expected no error for valid task, got: %v", err)
	}

	tasks := service.GetAllTasks()
	if len(tasks) != 1 {
		t.Fatalf("Expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Task != "Buy groceries" {
		t.Errorf("Expected task text 'Buy groceries', got %q", tasks[0].Task)
	}
	if tasks[0].Status != 0 {
		t.Errorf("Expected status 0 (Not Done), got %d", tasks[0].Status)
	}
}

// Тест: добавление задачи с коротким описанием (<3 символов) → ошибка
func TestAddTask_TooShort(t *testing.T) {
	service, _ := newTestService()

	err := service.CreateTask("Hi")
	if err == nil {
		t.Error("Expected error for short task description, got none")
	}
}


// func TestTaskService_CreateTask(t *testing.T) {
// 	tests := []struct {
// 		name string // description of this test case
// 		// Named input parameters for receiver constructor.
// 		store         store.TaskStore
// 		localeManager *locale.Manager
// 		// Named input parameters for target function.
// 		name    string
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ta, err := NewTaskService(tt.store, tt.localeManager)
// 			if err != nil {
// 				t.Fatalf("could not construct receiver type: %v", err)
// 			}
// 			gotErr := ta.CreateTask(tt.name)
// 			if gotErr != nil {
// 				if !tt.wantErr {
// 					t.Errorf("CreateTask() failed: %v", gotErr)
// 				}
// 				return
// 			}
// 			if tt.wantErr {
// 				t.Fatal("CreateTask() succeeded unexpectedly")
// 			}
// 		})
// 	}
// }
