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

func TestTaskService_CreateTask(t *testing.T) {
	testLocaleManager := locale.NewManager()

	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		store         store.TaskStore
		localeManager *locale.Manager
		// Named input parameters for target function.
		taskName string
		wantErr  bool
	}{
		{
			name:          "Empty name - should fail validation",
			store:         &InMemoryStore{},
			localeManager: testLocaleManager,
			taskName:      "",
			wantErr:       true,
		},
		{
			name:          "Invalid(short) name - should fail validation",
			store:         &InMemoryStore{},
			localeManager: testLocaleManager,
			taskName:      "ab",
			wantErr:       true,
		},
		{
			name:          "Valid name - should succeed",
			store:         &InMemoryStore{},
			localeManager: testLocaleManager,
			taskName:      "test",
			wantErr:       false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ta, err := NewTaskService(tt.store, tt.localeManager)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			gotErr := ta.CreateTask(tt.taskName)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CreateTask() error = %v, wantErr %v", gotErr, tt.wantErr)
				}
				return
			}

			if tt.wantErr {
				t.Fatal("CreateTask() succeeded unexpectedly")
			}
		})
	}
}

func TestTaskService_DeleteTask(t *testing.T) {
	testLocaleManager := locale.NewManager()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		initialTasks  []store.Task // to init inmemory store
		localeManager *locale.Manager
		// Named input parameters for target function.
		num           int
		wantErr       bool
		expectedTasks []store.Task // if no err
	}{
		{
			name: "Num out of slice - too high - should fail validation",
			initialTasks: []store.Task{
				{Task: "first", Status: store.StatusDone},
				{Task: "second", Status: store.StatusNotDone},
			},
			localeManager: testLocaleManager,
			num:           3,
			wantErr:       true,
			expectedTasks: nil,
		},
		{
			name: "Num out of slice - too low (0) - should fail validation",
			initialTasks: []store.Task{
				{Task: "first", Status: store.StatusDone},
				{Task: "second", Status: store.StatusNotDone},
			},
			localeManager: testLocaleManager,
			num:           0,
			wantErr:       true,
			expectedTasks: nil,
		},
		{
			name: "Valid num - should succeed",
			initialTasks: []store.Task{
				{Task: "first", Status: store.StatusDone},
				{Task: "second", Status: store.StatusNotDone},
			},
			localeManager: testLocaleManager,
			num:           1,
			wantErr:       false,
			expectedTasks: []store.Task{
				{Task: "second", Status: store.StatusNotDone},
			},
		},
		{
			name: "Delete only task - should succeed",
			initialTasks: []store.Task{
				{Task: "only", Status: store.StatusNotDone},
			},
			num:           1,
			wantErr:       false,
			expectedTasks: []store.Task{},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inMemoryStore := &InMemoryStore{tasks: tt.initialTasks}
			ta, err := NewTaskService(inMemoryStore, tt.localeManager)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			gotErr := ta.DeleteTask(tt.num)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("DeleteTask(%d) error = %v, wantErr %v", tt.num, gotErr, tt.wantErr)
				}
				return
			}

			if tt.wantErr {
				t.Fatal("DeleteTask() succeeded unexpectedly")
			}

			// Если ошибка не ожидалась, проверим результат
			if !tt.wantErr {
				// Проверяем длину списка задач в сервисе
				if len(ta.Tasks) != len(tt.expectedTasks) {
					t.Errorf("DeleteTask(%d) resulted in %d tasks, expected %d", tt.num, len(ta.Tasks), len(tt.expectedTasks))
					return // Если длина не совпадает, дальше проверять бессмысленно
				}

				// Проверяем длину списка задач, возвращаемого GetAllTasks
				allTasks := ta.GetAllTasks() // Вызываем тестируемый метод
				if len(allTasks) != len(tt.expectedTasks) {
					t.Errorf("DeleteTask(%d) - GetAllTasks() returned %d tasks, expected %d", tt.num, len(allTasks), len(tt.expectedTasks))
					return // Если длина не совпадает, дальше проверять бессмысленно
				}

				// Проверяем содержимое списка задач, возвращаемого GetAllTasks, поэлементно
				for i, expectedTask := range tt.expectedTasks {
					if i >= len(allTasks) {
						// Это условие не должно сработать, если длина совпала, но на всякий случай
						t.Errorf("DeleteTask(%d) - index %d out of bounds in GetAllTasks() result", tt.num, i)
						break
					}
					actualTaskFromGetAll := allTasks[i]
					if actualTaskFromGetAll.Task != expectedTask.Task || actualTaskFromGetAll.Status != expectedTask.Status {
						t.Errorf("DeleteTask(%d) - task at index %d from GetAllTasks() mismatch: got {Task: %s, Status: %v}, want {Task: %s, Status: %v}",
							tt.num, i, actualTaskFromGetAll.Task, actualTaskFromGetAll.Status, expectedTask.Task, expectedTask.Status)
					}
				}

				// Проверяем длину списка задач в хранилище
				if len(inMemoryStore.tasks) != len(tt.expectedTasks) {
					t.Errorf("DeleteTask(%d) - store has %d tasks after Save, expected %d", tt.num, len(inMemoryStore.tasks), len(tt.expectedTasks))
					return // Если длина не совпадает, дальше проверять бессмысленно
				}
			}
		})
	}
}
