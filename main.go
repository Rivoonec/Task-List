package main

import (
	"ToDoList/locale"
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
	// Initialize locale manager
	localeManager := locale.NewManager()

	// Initialize storage and service
	taskStore := store.NewJSONFileStore(store.StorageFileName)
	taskService, err := service.NewTaskService(taskStore, localeManager)
	if err != nil {
		log.Fatal(localeManager.GetFormatted("service_init_error", err))
	}

	// Signal handling for graceful shutdown
	setupSignalHandling(taskService, localeManager)


	// Initialize and run CLI
	reader := bufio.NewReader(os.Stdin)
	appCLI := cli.NewCLI(taskService, reader, localeManager)

	if err := appCLI.Run(); err != nil {
		log.Fatal(localeManager.GetFormatted("service_init_error", err))
	}
}

func setupSignalHandling(service *service.TaskService, localeManager *locale.Manager) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Printf("\n%s\n", localeManager.Get("signal_received"))
		// Data is saved automatically with every change
		// But we call explicit save just in case
		if err := service.Store.Save(service.Tasks); err != nil {
			fmt.Printf(localeManager.Get("file_save_error"), err)
		}
		os.Exit(0)
	}()
}
