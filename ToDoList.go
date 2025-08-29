package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"unicode/utf8"
)

func createTask(task string)() {
	if task == "" || utf8.RuneCountInString(task) <= 3 {
		fmt.Println("Имя задачи должно быть не менее трёх символов.") 
		return 
	}

	taskList = append(taskList, task)
	return
}

func deleteTask(task string)() {
	if task == "" || utf8.RuneCountInString(task) < 3 {
		fmt.Println("Ошибка ввода или имя менее трёх символов.")
		return
	}
	for index, value := range taskList {
		if value == task {
			taskList = append(taskList[:index], taskList[index + 1:]...)
			fmt.Printf("Удалили задачу %v .\n", task)
			return
		}
	}
	fmt.Printf("Не нашли такую задачу.\nПовторите ввод.\n")
	input, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
    		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
   			os.Exit(1)
		}
	deleteTask(input)	
	return
}

func updateTask(task string)() {
	if task == "" || utf8.RuneCountInString(task) < 3 {
		fmt.Println("Ошибка ввода или имя менее трёх символов.")
		return
	}
	
	for index, value := range taskList {
		if value == task {
			fmt.Println("Нашёл задачу.\nВведите новое имя.")
			input, err := reader.ReadString('\n')
			if err != nil && err != io.EOF {
    			fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
   				os.Exit(1)
			}
			if utf8.RuneCountInString(input) >= 3 || input != "" {
				taskList[index] = input
				return
			} else {
				fmt.Println("Ошибка ввода или имя менее трёх символов.")
				updateTask(task)
				return
			}
		}
	}
	fmt.Printf("Не нашли такую задачу.\nПовторите ввод.\n")
	input, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
    		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
   			os.Exit(1)
		}
	updateTask(input)	
	return
}

func readTaskList()() {
	fmt.Println("Список задач:")
	for index, value := range taskList {
		fmt.Printf("№%v. %v\n", index, value)
	}
}

func hello()() {
	fmt.Printf("\n\nДобро пожаловать в терминал взаимодействия со списком дел.\nВзаимодействуй с терминалом через ввод цифр.\n\n")
	fmt.Printf("1. Посмотри список дел.\n\n")
	fmt.Printf("2. Перейти в меню редактирования задач.( Имя задачи >= 3 символов)\n\n")
	fmt.Printf("3. Ещё раз открыть инструкцию.\n\n")
	fmt.Printf("4. Закрыть терминал.\n\n")
	readTaskList()
}

func getMenuChoice()( int) {
	fmt.Print(">")
	r, _, _ :=reader.ReadRune()

	if reader.Buffered() > 0 {
		reader.Discard(reader.Buffered())
	}

	if r >= '0' && r <= '9' {
		return int(r -'0')
	}
	return -1
}
var (
	taskList [] string
	reader = bufio.NewReader(os.Stdin)
)

func main() {
	hello()
	outerloop:
	for {
		fmt.Println("Ожидаю команды...\nЧтобы открыть инструкцию нажми 3.")

		switch selektorMenu:= getMenuChoice(); selektorMenu {
		case 1:
			readTaskList()
		case 2:
			fmt.Println("1. Добавить задачу.\n2. Удалить задачу.\n3. Изменить задачу.\n4. Выйти в главное меню.")

			switch selektorPodMenu:= getMenuChoice(); selektorPodMenu {
			case 1:
				fmt.Println("Введите задачу.")
				input, err := reader.ReadString('\n')
				if err != nil && err != io.EOF {
    				fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
   					os.Exit(1)
				}
				createTask(input)
			case 2:
				fmt.Println("Введите имя задачи, которую хотите удалить.")
				input, err := reader.ReadString('\n')
				if err != nil && err != io.EOF {
    				fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
   					os.Exit(1)
				}
				deleteTask(input)
			case 3:
				fmt.Println("Введите имя задачи, которую хотите изменить.")
				input, err := reader.ReadString('\n')
				if err != nil && err != io.EOF {
    				fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
   					os.Exit(1)
				}
				updateTask(input)
			case 4:
				break
			default:
				fmt.Println("Неверная команда.\nПопробуй ещё раз.\nВозвращаю в главное меню.")
				break
			}			
		case 3:
			hello()
		case 4:
			fmt.Println("Пока-пока")
			break outerloop
		default:
			fmt.Println("Неверная команда.\nПопробуй ещё разок.")
			hello()
		}
	}
}