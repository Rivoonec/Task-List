package cli

import (
	"ToDoList/service"
	"ToDoList/store"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func hello(taskList []store.Task) {
	fmt.Printf("\nИнструкция:\n\n")
	fmt.Printf("1. Посмотри список дел.\n")
	fmt.Printf("2. Перейти в меню редактирования задач.( Имя задачи >= 3 символов)\n")
	fmt.Printf("3. Сохранить изменения в файле.\n")
	fmt.Printf("4. Ещё раз открыть инструкцию.\n")
	fmt.Printf("5. Закрыть терминал.\n\n")
	service.ReadTaskList(taskList)
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
