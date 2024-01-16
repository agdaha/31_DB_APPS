package main

import (
	"fmt"
	"log"
	"os"
	"skillfactory/31_DB_APPS/pkg/storage"
	"skillfactory/31_DB_APPS/pkg/storage/postgres"
)

func main() {
	var db storage.Data
	pwd := os.Getenv("dbpass")
	pwd = "postgres"
	db, err := postgres.New("postgres://postgres:" + pwd + "@localhost/tasks")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	defer db.Close()

	// Вывод всех задач на печать
	printAllTask(db)

	// Создание и удаление новой задачи
	createAndDelNewTask(db)

	// Вывод задач одного автора
	printTaskAuthor1(db)

	// Вывод задач определенной метки
	printTaskLabel2(db)

	//Обновление задачи
	db.UpdateTask(1, postgres.Task{
		AuthorID:   2,
		AssignedID: 1,
		Title:      "Updated task",
		Content:    "Updated content",
	})
	fmt.Println("***Обновление задачи с ИД=1***")
	printAllTask(db)

}

func printTaskLabel2(db storage.Data) {
	tasks, err := db.TasksWithLabel(2)
	if err != nil {
		log.Println("tasks labels read error: ", err)
	}
	fmt.Println("*** Вывод задач с тегом = 2***")
	printTasks(tasks)
}

func printTaskAuthor1(db storage.Data) {
	tasks, err := db.TasksWithAuthor(1)
	if err != nil {
		log.Println("tasks authors read error: ", err)
	}
	fmt.Println("*** Вывод задач автора с ИД = 1***")
	printTasks(tasks)
}

func createAndDelNewTask(db storage.Data) {
	id, err := db.NewTask(postgres.Task{
		AuthorID:   2,
		AssignedID: 1,
		Title:      "New Task 3",
		Content:    "New Task 3 content",
	})
	if err != nil {
		log.Println("create task error: ", err)
	}
	fmt.Println("*** Создана новая задача с id ***", id)
	printAllTask(db)
	db.DeleteTask(id)
	fmt.Println("*** Удалена задача с id ***", id)
	printAllTask(db)
}

func printAllTask(db storage.Data) {
	tasks, err := db.Tasks()
	if err != nil {
		log.Println("all tasks  read error: ", err)
	}
	fmt.Println("*** Вывод всех задач***")
	printTasks(tasks)
}

func printTasks(tasks []postgres.Task) {
	for _, t := range tasks {
		fmt.Printf("Ид: %v Заголовок: %v Содержание: %v ИД автора:%v \n", t.ID, t.Title, t.Content, t.AuthorID)
	}
}
