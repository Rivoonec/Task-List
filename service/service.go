package service

import (
	"ToDoList/store"
	"fmt"
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

func (t *TaskService) CreateTask(taskName string) error {

	// 1. Валидация (бизнес логика)
	if taskName == "" || utf8.RuneCountInString(taskName) <= 3 {
		return fmt.Errorf("имя задачи должно быть не менее трёх символов")
	}

	// 2. Добавляем задачу в локальный список
	t.Tasks = append(t.Tasks, store.Task{
		Task:   taskName,
		Status: store.StatusNotDone,
	})

	// 3. Сохраняем через интерфейс
	return t.Store.Save(t.Tasks)
}

func (t *TaskService) DeleteTask(num int) error {
	num--

	// Валидация (достаточная проверка)
	if num < 0 || num >= len(t.Tasks) {
		return fmt.Errorf("неверный номер задачи")
	}

	// Прямой доступ к элементу (без цикла)
	t.Tasks = append(t.Tasks[:num], t.Tasks[num+1:]...)
	fmt.Printf("Удалили задачу номер %v.\n", num+1) // +1 чтобы показать исходный номер

	// Сохраняем изменения
	return t.Store.Save(t.Tasks)
}

func (t *TaskService) UpdateTask(num int, name string) error {
	num--

	// 1. Валидация
	if num >= len(t.Tasks) || num < 0 {
		return fmt.Errorf("неверный номер задачи")
	}
	if name == "" || utf8.RuneCountInString(name) <= 3 {
		return fmt.Errorf("имя задачи должно быть не менее трёх символов")
	}

	// 2. Поиск задачи и изменение имени
	t.Tasks[num].Task = name

	// 3. Сохранение изменений
	return t.Store.Save(t.Tasks)
}

func (t *TaskService) UpdateTaskStatus(num int, statusCode store.TaskStatus) error {
	num--

	// 1. Валидация
	if num >= len(t.Tasks) || num < 0 {
		return fmt.Errorf("неверный номер задачи")
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

/*func (t *TaskService) ReadTaskList() {
	fmt.Println("Список задач:")
	for index, task := range t.Tasks {
		var taskStatusString string
		switch task.Status {
		case 0:
			taskStatusString = "Not done"
		case 1:
			taskStatusString = "In progress"
		case 2:
			taskStatusString = "Done"
		default:
			fmt.Println("Какая-то ошибка.")
		}
		fmt.Printf("№%v. Прогресс: %v  Задача: %v\n", index+1, taskStatusString, task.Task)
	}
}*/
