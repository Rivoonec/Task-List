# 📝 Task Manager CLI

A powerful, multilingual command-line task management application built with Go. Manage your tasks efficiently with a beautiful terminal interface supporting multiple languages.

![Go Version](https://img.shields.io/badge/Go-1.24.4-blue)

## ✨ Features

- 🎯 **Simple Task Management** - Add, edit, delete, and track tasks
- 🌍 **Multilingual Support** - Switch between English and Russian seamlessly
- 💾 **Persistent Storage** - Automatic saving to JSON file
- 🎨 **Clean Interface** - Intuitive terminal-based menu system
- ⚡ **Fast & Lightweight** - Built with Go for optimal performance
- 🔄 **Status Tracking** - Track tasks as Not Done, In Progress, or Done
- 🛡 **Graceful Shutdown** - Automatic data preservation on exit

## 🏗 Project Structure

```
ToDoList/
├── 📁 locale/           # Internationalization
│   ├── locale.go       # Locale manager
│   └── messages.go     # Language messages
├── 📁 store/           # Data persistence
│   └── store.go        # JSON file storage
├── 📁 service/         # Business logic
│   └── service.go      # Task service layer
├── 📁 cli/             # User interface
│   └── cli.go          # Command-line interface
└── main.go             # Application entry point
```

## 🚀 Installation

### Prerequisites
- Go 1.24 or higher

### Build from Source

```bash
# Clone the repository
git clone https://github.com/Rivoonec/Task-List.git
cd task-manager-cli

# Build the application
go build -o task-manager main.go

# Run the application
./task-manager
```

### Using Go Run

```bash
go run main.go
```

## 📖 Usage

### Starting the Application

```bash
./task-manager
```

You'll be greeted with a welcome message and the main menu:

```
Welcome to Task Manager Terminal
Manage your tasks efficiently

Instructions:
1. View task list
2. Go to task editing menu (Task description >= 3 characters)
3. Show instructions again
4. Change language
5. Exit the application
```

### Main Menu Options

| Option | Description |
|--------|-------------|
| **1. View tasks** | Display all tasks with their current status |
| **2. Edit tasks** | Enter task management submenu |
| **3. Help** | Show instructions and usage guide |
| **4. Change language** | Switch between English and Russian |
| **5. Exit** | Quit the application |

### Task Management

#### Adding a Task
1. Select option 2 from main menu → "1. Add task"
2. Enter task description (minimum 3 characters)
3. Task is automatically saved with "Not Done" status

#### Editing Tasks
- **Change Description**: Update task text
- **Change Status**: 
  - 🔴 Not Done
  - 🟡 In Progress  
  - 🟢 Done

#### Deleting Tasks
- Confirmation prompt before deletion
- Tasks are permanently removed from storage

## 🌍 Language Support

The application supports multiple languages:

- **English** (default)
- **Russian** (Русский)

Switch languages anytime from the main menu (Option 4).

## 💾 Data Storage

Tasks are automatically saved to `list.json` in the current directory:

```json
[
  {
    "task": "Complete project documentation",
    "status": 1
  },
  {
    "task": "Review code changes", 
    "status": 2
  }
]
```

### Status Codes
- `0` - Not Done
- `1` - In Progress  
- `2` - Done

## 🛠 Technical Details

### Architecture

The application follows a clean architecture pattern:

```
CLI Layer (cli/) → Service Layer (service/) → Storage Layer (store/)
       ↓
  Localization (locale/)
```

### Key Components

- **`TaskService`**: Core business logic and validation
- **`TaskStore`**: Storage interface (JSON implementation)
- **`LocaleManager`**: Internationalization and text management  
- **`CLI`**: User interface and input handling

### Dependencies

- Pure Go standard library
- No external dependencies
- Cross-platform compatibility

## 🔧 Development

### Adding New Languages

1. Add new locale function in `messages.go`:

```go
func Spanish() *Locale {
    return &Locale{
        Language: "es",
        Messages: map[string]string{
            "app_name": "Gestor de Tareas",
            "welcome": "Bienvenido al Gestor de Tareas",
            // ... other messages
        },
    }
}
```

2. Register in `locale.go`:

```go
manager.RegisterLocale("es", Spanish())
```

### Extending Storage

Implement the `TaskStore` interface for different backends:

```go
type TaskStore interface {
    Save(tasks []Task) error
    Load() ([]Task, error)
}
```

## 📋 Requirements

- Task descriptions must be at least 3 characters long
- All changes are automatically persisted
- Application handles interrupt signals gracefully

## 🐛 Troubleshooting

### Common Issues

**"File load error"**: The storage file might be corrupted. Check `list.json` format.

**"Invalid task number"**: Ensure you're selecting valid task numbers from the list.

**Input issues**: The application expects numeric input for menu selections.

### Data Recovery

If the data file becomes corrupted, you can:
1. Delete `list.json` to start fresh
2. Manually fix the JSON format if familiar

**Happy Task Managing!** 🎉