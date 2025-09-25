package main

import (
	"ToDoList/cli"
	"ToDoList/store"
	"ToDoList/service"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
)





func main() {

	//Хранилище
	s := store.NewJSONFileStore(store.StorageFileName)

	taskList, err := s.Load()
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
		if err := s.Save(taskList); err != nil {
			fmt.Println("Ошибка сохранения файла: ", err)
		}
		os.Exit(0)
	}()
	fmt.Printf("\nДобро пожаловать в терминал взаимодействия со списком дел.\nВзаимодействуй с терминалом через ввод цифр.\n\n")
	cli.Hello(taskList)

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
								service.updateTaskStatus(taskList, num, store.StatusInProgress)
							case 2:
								updateTaskStatus(taskList, num, store.StatusDone)
							case 3:
								updateTaskStatus(taskList, num, store.StatusNotDone)
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
			err = s.Save(taskList)
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
	err = s.Save(taskList)
	if err != nil {
		fmt.Println("Ошибка сохранения файла: ", err)
	}
}
