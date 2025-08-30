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

func createTask(task string, taskList []string) []string {
	if task == "" || utf8.RuneCountInString(task) <= 3 {
		fmt.Println("Имя задачи должно быть не менее трёх символов.")
		return taskList
	}

	taskList = append(taskList, task)
	fmt.Println("Добавил задачу.")
	return taskList
}

func deleteTask(num int, taskList []string) []string {
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

func updateTask(num int, taskList []string) []string {
	for {
		num = num - 1
		for index := range taskList {
			if index == num {
				for {
					fmt.Println("Нашёл задачу.\nВведите новое имя.")
					input, err := reader.ReadString('\n')
					if err != nil && err != io.EOF {
						fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
						continue
					}
					if utf8.RuneCountInString(input) >= 3 || input != "" {
						taskList[index] = input
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

func readTaskList(taskList []string) {
	fmt.Println("Список задач:")
	for index, value := range taskList {
		fmt.Printf("№%v. %v\n", index+1, value)
	}
}

func hello(taskList []string) {
	fmt.Printf("\nДобро пожаловать в терминал взаимодействия со списком дел.\nВзаимодействуй с терминалом через ввод цифр.\n\n")
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

func saveToFile(taskList []string) error {
	data, err := json.Marshal(taskList)
	if err != nil {
		return err
	}

	return os.WriteFile("list.json", data, 0644)
}

func loadFromFile() ([]string, error) {
	data, err := os.ReadFile("list.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}

	var todos []string
	err = json.Unmarshal(data, &todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
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

	hello(taskList)

outerloop:
	for {
		fmt.Println("Ожидаю команды...\nЧтобы открыть инструкцию нажми 3.")

		switch selektorMenu := getMenuChoice(); selektorMenu {
		case 1:
			readTaskList(taskList)
		case 2:
		loop:
			for {
				fmt.Println("1. Добавить задачу.\n2. Удалить задачу.\n3. Изменить задачу.\n4. Выйти в главное меню.")

				switch selektorPodMenu := getMenuChoice(); selektorPodMenu {
				case 1:
					fmt.Println("Введите задачу.")
					input, err := reader.ReadString('\n')
					if err != nil && err != io.EOF {
						fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
						os.Exit(1)
					}
					taskList = createTask(input, taskList)
				case 2:

					fmt.Println("Введите номер задачи, которую хотите удалить.")
					num := getNum()
					taskList = deleteTask(num, taskList)
				case 3:
					fmt.Println("Введите имя задачи, которую хотите изменить.")
					num := getNum()
					taskList = updateTask(num, taskList)
				case 4:
					break loop
				default:
					fmt.Println("Неверная команда.\nПопробуй ещё раз.\nВозвращаю в главное меню.")
				}
			}

		case 3:
			err = saveToFile(taskList)
			if err != nil {
				fmt.Println("Ошибка сохранения файла: ", err)
			}
		case 4:
			hello(taskList)
		case 5:
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
