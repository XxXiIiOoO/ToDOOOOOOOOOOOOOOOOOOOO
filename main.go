package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID   int
	Name string
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		fmt.Println("Ашибка при сАздании БЭДЭ:", err)
		os.Exit(1)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		fmt.Println("Ашибка создания таблицы:", err)
		os.Exit(1)
	}
}

func addTask(name string) {
	insertQuery := "INSERT INTO tasks (name) VALUES (?)"
	_, err := db.Exec(insertQuery, name)
	if err != nil {
		fmt.Println("Ашибка дАбавления таблыцыыыы:", err)
	} else {
		fmt.Println("Ай маладэс таска создана.")
	}
}

func listTasks() {
	rows, err := db.Query("SELECT id, name FROM tasks")
	if err != nil {
		fmt.Println("EАшибка списка тасков:", err)
		return
	}
	defer rows.Close()

	if !rows.Next() {
		fmt.Println("Прикинь, нету доступных тасков братец.")
		return
	}

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Name)
		if err != nil {
			fmt.Println("Ашибка сканирования тастокв:", err)
			continue
		}
		fmt.Printf("%d: %s\n", task.ID, task.Name)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("Ашибка при повторном выполнении задачыыыыы:", err)
	}
}

func deleteTask(id int) {
	deleteQuery := "DELETE FROM tasks WHERE id = ?"
	_, err := db.Exec(deleteQuery, id)
	if err != nil {
		fmt.Println("Пойман на Ашибке при удалении таска:", err)
	} else {
		fmt.Printf("Таска %d с кайфом удалена.\n", id)
	}
}

func main() {
	initDB()
	defer db.Close()

	for {
		fmt.Println("ТУДУДУДУДУДУДУ ЛИСТИК МЕНЬЮЮЮЮС:")
		fmt.Println("1. дАбавить таску")
		fmt.Println("2. списОЧКА таска")
		fmt.Println("3. УДАЛЯЯЯЯЯЯЯЯЯЯЯЯЯЯЯЯЙ")
		fmt.Println("4. ВЫЫЫЫЫЫЫЫЫЫЫЫЫЫЙТИ")
		fmt.Print("Выбери па брацки: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var taskName string
			fmt.Print("Введи название таскиииии: ")
			fmt.Scanln(&taskName)
			addTask(taskName)
		case 2:
			listTasks()
		case 3:
			var taskID int
			fmt.Print("по АЙДИШНИКУ удали таску: ")
			fmt.Scan(&taskID)
			deleteTask(taskID)
		case 4:
			os.Exit(0)
		default:
			fmt.Println("ИНВЕЙЛИИИИД параметр, пж попробуй эгееейн.")
		}
	}
}
