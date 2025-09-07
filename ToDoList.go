package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"unicode/utf8"
)

type Task struct {
	Task   string `json:"task"`
	Status uint8  `json:"status"` // 0 - not done, 1 - in progress, 2 - done
}

func createTask(taskName string, taskList []Task) []Task {
	if taskName == "" || utf8.RuneCountInString(taskName) <= 3 {
		fmt.Println("Имя задачи должно быть не менее трёх символов.")
		return taskList
	}

	taskList = append(taskList, Task{Task: taskName})
	fmt.Println("Добавил задачу.")
	return taskList
}

func deleteTask(num int, taskList []Task) []Task {
	for {
		num = num - 1
		for index := range taskList {
			if index == num {
				taskList = append(taskList[:index], taskList[index+1:]...)
				fmt.Printf("Удалили задачу номер %v .\n", num)
				return taskList
			}
		}
		fmt.Printf("Не нашли такую задачу.\nПовторите ввод.\nВведите номер задачи, которую хотите удалить.\n")
		num = getNum()
	}
}

func updateTask(num int, taskList []Task) []Task {
	for {
		num = num - 1
		for index := range taskList {
			if index == num {
				for {
					fmt.Println("Нашёл задачу.\nВведите новое содержание.")
					input, err := reader.ReadString('\n')
					if err != nil && err != io.EOF {
						fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
						continue
					}
					if utf8.RuneCountInString(input) >= 3 || input != "" {
						taskList[index].Task = input
						return taskList
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
func updateTaskStatus(taskList []Task, num, statusCode int) []Task {
	for {
		num = num - 1
		for index := range taskList {
			if index == num {
				fmt.Println("Нашёл задачу.")
				taskList[index].Status = uint8(statusCode)
				return taskList
			}
		}
		fmt.Printf("Не нашли такую задачу.\nПовторите ввод.\nВведите номер задачи, которую хотите изменить.\n")
		num = getNum()
	}
}
func readTaskList(taskList []Task) {
	fmt.Println("Список задач:")
	for index, task := range taskList {
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

func hello(taskList []Task) {
	fmt.Printf("\nИнструкция:\n\n")
	fmt.Printf("1. Посмотри список дел.\n")
	fmt.Printf("2. Перейти в меню редактирования задач.( Имя задачи >= 3 символов)\n")
	fmt.Printf("3. Сохранить изменения в файле.\n")
	fmt.Printf("4. Ещё раз открыть инструкцию.\n")
	fmt.Printf("5. Закрыть терминал.\n\n")
	readTaskList(taskList)
}

func getMenuChoice() int {
	fmt.Print(">")
	r, _, _ := reader.ReadRune()

	if reader.Buffered() > 0 {
		reader.Discard(reader.Buffered())
	}

	if r >= '0' && r <= '9' {
		return int(r - '0')
	}
	return -1
}

func getNum() (num int) {
	for {
		input, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
			fmt.Println("Попробуйте ещё раз.")
			continue
		}
		input = strings.TrimSpace(input)

		num, err = strconv.Atoi(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка конвертации строк:", err)
			fmt.Println("Попробуйте ещё раз.")
			continue
		}
		return num
	}
}

func saveToFile(taskList []Task) error {
	data, err := json.Marshal(taskList)
	if err != nil {
		return err
	}

	return os.WriteFile("list.json", data, 0644)
}

func loadFromFile() ([]Task, error) {
	data, err := os.ReadFile("list.json")
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

var (
	reader = bufio.NewReader(os.Stdin)
)

func main() {

	taskList, err := loadFromFile()
	if err != nil {
		fmt.Println("Ошибка загрузки: ", err)
		return
	}
	// Спасибо DeepSeek за контрибьют Kappa
	// Создаем канал для перехвата сигналов ОС
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Запускаем горутину для обработки сигналов
	go func() {
		<-sigChan // Ждем сигнал
		fmt.Println("\nПолучен сигнал завершения. Сохраняем данные...")
		if err := saveToFile(taskList); err != nil {
			fmt.Println("Ошибка сохранения файла: ", err)
		}
		os.Exit(0)
	}()
	fmt.Printf("\nДобро пожаловать в терминал взаимодействия со списком дел.\nВзаимодействуй с терминалом через ввод цифр.\n\n")
	hello(taskList)

outerloop:
	for {
		fmt.Println("Ожидаю команды...\nЧтобы открыть инструкцию нажми 4.")

		switch selektorMenu := getMenuChoice(); selektorMenu {
		case 1: // Посмотри список дел
			readTaskList(taskList)
		case 2: // Перейти в меню редактирования задач.( Имя задачи >= 3 символов)
		loop:
			for {
				fmt.Println("1. Добавить задачу.\n2. Удалить задачу.\n3. Изменить задачу или её статус.\n4. Выйти в главное меню.")

				switch selektorPodMenu := getMenuChoice(); selektorPodMenu {
				case 1: // Добавить задачу
					fmt.Println("Введите задачу.")
					input, err := reader.ReadString('\n')
					if err != nil && err != io.EOF {
						fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
						os.Exit(1)
					}
					taskList = createTask(input, taskList)
				case 2: // Удалить задачу

					fmt.Println("Введите номер задачи, которую хотите удалить.")
					num := getNum()
					taskList = deleteTask(num, taskList)
				case 3: // Изменить задачу или её статус
					fmt.Println("Введите номер задачи, которую хотите изменить.")
					num := getNum()
					for {
						fmt.Println("1. Изменить задачу.\n2. Изменить статус задачи.\n3. Выйти из подменю.")

						switch selektorPodMenu := getMenuChoice(); selektorPodMenu {
						case 1: // Изменить задачу
							taskList = updateTask(num, taskList)
						case 2: // Изменить статус задачи.
							fmt.Println("Введите число,чтобы изменить статус задачи.\n1 - задача в процессе,\n2 - задача выполнена,\n3 - задача не выполнена.")
							switch selektorPodMenu := getMenuChoice(); selektorPodMenu {
							case 1:
								updateTaskStatus(taskList, num, 1)
							case 2:
								updateTaskStatus(taskList, num, 2)
							case 3:
								updateTaskStatus(taskList, num, 0)
							default:
								fmt.Println("Неверная команда.\nПопробуй ещё раз.\nВозвращаю в предыдущее меню.")
							}
						case 3: // Выйти из подменю
							break loop
						default:
							fmt.Println("Неверная команда.\nПопробуй ещё раз.")
						}
					}
				case 4: //Выйти в главное меню.
					break loop
				default:
					fmt.Println("Неверная команда.\nПопробуй ещё раз.")
				}
			}

		case 3: // Сохранить изменения в файле.
			err = saveToFile(taskList)
			if err != nil {
				fmt.Println("Ошибка сохранения файла: ", err)
			}
		case 4: // Ещё раз открыть инструкцию.
			hello(taskList)
		case 5: // Закрыть терминал.
			fmt.Println("Пока-пока")
			break outerloop
		default:
			fmt.Println("Неверная команда.\nПопробуй ещё разок.")
			hello(taskList)
		}
	}
	err = saveToFile(taskList)
	if err != nil {
		fmt.Println("Ошибка сохранения файла: ", err)
	}
}
