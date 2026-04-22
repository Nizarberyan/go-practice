package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"practice/task-manager/db/db"
	"practice/task-manager/tasks"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		fmt.Println("env Load error", err)
	}
	dbUrl := os.Getenv("GOOSE_DBSTRING")
	pool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		fmt.Println("db connection failed", err)
	}
	defer pool.Close()

	queries := db.New(pool)
	taskService := tasks.NewTaskService(queries)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("what would you like to do")
	fmt.Println("\n>")

	for {

		fmt.Println("what do you want to do \n")
		fmt.Println("add, list, quit")
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		switch input {
		case "add":
			fmt.Println("enter a title")
			scanner.Scan()
			title := scanner.Text()
			if title == "" {
				fmt.Println("title is required")
				return
			}
			fmt.Println("enter a description")
			scanner.Scan()
			descInput := scanner.Text()
			var pgDesc pgtype.Text
			if descInput == "" {
				pgDesc = pgtype.Text{
					Valid: false,
				}
			} else {
				pgDesc = pgtype.Text{
					String: descInput,
					Valid:  true,
				}
			}
			params := db.CreateTaskParams{
				Title:       title,
				Description: pgDesc,
			}
			taskService.CreateTask(ctx, params)
			fmt.Println(taskService.GetTasks(ctx))
		case "list":
			fmt.Println(taskService.GetTasks(ctx))
		case "quit":
			return

		default:
			fmt.Println("you need to choose an option")
		}

	}
}
