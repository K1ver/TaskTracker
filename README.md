# 🚀 Task Tracker

A beautiful CLI task management application built for the [task-tracker](https://roadmap.sh/projects/task-tracker) challenge from [roadmap.sh](https://roadmap.sh/).

---

## ✨ Features

- ✅ Add, update, and delete tasks
- 🏷️ Change task statuses (todo, in-progress, done)
- 📋 Filter tasks by status
- 💾 Local data storage

---

## 🚀 Quick Start

### Clone the Repository

```bash
git clone https://github.com/K1ver/TaskTracker.git
cd TaskTracker
```

### Build and Run

```bash
# Build the application
go build -o task-cli
```

---

## 🎯 Usage

### Add a Task
```bash
./task-cli add "Buy groceries"
```

### Update a Task
```bash
./task-cli update 1 "Buy groceries and cook dinner"
```

### Delete a Task
```bash
./task-cli delete 1
```

### Change Task Status
```bash
./task-cli mark-in-progress 1
./task-cli mark-done 1
./task-cli mark-todo 1
```

### List Tasks
```bash
./task-cli list           # All tasks
./task-cli list done      # Completed tasks
./task-cli list todo      # Planned tasks
./task-cli list in-progress # Tasks in progress
```

---

## 📦 Requirements

- Go 1.16+
- Compatible with Linux, macOS, Windows

---

## 🤝 Contributing

Pull requests and suggestions are welcome! For major changes, please open an issue first to discuss what you'd like to change.

*Built with ❤️ for the roadmap.sh community*