package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID          int
	Name        string
	Description string
}

var (
	tasks      []Task
	nextTaskID int = 1
	reader         = bufio.NewReader(os.Stdin)
)

func addTask(name, descr string) {
	t := Task{
		ID:          nextTaskID,
		Name:        name,
		Description: descr,
	}
	tasks = append(tasks, t)
	nextTaskID++
	fmt.Println("Задача добавлена.")
}

func getTask(id int) (*Task, error) {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i], nil
		}
	}
	return nil, fmt.Errorf("Задача с ID %d не найдена", id)
}

func editTask(id int, newName, newDescr string) error {
	task, err := getTask(id)
	if err != nil {
		return err
	}
	task.Name = newName
	task.Description = newDescr
	fmt.Println("Задача обновлена.")
	return nil
}

func deleteTask(id int) error {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Задача удалена.")
			return nil
		}
	}
	return fmt.Errorf("Задача с ID %d не найдена", id)
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("Список задач пуст.")
		return
	}
	fmt.Println("\nСписок задач:")
	for _, t := range tasks {
		fmt.Printf("ID: %d | %s — %s\n", t.ID, t.Name, t.Description)
	}
	fmt.Println()
}

func input(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func cliMenu() {
	for {
		fmt.Println("\n===== Меню задач =====")
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Просмотреть все задачи")
		fmt.Println("3. Редактировать задачу")
		fmt.Println("4. Удалить задачу")
		fmt.Println("5. Выйти")
		fmt.Print("Выбери действие: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, _ := strconv.Atoi(choiceStr)

		switch choice {
		case 1:
			name := input("Введите название задачи: ")
			descr := input("Введите описание: ")
			addTask(name, descr)
		case 2:
			listTasks()
		case 3:
			idStr := input("Введите ID задачи: ")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Ошибка: нужно ввести число.")
				continue
			}
			newName := input("Введите новое имя: ")
			newDescr := input("Введите новое описание: ")
			if err := editTask(id, newName, newDescr); err != nil {
				fmt.Println(err)
			}
		case 4:
			idStr := input("Введите ID задачи: ")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Ошибка: нужно ввести число.")
				continue
			}
			if err := deleteTask(id); err != nil {
				fmt.Println(err)
			}
		case 5:
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Неизвестная команда.")
		}
	}
}

func main() {
	fmt.Println("Добро пожаловать в Task Manager CLI!")
	cliMenu()
}
