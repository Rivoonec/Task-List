package service

import (
	"ToDoList/store"
	"fmt"
	"strings"
	"unicode/utf8"
)

//Переработать код, перенести валидациаю куда-нибудь?

type TaskService struct {
	// store - это НЕ конкретная реализация, а ИНТЕРФЕЙС
	// Это контракт, который гарантирует,
	// что у store есть методы Load() и Save()
	Store store.TaskStore
	Tasks []store.Task
}

// NewTaskService - конструктор, который принимает ЛЮБОЕ хранилище,
// удовлетворяющее интерфейсу TaskStore
func NewTaskService(store store.TaskStore) (*TaskService, error) {
	// Не знаем, что под капотом в store.Load(), просто хотим задачи
	tasks, err := store.Load()
	if err != nil {
		return nil, err
	}
	return &TaskService{
		Store: store, // Сохраняем интерфейс, а не реализацию
		Tasks: tasks, // Локальная копия задач в памяти
	}, nil
}

func (t *TaskService) CreateTask(name string) error {

	// 1. Валидация 
	if err := t.validateTaskName(name); err != nil {
		return fmt.Errorf("данные не прошли валидацию %w", err)
	}

	// 2. Добавляем задачу в локальный список
	t.Tasks = append(t.Tasks, store.Task{
		Task:   name,
		Status: store.StatusNotDone,
	})

	// 3. Сохраняем через интерфейс
	return t.Store.Save(t.Tasks)
}

func (t *TaskService) DeleteTask(num int) error {
	num--

	// Валидация (достаточная проверка)
	if err := t.validateTaskNumber(num); err != nil {
		return fmt.Errorf("данные не прошли валидацию %w", err)
	}


	// Прямой доступ к элементу (без цикла)
	t.Tasks = append(t.Tasks[:num], t.Tasks[num+1:]...)

	// Сохраняем изменения
	return t.Store.Save(t.Tasks)
}

func (t *TaskService) UpdateTask(num int, name string) error {
	num--

	// 1. Валидация
	if err := t.validateTaskNumber(num); err != nil {
		return fmt.Errorf("данные не прошли валидацию %w", err)
	}

	if err := t.validateTaskName(name); err != nil {
		return fmt.Errorf("данные не прошли валидацию %w", err)
	}

	// 2. Поиск задачи и изменение имени
	t.Tasks[num].Task = name

	// 3. Сохранение изменений
	return t.Store.Save(t.Tasks)
}

func (t *TaskService) UpdateTaskStatus(num int, statusCode store.TaskStatus) error {
	num--

	// 1. Валидация
	if err := t.validateTaskNumber(num); err != nil {
		return fmt.Errorf("данные не прошли валидацию %w", err)
	}

	// 2. Поиск задачи и изменение статуса
	t.Tasks[num].Status = statusCode

	// 3. Сохранение изменений
	return t.Store.Save(t.Tasks)
}

func (t *TaskService) GetStatusText(status store.TaskStatus) string {
		switch status {
		case store.StatusNotDone:
			return "Not done"
		case store.StatusInProgress:
			return "In progress"
		case store.StatusDone:
			return  "Done"
		default:
			return "Unknown"
		}
}

func (t *TaskService) GetAllTasks() []store.Task {
	return t.Tasks
}

// Отдельно валидация
func (t *TaskService) validateTaskNumber(num int) error {
    if num < 0 || num >= len(t.Tasks) {
        return fmt.Errorf("неверный номер задачи: %d", num+1)
    }
    return nil
}

func (t *TaskService) validateTaskName(name string) error {
    name = strings.TrimSpace(name)
    if name == "" {
        return fmt.Errorf("имя задачи не может быть пустым")
    }
    if utf8.RuneCountInString(name) < 3 {
        return fmt.Errorf("имя задачи должно быть не менее трёх символов")
    }
    return nil
}

