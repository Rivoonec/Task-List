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

func TestTaskService_UpdateTask(t *testing.T) {
	testLocaleManager := locale.NewManager()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		initialTasks  []store.Task // to init inmemory store
		localeManager *locale.Manager
		// Named input parameters for target function.
		taskName      string
		num           int
		wantErr       bool
		expectedTasks []store.Task // if no err
	}{
		{
			name: "Num out of slice - too high, valid taskName - should fail validation",
			initialTasks: []store.Task{
				{Task: "first", Status: store.StatusDone},
				{Task: "second", Status: store.StatusNotDone},
			},
			localeManager: testLocaleManager,
			taskName:      "test",
			num:           3,
			wantErr:       true,
			expectedTasks: nil,
		},
		{
			name: "Num out of slice - too low (0), valid taskName - should fail validation",
			initialTasks: []store.Task{
				{Task: "first", Status: store.StatusDone},
				{Task: "second", Status: store.StatusNotDone},
			},
			localeManager: testLocaleManager,
			taskName:      "test",
			num:           0,
			wantErr:       true,
			expectedTasks: nil,
		},
		{
			name: "Valid num, valid taskName - should succeed",
			initialTasks: []store.Task{
				{Task: "first", Status: store.StatusDone},
				{Task: "second", Status: store.StatusNotDone},
			},
			localeManager: testLocaleManager,
			taskName:      "test",
			num:           1,
			wantErr:       false,
			expectedTasks: []store.Task{
				{Task: "test", Status: store.StatusDone},
				{Task: "second", Status: store.StatusNotDone},
			},
		},
		{
			name: "Valid num, invalid taskName - should fail validation",
			initialTasks: []store.Task{
				{Task: "only", Status: store.StatusNotDone},
			},
			localeManager: testLocaleManager,
			taskName:      "ab",
			num:           1,
			wantErr:       true,
			expectedTasks: nil,
		},
		{
			name: "Valid num, empy taskName - should fail validation",
			initialTasks: []store.Task{
				{Task: "only", Status: store.StatusNotDone},
			},
			localeManager: testLocaleManager,
			taskName:      "",
			num:           1,
			wantErr:       true,
			expectedTasks: nil,
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
			gotErr := ta.UpdateTask(tt.num, tt.taskName)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("UpdateTask() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UpdateTask() succeeded unexpectedly")
			}
			if !tt.wantErr{
				actualTasks := ta.GetAllTasks()
    			if len(actualTasks) != len(tt.expectedTasks) {
    			    t.Errorf("UpdateTask() result length = %d, want %d", len(actualTasks), len(tt.expectedTasks))
    			}
			
    			for i := range actualTasks {
    			    if actualTasks[i].Task != tt.expectedTasks[i].Task {
    			        t.Errorf("UpdateTask() task[%d].Task = %q, want %q", i, actualTasks[i].Task, tt.expectedTasks[i].Task)
    			    }
    			    if actualTasks[i].Status != tt.expectedTasks[i].Status {
    			        t.Errorf("UpdateTask() task[%d].Status = %v, want %v", i, actualTasks[i].Status, tt.expectedTasks[i].Status)
    			    }
    			}
			}
		})
	}
}
func TestTaskService_UpdateTaskStatus(t *testing.T) {
	testLocaleManager := locale.NewManager()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		initialTasks  []store.Task // to init inmemory store
		localeManager *locale.Manager
		// Named input parameters for target function.
		taskStatus    store.TaskStatus
		num           int
		wantErr       bool
		expectedTasks []store.Task // if no err
	}{
		{
			name: "Num out of slice - too high, valid statusCode - should fail validation",
			initialTasks: []store.Task{
				{Task: "first", Status: store.StatusDone},
				{Task: "second", Status: store.StatusNotDone},
			},
			localeManager: testLocaleManager,
			taskStatus:    store.StatusInProgress,
			num:           3,
			wantErr:       true,
			expectedTasks: nil,
		},
		{
			name: "Num out of slice - too low (0), valid statusCode - should fail validation",
			initialTasks: []store.Task{
				{Task: "first", Status: store.StatusDone},
				{Task: "second", Status: store.StatusNotDone},
			},
			localeManager: testLocaleManager,
			taskStatus:    store.StatusInProgress,
			num:           0,
			wantErr:       true,
			expectedTasks: nil,
		},
		{
			name: "Valid num, valid statusCode - should succeed",
			initialTasks: []store.Task{
				{Task: "first", Status: store.StatusDone},
				{Task: "second", Status: store.StatusNotDone},
			},
			localeManager: testLocaleManager,
			taskStatus:    store.StatusInProgress,
			num:           1,
			wantErr:       false,
			expectedTasks: []store.Task{
				{Task: "first", Status: store.StatusDone},
				{Task: "second", Status: store.StatusInProgress},
			},
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
			gotErr := ta.UpdateTaskStatus(tt.num, tt.taskStatus)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("UpdateTaskStatus() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("UpdateTaskStatus() succeeded unexpectedly")
			}
		})
	}
}

func TestTaskService_validateTaskName(t *testing.T) {
	testLocaleManager := locale.NewManager()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		localeManager *locale.Manager
		// Named input parameters for target function.
		taskName   string
		wantErr bool
	}{
		{
			name:          "Empty name - should fail validation",
			localeManager: testLocaleManager,
			taskName:      "",
			wantErr:       true,
		},
		{
			name:          "Invalid(short) name - should fail validation",
			localeManager: testLocaleManager,
			taskName:      "ab",
			wantErr:       true,
		},
		{
			name:          "Valid name - should succeed",
			localeManager: testLocaleManager,
			taskName:      "test",
			wantErr:       false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inMemoryStore := &InMemoryStore{}
			ta, err := NewTaskService(inMemoryStore, tt.localeManager)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			gotErr := ta.validateTaskName(tt.taskName)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("validateTaskName() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("validateTaskName() succeeded unexpectedly")
			}
		})
	}
}

func TestTaskService_validateTaskNumber(t *testing.T) {
	testLocaleManager := locale.NewManager()
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		localeManager *locale.Manager
		// Named input parameters for target function.
		num     int 
		wantErr bool
	}{
		{
			name: "Num out of slice - too high - should fail validation",
			localeManager: testLocaleManager,
			num:           2,
			wantErr:       true,
		},
		{
			name: "Num out of slice - too low (0) - should fail validation",
			localeManager: testLocaleManager,
			num:           -1,
			wantErr:       true,
		},
		{
			name: "Valid num - should succeed",
			localeManager: testLocaleManager,
			num:           0,
			wantErr:       false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inMemoryStore := &InMemoryStore{
			    tasks: []store.Task{
			        {Task: "first", Status: store.StatusDone},
			        {Task: "second", Status: store.StatusNotDone},
			    },
			}
			ta, err := NewTaskService(inMemoryStore, tt.localeManager)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			gotErr := ta.validateTaskNumber(tt.num)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("validateTaskNumber() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("validateTaskNumber() succeeded unexpectedly")
			}
		})
	}
}
