package routes

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/spoonboy-io/koan"
	"github.com/spoonboy-io/todo-deployment-app/internal/templates"
	"html/template"
	"log"
)

func IndexHandler(c *fiber.Ctx, db *sql.DB, logger *koan.Logger) error {
	var res string
	var todos []string
	rows, err := db.Query("SELECT * FROM todos")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
	for rows.Next() {
		rows.Scan(&res)
		todos = append(todos, res)
	}
	var page bytes.Buffer
	templ, err := template.New("index").Parse(templates.Index())
	if err != nil {
		c.JSON("An error occurred")
	}

	_ = templ.Execute(&page, fiber.Map{
		"Todos": todos,
	})

	c.Set("Content-Type", "text/html")
	c.Response().BodyWriter().Write(page.Bytes())

	return nil
}

type Todo struct {
	Item string
}

func PostHandler(c *fiber.Ctx, db *sql.DB, logger *koan.Logger) error {
	newTodo := Todo{}
	if err := c.BodyParser(&newTodo); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}
	fmt.Printf("%v", newTodo)
	if newTodo.Item != "" {
		_, err := db.Exec("INSERT into todos VALUES ($1)", newTodo.Item)
		if err != nil {
			log.Fatalf("An error occured while executing query: %v", err)
		}
	}

	return c.Redirect("/")
}

func PutHandler(c *fiber.Ctx, db *sql.DB, logger *koan.Logger) error {
	olditem := c.Query("olditem")
	newitem := c.Query("newitem")
	db.Exec("UPDATE todos SET item=$1 WHERE item=$2", newitem, olditem)

	return c.Redirect("/")
}

func DeleteHandler(c *fiber.Ctx, db *sql.DB, logger *koan.Logger) error {
	todoToDelete := c.Query("item")
	db.Exec("DELETE from todos WHERE item=$1", todoToDelete)

	return c.SendString("deleted")
}