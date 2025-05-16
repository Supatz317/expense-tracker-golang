# 💸 Expense Tracker CLI (Go)

This repository is an **Expense Tracker CLI application** written in Go. It allows users to manage their expenses through commands such as **adding, listing, updating, deleting, and summarizing expenses**. The application uses the [Cobra](https://github.com/spf13/cobra) library for building the CLI and supports persistent storage of expenses in a JSON file.

---

## ✨ Key Features

### 📝 Expense Management

- ➕ **Add** new expenses with descriptions and amounts.
- 📋 **List** all recorded expenses in a tabular format.
- ✏️ **Update** existing expenses by ID.
- ❌ **Delete** expenses by ID.
- **Summary** expenses by month or year

### 📊 Expense Summaries

- 💰 **View total expenses**.
- 📅 **View expenses filtered by a specific month**.

### 💾 Persistent Storage

- Expenses are stored in a JSON file (`data.json`) using a generic storage system implemented in `storage.go`.

### 🧩 Modular Design

- Commands are organized in the `cmd` package, making it easy to extend functionality.
- Expense-related logic is encapsulated in the `expense.go` file.

---

## 📦 Dependencies

- [`github.com/spf13/cobra`](https://github.com/spf13/cobra) for CLI functionality.
- [`github.com/aquasecurity/table`](https://github.com/aquasecurity/table) for rendering tabular data.

---

## 🗂️ File Structure

- **`main.go`**: Entry point of the application. Loads and saves expenses using the storage system.
- **`cmd/`**: Contains CLI commands (`add`, `list`, `update`, `delete`, `summary`).
- **`model/`**: Defines the `Expense` struct and related methods.
- **`repository/`**: Implements a generic storage system for saving and loading data.

---

## 🚀 How to Use
1. **Install dependencies** 
    ```sh
    go mod tidy
    ```

2. **Run the application** with commands like:

    ```sh
    go run main.go add --description "groceries" --amount 200
    go run main.go list
    go run main.go summary --month 4
    ```

---

