package cli

import (
	"ToDoList/locale"
	"ToDoList/service"
	"ToDoList/store"
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type CLI struct {
	service *service.TaskService
	reader  *bufio.Reader
	locale *locale.Manager
}

func NewCLI(service *service.TaskService, reader *bufio.Reader, localeManager *locale.Manager) *CLI {
	return &CLI{
		service: service,
		reader:  reader,
		locale:  localeManager,
	}
}

func (c *CLI) Run() error {
	c.showWelcome()

	for {
		c.showMainMenu()

		switch choice := c.getMenuChoice(); choice {
		case 1: // View task list
			c.listTasks()
		case 2: // Go to task editing menu ( Name >= 3 symbols)
			c.showEditMenu()
		case 3: // Show help again
			c.showHelp()
		case 4: // Change language
			c.showLanguageMenu()
		case 5: // Exit
			return c.exit()
		default:
			c.showError(c.locale.Get("invalid_choice"))
		}
	}
}

// ========== MAIN MENU METHODS ==========

// Shows main menu
func (c *CLI) showMainMenu() {

	fmt.Printf("\n%s\n", c.locale.Get("main_menu_title"))
	fmt.Println(c.locale.Get("main_menu_option1"))
	fmt.Println(c.locale.Get("main_menu_option2"))
	fmt.Println(c.locale.Get("main_menu_option3"))
	fmt.Println(c.locale.Get("main_menu_option4"))
	fmt.Println(c.locale.Get("main_menu_option5"))
	fmt.Print(c.locale.Get("choose_option"))

}

// ========== EDIT MENU METHODS ==========

// Shows task editing menu
func (c *CLI) showEditMenu() {
	for {

		fmt.Printf("\n%s\n", c.locale.Get("edit_menu_title"))
		fmt.Println(c.locale.Get("edit_menu_option1"))
		fmt.Println(c.locale.Get("edit_menu_option2"))
		fmt.Println(c.locale.Get("edit_menu_option3"))
		fmt.Println(c.locale.Get("edit_menu_option4"))
		fmt.Print(c.locale.Get("choose_option"))

		switch choice := c.getMenuChoice(); choice {
		case 1: // Add task
			c.createTask()
		case 2: // Delete task
			c.deleteTask()
		case 3: /// Edit task or its status
			c.editTask()
		case 4: // Back to main menu
			return
		default:
			c.showError(c.locale.Get("invalid_choice"))
		}
	}
}

// ========== TASK MANAGEMENT METHODS ==========

func (c *CLI) createTask() {
	fmt.Print(c.locale.Get("enter_new_task"))
	taskName, err := c.ReadString()
	if err != nil {
		c.showError(c.locale.GetFormatted("input_error", err))
		return
	}

	if err := c.service.CreateTask(taskName); err != nil {
		c.showError(err.Error())
		return
	}

	c.showSuccess(c.locale.Get("task_added"))
}

func (c *CLI) deleteTask() {

    tasks := c.service.GetAllTasks()
    if len(tasks) == 0 {
		c.showError(c.locale.Get("no_tasks_to_delete"))
        return 
    }

	c.listTasks()

	fmt.Print(c.locale.Get("enter_task_number"))
	taskNum, err := c.readInt()
	if err != nil {
		c.showError(c.locale.GetFormatted("number_error", err))
		return
	}

	if err := c.service.DeleteTask(taskNum); err != nil {
		c.showError(err.Error())
		return
	}

	c.showSuccess(c.locale.Get("task_deleted"))
}

func (c *CLI) editTask() {
	c.listTasks()

	fmt.Print(c.locale.Get("enter_task_number"))
	taskNum, err := c.readInt()
	if err != nil {
		c.showError(c.locale.GetFormatted("number_error", err))
		return
	}

	c.showEditSubMenu(taskNum)
}

// ========== EDIT SUBMENU METHODS ==========

func (c *CLI) showEditSubMenu(taskNum int) {
	for {
		fmt.Printf("\n%s\n", c.locale.Get("edit_submenu_title"))
		fmt.Println(c.locale.Get("edit_submenu_option1"))
		fmt.Println(c.locale.Get("edit_submenu_option2"))
		fmt.Println(c.locale.Get("edit_submenu_option3"))
		fmt.Print(c.locale.Get("choose_option"))

		switch choice := c.getMenuChoice(); choice {
		case 1:
			c.updateTaskText(taskNum)
		case 2:
			c.updateTaskStatus(taskNum)
		case 3:
			return
		default:
			c.showError(c.locale.Get("invalid_choice"))
		}
	}
}

func (c *CLI) updateTaskText(taskNum int) {

	fmt.Print(c.locale.Get("enter_new_description"))

	taskName, err := c.ReadString()
	if err != nil {
		c.showError(c.locale.GetFormatted("input_error", err))
		return
	}

	if err := c.service.UpdateTask(taskNum, taskName); err != nil {
		c.showError(err.Error())
		return
	}

	c.showSuccess(c.locale.Get("description_updated"))
}

func (c *CLI) updateTaskStatus(taskNum int) {

	fmt.Println(c.locale.Get("choose_status"))
	fmt.Println(c.locale.Get("status_option1"))
	fmt.Println(c.locale.Get("status_option2"))
	fmt.Println(c.locale.Get("status_option3"))
	fmt.Print(c.locale.Get("choose_option"))

	choice := c.getMenuChoice()
	var newStatus store.TaskStatus

	switch choice {
	case 1:
		newStatus = store.StatusNotDone
	case 2:
		newStatus = store.StatusInProgress
	case 3:
		newStatus = store.StatusDone
	default:
		c.showError(c.locale.Get("invalid_status"))
		return
	}

	if err := c.service.UpdateTaskStatus(taskNum, newStatus); err != nil {
		c.showError(err.Error())
		return
	}

	c.showSuccess(c.locale.Get("status_updated"))
}

// ========== LANGUAGE SETTINGS METHODS ==========

func (c *CLI) showLanguageMenu() {
	fmt.Printf("\n%s\n", c.locale.Get("language_menu"))
	fmt.Printf("%s\n", c.locale.GetFormatted("current_language", c.locale.CurrentLocale()))
	fmt.Println(c.locale.Get("available_languages"))
	
	availableLocales := c.locale.AvailableLocales()
	for i, lang := range availableLocales {
		fmt.Printf("%d. %s\n", i+1, lang)
	}
	fmt.Print(c.locale.Get("choose_option"))

	choice := c.getMenuChoice()
	if choice >= 1 && choice <= len(availableLocales) {
		selectedLang := availableLocales[choice-1]
		if err := c.locale.SetLocale(selectedLang); err != nil {
			c.showError(err.Error())
			return
		}
		// Update service locale as well
		c.service.SetLocale(c.locale)
		c.showSuccess(c.locale.GetFormatted("language_changed", selectedLang))
	} else {
		c.showError(c.locale.Get("invalid_choice"))
	}
}

// ========== DISPLAY METHODS ==========

func (c *CLI) showWelcome() {
	fmt.Printf("\n%s\n", c.locale.Get("welcome"))
	fmt.Printf("%s\n", c.locale.Get("welcome_description"))
	fmt.Printf("%s\n\n", c.locale.Get("interaction"))
	c.showHelp()
}

func (c *CLI) showHelp() {
	fmt.Printf("\n%s\n", c.locale.Get("instructions"))
	fmt.Println(c.locale.Get("help_option_1"))
	fmt.Println(c.locale.Get("help_option_2"))
	fmt.Println(c.locale.Get("help_option_3"))
	fmt.Println(c.locale.Get("help_option_4"))
	fmt.Println(c.locale.Get("help_option_5"))
	fmt.Println()
}

func (c *CLI) listTasks() {
	tasks := c.service.GetAllTasks()
	if len(tasks) == 0 {
		fmt.Println(c.locale.Get("empty_task_list"))
		return
	}

	fmt.Printf("\n%s\n", c.locale.Get("task_list_title"))
	for i, task := range tasks {
		statusText := c.service.GetStatusText(task.Status)
		fmt.Printf(c.locale.Get("task_format")+"\n", i+1, statusText, task.Task)
	}
	fmt.Printf(c.locale.Get("task_count")+"\n", len(tasks))
}

// ========== HELPER METHODS ==========

func (c *CLI) getMenuChoice() int {
	fmt.Print(">")

	r, _, _ := c.reader.ReadRune()

	if c.reader.Buffered() > 0 {
		c.reader.Discard(c.reader.Buffered())
	}

	if r >= '0' && r <= '9' {
		return int(r - '0')
	}
	return -1
}

func (c *CLI) ReadString() (string, error) {
	input, err := c.reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", fmt.Errorf(c.locale.Get("input_error"), err)
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
		return 0, fmt.Errorf(c.locale.Get("number_error"), err)
	}
	return num, nil
}

func (c *CLI) showError(message string) {
	fmt.Printf(c.locale.Get("error"), message)
}

func (c *CLI) showSuccess(message string) {
	fmt.Printf(c.locale.Get("success"), message)
	fmt.Println("")
}

func (c *CLI) exit() error {
	fmt.Println(c.locale.Get("goodbye"))
	return nil
}
