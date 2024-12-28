# Task Management Bot 📋🤖

This bot was written with a goal to learn Go a little bit. It is a simple task management bot built using **Go**. It integrates with the Telegram Bot API and provides users with functionality to manage their tasks directly from a Telegram chat.

## Features 🎯

-   **List Tasks**: Displays a list of all your current tasks.
-   **Add Tasks**: Add new tasks to your task list.
-   **Mark Tasks as Done**: Mark tasks as completed.
-   **Delete Tasks**: Remove tasks from your list.

## Installation 🚀

### Prerequisites

1.  **Go** (1.20 or higher): Ensure Go is installed. Download it [here](https://golang.org/dl/).
2.  **MongoDB**: The project uses MongoDB for task storage. Install MongoDB and ensure it's running.
3.  **Telegram Bot Token**: Create a bot using the BotFather on Telegram and obtain your bot token.

### Steps

1.  Clone the repository:
    
    `git clone https://github.com/yourusername/task-management-bot.git
    cd task-management-bot` 
    
2.  Set up environment variables: Create a `.env` file in the project root and add the following:
    
    `TELEGRAM_BOT_TOKEN=your_telegram_bot_token
    MONGO_URI=mongodb://localhost:27017
    DATABASE_NAME=taskdb` 
    
3.  Install dependencies:
    
    `go mod tidy` 
    
4.  Run the application:
    
    `go run main.go` 
    

----------

## Usage 🛠️

### Commands

-   `/start`: Initialize the bot and get a welcome message.
-   `/add [task]`: Add a new task (e.g., `/add Buy groceries`).
-   `/list`: List all tasks.
-   `/done [task_id]`: Mark a task as done (e.g., `/done 123`).
-   `/delete [task_id]`: Delete a task (e.g., `/delete 123`).

----------

## Project Structure 📂

```
telegram-tasks-bot/
├── cmd/                    # Entry points for applications
│   └── bot/                # Main bot application
│       └── main.go         # Entry point for the Telegram bot
├── internal/               # Internal application logic (cannot be imported by other projects)
│   ├── app/                # Core application logic (orchestration layer)
│   │   └── app.go          # Main app struct and dependencies
│   ├── bot/                # Bot-specific logic
│   │   ├── handlers/       # Handlers for commands like /add, /list, etc.
│   │   │   └── tasks.go    # Task-related handlers
│   │   ├── bot.go          # Telegram Bot initialization
│   ├── config/             # Configuration utilities
│   │   ├── config.go       # Load and parse configuration
│   │   └── env.go          # Environment variable handling
│   ├── model/              # Data models
│   │   └── task.go         # Task struct definition
│   ├── repository/         # Database access layer
│   │   └── task_repository.go  # CRUD operations for tasks
├── .env                    # Environment variables file
├── .gitignore              # Ignored files for Git
├── go.mod                  # Go module file
├── go.sum                  # Go dependencies file
├── README.md               # Project documentation
└── tasks.db                # SQLite database file (if applicable)
```


----------

## Task Representation 🗂️

### MongoDB Collection

Each task is stored in the `tasks` collection with the following schema:

`{
  "_id": "ObjectId",
  "task": "string",
  "done": "boolean"
}` 

Example:
`{
  "_id": "676fe000d92f53f6b92b223c",
  "task": "Buy milk",
  "done": false
}` 

----------

## Contribution 🤝

Contributions are welcome! Follow these steps:

1.  Fork the repository.
2.  Create a new branch:
    
    `git checkout -b feature/your-feature` 
    
3.  Commit your changes:
    
    `git commit -m "Add your feature"` 
    
4.  Push to the branch:
    
    ```git push origin feature/your-feature```
    
5.  Open a pull request.

----------

## License 📜

This project is licensed under the MIT License. See the `LICENSE` file for details.

----------

## Contact 📧

For issues or feature requests, feel free to open an issue or contact me via:

-   Email: gleymivan@icloud.com
-   Telegram: @GGleym

----------
