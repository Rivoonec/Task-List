package service

import (
	"fmt"
	"unicode/utf8"
	"ToDoList/store"
)
//Переработать код, перенести валидатицаю куда-нибудь

// TaskService содержит бизнес-логику приложения
type TaskService struct {
	// store - это НЕ конкретная реализация, а ИНТЕРФЕЙС
    // Это контракт, который гарантирует,
	// что у store есть методы Load() и Save()
	store store.TaskStore
	tasks []store.Task
}

// NewTaskService - конструктор, который принимает ЛЮБОЕ хранилище, 
// удовлетворяющее интерфейсу TaskStore
func NewTaskService(store store.TaskStore)(*TaskService, error){

	// Не знаем, что под капотом в store.Load(), просто хотим задачи
	tasks, err := store.Load()
	if err != nil {
		return nil, err
	}
	return &TaskService{
		store : store, // Сохраняем интерфейс, а не реализацию
		tasks : tasks, // Локальная копия задач в памяти
	}, nil
}
func (t *TaskService) CreateTask(taskName string) error {

	// 1. Валидация (бизнес логика)
	if taskName == "" || utf8.RuneCountInString(taskName) <= 3 {
		return fmt.Errorf("Имя задачи должно быть не менее трёх символов.")
	}

	// 2. Добавляем задачу в локальный список
	t.tasks = append(t.tasks, store.Task{
		Task: taskName,
		Status: store.StatusNotDone,
		})

	// 3. Сохраняем через интерфейс
	return t.store.Save(t.tasks)
}


func (t *TaskService) DeleteTask(num int) error {
	num--

	// 1. Валидация
	if num >= len(t.tasks) || num < 0 {
		return fmt.Errorf("Неверный номер задачи.")
	}
	
	// 2. Поиск нужной задачи и её удаление.
	for index := range t.tasks {
		if index == num {
			t.tasks= append(t.tasks[:index], t.tasks[index+1:]...)
			fmt.Printf("Удалили задачу номер %v .\n", num)

			// 3. Сохраняем изменения через интерфейс.
			return t.store.Save(t.tasks)
		}
	}

	return fmt.Errorf("Не нашли задачу с таким номером.")
}

// Переписать всё, что ниже в методы
func (t *TaskService) UpdateTask(num int, name string) error {
	num--

	// 1. Валидация
	if num >= len(t.tasks) || num < 0 {
		return fmt.Errorf("Неверный номер задачи.")
	}
	if name == "" || utf8.RuneCountInString(name) <= 3 {
		return fmt.Errorf("Имя задачи должно быть не менее трёх символов.")
	}
	// !!! Изменить на прямой доступ
	// 2. Поиск задачи и изменение имени
	for index := range t.tasks {
		if index == num {
			t.tasks[index].Task = name
			// 3. Сохранение изменений
			return t.store.Save(t.tasks)
		}
	}

	return fmt.Errorf("Не нашли задачу с таким номером.")
}
func UpdateTaskStatus(t []Task, num int, statusCode TaskStatus) []Task {
	for {
		num = num - 1
		for index := range t.tasks {
			if index == num {
				fmt.Println("Нашёл задачу.")
				t.tasks[index].Status = statusCode
				return t.tasks
			}
		}
		fmt.Printf("Не нашли такую задачу.\nПовторите ввод.\nВведите номер задачи, которую хотите изменить.\n")
		num = getNum()
	}
}
func ReadTaskList(t []Task) {
	fmt.Println("Список задач:")
	for index, task := range t.tasks {
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
}
