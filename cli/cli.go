package cli

import (
	"ToDoList/service"
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type CLI struct {
	service  *service.TaskService
	reader   *bufio.Reader
}

func NewCli(service *service.TaskService, reader *bufio.Reader) *CLI{
	return &CLI{
			service: service,
			reader: reader,
	}
}

func (c *CLI) Run() error {
	c.showWelcome()

	for{
		c.showMainMenu()

		switch choice:= c.getMenuChoice(); choice{
		case 1:// Посмотри список дел
			c.listTasks()
		case 2: // Перейти в меню редактирования задач.( Имя задачи >= 3 символов)
			c.showEditMenu()
		case 3: // Ещё раз открыть инструкцию.
			c.showMainMenu()
		case 4: // Закрыть терминал.
			return c.exit()
		default:
			c.showError("Неверная команда. Попробуйте снова.")
		}
		// Перенести запуск основных меню и выборов сюда
		// Обобщить их до методов CLI
	}
}

//Показывает главное меню
func (c *CLI) showMainMenu() {
	fmt.Println("\nГлавное меню:")
	fmt.Println("1. Посмотреть задачи")
	fmt.Println("2. Редактировать задачи")
	fmt.Println("4. Помощь")
	fmt.Println("4. Закрыть терминал.\n\n")
	fmt.Print("Выберите вариант: ")
}

//Показывает меню редактирования задач
func (c *CLI) showEditMenu(){
	for {
		fmt.Println("\nМеню редактирования:")
		fmt.Println("1. Добавить задачу.")
		fmt.Println("2. Удалить задачу.")
		fmt.Println("3. Изменить задачу или её статус.")
		fmt.Println("4. Назад в главное меню.")
		fmt.Print("Выберите вариант: > ")

		switch choice := c.getMenuChoice(); choice{
		case 1: // Добавить задачу
			c.createTask()
		case 2: // Удалить задачу
			c.deleteTask()
		case 3: // Изменить задачу или её статус
			c.editTask()
		case 4: //Выйти в главное меню
			return 
		default:
			c.showError("Неверная команда. Попробуйте снова.")
		}
	}
}

//Показывает подменю редактирования конкретной задачи
func (c *CLI) showEditSubMenu(taskNum int) {
	for {
		fmt.Println("\nЧто вы хотите сделать с задачей?")
		fmt.Println("1. Изменить текст задачи")
		fmt.Println("2. Изменить статус задачи")
		fmt.Println("3. Назад")
		fmt.Print("Выберите вариант: ")
		
		switch choice := c.getMenuChoice();choice {
		case 1:
			c.updateTaskText(taskNum)
		case 2:
			c.updateTaskStatus(taskNum)
		case 3:
			return
		default:
			c.showError("Неверная команда. Попробуйте снова.")
		}
	}
}

// Показывает приветственное сообщение
func (c *CLI) showWelcome() {
	fmt.Printf("\nДобро пожаловать в терминал списка дел.\n")
	fmt.Printf("Взаимодействуй с терминалом через ввод цифр.\n\n")
	c.showHelp()

}

func (c *CLI) listTasks(){
	tasks := c.service.GetAllTasks()
	if len(tasks) == 0 {
		fmt.Println("Список задач пуст.")
		return
	}

	fmt.Println("\nСписок задача:")
	for i, task := range tasks {
		statusText := c.service.GetStatusText(task.Status)
		fmt.Printf("№%v. [%s]  %s\n", i+1,statusText, task.Task)
	}
}

// Вспомогательные методы для ввода/вывода

func (c *CLI) getMenuChoice() int {
	fmt.Print(">")

	r, _, _ := c.reader.ReadRune()

	if c.reader.Buffered() > 0 {
		c.reader.Discard(c.reader.Buffered())
	}

	// Валидация
	if r >= '0' && r <= '9' {
		return int(r - '0')
	}
	return -1
}

func (c *CLI) ReadString()(string, error){
	input, err := c.reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("ошибка ввода %w", err)
	}
	return strings.TrimSpace(input), nil
}

func (c *CLI) readInt() (int, error) {
	input, err := c.ReadString()
	if err != nil {
		return 0, err
	}
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("ошибка конвертации строк", err)
	}
	return num, nil
}

func (c *CLI) showError(message string, args ...any){
	fmt.Printf("❌ Ошибка: "+message+"\n", args...)
}

func (c *CLI) showSuccess(message string) {
	fmt.Printf("✅ %s\n", message)	
}

func(c *CLI) exit() error {
	fmt.Println("До свидания!")
	return nil
}