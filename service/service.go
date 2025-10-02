package service

import (
	"ToDoList/locale"
	"ToDoList/store"
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)


type TaskService struct {
	// store is NOT a concrete implementation, but an INTERFACE
	// This is a contract that guarantees
	// that store has Load() and Save() methods
	Store store.TaskStore
	Tasks []store.Task
	locale *locale.Manager
}

// NewTaskService - constructor that accepts ANY storage
// that satisfies TaskStore interface
func NewTaskService(store store.TaskStore, localeManager *locale.Manager) (*TaskService, error) {
	// We don't know what's under the hood in store.Load(), we just want tasks
	tasks, err := store.Load()
	if err != nil {
		return nil, fmt.Errorf(localeManager.Get("file_load_error"), err)
	}
	return &TaskService{
		Store:  store, // We save interface, not implementation
		Tasks:  tasks, // Local copy of tasks in memory
		locale: localeManager,
	}, nil
}

func (t *TaskService) CreateTask(name string) error {

	// 1. Validation
	if err := t.validateTaskName(name); err != nil {
		return fmt.Errorf(t.locale.Get("validation_error"), err)
	}

	// 2. Add task to local list
	t.Tasks = append(t.Tasks, store.Task{
		Task:   name,
		Status: store.StatusNotDone,
	})

	// 3. Save through interface
	if err := t.Store.Save(t.Tasks); err != nil {
		// Wrap store error with localized message
		return fmt.Errorf(t.locale.Get("file_save_error"), err)
	}
	return nil
}

func (t *TaskService) DeleteTask(num int) error {
	num--

	// Validation
	if err := t.validateTaskNumber(num); err != nil {
		return fmt.Errorf(t.locale.Get("validation_error"), err)
	}


	// Direct access to element (without loop)
	t.Tasks = append(t.Tasks[:num], t.Tasks[num+1:]...)

	// 3. Save through interface
	if err := t.Store.Save(t.Tasks); err != nil {
		// Wrap store error with localized message
		return fmt.Errorf(t.locale.Get("file_save_error"), err)
	}
	return nil
}

func (t *TaskService) UpdateTask(num int, name string) error {
	num--

	// 1. Validation
	if err := t.validateTaskNumber(num); err != nil {
		return fmt.Errorf(t.locale.Get("validation_error"), err)
	}

	if err := t.validateTaskName(name); err != nil {
		return fmt.Errorf(t.locale.Get("validation_error"), err)
	}

	// 2. Find task and update description
	t.Tasks[num].Task = name

	// 3. Save through interface
	if err := t.Store.Save(t.Tasks); err != nil {
		// Wrap store error with localized message
		return fmt.Errorf(t.locale.Get("file_save_error"), err)
	}
	return nil
}

func (t *TaskService) UpdateTaskStatus(num int, statusCode store.TaskStatus) error {
	num--

	// 1. Validation
	if err := t.validateTaskNumber(num); err != nil {
		return fmt.Errorf(t.locale.Get("validation_error"), err)
	}

	// 2. Find task and update status
	t.Tasks[num].Status = statusCode

	// 3. Save through interface
	if err := t.Store.Save(t.Tasks); err != nil {
		// Wrap store error with localized message
		return fmt.Errorf(t.locale.Get("file_save_error"), err)
	}
	return nil
}

func (t *TaskService) GetStatusText(status store.TaskStatus) string {
		switch status {
		case store.StatusNotDone:
			return t.locale.Get("status_option1")
		case store.StatusInProgress:
			return t.locale.Get("status_option2")
		case store.StatusDone:
			return t.locale.Get("status_option3")
		default:
			return "Unknown"
		}
}

func (t *TaskService) GetAllTasks() []store.Task {
	return t.Tasks
}

// Separate validation methods
func (t *TaskService) validateTaskNumber(num int) error {
    if num < 0 || num >= len(t.Tasks) {
		return fmt.Errorf("%s: %d", t.locale.Get("task_validation_number"), num+1)
    }
    return nil
}

func (t *TaskService) validateTaskName(name string) error {
    name = strings.TrimSpace(name)
    if name == "" {
		return errors.New(t.locale.Get("task_validation_empty"))
    }
    if utf8.RuneCountInString(name) < 3 {
		return errors.New(t.locale.Get("task_validation_short"))
    }
    return nil
}

// SetLocale allows changing locale at runtime
func (t *TaskService) SetLocale(localeManager *locale.Manager) {
	t.locale = localeManager
}

// SaveTasks allows explicit saving of tasks (useful for signal handling)
func (t *TaskService) SaveTasks() error {
	if err := t.Store.Save(t.Tasks); err != nil {
		return fmt.Errorf(t.locale.Get("service_save_error"), err)
	}
	return nil
}