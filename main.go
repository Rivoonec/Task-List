package main

import (
	"ToDoList/cli"
	"ToDoList/service"
	"ToDoList/store"
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Инициализация хранилища и сервиса
	taskStore := store.NewJSONFileStore(store.StorageFileName)
	taskService, err := service.NewTaskService(taskStore)
	if err != nil {
		log.Fatal("Ошибка при создании сервиса задач:", err)
	}

	// Обработка сигналов для graceful shutdown
	setupSignalHandling(taskService)

	// Инициализация и запуск CLI
	reader := bufio.NewReader(os.Stdin)
	appCLI := cli.NewCli(taskService, reader)

	if err := appCLI.Run(); err != nil {
		log.Fatal("Ошибка при работе приложения", err)
	}
}

func setupSignalHandling(service *service.TaskService) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nПолучен сигнал завершения. Сохраняем данные...")
		// Данные сохраняются автоматически при каждом изменении
		// Но на всякий случай можно вызвать явное сохранение
		if err := service.Store.Save(service.Tasks); err != nil {
			fmt.Println("Ошибка сохранения файла:", err)
		}
		os.Exit(0)
	}()
}
