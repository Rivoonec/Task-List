package service

import (
	"fmt"
	"unicode/utf8"
	"ToDoList/store"
)

type TaskService struct {
	store store.TaskStore
	tasks []store.Task
}


func (t *TaskService) CreateTask(taskName string) error {
	if taskName == "" || utf8.RuneCountInString(taskName) <= 3 {
		return fmt.Errorf("Имя задачи должно быть не менее трёх символов.")
	}

	t.tasks = append(t.tasks, store.Task{Task: taskName})

	return nil
}

// Переписать всё, что ниже в методы
func (t *TaskService) DeleteTask(num int) error {
	for {
		num = num - 1
		for index := range t.tasks {
			if index == num {
				t.tasks= append(t.tasks[:index], t.tasks[index+1:]...)
				fmt.Printf("Удалили задачу номер %v .\n", num)
				return nil
			}
		}
		// Превратить в ошибку и завершать функцию, а не бесконечный ввод?
		fmt.Printf("Не нашли такую задачу.\nПовторите ввод.\nВведите номер задачи, которую хотите удалить.\n")
		num = getNum()
	}
}

func (t *TaskService) UpdateTask(num int) []Task {
	for {
		num = num - 1
		for index := range t.tasks {
			if index == num {
				for {
					fmt.Println("Нашёл задачу.\nВведите новое содержание.")
					input, err := reader.ReadString('\n')
					if err != nil && err != io.EOF {
						fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
						continue
					}
					if utf8.RuneCountInString(input) >= 3 || input != "" {
						t.tasks[index].Task = input
						return t.tasks
					} else {
						fmt.Println("Ошибка ввода или имя менее трёх символов.")
					}
				}
			}
		}
		fmt.Printf("Не нашли такую задачу.\nПовторите ввод.\nВведите номер задачи, которую хотите изменить.\n")
		num = getNum()
	}
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
