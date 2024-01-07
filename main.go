package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	database "github.com/wiratkhamphan/go-env-0/database"
)

// User represents a user in your application
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"username"`
	Email string `json:"email"`
	// Add other fields as needed
}

// func handler(c *fiber.Ctx) error {
// 	c.SendString("Hello, this is the handler function!")
// 	return nil
// }

func getUsersHandler(db *sql.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Execute the SELECT query to fetch all users
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			log.Fatal(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		defer rows.Close()

		// Process the query result
		var users []User
		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.Name, &user.Email)
			if err != nil {
				log.Fatal(err)
				return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
			users = append(users, user)
		}

		// Return the users as JSON
		return c.JSON(users)
	}
}
func getUsersHandlerID(db *sql.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// Retrieve the username from the URL parameter
		username := c.Params("username")

		// Execute the SELECT query to fetch users based on the specified username
		rows, err := db.Query("SELECT * FROM users WHERE username = ?", username)
		if err != nil {
			log.Fatal(err)
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		defer rows.Close()

		// Process the query result
		var users []User
		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.Name, &user.Email)
			if err != nil {
				log.Fatal(err)
				return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
			users = append(users, user)
		}

		// Return the users as JSON
		return c.JSON(users)
	}
}
func main() {
	// Connect to the database (replace 'your_dsn' with your actual database connection string)
	db, err := database.DBConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Test App v1.0.1",
	})

	app.Route("/shop", func(api fiber.Router) {
		api.Get("/users", getUsersHandler(db)).Name("foo")
		api.Get("/users/:username", getUsersHandlerID(db)).Name("bar")
	}, "test.")

	log.Fatal(app.Listen(":8080"))
}
