package rest

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"

	"github.com/mecitsemerci/go-todo-app/internal/infra/sqlitedb"
	"github.com/mecitsemerci/go-todo-app/internal/rest/handlers/auth"
	"github.com/mecitsemerci/go-todo-app/internal/rest/handlers/todo"
	"github.com/mecitsemerci/go-todo-app/internal/rest/middleware"
	"github.com/mecitsemerci/go-todo-app/internal/service"
)

func Register(app *fiber.App, logger *zap.Logger) error {
	api := app.Group("/api")

	db, err := sql.Open("sqlite3", "./tododb.db")

	if err != nil {
		return err
	}

	MakeMigrations(db)

	v1 := api.Group("v1")

	authorize := middleware.NewAuthorize(logger)

	registerAuth(db, logger, v1, authorize)
	registerTodo(db, logger, v1, authorize)

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})

	api.Get("/metrics", monitor.New())

	return nil
}

func registerAuth(db *sql.DB, logger *zap.Logger, v1 fiber.Router, authorize func(*fiber.Ctx) error) {
	repo := sqlitedb.NewUserRepository(db)
	authSrv := service.NewAuth(repo, logger)
	authHandler := auth.NewHandler(authSrv, logger)

	ag := v1.Group("/auth")
	ag.Post("/register", authHandler.Register)
	ag.Post("/login", authHandler.Login)

	protected := ag.Group("/", authorize)

	protected.Get("/profile", authHandler.GetProfile)
}

func registerTodo(db *sql.DB, logger *zap.Logger, v1 fiber.Router, authorize func(*fiber.Ctx) error) {
	// Implement todo handlers
	taskRepo := sqlitedb.NewTaskRepository(db)
	todoSrv := service.NewTodo(taskRepo, logger)
	todoHandler := todo.NewHandler(todoSrv, logger)
	todoGroup := v1.Group("/tasks")
	todoGroup.Post("/", authorize, todoHandler.Create)
	todoGroup.Get("/:id", authorize, todoHandler.Find)
	todoGroup.Put("/:id", authorize, todoHandler.Update)
	todoGroup.Delete("/:id", authorize, todoHandler.Delete)
	todoGroup.Get("/", authorize, todoHandler.GetAll)
}

func MakeMigrations(db *sql.DB) {
	stmtUsers := `CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		role VARCHAR(50) NOT NULL DEFAULT 'member',		
		last_login DATETIME DEFAULT CURRENT_TIMESTAMP,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(stmtUsers)

	if err != nil {
		log.Fatal(err)
	}

	stmtTasks := `CREATE TABLE IF NOT EXISTS tasks (
        id TEXT PRIMARY KEY NOT NULL,
        title TEXT NOT NULL,
        description TEXT,
        completed BOOLEAN NOT NULL DEFAULT FALSE,
		priority INTEGER NOT NULL CHECK (priority >= 1 AND priority <= 10),
        deleted BOOLEAN NOT NULL DEFAULT FALSE,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        user_id TEXT NOT NULL,
		FOREIGN KEY (user_id) REFERENCES users(id)
		ON DELETE CASCADE
		ON UPDATE CASCADE
	);`

	_, err = db.Exec(stmtTasks)

	if err != nil {
		log.Fatal(err)
	}
}
