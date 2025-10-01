package locale

// English returns English locale messages
func English() *Locale {
    return &Locale{
        Language: "en",
        Messages: map[string]string{
            // ========== APPLICATION & SYSTEM MESSAGES ==========
            "app_name":        "Task Manager",
            "app_version":     "1.0.0",
            
            // ========== COMMON UI ELEMENTS ==========
            "error":           "❌ Error: %w",
            "success":         "✅ %s",
            "info":            "ℹ️ %s",
            "warning":         "⚠️ %s",
            "goodbye":         "Goodbye!",
            "invalid_choice":  "Invalid choice. Please try again.",
            "press_enter":     "Press Enter to continue...",
            
            // ========== WELCOME & HELP SECTION ==========
            "welcome":              "Welcome to Task Manager Terminal",
            "welcome_description":  "Manage your tasks efficiently",
            "interaction":          "Interact with the terminal by entering numbers",
            "instructions":         "Instructions:",
            "help_option_1":        "1. View task list",
            "help_option_2":        "2. Go to task editing menu (Task description >= 3 characters)",
            "help_option_3":        "3. Show instructions again", 
            "help_option_4":        "4. Change language",
            "help_option_5":        "5. Exit the application",
            
            // ========== LANGUAGE SETTINGS ==========
            "language_menu":        "Language Settings:",
            "current_language":     "Current language: %s",
            "select_language":      "Select language:",
            "language_changed":     "Language changed to: %s",
            "available_languages":  "Available languages:",
            
            // ========== MAIN MENU SECTION ==========
            "main_menu_title":      "Main Menu:",
            "main_menu_option1":    "1. View tasks",
            "main_menu_option2":    "2. Edit tasks", 
            "main_menu_option3":    "3. Help",
            "main_menu_option4":    "4. Change language",
            "main_menu_option5":    "5. Exit",
            "choose_option":        "Choose an option:",
            
            // ========== EDIT MENU SECTION ==========
            "edit_menu_title":      "Edit Menu:",
            "edit_menu_option1":    "1. Add task",
            "edit_menu_option2":    "2. Delete task",
            "edit_menu_option3":    "3. Edit task or its status",
            "edit_menu_option4":    "4. Back to main menu",
            
            // ========== TASK MANAGEMENT SECTION ==========
            "enter_new_task":       "Enter new task:",
            "task_added":           "Task successfully added!",
            "enter_task_number":    "Enter task number:",
            "task_deleted":         "Task successfully deleted!",
            "no_tasks_to_delete":   "No tasks to delete.",
            "confirm_delete":       "Are you sure you want to delete task: \"%s\"? (y/n):",
            "deletion_cancelled":   "Deletion cancelled",
            
            // ========== TASK LIST DISPLAY SECTION ==========
            "task_list_title":      "Task List:",
            "empty_task_list":      "Task list is empty.",
            "task_format":          "%d. [%s] %s",
            "task_count":           "Total tasks: %d",
            
            // ========== EDIT SUBMENU SECTION ==========
            "edit_submenu_title":    "What do you want to do with the task?",
            "edit_submenu_option1":  "1. Change task description",
            "edit_submenu_option2":  "2. Change task status", 
            "edit_submenu_option3":  "3. Back",
            
            // ========== STATUS MANAGEMENT SECTION ==========
            "choose_status":         "Choose new status:",
            "status_option1":        "1. Not done",
            "status_option2":        "2. In progress",
            "status_option3":        "3. Done",
            "invalid_status":        "Invalid status choice.",
            "status_updated":        "Task status successfully changed!",
            "description_updated":   "Task description successfully updated!",
            "enter_new_description": "Enter new task description:",
            
            // ========== VALIDATION ERRORS SECTION ==========
            "invalid_task_number":  "Invalid task number. Available numbers: 1-%d",
            "input_error":          "Input error: %w",
            "number_error":         "Number reading error: %w",
            "validation_error":     "Validation error: %w",
            
            // ========== FILE OPERATIONS SECTION ==========
            "file_save_success":    "Data successfully saved to file",
            "file_save_error":      "File save error: %w",
            "file_load_error":      "File load error: %w",
            "auto_save_enabled":    "Auto-save enabled",
            
            // ========== SIGNAL HANDLING SECTION ==========
            "signal_received":      "Received termination signal. Saving data...",
            "graceful_shutdown":    "Performing graceful shutdown...",
            
            // ========== SERVICE ERRORS SECTION ==========
            "service_init_error":   "Service initialization error: %%w",
            "service_save_error":   "Service save error: %w",
            "service_load_error":   "Service load error: %w",
            
            // ========== TASK VALIDATION SECTION ==========
            "task_validation_empty":    "Task description cannot be empty",
            "task_validation_short":    "Task description must be at least three characters long",
            "task_validation_number":   "Invalid task number: %d",
        },
    }
}

// Russian returns Russian locale messages
func Russian() *Locale {
    return &Locale{
        Language: "ru",
        Messages: map[string]string{
            // ========== APPLICATION & SYSTEM MESSAGES ==========
            "app_name":        "Менеджер задач",
            "app_version":     "1.0.0",
            
            // ========== COMMON UI ELEMENTS ==========
            "error":           "❌ Ошибка: %w",
            "success":         "✅ %s",
            "info":            "ℹ️ %s",
            "warning":         "⚠️ %s",
            "goodbye":         "До свидания!",
            "invalid_choice":  "Неверная команда. Попробуйте снова.",
            "press_enter":     "Нажмите Enter для продолжения...",
            
            // ========== WELCOME & HELP SECTION ==========
            "welcome":              "Добро пожаловать в терминал менеджера задач",
            "welcome_description":  "Эффективно управляйте вашими задачами",
            "interaction":          "Взаимодействуйте с терминалом через ввод цифр",
            "instructions":         "Инструкция:",
            "help_option_1":        "1. Посмотреть список задач",
            "help_option_2":        "2. Перейти в меню редактирования (Описание задачи >= 3 символов)",
            "help_option_3":        "3. Показать инструкцию снова",
            "help_option_4":        "4. Изменить язык",
            "help_option_5":        "5. Выйти из приложения",
            
            // ========== LANGUAGE SETTINGS ==========
            "language_menu":        "Настройки языка:",
            "current_language":     "Текущий язык: %s",
            "select_language":      "Выберите язык:",
            "language_changed":     "Язык изменен на: %s",
            "available_languages":  "Доступные языки:",
            
            // ========== MAIN MENU SECTION ==========
            "main_menu_title":      "Главное меню:",
            "main_menu_option1":    "1. Посмотреть задачи",
            "main_menu_option2":    "2. Редактировать задачи",
            "main_menu_option3":    "3. Помощь", 
            "main_menu_option4":    "4. Изменить язык",
            "main_menu_option5":    "5. Выйти",
            "choose_option":        "Выберите вариант: >",
            
            // ========== EDIT MENU SECTION ==========
            "edit_menu_title":      "Меню редактирования:",
            "edit_menu_option1":    "1. Добавить задачу",
            "edit_menu_option2":    "2. Удалить задачу",
            "edit_menu_option3":    "3. Изменить задачу или её статус",
            "edit_menu_option4":    "4. Назад в главное меню",
            
            // ========== TASK MANAGEMENT SECTION ==========  
            "enter_new_task":       "Введите новую задачу:",
            "task_added":           "Задача успешно добавлена!",
            "enter_task_number":    "Введите номер задачи:",
            "task_deleted":         "Задача успешно удалена!",
            "no_tasks_to_delete":   "Нет задач для удаления.",
            "confirm_delete":       "Вы уверены, что хотите удалить задачу: \"%s\"? (y/n):",
            "deletion_cancelled":   "Удаление отменено",
            
            // ========== TASK LIST DISPLAY SECTION ==========
            "task_list_title":      "Список задач:",
            "empty_task_list":      "Список задач пуст.",
            "task_format":          "%d. [%s] %s",
            "task_count":           "Всего задач: %d",
            
            // ========== EDIT SUBMENU SECTION ==========
            "edit_submenu_title":    "Что вы хотите сделать с задачей?",
            "edit_submenu_option1":  "1. Изменить описание задачи",
            "edit_submenu_option2":  "2. Изменить статус задачи",
            "edit_submenu_option3":  "3. Назад",
            
            // ========== STATUS MANAGEMENT SECTION ==========
            "choose_status":         "Выберите новый статус:",
            "status_option1":        "1. Не выполнена",
            "status_option2":        "2. В процессе", 
            "status_option3":        "3. Выполнена",
            "invalid_status":        "Неверный выбор статуса.",
            "status_updated":        "Статус задачи успешно изменен!",
            "description_updated":   "Описание задачи успешно обновлено!",
            "enter_new_description": "Введите новое описание задачи:",
            
            // ========== VALIDATION ERRORS SECTION ==========
            "invalid_task_number":  "Неверный номер задачи. Доступные номера: 1-%d",
            "input_error":          "Ошибка ввода: %w",
            "number_error":         "Ошибка чтения номера: %w",
            "validation_error":     "Ошибка валидации: %w",
            
            // ========== FILE OPERATIONS SECTION ==========
            "file_save_success":    "Данные успешно сохранены в файл",
            "file_save_error":      "Ошибка сохранения файла: %w",
            "file_load_error":      "Ошибка загрузки файла: %w",
            "auto_save_enabled":    "Автосохранение включено",
            
            // ========== SIGNAL HANDLING SECTION ==========
            "signal_received":      "Получен сигнал завершения. Сохраняем данные...",
            "graceful_shutdown":    "Выполняется плавное завершение...",
            
            // ========== SERVICE ERRORS SECTION ==========
            "service_init_error":   "Ошибка инициализации сервиса: %w",
            "service_save_error":   "Ошибка сохранения сервиса: %s",
            "service_load_error":   "Ошибка загрузки сервиса: %w",
            
            // ========== TASK VALIDATION SECTION ==========
            "task_validation_empty":    "Описание задачи не может быть пустым",
            "task_validation_short":    "Описание задачи должно быть не менее трёх символов",
            "task_validation_number":   "Неверный номер задачи: %d",
        },
    }
}